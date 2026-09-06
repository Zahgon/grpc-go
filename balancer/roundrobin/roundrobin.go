package roundrobin

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const Name = "round_robin"

var logger = grpclog.Component("roundrobin")

func init() {
	balancer.Register(builder{})
}

type builder struct{}

func (bb builder) Name() string { _ = "STUB: not implemented"; return "" }

func (bb builder) Build(cc balancer.ClientConn, opts balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

type rrBalancer struct {
	balancer.Balancer
	cc     balancer.ClientConn
	logger *internalgrpclog.PrefixLogger
}

func (b *rrBalancer) UpdateClientConnState(ccs balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}
