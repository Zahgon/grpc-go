package grpc

import (
	"context"
	"sync"

	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

type ccResolverWrapper struct {
	cc                  *ClientConn
	ignoreServiceConfig bool
	serializer          *grpcsync.CallbackSerializer
	serializerCancel    context.CancelFunc

	resolver resolver.Resolver

	mu       sync.Mutex
	curState resolver.State
	closed   bool
}

func newCCResolverWrapper(cc *ClientConn) *ccResolverWrapper { _ = "STUB: not implemented"; return nil }

func (ccr *ccResolverWrapper) start() error { _ = "STUB: not implemented"; return nil }

func (ccr *ccResolverWrapper) resolveNow(o resolver.ResolveNowOptions) {
	_ = "STUB: not implemented"
	return
}

func (ccr *ccResolverWrapper) close() { _ = "STUB: not implemented"; return }

func (ccr *ccResolverWrapper) UpdateState(s resolver.State) error {
	_ = "STUB: not implemented"
	return nil
}

func (ccr *ccResolverWrapper) ReportError(err error) { _ = "STUB: not implemented"; return }

func (ccr *ccResolverWrapper) NewAddress(addrs []resolver.Address) {
	_ = "STUB: not implemented"
	return
}

func (ccr *ccResolverWrapper) ParseServiceConfig(scJSON string) *serviceconfig.ParseResult {
	_ = "STUB: not implemented"
	return nil
}

func (ccr *ccResolverWrapper) addChannelzTraceEvent(s resolver.State) {
	_ = "STUB: not implemented"
	return
}

func addressesToEndpoints(addrs []resolver.Address) []resolver.Endpoint {
	_ = "STUB: not implemented"
	return nil
}
