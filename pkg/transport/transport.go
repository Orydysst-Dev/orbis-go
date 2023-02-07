package transport

import (
	"context"

	"github.com/libp2p/go-libp2p/core/protocol"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/sourcenetwork/orbis-go/pkg/crypto"
)

type Handler func(Message) error

type Transport interface {
	Send(ctx context.Context, node Node, msg Message) error
	Gossip(ctx context.Context, topic string, msg Message) error
	Connect(ctx context.Context, node Node) error
	Host() Host
	NewMessage(id string, gossip bool, payload []byte, msgType string) (Message, error)
	AddHandler(pid protocol.ID, handler Handler)
	RemoveHandler(pid protocol.ID)
}

type Node interface {
	ID() string
	PublicKey() crypto.PublicKey
	Address() ma.Multiaddr
}

type Host interface {
	Node
	Sign(msg []byte) ([]byte, error)
}