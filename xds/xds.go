package xds

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	internaladmin "google.golang.org/grpc/internal/admin"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/xds/csds"

	_ "google.golang.org/grpc/credentials/tls/certprovider/pemfile"
	_ "google.golang.org/grpc/internal/xds/balancer"
	_ "google.golang.org/grpc/internal/xds/clusterspecifier/rls"
	_ "google.golang.org/grpc/internal/xds/httpfilter/extproc"
	_ "google.golang.org/grpc/internal/xds/httpfilter/fault"
	_ "google.golang.org/grpc/internal/xds/httpfilter/rbac"
	_ "google.golang.org/grpc/internal/xds/httpfilter/router"
	_ "google.golang.org/grpc/internal/xds/resolver"
	_ "google.golang.org/grpc/internal/xds/xdsclient/xdslbregistry/converter"

	v3statusgrpc "github.com/envoyproxy/go-control-plane/envoy/service/status/v3"
)

var logger = grpclog.Component("xds")

func init() {
	internaladmin.AddService(func(registrar grpc.ServiceRegistrar) (func(), error) {
		var grpcServer *grpc.Server
		switch ss := registrar.(type) {
		case *grpc.Server:
			grpcServer = ss
		case *GRPCServer:
			sss, ok := ss.gs.(*grpc.Server)
			if !ok {
				logger.Warning("grpc server within xds.GRPCServer is not *grpc.Server, CSDS will not be registered")
				return nil, nil
			}
			grpcServer = sss
		default:

			logger.Error("Server to register service on is neither a *grpc.Server or a *xds.GRPCServer, CSDS will not be registered")
			return nil, nil
		}

		csdss, err := csds.NewClientStatusDiscoveryServer()
		if err != nil {
			return nil, fmt.Errorf("failed to create csds server: %v", err)
		}
		v3statusgrpc.RegisterClientStatusDiscoveryServiceServer(grpcServer, csdss)
		return csdss.Close, nil
	})
}

func NewXDSResolverWithConfigForTesting(bootstrapConfig []byte) (resolver.Builder, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Builder), nil
}
