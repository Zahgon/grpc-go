package orca

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc/balancer"

	v3orcapb "github.com/cncf/xds/go/xds/data/orca/v3"
	v3orcaservicegrpc "github.com/cncf/xds/go/xds/service/orca/v3"
)

type producerBuilder struct{}

func (*producerBuilder) Build(cci any) (balancer.Producer, func()) {
	_ = "STUB: not implemented"
	return *new(balancer.Producer), nil
}

var producerBuilderSingleton = &producerBuilder{}

type OOBListener interface {
	OnLoadReport(*v3orcapb.OrcaLoadReport)
}

type OOBListenerOptions struct {
	ReportInterval time.Duration
}

func RegisterOOBListener(sc balancer.SubConn, l OOBListener, opts OOBListenerOptions) (stop func()) {
	_ = "STUB: not implemented"
	return nil
}

type producer struct {
	client v3orcaservicegrpc.OpenRcaServiceClient

	backoff func(int) time.Duration
	stopped chan struct{}

	mu          sync.Mutex
	intervals   map[time.Duration]int
	listeners   map[OOBListener]struct{}
	minInterval time.Duration
	stop        func()
}

func (p *producer) registerListener(l OOBListener, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (p *producer) unregisterListener(l OOBListener, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (p *producer) recomputeMinInterval() { _ = "STUB: not implemented"; return }

func (p *producer) updateRunLocked() { _ = "STUB: not implemented"; return }

func (p *producer) run(ctx context.Context, done chan struct{}, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (p *producer) runStream(ctx context.Context, interval time.Duration) (resetBackoff bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}
