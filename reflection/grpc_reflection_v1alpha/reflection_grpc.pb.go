package grpc_reflection_v1alpha

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	ServerReflection_ServerReflectionInfo_FullMethodName = "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
)

type ServerReflectionClient interface {
	ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ServerReflectionRequest, ServerReflectionResponse], error)
}

type serverReflectionClient struct {
	cc grpc.ClientConnInterface
}

func NewServerReflectionClient(cc grpc.ClientConnInterface) ServerReflectionClient {
	_ = "STUB: not implemented"
	return *new(ServerReflectionClient)
}

func (c *serverReflectionClient) ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ServerReflectionRequest, ServerReflectionResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ServerReflection_ServerReflectionInfoClient = grpc.BidiStreamingClient[ServerReflectionRequest, ServerReflectionResponse]

type ServerReflectionServer interface {
	ServerReflectionInfo(grpc.BidiStreamingServer[ServerReflectionRequest, ServerReflectionResponse]) error
}

type UnimplementedServerReflectionServer struct{}

func (UnimplementedServerReflectionServer) ServerReflectionInfo(grpc.BidiStreamingServer[ServerReflectionRequest, ServerReflectionResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedServerReflectionServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeServerReflectionServer interface {
	mustEmbedUnimplementedServerReflectionServer()
}

func RegisterServerReflectionServer(s grpc.ServiceRegistrar, srv ServerReflectionServer) {
	_ = "STUB: not implemented"
	return
}

func _ServerReflection_ServerReflectionInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type ServerReflection_ServerReflectionInfoServer = grpc.BidiStreamingServer[ServerReflectionRequest, ServerReflectionResponse]

var ServerReflection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.reflection.v1alpha.ServerReflection",
	HandlerType: (*ServerReflectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerReflectionInfo",
			Handler:       _ServerReflection_ServerReflectionInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/reflection/v1alpha/reflection.proto",
}
