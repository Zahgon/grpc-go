package endpointsharding

import (
	rand "math/rand/v2"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/resolver"
)

var randIntN = rand.IntN

type ChildState struct {
	Endpoint resolver.Endpoint
	State    balancer.State
	ExitIdle func()
}

type Options struct {
	DisableAutoReconnect bool
}

type ChildBuilderFunc func(cc balancer.ClientConn, opts balancer.BuildOptions) balancer.Balancer

func NewBalancer(cc balancer.ClientConn, opts balancer.BuildOptions, childBuilder ChildBuilderFunc, esOpts Options) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

type endpointSharding struct {
	cc           balancer.ClientConn
	bOpts        balancer.BuildOptions
	esOpts       Options
	childBuilder ChildBuilderFunc

	mu                  sync.Mutex
	endpoints           *resolver.EndpointMap[*endpointState]
	inhibitChildUpdates bool
}

func rotateEndpoints(es []resolver.Endpoint) []resolver.Endpoint {
	_ = "STUB: not implemented"
	return nil
}

func (es *endpointSharding) UpdateClientConnState(state balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *endpointSharding) ResolverError(err error) { _ = "STUB: not implemented"; return }

func (es *endpointSharding) UpdateSubConnState(balancer.SubConn, balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (es *endpointSharding) Close() { _ = "STUB: not implemented"; return }

func (es *endpointSharding) ExitIdle() { _ = "STUB: not implemented"; return }

func (es *endpointSharding) inhibitUpdatesFromChildren() { _ = "STUB: not implemented"; return }

func (es *endpointSharding) allowUpdatesFromChildren() { _ = "STUB: not implemented"; return }

func (es *endpointSharding) updateStateLocked() { _ = "STUB: not implemented"; return }

type pickerWithChildStates struct {
	pickers     []balancer.Picker
	childStates []ChildState
	next        uint32
}

func (p *pickerWithChildStates) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

func ChildStatesFromPicker(picker balancer.Picker) []ChildState {
	_ = "STUB: not implemented"
	return nil
}

type endpointState struct {
	balancer.ClientConn

	parent               *endpointSharding
	endpoint             resolver.Endpoint
	state                balancer.State
	disableAutoReconnect bool

	childMu sync.Mutex
	childLB balancer.Balancer
	closed  bool
}

func (es *endpointState) UpdateState(state balancer.State) { _ = "STUB: not implemented"; return }

func (es *endpointState) updateClientConnState(state balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *endpointState) resolverError(err error) { _ = "STUB: not implemented"; return }

func (es *endpointState) close() { _ = "STUB: not implemented"; return }

func (es *endpointState) exitIdle() { _ = "STUB: not implemented"; return }
