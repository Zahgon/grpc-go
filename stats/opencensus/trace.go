package opencensus

import (
	"context"

	"go.opencensus.io/trace"

	"google.golang.org/grpc/stats"
)

type traceInfo struct {
	span         *trace.Span
	countSentMsg uint32
	countRecvMsg uint32
}

func (csh *clientStatsHandler) traceTagRPC(ctx context.Context, rti *stats.RPCTagInfo) (context.Context, *traceInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (ssh *serverStatsHandler) traceTagRPC(ctx context.Context, rti *stats.RPCTagInfo) (context.Context, *traceInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func populateSpan(_ context.Context, rs stats.RPCStats, ti *traceInfo) {
	_ = "STUB: not implemented"
	return
}
