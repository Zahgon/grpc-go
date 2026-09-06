package benchmark

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var logger = grpclog.Component("benchmark")

func setPayload(p *testpb.Payload, t testpb.PayloadType, size int) {
	_ = "STUB: not implemented"
	return
}

func NewPayload(t testpb.PayloadType, size int) *testpb.Payload {
	_ = "STUB: not implemented"
	return nil
}

type testServer struct {
	testgrpc.UnimplementedBenchmarkServiceServer
}

func (s *testServer) UnaryCall(_ context.Context, in *testpb.SimpleRequest) (*testpb.SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const UnconstrainedStreamingHeader = "unconstrained-streaming"

const UnconstrainedStreamingDelayHeader = "unconstrained-streaming-delay"

const PreloadMsgSizeHeader = "preload-msg-size"

func (s *testServer) StreamingCall(stream testgrpc.BenchmarkService_StreamingCallServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *testServer) UnconstrainedStreamingCall(stream testgrpc.BenchmarkService_StreamingCallServer, preloadMsgSize int) error {
	_ = "STUB: not implemented"
	return nil
}

type byteBufServer struct {
	testgrpc.UnimplementedBenchmarkServiceServer
	respSize int32
}

func (s *byteBufServer) UnaryCall(context.Context, *testpb.SimpleRequest) (*testpb.SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *byteBufServer) StreamingCall(stream testgrpc.BenchmarkService_StreamingCallServer) error {
	_ = "STUB: not implemented"
	return nil
}

type ServerInfo struct {
	Type string

	Metadata any

	Listener net.Listener
}

func StartServer(info ServerInfo, opts ...grpc.ServerOption) func() {
	_ = "STUB: not implemented"
	return nil
}

func DoUnaryCall(tc testgrpc.BenchmarkServiceClient, reqSize, respSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func DoStreamingRoundTrip(stream testgrpc.BenchmarkService_StreamingCallClient, reqSize, respSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func DoStreamingRoundTripPreloaded(stream testgrpc.BenchmarkService_StreamingCallClient, req any) error {
	_ = "STUB: not implemented"
	return nil
}

func DoByteBufStreamingRoundTrip(stream testgrpc.BenchmarkService_StreamingCallClient, reqSize, _ int) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClientConn(addr string, opts ...grpc.DialOption) *grpc.ClientConn {
	_ = "STUB: not implemented"
	return nil
}

func NewClientConnWithContext(_ context.Context, addr string, opts ...grpc.DialOption) *grpc.ClientConn {
	_ = "STUB: not implemented"
	return nil
}
