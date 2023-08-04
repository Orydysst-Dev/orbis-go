package app

import (
	"context"
	"fmt"

	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/samber/do"

	ringv1alpha1 "github.com/sourcenetwork/orbis-go/gen/proto/orbis/ring/v1alpha1"
	"github.com/sourcenetwork/orbis-go/pkg/authn"
	"github.com/sourcenetwork/orbis-go/pkg/authz"
	"github.com/sourcenetwork/orbis-go/pkg/bulletin"
	"github.com/sourcenetwork/orbis-go/pkg/crypto"
	"github.com/sourcenetwork/orbis-go/pkg/crypto/proof"
	"github.com/sourcenetwork/orbis-go/pkg/db"
	"github.com/sourcenetwork/orbis-go/pkg/dkg"
	"github.com/sourcenetwork/orbis-go/pkg/pre"
	"github.com/sourcenetwork/orbis-go/pkg/pss"
	"github.com/sourcenetwork/orbis-go/pkg/transport"
	p2ptransport "github.com/sourcenetwork/orbis-go/pkg/transport/p2p"
	"github.com/sourcenetwork/orbis-go/pkg/types"
)

var log = logging.Logger("orbis/ring")

type Ring struct {
	ID types.RingID

	DKG dkg.DKG
	PSS pss.PSS
	PRE pre.PRE

	Authz    authz.Authz
	Authn    authn.CredentialService
	Resolver authn.KeyResolver

	// collection of registered services
	// that require startup/shutdown and
	// expose hooks.
	services []service

	Transport transport.Transport
	Bulletin  bulletin.Bulletin
	DB        *db.DB

	N int
	T int

	nodes []transport.Node

	inj *do.Injector
}

type State map[string]string

type service interface {
	Start(context.Context) error
	Close(context.Context) error
	State() string
	Name() string
}

/*
ring1
manifest := {
	"N": 9,
	"T": 7,
	"curve": "Ed25519",

	"dkg": "rabin",
	"pss": "avpss",
	"pre": "elgamal",
	"bulletin": "sourcehub",
	"transport": "libp2p"
}

ring2
manifest := {
	"N": 9,
	"T": 7,
	"curve": "Ed25519",

	"dkg": "rabin",
	"pss": "avpss",
	"pre": "elgamal",
	"bulletin": "sourcehub",
	"transport": "libp2p"
}

ring3
manifest := {
	"N": 9,
	"T": 7,
	"curve": "Ed25519",

	"dkg": "rabin",
	"pss": "avpss",
	"pre": "elgamal",
	"bulletin": "cosmos:sourcehub",
	"transport": "libp2p"
}
*/

/*

nodeconfig {
	bulletin:
		cosmos:
			sourcehub:
				rpcURL: 'localhost:1234/rpc'
				chainID: sourcehub-1
}

*/

func (app *App) JoinRing(ctx context.Context, ring *ringv1alpha1.Ring) (*Ring, error) {
	app.mu.Lock()
	defer app.mu.Unlock()
	r, err := app.joinRing(ctx, ring, false /* fromState */)
	if err != nil {
		return nil, err
	}
	app.rings[r.ID] = r
	err = app.ringRepo.Create(ctx, ring)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (app *App) joinRing(ctx context.Context, ring *ringv1alpha1.Ring, fromState bool) (*Ring, error) {
	rid := types.RingID(ring.Id)
	app.mu.Lock()
	defer app.mu.Unlock()

	if _, exists := app.rings[rid]; exists {
		return nil, fmt.Errorf("already joined ring %s", rid)
	}

	rs := &Ring{}

	// rings get their own cloned dependency injector handler,
	// since we actually create and initialize services which
	// are only scoped to rings, compared to the factories
	// which are global.
	inj := app.inj.Clone()

	// factories
	authnFactory, err := do.InvokeNamed[types.Factory[authn.CredentialService]](inj, ring.Authentication)
	if err != nil {
		return nil, fmt.Errorf("invoke authn credential service: %w", err)
	}

	dkgFactory, err := do.InvokeNamed[types.Factory[dkg.DKG]](inj, ring.Dkg)
	if err != nil {
		return nil, fmt.Errorf("invoke dkg factory: %w", err)
	}

	pssFactory, err := do.InvokeNamed[types.Factory[pss.PSS]](inj, ring.Pss)
	if err != nil {
		return nil, fmt.Errorf("invoke pss factory: %w", err)
	}

	preFactory, err := do.InvokeNamed[types.Factory[pre.PRE]](inj, ring.Pre)
	if err != nil {
		return nil, err
	}

	// get global services
	d, err := do.Invoke[*db.DB](inj)
	if err != nil {
		return nil, fmt.Errorf("invoke db: %w", err)
	}

	// register configured generic transport locally
	tp, err := do.InvokeNamed[transport.Transport](inj, ring.Transport)
	if err != nil {
		return nil, fmt.Errorf("invoke transport: %w", err)
	}
	do.ProvideValue(inj, tp)

	// register configured generic bulletin locally
	bb, err := do.InvokeNamed[bulletin.Bulletin](inj, ring.Bulletin)
	if err != nil {
		return nil, fmt.Errorf("invoke bulletin: %w", err)
	}
	do.ProvideValue(inj, bb)

	authz, err := do.InvokeNamed[authz.Authz](inj, ring.Authorization)
	if err != nil {
		return nil, fmt.Errorf("invoke authz: %w", err)
	}
	do.ProvideValue(inj, authz)

	// setup and register local services
	authn, err := authnFactory.New(inj, []db.RepoKey{}) // empty repo keys
	if err != nil {
		return nil, fmt.Errorf("create authn: %w", err)
	}

	dkgRepoKeys := app.repoKeysForService(dkgFactory.Name())
	log.Debugf("dkg repo keys: %v", dkgRepoKeys)
	dkgSrv, err := dkgFactory.New(inj, dkgRepoKeys)
	if err != nil {
		return nil, fmt.Errorf("create dkg service: %w", err)
	}

	nodes, err := nodesFromIDs(ring.Nodes)
	if err != nil {
		return nil, fmt.Errorf("conver nodes from ring ids")
	}

	err = dkgSrv.Init(ctx, app.privateKey, rid, nodes, ring.N, ring.T, fromState)
	if err != nil {
		return nil, fmt.Errorf("initializing dkg: %w", err)
	}
	do.ProvideValue(inj, dkgSrv)

	err = rs.registerService(dkgSrv)
	if err != nil {
		return nil, fmt.Errorf("starting dkg: %w", err)
	}

	preRepoKeys := app.repoKeysForService(preFactory.Name())
	preSrv, err := preFactory.New(inj, preRepoKeys)
	if err != nil {
		return nil, fmt.Errorf("create pre service: %w", err)
	}

	err = preSrv.Init(rid, ring.N, ring.T, []types.Node{})
	if err != nil {
		return nil, fmt.Errorf("initialize pre: %w", err)
	}
	do.ProvideValue(inj, preSrv)

	pssRepoKeys := app.repoKeysForService(pssFactory.Name())
	pssSrv, err := pssFactory.New(inj, pssRepoKeys)
	if err != nil {
		return nil, fmt.Errorf("create pss service: %w", err)
	}

	err = pssSrv.Init(rid, ring.N, ring.T, []types.Node{})
	if err != nil {
		return nil, fmt.Errorf("create pss service: %w", err)
	}

	rs = &Ring{
		ID:        rid,
		DKG:       dkgSrv,
		PSS:       pssSrv,
		PRE:       preSrv,
		Transport: tp,
		Bulletin:  bb,
		DB:        d,
		inj:       inj,
		N:         int(ring.N),
		T:         int(ring.T),

		services: rs.services, // this is dumb, but im being lazy, sorry.
		Authz:    authz,
		Authn:    authn,
	}

	app.rings[rs.ID] = rs

	// called in ring.Join() - go rs.handleEvents()

	return rs, nil
}

func (app *App) GetRing(ctx context.Context, rid types.RingID) (*Ring, error) {
	return nil, nil
}

func nodesFromIDs(nodes []*ringv1alpha1.Node) ([]transport.Node, error) {

	var tNodes []transport.Node
	for _, n := range nodes {

		id, err := peer.Decode(n.Id)
		if err != nil {
			return nil, fmt.Errorf("invalid peer id: %w", err)
		}

		pubKey, err := id.ExtractPublicKey()
		if err != nil {
			return nil, fmt.Errorf("extract publick key from id: %w", err)
		}

		key, err := crypto.PublicKeyFromLibP2P(pubKey)
		if err != nil {
			return nil, fmt.Errorf("invalid public key: %w", err)
		}

		addr, err := ma.NewMultiaddr(n.Address)
		if err != nil {
			return nil, fmt.Errorf("invalid address: %w", err)
		}

		node := p2ptransport.NewNode(n.Id, key, addr)
		tNodes = append(tNodes, node)
	}

	return tNodes, nil
}

func (r *Ring) Store(context.Context, types.SecretID, *types.Secret, proof.VerifiableEncryption) error {
	return nil
}

func (r *Ring) Get(context.Context, types.SecretID) (types.Secret, error) {
	return types.Secret{}, nil
}

func (r *Ring) GetShares(context.Context, types.SecretID) ([]types.PrivSecretShare, error) {
	return nil, nil
}

func (r *Ring) Delete(context.Context, types.SecretID) error {
	return nil
}

func (r *Ring) PublicKey() (crypto.PublicKey, error) {
	return r.DKG.PublicKey()
}

func (r *Ring) Refresh(context.Context, pss.Config) (pss.RefreshState, error) {
	return pss.RefreshState{}, nil
}

func (r *Ring) Threshold() int {
	return 0
}

func (r *Ring) State() State {
	state := make(State)
	for _, s := range r.services {
		state[s.Name()] = s.State()
	}
	return state
}

func (r *Ring) Nodes() []pss.Node {
	return nil
}

func (r *Ring) registerService(srv service) error {
	if srv == nil {
		return fmt.Errorf("service can't be nil")
	}
	r.services = append(r.services, srv)
	return nil
}

func (r *Ring) Start(ctx context.Context) error {
	var err error
	for _, srv := range r.services {
		err = srv.Start(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// LoadRings loads any existing rings into state
func (app *App) LoadRings(ctx context.Context) error {
	app.mu.Lock()
	defer app.mu.Unlock()

	rings, err := app.ringRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, r := range rings {
		ring, err := app.joinRing(ctx, r, true /* fromState */)
		if err != nil {
			return err
		}

		app.rings[ring.ID] = ring
	}

	return nil
}
