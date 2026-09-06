package base

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
)

var logger = grpclog.Component("balancer")

type baseBuilder struct {
	name          string
	pickerBuilder PickerBuilder
	config        Config
}

func (bb *baseBuilder) Build(cc balancer.ClientConn, _ balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

func (bb *baseBuilder) Name() string { _ = "STUB: not implemented"; return "" }

type baseBalancer struct {
	cc            balancer.ClientConn
	pickerBuilder PickerBuilder

	csEvltr *balancer.ConnectivityStateEvaluator
	state   connectivity.State

	subConns *resolver.AddressMapV2[balancer.SubConn]
	scStates map[balancer.SubConn]connectivity.State
	picker   balancer.Picker
	config   Config

	resolverErr error
	connErr     error
}

func (b *baseBalancer) ResolverError(err error) { _ = "STUB: not implemented"; return }

func (b *baseBalancer) UpdateClientConnState(s balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *baseBalancer) mergeErrors() error { _ = "STUB: not implemented"; return nil }

func (b *baseBalancer) regeneratePicker() { _ = "STUB: not implemented"; return }

func (b *baseBalancer) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (b *baseBalancer) updateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (b *baseBalancer) Close() { _ = "STUB: not implemented"; return }

func (b *baseBalancer) ExitIdle() { _ = "STUB: not implemented"; return }

func NewErrPicker(err error) balancer.Picker {
	_ = "STUB: not implemented"
	return *new(balancer.Picker)
}

var NewErrPickerV2 = NewErrPicker

type errPicker struct {
	err error
}

func (p *errPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}
