package grpc_lb_v1

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	LoadBalancer_BalanceLoad_FullMethodName = "/grpc.lb.v1.LoadBalancer/BalanceLoad"
)

type LoadBalancerClient interface {
	BalanceLoad(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[LoadBalanceRequest, LoadBalanceResponse], error)
}

type loadBalancerClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadBalancerClient(cc grpc.ClientConnInterface) LoadBalancerClient {
	_ = "STUB: not implemented"
	return *new(LoadBalancerClient)
}

func (c *loadBalancerClient) BalanceLoad(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[LoadBalanceRequest, LoadBalanceResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type LoadBalancer_BalanceLoadClient = grpc.BidiStreamingClient[LoadBalanceRequest, LoadBalanceResponse]

type LoadBalancerServer interface {
	BalanceLoad(grpc.BidiStreamingServer[LoadBalanceRequest, LoadBalanceResponse]) error
}

type UnimplementedLoadBalancerServer struct{}

func (UnimplementedLoadBalancerServer) BalanceLoad(grpc.BidiStreamingServer[LoadBalanceRequest, LoadBalanceResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedLoadBalancerServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeLoadBalancerServer interface {
	mustEmbedUnimplementedLoadBalancerServer()
}

func RegisterLoadBalancerServer(s grpc.ServiceRegistrar, srv LoadBalancerServer) {
	_ = "STUB: not implemented"
	return
}

func _LoadBalancer_BalanceLoad_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type LoadBalancer_BalanceLoadServer = grpc.BidiStreamingServer[LoadBalanceRequest, LoadBalanceResponse]

var LoadBalancer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.lb.v1.LoadBalancer",
	HandlerType: (*LoadBalancerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BalanceLoad",
			Handler:       _LoadBalancer_BalanceLoad_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/lb/v1/load_balancer.proto",
}
