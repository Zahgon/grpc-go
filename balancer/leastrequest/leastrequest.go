package leastrequest

import (
	"encoding/json"
	rand "math/rand/v2"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

const Name = "least_request_experimental"

var (
	randuint32 = rand.Uint32
	logger     = grpclog.Component("least-request")
)

func init() {
	balancer.Register(bb{})
}

type LBConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`

	ChoiceCount uint32 `json:"choiceCount,omitempty"`
}

type bb struct{}

func (bb) ParseConfig(s json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	_ = "STUB: not implemented"
	return *new(serviceconfig.LoadBalancingConfig), nil
}

func (bb) Name() string { _ = "STUB: not implemented"; return "" }

func (bb) Build(cc balancer.ClientConn, bOpts balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

type leastRequestBalancer struct {
	balancer.ClientConn
	child  balancer.Balancer
	logger *internalgrpclog.PrefixLogger

	mu          sync.Mutex
	choiceCount uint32

	endpointRPCCounts *resolver.EndpointMap[*atomic.Int32]
}

func (lrb *leastRequestBalancer) Close() { _ = "STUB: not implemented"; return }

func (lrb *leastRequestBalancer) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (lrb *leastRequestBalancer) ResolverError(err error) { _ = "STUB: not implemented"; return }

func (lrb *leastRequestBalancer) ExitIdle() { _ = "STUB: not implemented"; return }

func (lrb *leastRequestBalancer) UpdateClientConnState(ccs balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

type endpointState struct {
	picker  balancer.Picker
	numRPCs *atomic.Int32
}

func (lrb *leastRequestBalancer) UpdateState(state balancer.State) {
	_ = "STUB: not implemented"
	return
}

type picker struct {
	choiceCount    uint32
	endpointStates []endpointState
}

func (p *picker) Pick(pInfo balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}
