package orca

import (
	"context"
	"sync"

	"google.golang.org/grpc"
	grpcinternal "google.golang.org/grpc/internal"
)

type CallMetricsRecorder interface {
	ServerMetricsRecorder

	SetRequestCost(name string, val float64)

	DeleteRequestCost(name string)

	SetNamedMetric(name string, val float64)

	DeleteNamedMetric(name string)
}

type callMetricsRecorderCtxKey struct{}

func CallMetricsRecorderFromContext(ctx context.Context) CallMetricsRecorder {
	_ = "STUB: not implemented"
	return *new(CallMetricsRecorder)
}

type recorderWrapper struct {
	once sync.Once
	r    CallMetricsRecorder
	smp  ServerMetricsProvider
}

func (rw *recorderWrapper) recorder() CallMetricsRecorder {
	_ = "STUB: not implemented"
	return *new(CallMetricsRecorder)
}

func (rw *recorderWrapper) setTrailerMetadata(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

var joinServerOptions = grpcinternal.JoinServerOptions.(func(...grpc.ServerOption) grpc.ServerOption)

func CallMetricsServerOption(smp ServerMetricsProvider) grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}

func unaryInt(smp ServerMetricsProvider) func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return nil
}

func streamInt(smp ServerMetricsProvider) func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func newContextWithRecorderWrapper(ctx context.Context, r *recorderWrapper) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
