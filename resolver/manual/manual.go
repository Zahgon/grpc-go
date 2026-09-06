package manual

import (
	"sync"

	"google.golang.org/grpc/resolver"
)

func NewBuilderWithScheme(scheme string) *Resolver { _ = "STUB: not implemented"; return nil }

type Resolver struct {
	BuildCallback func(resolver.Target, resolver.ClientConn, resolver.BuildOptions)

	UpdateStateCallback func(err error)

	ResolveNowCallback func(resolver.ResolveNowOptions)

	CloseCallback func()
	scheme        string

	mu sync.Mutex
	cc resolver.ClientConn

	lastSeenState *resolver.State
}

func (r *Resolver) InitialState(s resolver.State) { _ = "STUB: not implemented"; return }

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Resolver), nil
}

func (r *Resolver) Scheme() string { _ = "STUB: not implemented"; return "" }

func (r *Resolver) ResolveNow(o resolver.ResolveNowOptions) { _ = "STUB: not implemented"; return }

func (r *Resolver) Close() { _ = "STUB: not implemented"; return }

func (r *Resolver) UpdateState(s resolver.State) { _ = "STUB: not implemented"; return }

func (r *Resolver) CC() resolver.ClientConn {
	_ = "STUB: not implemented"
	return *new(resolver.ClientConn)
}
