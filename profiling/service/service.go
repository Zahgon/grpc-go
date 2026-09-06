package service

import (
	"context"
	"errors"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/profiling"
	ppb "google.golang.org/grpc/profiling/proto"
)

var logger = grpclog.Component("profiling")

type ProfilingConfig struct {
	Enabled bool

	StreamStatsSize uint32

	Server *grpc.Server
}

var errorNilServer = errors.New("profiling: no grpc.Server provided")

func Init(pc *ProfilingConfig) error { _ = "STUB: not implemented"; return nil }

type profilingServer struct {
	ppb.UnimplementedProfilingServer
	drainMutex sync.Mutex
}

var profilingServerInstance *profilingServer
var profilingServerOnce sync.Once

func getProfilingServerInstance() *profilingServer { _ = "STUB: not implemented"; return nil }

func (s *profilingServer) Enable(_ context.Context, req *ppb.EnableRequest) (*ppb.EnableResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func timerToProtoTimer(timer *profiling.Timer) *ppb.Timer { _ = "STUB: not implemented"; return nil }

func statToProtoStat(stat *profiling.Stat) *ppb.Stat { _ = "STUB: not implemented"; return nil }

func (s *profilingServer) GetStreamStats(context.Context, *ppb.GetStreamStatsRequest) (*ppb.GetStreamStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
