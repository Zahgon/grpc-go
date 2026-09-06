package opentelemetry

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
)

const (
	delayedResolutionEventName = "Delayed name resolution complete"
	tracerName                 = "grpc-go"
)

type clientTracingHandler struct {
	options Options
}

func (h *clientTracingHandler) initializeTraces() { _ = "STUB: not implemented"; return }

func (h *clientTracingHandler) unaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *clientTracingHandler) streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream), nil
}

func (h *clientTracingHandler) finishTrace(err error, ts trace.Span) {
	_ = "STUB: not implemented"
	return
}

func (h *clientTracingHandler) traceTagRPC(ctx context.Context, ai *attemptInfo, nameResolutionDelayed bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *clientTracingHandler) createCallTraceSpan(ctx context.Context, method string) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

func (h *clientTracingHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *clientTracingHandler) HandleConn(context.Context, stats.ConnStats) {
	_ = "STUB: not implemented"
	return
}

func (h *clientTracingHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *clientTracingHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}
