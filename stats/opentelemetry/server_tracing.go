package opentelemetry

import (
	"context"

	"google.golang.org/grpc/stats"
)

type serverTracingHandler struct {
	options Options
}

func (h *serverTracingHandler) initializeTraces() { _ = "STUB: not implemented"; return }

func (h *serverTracingHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *serverTracingHandler) traceTagRPC(ctx context.Context, ai *attemptInfo) (context.Context, *attemptInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *serverTracingHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}

func (h *serverTracingHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *serverTracingHandler) HandleConn(context.Context, stats.ConnStats) {
	_ = "STUB: not implemented"
	return
}
