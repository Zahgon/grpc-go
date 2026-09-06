package helloworld

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	Greeter_SayHello_FullMethodName = "/helloworld.Greeter/SayHello"
)

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	_ = "STUB: not implemented"
	return *new(GreeterClient)
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}

type UnimplementedGreeterServer struct{}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedGreeterServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	_ = "STUB: not implemented"
	return
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/helloworld/helloworld/helloworld.proto",
}
