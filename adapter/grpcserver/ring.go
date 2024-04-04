package grpcserver

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/orbis-go/app"
	ringv1alpha1 "github.com/sourcenetwork/orbis-go/gen/proto/orbis/ring/v1alpha1"
	"github.com/sourcenetwork/orbis-go/pkg/crypto"
	"github.com/sourcenetwork/orbis-go/pkg/crypto/proof"
	"github.com/sourcenetwork/orbis-go/pkg/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ringService wraps application to provides gRPCs.
type ringService struct {
	ringv1alpha1.UnimplementedRingServiceServer

	app *app.App
}

func newRingService(a *app.App) *ringService {
	return &ringService{
		app: a,
	}
}

func (s *ringService) ListRings(ctx context.Context, req *ringv1alpha1.ListRingsRequest) (*ringv1alpha1.ListRingsResponse, error) {

	rings, err := s.app.ListRing(ctx)
	if err != nil {
		return nil, fmt.Errorf("list rings: %w", err)
	}

	ringResp := make([]*ringv1alpha1.Ring, len(rings))
	for i, r := range rings {
		ringResp[i] = &ringv1alpha1.Ring{
			Id:       string(r.ID),
			Manifest: r.Manifest(),
		}
	}

	return &ringv1alpha1.ListRingsResponse{
		Rings: ringResp,
	}, nil
}

func (s *ringService) CreateRing(ctx context.Context, req *ringv1alpha1.CreateRingRequest) (*ringv1alpha1.CreateRingResponse, error) {

	bgctx := context.Background()
	r, err := s.app.JoinRing(bgctx, req.Manifest)
	if err != nil {
		return nil, fmt.Errorf("create ring: %w", err)
	}

	resp := &ringv1alpha1.CreateRingResponse{
		Id: string(r.ID),
	}

	go func() {
		err = r.Start(bgctx)
		if err != nil {
			log.Error("start ring: %w", err)
		}
	}()

	return resp, nil
}

func (s *ringService) GetRing(ctx context.Context, req *ringv1alpha1.GetRingRequest) (*ringv1alpha1.GetRingResponse, error) {

	ring, err := s.app.GetRing(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "ring not found")
	}

	resp := &ringv1alpha1.GetRingResponse{
		Ring: &ringv1alpha1.Ring{
			Id:       string(ring.ID),
			Manifest: ring.Manifest(),
		},
	}

	return resp, nil
}

func (s *ringService) DeleteRing(ctx context.Context, req *ringv1alpha1.DeleteRingRequest) (*emptypb.Empty, error) {

	return nil, errUnimplemented
}

func (s *ringService) PublicKey(ctx context.Context, req *ringv1alpha1.PublicKeyRequest) (*ringv1alpha1.PublicKeyResponse, error) {

	ring, err := s.app.GetRing(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "ring not found")
	}

	publicKey, err := ring.DKG.PublicKey()
	if err != nil {
		return nil, err
	}

	protoPublicKey, err := crypto.PublicKeyToProto(publicKey)
	if err != nil {
		return nil, status.Error(codes.Internal, "can't get public key")
	}

	resp := &ringv1alpha1.PublicKeyResponse{
		PublicKey: protoPublicKey,
	}

	return resp, nil
}

func (s *ringService) Refresh(ctx context.Context, req *ringv1alpha1.RefreshRequest) (*ringv1alpha1.RefreshResponse, error) {

	resp := &ringv1alpha1.RefreshResponse{}

	return resp, errUnimplemented
}

func (s *ringService) State(ctx context.Context, req *ringv1alpha1.StateRequest) (*ringv1alpha1.StateResponse, error) {

	r, err := s.app.GetRing(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "ring not found")
	}

	states := r.State()
	services := make([]*ringv1alpha1.ServiceState, len(states))
	i := 0
	for name, state := range states {
		services[i] = &ringv1alpha1.ServiceState{
			Name:  name,
			State: state,
		}
		i++
	}
	resp := &ringv1alpha1.StateResponse{
		Services: services,
	}

	return resp, nil
}

func (s *ringService) ListSecrets(ctx context.Context, req *ringv1alpha1.ListSecretsRequest) (*ringv1alpha1.ListSecretsResponse, error) {
	return nil, errUnimplemented
}

func (s *ringService) StoreSecret(ctx context.Context, req *ringv1alpha1.StoreSecretRequest) (*ringv1alpha1.StoreSecretResponse, error) {

	r, err := s.app.GetRing(ctx, req.RingId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "ring not found")
	}

	secret := &types.Secret{
		Secret: &ringv1alpha1.Secret{
			EncCmt:   req.Secret.EncCmt,
			EncScrt:  req.Secret.EncScrt,
			AuthzCtx: req.Secret.AuthzCtx,
		},
	}

	sid, err := r.StoreSecret(ctx, r.ID, secret)
	if err != nil {
		return nil, fmt.Errorf("store secret: %w", err)
	}

	resp := &ringv1alpha1.StoreSecretResponse{
		SecretId: string(sid),
	}

	return resp, nil
}

func (s *ringService) ReencryptSecret(ctx context.Context, req *ringv1alpha1.ReencryptSecretRequest) (*ringv1alpha1.ReencryptSecretResponse, error) {
	log.Infof("ReencryptSecret(): request: ringid=%s secretid=%s", req.RingId, req.SecretId)
	r, err := s.app.GetRing(ctx, req.RingId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "ring not found")
	}

	token, err := r.Authn.GetRequestToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "missing authentication token")
	}

	authInfo, err := r.Authn.VerifyRequestSubject(ctx, token)
	if err != nil {
		log.Error(err)
		return nil, status.Errorf(codes.Unknown, "failed to verify token: %s", err)
	}

	log.Infof("ReencryptSecret(): get secret: secretid=%s", req.SecretId)
	scrt, err := r.GetSecret(ctx, req.SecretId)
	if err != nil {
		return nil, err
	}

	log.Infof("ReencryptSecret(): authz.Check(): perm='%s' subject='%s'", scrt.AuthzCtx, authInfo.Subject)
	ok, err := r.Authz.Check(ctx, scrt.AuthzCtx, "user:"+authInfo.Subject)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "permission denied: "+err.Error())
	}
	if !ok {
		return nil, errUnAuthorized
	}

	var p proof.VerifiableEncryption
	log.Infof("ReencryptSecret(): running reencryption")
	xncCmt, encScrt, err := r.ReencryptSecret(ctx, authInfo.PubKey, types.SecretID(req.SecretId), p)
	if err != nil {
		return nil, fmt.Errorf("reencrypt secret: %w", err)
	}

	log.Infof("ReencryptSecret(): completed reencryption succesfully")
	resp := &ringv1alpha1.ReencryptSecretResponse{
		XncCmt:  xncCmt,
		EncScrt: encScrt,
	}

	return resp, nil
}

func (s *ringService) DeleteSecret(ctx context.Context, req *ringv1alpha1.DeleteSecretRequest) (*emptypb.Empty, error) {
	return nil, errUnimplemented
}
