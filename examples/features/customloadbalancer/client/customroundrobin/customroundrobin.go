package customroundrobin

import (
	"encoding/json"
	"sync/atomic"

	_ "google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/serviceconfig"
)

func init() {
	balancer.Register(customRoundRobinBuilder{})
}

const customRRName = "custom_round_robin"

type customRRConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`

	ChooseSecond uint32 `json:"chooseSecond,omitempty"`
}

type customRoundRobinBuilder struct{}

func (customRoundRobinBuilder) ParseConfig(s json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	_ = "STUB: not implemented"
	return *new(serviceconfig.LoadBalancingConfig), nil
}

func (customRoundRobinBuilder) Name() string { _ = "STUB: not implemented"; return "" }

func (customRoundRobinBuilder) Build(cc balancer.ClientConn, bOpts balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

type customRoundRobin struct {
	balancer.Balancer
	balancer.ClientConn
	bOpts balancer.BuildOptions

	cfg atomic.Pointer[customRRConfig]
}

func (crr *customRoundRobin) UpdateClientConnState(state balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (crr *customRoundRobin) UpdateState(state balancer.State) { _ = "STUB: not implemented"; return }

type customRoundRobinPicker struct {
	pickers      []balancer.Picker
	chooseSecond uint32
	next         uint32
}

func (crrp *customRoundRobinPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}
