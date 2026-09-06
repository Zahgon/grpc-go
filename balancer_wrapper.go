package grpc

import (
	"context"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/experimental/stats"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/internal/balancer/gracefulswitch"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/resolver"
)

var (
	noOpRegisterHealthListenerFn = func(_ context.Context, listener func(balancer.SubConnState)) func() {
		listener(balancer.SubConnState{ConnectivityState: connectivity.Ready})
		return func() {}
	}
)

type ccBalancerWrapper struct {
	internal.EnforceClientConnEmbedding

	cc               *ClientConn
	opts             balancer.BuildOptions
	serializer       *grpcsync.CallbackSerializer
	serializerCancel context.CancelFunc

	curBalancerName string
	balancer        *gracefulswitch.Balancer

	mu     sync.Mutex
	closed bool
}

func newCCBalancerWrapper(cc *ClientConn) *ccBalancerWrapper { _ = "STUB: not implemented"; return nil }

func (ccb *ccBalancerWrapper) MetricsRecorder() stats.MetricsRecorder {
	_ = "STUB: not implemented"
	return *new(stats.MetricsRecorder)
}

func (ccb *ccBalancerWrapper) updateClientConnState(ccs *balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (ccb *ccBalancerWrapper) resolverError(err error) { _ = "STUB: not implemented"; return }

func (ccb *ccBalancerWrapper) close() { _ = "STUB: not implemented"; return }

func (ccb *ccBalancerWrapper) exitIdle() { _ = "STUB: not implemented"; return }

func (ccb *ccBalancerWrapper) NewSubConn(addrs []resolver.Address, opts balancer.NewSubConnOptions) (balancer.SubConn, error) {
	_ = "STUB: not implemented"
	return *new(balancer.SubConn), nil
}

func (ccb *ccBalancerWrapper) RemoveSubConn(balancer.SubConn) { _ = "STUB: not implemented"; return }

func (ccb *ccBalancerWrapper) UpdateAddresses(sc balancer.SubConn, addrs []resolver.Address) {
	_ = "STUB: not implemented"
	return
}

func (ccb *ccBalancerWrapper) UpdateState(s balancer.State) { _ = "STUB: not implemented"; return }

func (ccb *ccBalancerWrapper) ResolveNow(o resolver.ResolveNowOptions) {
	_ = "STUB: not implemented"
	return
}

func (ccb *ccBalancerWrapper) Target() string { _ = "STUB: not implemented"; return "" }

type acBalancerWrapper struct {
	internal.EnforceSubConnEmbedding
	ac            *addrConn
	ccb           *ccBalancerWrapper
	stateListener func(balancer.SubConnState)

	producersMu sync.Mutex
	producers   map[balancer.ProducerBuilder]*refCountedProducer

	healthMu sync.Mutex

	healthData *healthData
}

type healthData struct {
	connectivityState connectivity.State

	closeHealthProducer func()
}

func newHealthData(s connectivity.State) *healthData { _ = "STUB: not implemented"; return nil }

func (acbw *acBalancerWrapper) updateState(s connectivity.State, err error) {
	_ = "STUB: not implemented"
	return
}

func (acbw *acBalancerWrapper) String() string { _ = "STUB: not implemented"; return "" }

func (acbw *acBalancerWrapper) UpdateAddresses(addrs []resolver.Address) {
	_ = "STUB: not implemented"
	return
}

func (acbw *acBalancerWrapper) Connect() { _ = "STUB: not implemented"; return }

func (acbw *acBalancerWrapper) Shutdown() { _ = "STUB: not implemented"; return }

func (acbw *acBalancerWrapper) NewStream(ctx context.Context, desc *StreamDesc, method string, opts ...CallOption) (ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

func (acbw *acBalancerWrapper) Invoke(ctx context.Context, method string, args any, reply any, opts ...CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

type refCountedProducer struct {
	producer balancer.Producer
	refs     int
	close    func()
}

func (acbw *acBalancerWrapper) GetOrBuildProducer(pb balancer.ProducerBuilder) (balancer.Producer, func()) {
	_ = "STUB: not implemented"
	return *new(balancer.Producer), nil
}

func (acbw *acBalancerWrapper) closeProducers() { _ = "STUB: not implemented"; return }

type healthProducerRegisterFn = func(context.Context, balancer.SubConn, string, func(balancer.SubConnState)) func()

func (acbw *acBalancerWrapper) healthListenerRegFn() func(context.Context, func(balancer.SubConnState)) func() {
	_ = "STUB: not implemented"
	return nil
}

func (acbw *acBalancerWrapper) RegisterHealthListener(listener func(balancer.SubConnState)) {
	_ = "STUB: not implemented"
	return
}
