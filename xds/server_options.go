package xds

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/internal/xds/xdsclient"
)

func ServingModeCallback(cb ServingModeCallbackFunc) grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}

type ServingModeCallbackFunc func(addr net.Addr, args ServingModeChangeArgs)

type ServingModeChangeArgs struct {
	Mode connectivity.ServingMode

	Err error
}

func BootstrapContentsForTesting(bootstrapContents []byte) grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}

func ClientPoolForTesting(pool *xdsclient.Pool) grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}
