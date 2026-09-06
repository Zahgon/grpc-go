package grpc_testing

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	MetricsService_GetAllGauges_FullMethodName = "/grpc.testing.MetricsService/GetAllGauges"
	MetricsService_GetGauge_FullMethodName     = "/grpc.testing.MetricsService/GetGauge"
)

type MetricsServiceClient interface {
	GetAllGauges(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GaugeResponse], error)

	GetGauge(ctx context.Context, in *GaugeRequest, opts ...grpc.CallOption) (*GaugeResponse, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	_ = "STUB: not implemented"
	return *new(MetricsServiceClient)
}

func (c *metricsServiceClient) GetAllGauges(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GaugeResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MetricsService_GetAllGaugesClient = grpc.ServerStreamingClient[GaugeResponse]

func (c *metricsServiceClient) GetGauge(ctx context.Context, in *GaugeRequest, opts ...grpc.CallOption) (*GaugeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MetricsServiceServer interface {
	GetAllGauges(*EmptyMessage, grpc.ServerStreamingServer[GaugeResponse]) error

	GetGauge(context.Context, *GaugeRequest) (*GaugeResponse, error)
	mustEmbedUnimplementedMetricsServiceServer()
}

type UnimplementedMetricsServiceServer struct{}

func (UnimplementedMetricsServiceServer) GetAllGauges(*EmptyMessage, grpc.ServerStreamingServer[GaugeResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedMetricsServiceServer) GetGauge(context.Context, *GaugeRequest) (*GaugeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedMetricsServiceServer) mustEmbedUnimplementedMetricsServiceServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedMetricsServiceServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeMetricsServiceServer interface {
	mustEmbedUnimplementedMetricsServiceServer()
}

func RegisterMetricsServiceServer(s grpc.ServiceRegistrar, srv MetricsServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _MetricsService_GetAllGauges_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type MetricsService_GetAllGaugesServer = grpc.ServerStreamingServer[GaugeResponse]

func _MetricsService_GetGauge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var MetricsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGauge",
			Handler:    _MetricsService_GetGauge_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllGauges",
			Handler:       _MetricsService_GetAllGauges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "interop/stress/grpc_testing/metrics.proto",
}
