package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	Profiling_Enable_FullMethodName         = "/grpc.go.profiling.v1alpha.Profiling/Enable"
	Profiling_GetStreamStats_FullMethodName = "/grpc.go.profiling.v1alpha.Profiling/GetStreamStats"
)

type ProfilingClient interface {
	Enable(ctx context.Context, in *EnableRequest, opts ...grpc.CallOption) (*EnableResponse, error)

	GetStreamStats(ctx context.Context, in *GetStreamStatsRequest, opts ...grpc.CallOption) (*GetStreamStatsResponse, error)
}

type profilingClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilingClient(cc grpc.ClientConnInterface) ProfilingClient {
	_ = "STUB: not implemented"
	return *new(ProfilingClient)
}

func (c *profilingClient) Enable(ctx context.Context, in *EnableRequest, opts ...grpc.CallOption) (*EnableResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *profilingClient) GetStreamStats(ctx context.Context, in *GetStreamStatsRequest, opts ...grpc.CallOption) (*GetStreamStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ProfilingServer interface {
	Enable(context.Context, *EnableRequest) (*EnableResponse, error)

	GetStreamStats(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error)
}

type UnimplementedProfilingServer struct{}

func (UnimplementedProfilingServer) Enable(context.Context, *EnableRequest) (*EnableResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedProfilingServer) GetStreamStats(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedProfilingServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeProfilingServer interface {
	mustEmbedUnimplementedProfilingServer()
}

func RegisterProfilingServer(s grpc.ServiceRegistrar, srv ProfilingServer) {
	_ = "STUB: not implemented"
	return
}

func _Profiling_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _Profiling_GetStreamStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var Profiling_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.go.profiling.v1alpha.Profiling",
	HandlerType: (*ProfilingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enable",
			Handler:    _Profiling_Enable_Handler,
		},
		{
			MethodName: "GetStreamStats",
			Handler:    _Profiling_GetStreamStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profiling/proto/service.proto",
}
