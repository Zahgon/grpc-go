package csds

import (
	"context"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"

	v3statusgrpc "github.com/envoyproxy/go-control-plane/envoy/service/status/v3"
	v3statuspb "github.com/envoyproxy/go-control-plane/envoy/service/status/v3"
)

var logger = grpclog.Component("xds")

const prefix = "[csds-server %p] "

func prefixLogger(s *ClientStatusDiscoveryServer) *internalgrpclog.PrefixLogger {
	_ = "STUB: not implemented"
	return nil
}

type ClientStatusDiscoveryServer struct {
	logger *internalgrpclog.PrefixLogger
}

func NewClientStatusDiscoveryServer() (*ClientStatusDiscoveryServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ClientStatusDiscoveryServer) StreamClientStatus(stream v3statusgrpc.ClientStatusDiscoveryService_StreamClientStatusServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClientStatusDiscoveryServer) FetchClientStatus(_ context.Context, req *v3statuspb.ClientStatusRequest) (*v3statuspb.ClientStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ClientStatusDiscoveryServer) buildClientStatusRespForReq(req *v3statuspb.ClientStatusRequest) (*v3statuspb.ClientStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ClientStatusDiscoveryServer) Close() { _ = "STUB: not implemented"; return }
