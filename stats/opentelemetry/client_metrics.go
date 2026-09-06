package opentelemetry

import (
	"context"
	"time"

	"google.golang.org/grpc"
	estats "google.golang.org/grpc/experimental/stats"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
)

type clientMetricsHandler struct {
	estats.MetricsRecorder
	options       Options
	clientMetrics clientMetrics
}

func (h *clientMetricsHandler) initializeMetrics() { _ = "STUB: not implemented"; return }

func getOrCreateCallInfo(ctx context.Context, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (context.Context, *callInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *clientMetricsHandler) unaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

func determineMethod(method string, opts ...grpc.CallOption) string {
	_ = "STUB: not implemented"
	return ""
}

func (h *clientMetricsHandler) streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream), nil
}

func (h *clientMetricsHandler) perCallMetrics(ctx context.Context, err error, startTime time.Time, ci *callInfo) {
	_ = "STUB: not implemented"
	return
}

func (h *clientMetricsHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *clientMetricsHandler) HandleConn(context.Context, stats.ConnStats) {
	_ = "STUB: not implemented"
	return
}

func (h *clientMetricsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *clientMetricsHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}

func (h *clientMetricsHandler) processRPCEvent(ctx context.Context, s stats.RPCStats, ai *attemptInfo) {
	_ = "STUB: not implemented"
	return
}

func (h *clientMetricsHandler) setLabelsFromPluginOption(ai *attemptInfo, incomingMetadata metadata.MD) {
	_ = "STUB: not implemented"
	return
}

func (h *clientMetricsHandler) processRPCEnd(ctx context.Context, ai *attemptInfo, e *stats.End) {
	_ = "STUB: not implemented"
	return
}

const (
	ClientAttemptStartedMetricName string = "grpc.client.attempt.started"

	ClientAttemptDurationMetricName string = "grpc.client.attempt.duration"

	ClientAttemptSentCompressedTotalMessageSizeMetricName string = "grpc.client.attempt.sent_total_compressed_message_size"

	ClientAttemptRcvdCompressedTotalMessageSizeMetricName string = "grpc.client.attempt.rcvd_total_compressed_message_size"

	ClientCallDurationMetricName string = "grpc.client.call.duration"
)
