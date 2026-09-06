package health

import (
	"context"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/internal"
)

func init() {
	producerBuilderSingleton = &producerBuilder{}
	internal.RegisterClientHealthCheckListener = registerClientSideHealthCheckListener
}

type producerBuilder struct{}

var producerBuilderSingleton *producerBuilder

func (*producerBuilder) Build(cci any) (balancer.Producer, func()) {
	_ = "STUB: not implemented"
	return *new(balancer.Producer), nil
}

type healthServiceProducer struct {
	cc grpc.ClientConnInterface

	mu     sync.Mutex
	cancel func()
}

func registerClientSideHealthCheckListener(ctx context.Context, sc balancer.SubConn, serviceName string, listener func(balancer.SubConnState)) func() {
	_ = "STUB: not implemented"
	return nil
}

func (p *healthServiceProducer) startHealthCheck(ctx context.Context, sc balancer.SubConn, serviceName string, listener func(balancer.SubConnState)) {
	_ = "STUB: not implemented"
	return
}
