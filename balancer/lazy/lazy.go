package lazy

import (
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/grpclog"

	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

var (
	logger = grpclog.Component("lazy-lb")
)

const (
	logPrefix = "[lazy-lb %p] "
)

type ChildBuilderFunc func(cc balancer.ClientConn, opts balancer.BuildOptions) balancer.Balancer

func NewBalancer(cc balancer.ClientConn, bOpts balancer.BuildOptions, childBuilder ChildBuilderFunc) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

type lazyBalancer struct {
	cc           balancer.ClientConn
	buildOptions balancer.BuildOptions
	logger       *internalgrpclog.PrefixLogger
	childBuilder ChildBuilderFunc

	mu                    sync.Mutex
	delegate              balancer.Balancer
	latestClientConnState *balancer.ClientConnState
	latestResolverError   error
}

func (lb *lazyBalancer) Close() { _ = "STUB: not implemented"; return }

func (lb *lazyBalancer) ResolverError(err error) { _ = "STUB: not implemented"; return }

func (lb *lazyBalancer) UpdateClientConnState(ccs balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (lb *lazyBalancer) UpdateSubConnState(balancer.SubConn, balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (lb *lazyBalancer) ExitIdle() { _ = "STUB: not implemented"; return }

type idlePicker struct {
	exitIdle func()
}

func (i *idlePicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}
