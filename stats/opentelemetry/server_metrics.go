package opentelemetry

import (
	"context"
	"sync/atomic"

	"google.golang.org/grpc"
	estats "google.golang.org/grpc/experimental/stats"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
)

type serverMetricsHandler struct {
	estats.MetricsRecorder
	options       Options
	serverMetrics serverMetrics
}

func (h *serverMetricsHandler) initializeMetrics() { _ = "STUB: not implemented"; return }

type attachLabelsTransportStream struct {
	grpc.ServerTransportStream

	attachedLabels         atomic.Bool
	metadataExchangeLabels metadata.MD
}

func (s *attachLabelsTransportStream) SetHeader(md metadata.MD) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *attachLabelsTransportStream) SendHeader(md metadata.MD) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *serverMetricsHandler) unaryInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type attachLabelsStream struct {
	grpc.ServerStream

	attachedLabels         atomic.Bool
	metadataExchangeLabels metadata.MD
}

func (s *attachLabelsStream) SetHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (s *attachLabelsStream) SendHeader(md metadata.MD) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *attachLabelsStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (h *serverMetricsHandler) streamInterceptor(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *serverMetricsHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *serverMetricsHandler) HandleConn(context.Context, stats.ConnStats) {
	_ = "STUB: not implemented"
	return
}

func (h *serverMetricsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *serverMetricsHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}

func (h *serverMetricsHandler) processRPCData(ctx context.Context, s stats.RPCStats, ai *attemptInfo) {
	_ = "STUB: not implemented"
	return
}

func (h *serverMetricsHandler) processRPCEnd(ctx context.Context, ai *attemptInfo, e *stats.End) {
	_ = "STUB: not implemented"
	return
}

const (
	ServerCallStartedMetricName string = "grpc.server.call.started"

	ServerCallSentCompressedTotalMessageSizeMetricName string = "grpc.server.call.sent_total_compressed_message_size"

	ServerCallRcvdCompressedTotalMessageSizeMetricName string = "grpc.server.call.rcvd_total_compressed_message_size"

	ServerCallDurationMetricName string = "grpc.server.call.duration"
)
