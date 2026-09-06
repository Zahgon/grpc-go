package xds

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/internal/xds/server"
	"google.golang.org/grpc/internal/xds/xdsclient"
)

const serverPrefix = "[xds-server %p] "

var (
	xdsClientPool = xdsclient.DefaultPool
	newGRPCServer = func(opts ...grpc.ServerOption) grpcServer {
		return grpc.NewServer(opts...)
	}
)

type grpcServer interface {
	RegisterService(*grpc.ServiceDesc, any)
	Serve(net.Listener) error
	Stop()
	GracefulStop()
	GetServiceInfo() map[string]grpc.ServiceInfo
}

type GRPCServer struct {
	gs             grpcServer
	quit           *grpcsync.Event
	logger         *internalgrpclog.PrefixLogger
	opts           *server.Options
	xdsC           xdsclient.XDSClient
	xdsClientClose func()
}

func NewGRPCServer(opts ...grpc.ServerOption) (*GRPCServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *GRPCServer) handleServerOptions(opts []grpc.ServerOption) {
	_ = "STUB: not implemented"
	return
}

func (s *GRPCServer) defaultServerOptions() *server.Options { _ = "STUB: not implemented"; return nil }

func (s *GRPCServer) loggingServerModeChangeCallback(addr net.Addr, mode connectivity.ServingMode, err error) {
	_ = "STUB: not implemented"
	return
}

func (s *GRPCServer) RegisterService(sd *grpc.ServiceDesc, ss any) {
	_ = "STUB: not implemented"
	return
}

func (s *GRPCServer) GetServiceInfo() map[string]grpc.ServiceInfo {
	_ = "STUB: not implemented"
	return nil
}

func (s *GRPCServer) Serve(lis net.Listener) error { _ = "STUB: not implemented"; return nil }

func (s *GRPCServer) Stop() { _ = "STUB: not implemented"; return }

func (s *GRPCServer) GracefulStop() { _ = "STUB: not implemented"; return }
