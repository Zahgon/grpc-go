package xds

import (
	"encoding/json"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/serviceconfig"
)

func init() {
	balancer.Register(rpcBehaviorBB{})
}

const name = "test.RpcBehaviorLoadBalancer"

type rpcBehaviorBB struct{}

func (rpcBehaviorBB) Build(cc balancer.ClientConn, bOpts balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

func (rpcBehaviorBB) ParseConfig(s json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	_ = "STUB: not implemented"
	return *new(serviceconfig.LoadBalancingConfig), nil
}

func (rpcBehaviorBB) Name() string { _ = "STUB: not implemented"; return "" }

type lbConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`
	RPCBehavior                       string `json:"rpcBehavior,omitempty"`
}

type rpcBehaviorLB struct {
	balancer.ClientConn

	balancer.Balancer

	mu  sync.Mutex
	cfg *lbConfig
}

func (b *rpcBehaviorLB) UpdateClientConnState(s balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *rpcBehaviorLB) UpdateState(state balancer.State) { _ = "STUB: not implemented"; return }

type rpcBehaviorPicker struct {
	childPicker balancer.Picker
	rpcBehavior string
}

func (p *rpcBehaviorPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

func newRPCBehaviorPicker(childPicker balancer.Picker, rpcBehavior string) *rpcBehaviorPicker {
	_ = "STUB: not implemented"
	return nil
}
