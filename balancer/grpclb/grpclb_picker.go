package grpclb

import (
	"sync"

	"google.golang.org/grpc/balancer"
	lbpb "google.golang.org/grpc/balancer/grpclb/grpc_lb_v1"
)

type rpcStats struct {
	numCallsStarted                        int64
	numCallsFinished                       int64
	numCallsFinishedWithClientFailedToSend int64
	numCallsFinishedKnownReceived          int64

	mu sync.Mutex

	numCallsDropped map[string]int64
}

func newRPCStats() *rpcStats { _ = "STUB: not implemented"; return nil }

func isZeroStats(stats *lbpb.ClientStats) bool { _ = "STUB: not implemented"; return false }

func (s *rpcStats) toClientStats() *lbpb.ClientStats { _ = "STUB: not implemented"; return nil }

func (s *rpcStats) drop(token string) { _ = "STUB: not implemented"; return }

func (s *rpcStats) failedToSend() { _ = "STUB: not implemented"; return }

func (s *rpcStats) knownReceived() { _ = "STUB: not implemented"; return }

type rrPicker struct {
	mu           sync.Mutex
	subConns     []balancer.SubConn
	subConnsNext int
}

func newRRPicker(readySCs []balancer.SubConn) *rrPicker { _ = "STUB: not implemented"; return nil }

func (p *rrPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

type lbPicker struct {
	mu             sync.Mutex
	serverList     []*lbpb.Server
	serverListNext int
	subConns       []balancer.SubConn
	subConnsNext   int

	stats *rpcStats
}

func newLBPicker(serverList []*lbpb.Server, readySCs []balancer.SubConn, stats *rpcStats) *lbPicker {
	_ = "STUB: not implemented"
	return nil
}

func (p *lbPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

func (p *lbPicker) updateReadySCs(readySCs []balancer.SubConn) { _ = "STUB: not implemented"; return }
