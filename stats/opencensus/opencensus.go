package opencensus

import (
	"context"
	"time"

	"go.opencensus.io/trace"

	"google.golang.org/grpc"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/stats"
)

var (
	joinDialOptions = internal.JoinDialOptions.(func(...grpc.DialOption) grpc.DialOption)
)

type TraceOptions struct {
	TS trace.Sampler

	DisableTrace bool
}

func DialOption(to TraceOptions) grpc.DialOption {
	_ = "STUB: not implemented"
	return *new(grpc.DialOption)
}

func ServerOption(to TraceOptions) grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}

func (csh *clientStatsHandler) createCallSpan(ctx context.Context, method string) (context.Context, *trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func perCallTracesAndMetrics(err error, span *trace.Span, startTime time.Time, method string) {
	_ = "STUB: not implemented"
	return
}

func (csh *clientStatsHandler) unaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (csh *clientStatsHandler) streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream), nil
}

type rpcInfo struct {
	mi *metricsInfo
	ti *traceInfo
}

type rpcInfoKey struct{}

func setRPCInfo(ctx context.Context, ri *rpcInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getRPCInfo(ctx context.Context) *rpcInfo { _ = "STUB: not implemented"; return nil }

func SpanContextFromContext(ctx context.Context) (trace.SpanContext, bool) {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext), false
}

type clientStatsHandler struct {
	to TraceOptions
}

func (csh *clientStatsHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (csh *clientStatsHandler) HandleConn(context.Context, stats.ConnStats) {
	_ = "STUB: not implemented"
	return
}

func (csh *clientStatsHandler) TagRPC(ctx context.Context, rti *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (csh *clientStatsHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}

type serverStatsHandler struct {
	to TraceOptions
}

func (ssh *serverStatsHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (ssh *serverStatsHandler) HandleConn(context.Context, stats.ConnStats) {
	_ = "STUB: not implemented"
	return
}

func (ssh *serverStatsHandler) TagRPC(ctx context.Context, rti *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (ssh *serverStatsHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}
