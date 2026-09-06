package echo

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	Echo_UnaryEcho_FullMethodName                  = "/grpc.examples.echo.Echo/UnaryEcho"
	Echo_ServerStreamingEcho_FullMethodName        = "/grpc.examples.echo.Echo/ServerStreamingEcho"
	Echo_ClientStreamingEcho_FullMethodName        = "/grpc.examples.echo.Echo/ClientStreamingEcho"
	Echo_BidirectionalStreamingEcho_FullMethodName = "/grpc.examples.echo.Echo/BidirectionalStreamingEcho"
)

type EchoClient interface {
	UnaryEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)

	ServerStreamingEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[EchoResponse], error)

	ClientStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[EchoRequest, EchoResponse], error)

	BidirectionalStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[EchoRequest, EchoResponse], error)
}

type echoClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	_ = "STUB: not implemented"
	return *new(EchoClient)
}

func (c *echoClient) UnaryEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *echoClient) ServerStreamingEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[EchoResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Echo_ServerStreamingEchoClient = grpc.ServerStreamingClient[EchoResponse]

func (c *echoClient) ClientStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[EchoRequest, EchoResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Echo_ClientStreamingEchoClient = grpc.ClientStreamingClient[EchoRequest, EchoResponse]

func (c *echoClient) BidirectionalStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[EchoRequest, EchoResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Echo_BidirectionalStreamingEchoClient = grpc.BidiStreamingClient[EchoRequest, EchoResponse]

type EchoServer interface {
	UnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error)

	ServerStreamingEcho(*EchoRequest, grpc.ServerStreamingServer[EchoResponse]) error

	ClientStreamingEcho(grpc.ClientStreamingServer[EchoRequest, EchoResponse]) error

	BidirectionalStreamingEcho(grpc.BidiStreamingServer[EchoRequest, EchoResponse]) error
	mustEmbedUnimplementedEchoServer()
}

type UnimplementedEchoServer struct{}

func (UnimplementedEchoServer) UnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedEchoServer) ServerStreamingEcho(*EchoRequest, grpc.ServerStreamingServer[EchoResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedEchoServer) ClientStreamingEcho(grpc.ClientStreamingServer[EchoRequest, EchoResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedEchoServer) BidirectionalStreamingEcho(grpc.BidiStreamingServer[EchoRequest, EchoResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedEchoServer) mustEmbedUnimplementedEchoServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedEchoServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeEchoServer interface {
	mustEmbedUnimplementedEchoServer()
}

func RegisterEchoServer(s grpc.ServiceRegistrar, srv EchoServer) { _ = "STUB: not implemented"; return }

func _Echo_UnaryEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _Echo_ServerStreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type Echo_ServerStreamingEchoServer = grpc.ServerStreamingServer[EchoResponse]

func _Echo_ClientStreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type Echo_ClientStreamingEchoServer = grpc.ClientStreamingServer[EchoRequest, EchoResponse]

func _Echo_BidirectionalStreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type Echo_BidirectionalStreamingEchoServer = grpc.BidiStreamingServer[EchoRequest, EchoResponse]

var Echo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.examples.echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryEcho",
			Handler:    _Echo_UnaryEcho_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamingEcho",
			Handler:       _Echo_ServerStreamingEcho_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamingEcho",
			Handler:       _Echo_ClientStreamingEcho_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamingEcho",
			Handler:       _Echo_BidirectionalStreamingEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "examples/features/proto/echo/echo.proto",
}
