package grpclb

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/internal/backoff"
	"google.golang.org/grpc/resolver"

	lbpb "google.golang.org/grpc/balancer/grpclb/grpc_lb_v1"
)

func serverListEqual(a, b []*lbpb.Server) bool { _ = "STUB: not implemented"; return false }

func (lb *lbBalancer) processServerList(l *lbpb.ServerList) { _ = "STUB: not implemented"; return }

func (lb *lbBalancer) refreshSubConns(backendAddrs []resolver.Address, fallback bool, pickFirst bool) {
	_ = "STUB: not implemented"
	return
}

type remoteBalancerCCWrapper struct {
	cc      *grpc.ClientConn
	lb      *lbBalancer
	backoff backoff.Strategy
	done    chan struct{}

	streamMu     sync.Mutex
	streamCancel func()

	wg sync.WaitGroup
}

func (lb *lbBalancer) newRemoteBalancerCCWrapper() error { _ = "STUB: not implemented"; return nil }

func (ccw *remoteBalancerCCWrapper) close() { _ = "STUB: not implemented"; return }

func (ccw *remoteBalancerCCWrapper) readServerList(s *balanceLoadClientStream) error {
	_ = "STUB: not implemented"
	return nil
}

func (ccw *remoteBalancerCCWrapper) sendLoadReport(s *balanceLoadClientStream, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (ccw *remoteBalancerCCWrapper) callRemoteBalancer(ctx context.Context) (backoff bool, _ error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ccw *remoteBalancerCCWrapper) cancelRemoteBalancerCall() { _ = "STUB: not implemented"; return }

func (ccw *remoteBalancerCCWrapper) watchRemoteBalancer() { _ = "STUB: not implemented"; return }
