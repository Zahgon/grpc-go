package grpc_testing

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	BenchmarkService_UnaryCall_FullMethodName           = "/grpc.testing.BenchmarkService/UnaryCall"
	BenchmarkService_StreamingCall_FullMethodName       = "/grpc.testing.BenchmarkService/StreamingCall"
	BenchmarkService_StreamingFromClient_FullMethodName = "/grpc.testing.BenchmarkService/StreamingFromClient"
	BenchmarkService_StreamingFromServer_FullMethodName = "/grpc.testing.BenchmarkService/StreamingFromServer"
	BenchmarkService_StreamingBothWays_FullMethodName   = "/grpc.testing.BenchmarkService/StreamingBothWays"
)

type BenchmarkServiceClient interface {
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)

	StreamingCall(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SimpleRequest, SimpleResponse], error)

	StreamingFromClient(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SimpleRequest, SimpleResponse], error)

	StreamingFromServer(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SimpleResponse], error)

	StreamingBothWays(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SimpleRequest, SimpleResponse], error)
}

type benchmarkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBenchmarkServiceClient(cc grpc.ClientConnInterface) BenchmarkServiceClient {
	_ = "STUB: not implemented"
	return *new(BenchmarkServiceClient)
}

func (c *benchmarkServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *benchmarkServiceClient) StreamingCall(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SimpleRequest, SimpleResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type BenchmarkService_StreamingCallClient = grpc.BidiStreamingClient[SimpleRequest, SimpleResponse]

func (c *benchmarkServiceClient) StreamingFromClient(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SimpleRequest, SimpleResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type BenchmarkService_StreamingFromClientClient = grpc.ClientStreamingClient[SimpleRequest, SimpleResponse]

func (c *benchmarkServiceClient) StreamingFromServer(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SimpleResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type BenchmarkService_StreamingFromServerClient = grpc.ServerStreamingClient[SimpleResponse]

func (c *benchmarkServiceClient) StreamingBothWays(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SimpleRequest, SimpleResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type BenchmarkService_StreamingBothWaysClient = grpc.BidiStreamingClient[SimpleRequest, SimpleResponse]

type BenchmarkServiceServer interface {
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)

	StreamingCall(grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]) error

	StreamingFromClient(grpc.ClientStreamingServer[SimpleRequest, SimpleResponse]) error

	StreamingFromServer(*SimpleRequest, grpc.ServerStreamingServer[SimpleResponse]) error

	StreamingBothWays(grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]) error
	mustEmbedUnimplementedBenchmarkServiceServer()
}

type UnimplementedBenchmarkServiceServer struct{}

func (UnimplementedBenchmarkServiceServer) UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedBenchmarkServiceServer) StreamingCall(grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedBenchmarkServiceServer) StreamingFromClient(grpc.ClientStreamingServer[SimpleRequest, SimpleResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedBenchmarkServiceServer) StreamingFromServer(*SimpleRequest, grpc.ServerStreamingServer[SimpleResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedBenchmarkServiceServer) StreamingBothWays(grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedBenchmarkServiceServer) mustEmbedUnimplementedBenchmarkServiceServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedBenchmarkServiceServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeBenchmarkServiceServer interface {
	mustEmbedUnimplementedBenchmarkServiceServer()
}

func RegisterBenchmarkServiceServer(s grpc.ServiceRegistrar, srv BenchmarkServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _BenchmarkService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _BenchmarkService_StreamingCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type BenchmarkService_StreamingCallServer = grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]

func _BenchmarkService_StreamingFromClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type BenchmarkService_StreamingFromClientServer = grpc.ClientStreamingServer[SimpleRequest, SimpleResponse]

func _BenchmarkService_StreamingFromServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type BenchmarkService_StreamingFromServerServer = grpc.ServerStreamingServer[SimpleResponse]

func _BenchmarkService_StreamingBothWays_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type BenchmarkService_StreamingBothWaysServer = grpc.BidiStreamingServer[SimpleRequest, SimpleResponse]

var BenchmarkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.BenchmarkService",
	HandlerType: (*BenchmarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _BenchmarkService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingCall",
			Handler:       _BenchmarkService_StreamingCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamingFromClient",
			Handler:       _BenchmarkService_StreamingFromClient_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamingFromServer",
			Handler:       _BenchmarkService_StreamingFromServer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamingBothWays",
			Handler:       _BenchmarkService_StreamingBothWays_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/testing/benchmark_service.proto",
}
