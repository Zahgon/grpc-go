package peer

import (
	"context"
	"net"

	"google.golang.org/grpc/credentials"
)

type Peer struct {
	Addr net.Addr

	LocalAddr net.Addr

	AuthInfo credentials.AuthInfo
}

func (p *Peer) String() string { _ = "STUB: not implemented"; return "" }

type peerKey struct{}

func NewContext(ctx context.Context, p *Peer) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) (p *Peer, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}
