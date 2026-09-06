package opencensus

import (
	"context"
	"time"

	"go.opencensus.io/stats/view"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/stats"
)

var logger = grpclog.Component("opencensus-instrumentation")

var canonicalString = internal.CanonicalString.(func(codes.Code) string)

var (
	bytesDistributionBounds  = []float64{1024, 2048, 4096, 16384, 65536, 262144, 1048576, 4194304, 16777216, 67108864, 268435456, 1073741824, 4294967296}
	bytesDistribution        = view.Distribution(bytesDistributionBounds...)
	millisecondsDistribution = view.Distribution(0.01, 0.05, 0.1, 0.3, 0.6, 0.8, 1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000)
	countDistributionBounds  = []float64{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536}
	countDistribution        = view.Distribution(countDistributionBounds...)
)

func removeLeadingSlash(mn string) string { _ = "STUB: not implemented"; return "" }

type metricsInfo struct {
	sentMsgs int64

	sentBytes int64

	sentCompressedBytes int64

	recvMsgs int64

	recvBytes int64

	recvCompressedBytes int64

	startTime time.Time
	method    string
}

func (csh *clientStatsHandler) statsTagRPC(ctx context.Context, info *stats.RPCTagInfo) (context.Context, *metricsInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (ssh *serverStatsHandler) statsTagRPC(ctx context.Context, info *stats.RPCTagInfo) (context.Context, *metricsInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func recordRPCData(ctx context.Context, s stats.RPCStats, mi *metricsInfo) {
	_ = "STUB: not implemented"
	return
}

func recordDataBegin(ctx context.Context, mi *metricsInfo, b *stats.Begin) {
	_ = "STUB: not implemented"
	return
}

func recordDataOutPayload(mi *metricsInfo, op *stats.OutPayload) { _ = "STUB: not implemented"; return }

func recordDataInPayload(mi *metricsInfo, ip *stats.InPayload) { _ = "STUB: not implemented"; return }

func recordDataEnd(ctx context.Context, mi *metricsInfo, e *stats.End) {
	_ = "STUB: not implemented"
	return
}
