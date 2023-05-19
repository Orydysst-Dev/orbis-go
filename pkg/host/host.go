package host

import (
	"context"
	"crypto/rand"
	"fmt"
	mrand "math/rand"
	"sync"
	"time"

	"github.com/sourcenetwork/orbis-go/config"

	logging "github.com/ipfs/go-log"
	libp2p "github.com/libp2p/go-libp2p"
	libp2pdht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	libp2pcrypto "github.com/libp2p/go-libp2p/core/crypto"
	libp2phost "github.com/libp2p/go-libp2p/core/host"
	libp2ppeer "github.com/libp2p/go-libp2p/core/peer"
	libp2pprotocol "github.com/libp2p/go-libp2p/core/protocol"
	libp2prouting "github.com/libp2p/go-libp2p/core/routing"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
	libp2pconnmgr "github.com/libp2p/go-libp2p/p2p/net/connmgr"
	libp2pnoise "github.com/libp2p/go-libp2p/p2p/security/noise"
	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
)

var log = logging.Logger("orbis/p2p")

type Host struct {
	libp2phost.Host
	idht   *libp2pdht.IpfsDHT
	pubsub *pubsub.PubSub
	topics map[string]*pubsub.Topic
}

func New(ctx context.Context, cfg config.Host) (*Host, error) {

	// Convert string to libp2p crypto type.
	// Invalid types and/or bits are handled by libp2p.
	cryptoType := libp2pcrypto.RSA
	switch cfg.Crypto.Type {
	case "ed25519":
		cryptoType = libp2pcrypto.Ed25519
	case "secp256k1":
		cryptoType = libp2pcrypto.Secp256k1
	case "ecdsa":
		cryptoType = libp2pcrypto.ECDSA
	}

	randomness := rand.Reader
	if seed := cfg.Crypto.Seed; seed != 0 {
		randomness = mrand.New(mrand.NewSource(int64(seed)))
	}

	priv, _, err := libp2pcrypto.GenerateKeyPairWithReader(cryptoType, cfg.Crypto.Bits, randomness)
	if err != nil {
		return nil, fmt.Errorf("generate key pair: %w", err)
	}

	connmgr, err := libp2pconnmgr.NewConnManager(
		100, // Lowwater
		400, // HighWater,
		libp2pconnmgr.WithGracePeriod(time.Minute),
	)
	if err != nil {
		return nil, fmt.Errorf("create connection manager: %w", err)
	}

	var idht *libp2pdht.IpfsDHT

	var dhtOptions []libp2pdht.Option
	if len(cfg.BootstrapPeers) == 0 {
		log.Infof("Host running as a bootsrap node")
		dhtOptions = append(dhtOptions, libp2pdht.Mode(libp2pdht.ModeServer))
	}

	h, err := libp2p.New(
		// Use the keypair we generated
		libp2p.Identity(priv),
		// Multiple listen addresses
		libp2p.ListenAddrStrings(cfg.ListenAddresses...),
		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),
		// support noise connections
		libp2p.Security(libp2pnoise.ID, libp2pnoise.New),
		// support any other default transports (TCP)
		libp2p.DefaultTransports,
		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		libp2p.ConnectionManager(connmgr),
		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),
		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h libp2phost.Host) (libp2prouting.PeerRouting, error) {
			idht, err = libp2pdht.New(ctx, h, dhtOptions...)
			return idht, err
		}),
		libp2p.EnableNATService(),
	)
	if err != nil {
		return nil, fmt.Errorf("create libp2p host: %w", err)
	}

	gossipSub, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}

	// log.Infof("listenaddress peers: %v", cfg.ListenAddresses)
	// log.Infof("bootstrap peers: %v", cfg.BootstrapPeers)

	host := &Host{
		Host:   h,
		idht:   idht,
		pubsub: gossipSub,
		topics: map[string]*pubsub.Topic{},
	}

	host.Bootstrap(ctx, cfg)

	return host, nil
}

func (h *Host) Subscribe(ctx context.Context, topic string) (*pubsub.Subscription, error) {

	t, err := h.join(topic)
	if err != nil {
		return nil, err
	}

	return t.Subscribe()

	// sub, err := t.Subscribe()
	// if err != nil {
	// 	return err
	// }

	// for {
	// 	msg, err := sub.Next(ctx)
	// 	if err == context.Canceled {
	// 		break
	// 	}
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	if msg.ReceivedFrom == h.ID() {
	// 		continue
	// 	}
	// 	h.log.Infof("%s | %s > %s", topic, msg.ReceivedFrom.Pretty(), msg.Data)
	// }

	// return nil
}

func (h *Host) Publish(ctx context.Context, topic string, data []byte) error {

	t, err := h.join(topic)
	if err != nil {
		return err
	}

	err = t.Publish(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (h *Host) join(topic string) (*pubsub.Topic, error) {

	t, exists := h.topics[topic]
	if exists {
		return t, nil
	}

	t, err := h.pubsub.Join(topic)
	if err != nil {
		return nil, err
	}
	h.topics[topic] = t

	return t, nil
}

func (h *Host) Bootstrap(ctx context.Context, cfg config.Host) {
	var wg sync.WaitGroup
	for _, peerAddr := range cfg.BootstrapPeers {
		pi, err := libp2ppeer.AddrInfoFromString(peerAddr)
		if err != nil {
			log.Warnf("Can't parse peer addr info string: %s, %s", pi, err)
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := h.Connect(ctx, *pi); err != nil {
				log.Warnf("Can't connect to peer: %s, %s", pi, err)
			} else {
				log.Infof("Connected to bootstrap node: %s", pi)
			}
		}()
	}
	wg.Wait()

}

func (h *Host) Discover(ctx context.Context, rendezvous string) {

	pi := libp2ppeer.AddrInfo{
		ID:    h.ID(),
		Addrs: h.Addrs(),
	}
	log.Infof("Announcing ourselves: %s", pi)
	d := drouting.NewRoutingDiscovery(h.idht)
	dutil.Advertise(ctx, d, rendezvous)

	log.Infof("Searching for other peers...")
	peerChan, err := d.FindPeers(ctx, rendezvous)
	if err != nil {
		panic(err)
	}

	go func() {
		defer log.Infof("Peer discovery finished")
		for peer := range peerChan {
			if peer.ID == h.ID() {
				continue
			}

			if len(peer.Addrs) == 0 {
				continue
			}

			err = h.Connect(ctx, peer)
			if err != nil {
				log.Warnf("Connection failed:", err)
				continue
			}

			// stream, err := host.NewStream(ctx, peer.ID, libp2pprotocol.ID(protocolID))
			// if err != nil {
			// 	log.Warnf("Connection failed:", err)
			// 	continue
			// }
			// rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

			// go writeData(rw)
			// go readData(rw)

			log.Infof("Connected to: %s", peer)
		}
	}()
}

func (h *Host) Peers() []string {

	var peers []string
	s := h.Network().Peerstore()
	for _, p := range h.Network().Peers() {
		a := s.PeerInfo(p)
		peers = append(peers, a.String())

	}

	return peers
}

func (h *Host) Send(ctx context.Context, pi libp2ppeer.AddrInfo, protocol string, data []byte) error {

	stream, err := h.NewStream(ctx, pi.ID, libp2pprotocol.ID(protocol))
	if err != nil {
		return fmt.Errorf("new stream: %w", err)
	}
	defer stream.Close()

	_, err = stream.Write(data)
	if err != nil {
		return fmt.Errorf("write data: %w", err)
	}

	return nil
}

func (h *Host) PubSub() *pubsub.PubSub {
	return h.pubsub
}
