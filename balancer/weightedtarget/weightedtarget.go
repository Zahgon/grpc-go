package weightedtarget

import (
	"encoding/json"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/weightedtarget/weightedaggregator"
	"google.golang.org/grpc/internal/balancergroup"
	"google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/internal/wrr"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

const Name = "weighted_target_experimental"

var NewRandomWRR = wrr.NewRandom

func init() {
	balancer.Register(bb{})
}

type bb struct{}

func (bb) Build(cc balancer.ClientConn, bOpts balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

func (bb) Name() string { _ = "STUB: not implemented"; return "" }

func (bb) ParseConfig(c json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	_ = "STUB: not implemented"
	return *new(serviceconfig.LoadBalancingConfig), nil
}

type weightedTargetBalancer struct {
	logger *grpclog.PrefixLogger

	bg              *balancergroup.BalancerGroup
	stateAggregator *weightedaggregator.Aggregator

	targets map[string]Target
}

type localityKeyType string

const localityKey = localityKeyType("locality")

func LocalityFromResolverState(state resolver.State) string { _ = "STUB: not implemented"; return "" }

func (b *weightedTargetBalancer) UpdateClientConnState(s balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *weightedTargetBalancer) ResolverError(err error) { _ = "STUB: not implemented"; return }

func (b *weightedTargetBalancer) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (b *weightedTargetBalancer) Close() { _ = "STUB: not implemented"; return }

func (b *weightedTargetBalancer) ExitIdle() { _ = "STUB: not implemented"; return }
