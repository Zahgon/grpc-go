package grpc_channelz_v1

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	Channelz_GetTopChannels_FullMethodName   = "/grpc.channelz.v1.Channelz/GetTopChannels"
	Channelz_GetServers_FullMethodName       = "/grpc.channelz.v1.Channelz/GetServers"
	Channelz_GetServer_FullMethodName        = "/grpc.channelz.v1.Channelz/GetServer"
	Channelz_GetServerSockets_FullMethodName = "/grpc.channelz.v1.Channelz/GetServerSockets"
	Channelz_GetChannel_FullMethodName       = "/grpc.channelz.v1.Channelz/GetChannel"
	Channelz_GetSubchannel_FullMethodName    = "/grpc.channelz.v1.Channelz/GetSubchannel"
	Channelz_GetSocket_FullMethodName        = "/grpc.channelz.v1.Channelz/GetSocket"
)

type ChannelzClient interface {
	GetTopChannels(ctx context.Context, in *GetTopChannelsRequest, opts ...grpc.CallOption) (*GetTopChannelsResponse, error)

	GetServers(ctx context.Context, in *GetServersRequest, opts ...grpc.CallOption) (*GetServersResponse, error)

	GetServer(ctx context.Context, in *GetServerRequest, opts ...grpc.CallOption) (*GetServerResponse, error)

	GetServerSockets(ctx context.Context, in *GetServerSocketsRequest, opts ...grpc.CallOption) (*GetServerSocketsResponse, error)

	GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*GetChannelResponse, error)

	GetSubchannel(ctx context.Context, in *GetSubchannelRequest, opts ...grpc.CallOption) (*GetSubchannelResponse, error)

	GetSocket(ctx context.Context, in *GetSocketRequest, opts ...grpc.CallOption) (*GetSocketResponse, error)
}

type channelzClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelzClient(cc grpc.ClientConnInterface) ChannelzClient {
	_ = "STUB: not implemented"
	return *new(ChannelzClient)
}

func (c *channelzClient) GetTopChannels(ctx context.Context, in *GetTopChannelsRequest, opts ...grpc.CallOption) (*GetTopChannelsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *channelzClient) GetServers(ctx context.Context, in *GetServersRequest, opts ...grpc.CallOption) (*GetServersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *channelzClient) GetServer(ctx context.Context, in *GetServerRequest, opts ...grpc.CallOption) (*GetServerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *channelzClient) GetServerSockets(ctx context.Context, in *GetServerSocketsRequest, opts ...grpc.CallOption) (*GetServerSocketsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *channelzClient) GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*GetChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *channelzClient) GetSubchannel(ctx context.Context, in *GetSubchannelRequest, opts ...grpc.CallOption) (*GetSubchannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *channelzClient) GetSocket(ctx context.Context, in *GetSocketRequest, opts ...grpc.CallOption) (*GetSocketResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ChannelzServer interface {
	GetTopChannels(context.Context, *GetTopChannelsRequest) (*GetTopChannelsResponse, error)

	GetServers(context.Context, *GetServersRequest) (*GetServersResponse, error)

	GetServer(context.Context, *GetServerRequest) (*GetServerResponse, error)

	GetServerSockets(context.Context, *GetServerSocketsRequest) (*GetServerSocketsResponse, error)

	GetChannel(context.Context, *GetChannelRequest) (*GetChannelResponse, error)

	GetSubchannel(context.Context, *GetSubchannelRequest) (*GetSubchannelResponse, error)

	GetSocket(context.Context, *GetSocketRequest) (*GetSocketResponse, error)
}

type UnimplementedChannelzServer struct{}

func (UnimplementedChannelzServer) GetTopChannels(context.Context, *GetTopChannelsRequest) (*GetTopChannelsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedChannelzServer) GetServers(context.Context, *GetServersRequest) (*GetServersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedChannelzServer) GetServer(context.Context, *GetServerRequest) (*GetServerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedChannelzServer) GetServerSockets(context.Context, *GetServerSocketsRequest) (*GetServerSocketsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedChannelzServer) GetChannel(context.Context, *GetChannelRequest) (*GetChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedChannelzServer) GetSubchannel(context.Context, *GetSubchannelRequest) (*GetSubchannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedChannelzServer) GetSocket(context.Context, *GetSocketRequest) (*GetSocketResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedChannelzServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeChannelzServer interface {
	mustEmbedUnimplementedChannelzServer()
}

func RegisterChannelzServer(s grpc.ServiceRegistrar, srv ChannelzServer) {
	_ = "STUB: not implemented"
	return
}

func _Channelz_GetTopChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _Channelz_GetServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _Channelz_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _Channelz_GetServerSockets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _Channelz_GetChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _Channelz_GetSubchannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _Channelz_GetSocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var Channelz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.channelz.v1.Channelz",
	HandlerType: (*ChannelzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopChannels",
			Handler:    _Channelz_GetTopChannels_Handler,
		},
		{
			MethodName: "GetServers",
			Handler:    _Channelz_GetServers_Handler,
		},
		{
			MethodName: "GetServer",
			Handler:    _Channelz_GetServer_Handler,
		},
		{
			MethodName: "GetServerSockets",
			Handler:    _Channelz_GetServerSockets_Handler,
		},
		{
			MethodName: "GetChannel",
			Handler:    _Channelz_GetChannel_Handler,
		},
		{
			MethodName: "GetSubchannel",
			Handler:    _Channelz_GetSubchannel_Handler,
		},
		{
			MethodName: "GetSocket",
			Handler:    _Channelz_GetSocket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/channelz/v1/channelz.proto",
}
