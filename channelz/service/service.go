package service

import (
	"context"

	channelzgrpc "google.golang.org/grpc/channelz/grpc_channelz_v1"
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/internal/channelz"
)

func init() {
	channelz.TurnOn()
}

func RegisterChannelzServiceToServer(s grpc.ServiceRegistrar) { _ = "STUB: not implemented"; return }

func newCZServer() channelzgrpc.ChannelzServer {
	_ = "STUB: not implemented"
	return *new(channelzgrpc.ChannelzServer)
}

type serverImpl struct {
	channelzgrpc.UnimplementedChannelzServer
}

func (s *serverImpl) GetChannel(_ context.Context, req *channelzpb.GetChannelRequest) (*channelzpb.GetChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serverImpl) GetTopChannels(_ context.Context, req *channelzpb.GetTopChannelsRequest) (*channelzpb.GetTopChannelsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serverImpl) GetServer(_ context.Context, req *channelzpb.GetServerRequest) (*channelzpb.GetServerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serverImpl) GetServers(_ context.Context, req *channelzpb.GetServersRequest) (*channelzpb.GetServersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serverImpl) GetSubchannel(_ context.Context, req *channelzpb.GetSubchannelRequest) (*channelzpb.GetSubchannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serverImpl) GetServerSockets(_ context.Context, req *channelzpb.GetServerSocketsRequest) (*channelzpb.GetServerSocketsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serverImpl) GetSocket(_ context.Context, req *channelzpb.GetSocketRequest) (*channelzpb.GetSocketResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
