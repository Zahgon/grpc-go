package rls

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/rls/internal/adaptive"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/internal/grpcsync"
	rlsgrpc "google.golang.org/grpc/internal/proto/grpc_lookup_v1"
	rlspb "google.golang.org/grpc/internal/proto/grpc_lookup_v1"
)

var newAdaptiveThrottler = func() adaptiveThrottler { return adaptive.New() }

var newConnectivityStateSubscriber = func(sub grpcsync.Subscriber) grpcsync.Subscriber {
	return sub
}

type adaptiveThrottler interface {
	ShouldThrottle() bool
	RegisterBackendResponse(throttled bool)
}

type controlChannel struct {
	rpcTimeout time.Duration

	backToReadyFunc func()

	throttler adaptiveThrottler

	cc                      *grpc.ClientConn
	client                  rlsgrpc.RouteLookupServiceClient
	logger                  *internalgrpclog.PrefixLogger
	dropConnStateSubscriber func()
	seenTransientFailure    bool
}

func newControlChannel(rlsServerName, serviceConfig string, rpcTimeout time.Duration, bOpts balancer.BuildOptions, backToReadyFunc func()) (*controlChannel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cc *controlChannel) OnMessage(msg any) { _ = "STUB: not implemented"; return }

func (cc *controlChannel) dialOpts(bOpts balancer.BuildOptions, serviceConfig string) ([]grpc.DialOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cc *controlChannel) close() { _ = "STUB: not implemented"; return }

type lookupCallback func(targets []string, headerData string, err error)

func (cc *controlChannel) lookup(reqKeys map[string]string, reason rlspb.RouteLookupRequest_Reason, staleHeaders string, cb lookupCallback) (throttled bool) {
	_ = "STUB: not implemented"
	return false
}
