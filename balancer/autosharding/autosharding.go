package autosharding

import (
	"encoding/json"

	"google.golang.org/grpc/balancer"
	iserviceconfig "google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/serviceconfig"
)

const Name = "autosharding_experimental"

func init() {
	balancer.Register(bb{})
}

type lbConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`

	ChannelFactoryKey        string                  `json:"channelFactoryKey,omitempty"`
	AutoShardingTarget       string                  `json:"autoshardingTarget,omitempty"`
	KeyHeaderName            string                  `json:"keyHeaderName,omitempty"`
	EnableFallback           bool                    `json:"enableFallback,omitempty"`
	InitialAssignmentTimeout iserviceconfig.Duration `json:"initialAssignmentTimeout,omitempty"`
}

type bb struct{}

func (bb) Name() string { _ = "STUB: not implemented"; return "" }

func (bb) ParseConfig(s json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	_ = "STUB: not implemented"
	return *new(serviceconfig.LoadBalancingConfig), nil
}

func (bb) Build(balancer.ClientConn, balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

type autoshardingBalancer struct {
	balancer.Balancer
}
