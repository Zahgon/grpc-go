package main

import (
	"context"
	"flag"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark/stats"
	"google.golang.org/grpc/internal/syscall"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"

	_ "google.golang.org/grpc/xds"
)

var caFile = flag.String("ca_file", "", "The file containing the CA root cert file")

type lockingHistogram struct {
	mu        sync.Mutex
	histogram *stats.Histogram
}

func (h *lockingHistogram) add(value int64) { _ = "STUB: not implemented"; return }

func (h *lockingHistogram) swap(o *stats.Histogram) *stats.Histogram {
	_ = "STUB: not implemented"
	return nil
}

func (h *lockingHistogram) mergeInto(merged *stats.Histogram) { _ = "STUB: not implemented"; return }

type benchmarkClient struct {
	closeConns        func()
	lastResetTime     time.Time
	histogramOptions  stats.HistogramOptions
	lockingHistograms []lockingHistogram
	rusageLastReset   *syscall.Rusage
}

func printClientConfig(config *testpb.ClientConfig) { _ = "STUB: not implemented"; return }

func setupClientEnv(config *testpb.ClientConfig) { _ = "STUB: not implemented"; return }

func createConns(config *testpb.ClientConfig) ([]*grpc.ClientConn, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func performRPCs(ctx context.Context, config *testpb.ClientConfig, conns []*grpc.ClientConn, bc *benchmarkClient) error {
	_ = "STUB: not implemented"
	return nil
}

func startBenchmarkClient(ctx context.Context, config *testpb.ClientConfig) (*benchmarkClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bc *benchmarkClient) unaryLoop(ctx context.Context, conns []*grpc.ClientConn, rpcCountPerConn int, reqSize int, respSize int, poissonLambda *float64) {
	_ = "STUB: not implemented"
	return
}

func (bc *benchmarkClient) streamingLoop(ctx context.Context, conns []*grpc.ClientConn, rpcCountPerConn int, reqSize int, respSize int, payloadType string, poissonLambda *float64) {
	_ = "STUB: not implemented"
	return
}

func (bc *benchmarkClient) poissonUnary(client testgrpc.BenchmarkServiceClient, idx int, reqSize int, respSize int, lambda float64) {
	_ = "STUB: not implemented"
	return
}

func (bc *benchmarkClient) poissonStreaming(stream testgrpc.BenchmarkService_StreamingCallClient, idx int, reqSize int, respSize int, lambda float64, doRPC func(testgrpc.BenchmarkService_StreamingCallClient, int, int) error) {
	_ = "STUB: not implemented"
	return
}

func (bc *benchmarkClient) getStats(reset bool) *testpb.ClientStats {
	_ = "STUB: not implemented"
	return nil
}

func (bc *benchmarkClient) shutdown() { _ = "STUB: not implemented"; return }
