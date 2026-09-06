package interop

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark/stats"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
)

type SoakWorkerResults struct {
	iterationsSucceeded int
	Failures            int
	Latencies           *stats.Histogram
}

type SoakIterationConfig struct {
	RequestSize  int
	ResponseSize int
	Client       testgrpc.TestServiceClient
	CallOptions  []grpc.CallOption
}

type SoakTestConfig struct {
	RequestSize                      int
	ResponseSize                     int
	PerIterationMaxAcceptableLatency time.Duration
	MinTimeBetweenRPCs               time.Duration
	OverallTimeout                   time.Duration
	ServerAddr                       string
	NumWorkers                       int
	Iterations                       int
	MaxFailures                      int
	ChannelForTest                   func() (*grpc.ClientConn, func())
}

func doOneSoakIteration(ctx context.Context, config SoakIterationConfig) (latency time.Duration, err error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func executeSoakTestInWorker(ctx context.Context, config SoakTestConfig, startTime time.Time, workerID int, soakWorkerResults *SoakWorkerResults) {
	_ = "STUB: not implemented"
	return
}

func DoSoakTest(ctx context.Context, soakConfig SoakTestConfig) { _ = "STUB: not implemented"; return }
