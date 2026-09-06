package grpclb

import (
	"encoding/json"

	"google.golang.org/grpc/balancer/pickfirst"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/serviceconfig"
)

const (
	roundRobinName = roundrobin.Name
	pickFirstName  = pickfirst.Name
)

type grpclbServiceConfig struct {
	serviceconfig.LoadBalancingConfig
	ChildPolicy *[]map[string]json.RawMessage
	ServiceName string
}

func (b *lbBuilder) ParseConfig(lbConfig json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	_ = "STUB: not implemented"
	return *new(serviceconfig.LoadBalancingConfig), nil
}

func childIsPickFirst(sc *grpclbServiceConfig) bool { _ = "STUB: not implemented"; return false }
