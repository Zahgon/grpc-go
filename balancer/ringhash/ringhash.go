package ringhash

import (
	"encoding/json"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/internal/grpclog"
	iringhash "google.golang.org/grpc/internal/ringhash"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

const Name = "ring_hash_experimental"

func lazyPickFirstBuilder(cc balancer.ClientConn, opts balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

func init() {
	balancer.Register(bb{})
}

type bb struct{}

func (bb) Build(cc balancer.ClientConn, opts balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

func (bb) Name() string { _ = "STUB: not implemented"; return "" }

func (bb) ParseConfig(c json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	_ = "STUB: not implemented"
	return *new(serviceconfig.LoadBalancingConfig), nil
}

type ringhashBalancer struct {
	balancer.ClientConn
	logger *grpclog.PrefixLogger
	child  balancer.Balancer

	mu                   sync.Mutex
	config               *iringhash.LBConfig
	inhibitChildUpdates  bool
	shouldRegenerateRing bool
	endpointStates       *resolver.EndpointMap[*endpointState]

	ring *ring
}

func hashKey(endpoint resolver.Endpoint) string { _ = "STUB: not implemented"; return "" }

func (b *ringhashBalancer) UpdateState(state balancer.State) { _ = "STUB: not implemented"; return }

func (b *ringhashBalancer) UpdateClientConnState(ccs balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *ringhashBalancer) ResolverError(err error) { _ = "STUB: not implemented"; return }

func (b *ringhashBalancer) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (b *ringhashBalancer) updatePickerLocked() { _ = "STUB: not implemented"; return }

func (b *ringhashBalancer) Close() { _ = "STUB: not implemented"; return }

func (b *ringhashBalancer) ExitIdle() { _ = "STUB: not implemented"; return }

func (b *ringhashBalancer) newPickerLocked() *picker { _ = "STUB: not implemented"; return nil }

func (b *ringhashBalancer) aggregatedStateLocked() connectivity.State {
	_ = "STUB: not implemented"
	return *new(connectivity.State)
}

func getWeightAttribute(e resolver.Endpoint) uint32 { _ = "STUB: not implemented"; return 0 }

type endpointState struct {
	hashKey  string
	weight   uint32
	exitIdle func()

	state balancer.State
}
