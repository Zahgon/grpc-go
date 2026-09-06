package grpc_testing

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	WorkerService_RunServer_FullMethodName  = "/grpc.testing.WorkerService/RunServer"
	WorkerService_RunClient_FullMethodName  = "/grpc.testing.WorkerService/RunClient"
	WorkerService_CoreCount_FullMethodName  = "/grpc.testing.WorkerService/CoreCount"
	WorkerService_QuitWorker_FullMethodName = "/grpc.testing.WorkerService/QuitWorker"
)

type WorkerServiceClient interface {
	RunServer(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ServerArgs, ServerStatus], error)

	RunClient(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ClientArgs, ClientStatus], error)

	CoreCount(ctx context.Context, in *CoreRequest, opts ...grpc.CallOption) (*CoreResponse, error)

	QuitWorker(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	_ = "STUB: not implemented"
	return *new(WorkerServiceClient)
}

func (c *workerServiceClient) RunServer(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ServerArgs, ServerStatus], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type WorkerService_RunServerClient = grpc.BidiStreamingClient[ServerArgs, ServerStatus]

func (c *workerServiceClient) RunClient(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ClientArgs, ClientStatus], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type WorkerService_RunClientClient = grpc.BidiStreamingClient[ClientArgs, ClientStatus]

func (c *workerServiceClient) CoreCount(ctx context.Context, in *CoreRequest, opts ...grpc.CallOption) (*CoreResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *workerServiceClient) QuitWorker(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type WorkerServiceServer interface {
	RunServer(grpc.BidiStreamingServer[ServerArgs, ServerStatus]) error

	RunClient(grpc.BidiStreamingServer[ClientArgs, ClientStatus]) error

	CoreCount(context.Context, *CoreRequest) (*CoreResponse, error)

	QuitWorker(context.Context, *Void) (*Void, error)
	mustEmbedUnimplementedWorkerServiceServer()
}

type UnimplementedWorkerServiceServer struct{}

func (UnimplementedWorkerServiceServer) RunServer(grpc.BidiStreamingServer[ServerArgs, ServerStatus]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedWorkerServiceServer) RunClient(grpc.BidiStreamingServer[ClientArgs, ClientStatus]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedWorkerServiceServer) CoreCount(context.Context, *CoreRequest) (*CoreResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedWorkerServiceServer) QuitWorker(context.Context, *Void) (*Void, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedWorkerServiceServer) mustEmbedUnimplementedWorkerServiceServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedWorkerServiceServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _WorkerService_RunServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type WorkerService_RunServerServer = grpc.BidiStreamingServer[ServerArgs, ServerStatus]

func _WorkerService_RunClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type WorkerService_RunClientServer = grpc.BidiStreamingServer[ClientArgs, ClientStatus]

func _WorkerService_CoreCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _WorkerService_QuitWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CoreCount",
			Handler:    _WorkerService_CoreCount_Handler,
		},
		{
			MethodName: "QuitWorker",
			Handler:    _WorkerService_QuitWorker_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunServer",
			Handler:       _WorkerService_RunServer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RunClient",
			Handler:       _WorkerService_RunClient_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/testing/worker_service.proto",
}
