package grpc_health_v1

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	Health_Check_FullMethodName = "/grpc.health.v1.Health/Check"
	Health_List_FullMethodName  = "/grpc.health.v1.Health/List"
	Health_Watch_FullMethodName = "/grpc.health.v1.Health/Watch"
)

type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)

	List(ctx context.Context, in *HealthListRequest, opts ...grpc.CallOption) (*HealthListResponse, error)

	Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HealthCheckResponse], error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	_ = "STUB: not implemented"
	return *new(HealthClient)
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *healthClient) List(ctx context.Context, in *HealthListRequest, opts ...grpc.CallOption) (*HealthListResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *healthClient) Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HealthCheckResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Health_WatchClient = grpc.ServerStreamingClient[HealthCheckResponse]

type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)

	List(context.Context, *HealthListRequest) (*HealthListResponse, error)

	Watch(*HealthCheckRequest, grpc.ServerStreamingServer[HealthCheckResponse]) error
}

type UnimplementedHealthServer struct{}

func (UnimplementedHealthServer) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedHealthServer) List(context.Context, *HealthListRequest) (*HealthListResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedHealthServer) Watch(*HealthCheckRequest, grpc.ServerStreamingServer[HealthCheckResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedHealthServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(s grpc.ServiceRegistrar, srv HealthServer) {
	_ = "STUB: not implemented"
	return
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _Health_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _Health_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type Health_WatchServer = grpc.ServerStreamingServer[HealthCheckResponse]

var Health_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.health.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Health_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Health_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/health/v1/health.proto",
}
