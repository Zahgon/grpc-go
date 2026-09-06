package grpclb

import (
	"sync"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/resolver"
)

const subConnCacheTime = time.Second * 10

type lbCacheClientConn struct {
	balancer.ClientConn

	timeout time.Duration

	mu sync.Mutex

	subConnCache  map[resolver.Address]*subConnCacheEntry
	subConnToAddr map[balancer.SubConn]resolver.Address
}

type subConnCacheEntry struct {
	sc balancer.SubConn

	cancel        func()
	abortDeleting bool
}

func newLBCacheClientConn(cc balancer.ClientConn) *lbCacheClientConn {
	_ = "STUB: not implemented"
	return nil
}

func (ccc *lbCacheClientConn) NewSubConn(addrs []resolver.Address, opts balancer.NewSubConnOptions) (balancer.SubConn, error) {
	_ = "STUB: not implemented"
	return *new(balancer.SubConn), nil
}

func (ccc *lbCacheClientConn) RemoveSubConn(sc balancer.SubConn) { _ = "STUB: not implemented"; return }

type lbCacheSubConn struct {
	balancer.SubConn
	ccc *lbCacheClientConn
}

func (sc *lbCacheSubConn) Shutdown() { _ = "STUB: not implemented"; return }

func (ccc *lbCacheClientConn) UpdateState(s balancer.State) { _ = "STUB: not implemented"; return }

func (ccc *lbCacheClientConn) close() { _ = "STUB: not implemented"; return }

type lbCachePicker struct {
	balancer.Picker
}

func (cp *lbCachePicker) Pick(i balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}
