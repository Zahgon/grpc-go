package interop

import (
	"context"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/resolver"

	v3orcapb "github.com/cncf/xds/go/xds/data/orca/v3"
)

var orcaLogger = grpclog.Component("orca")

func init() {
	balancer.Register(orcabb{})
}

type orcabb struct{}

func (orcabb) Build(cc balancer.ClientConn, bOpts balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

func (orcabb) Name() string { _ = "STUB: not implemented"; return "" }

type orcab struct {
	balancer.ClientConn
	child    balancer.Balancer
	oobState *oobState
	logger   *internalgrpclog.PrefixLogger

	mu               sync.Mutex
	stopOOBListeners map[balancer.SubConn]func()
}

func (b *orcab) UpdateClientConnState(s balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *orcab) ResolverError(err error) { _ = "STUB: not implemented"; return }

func (b *orcab) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (b *orcab) ExitIdle() { _ = "STUB: not implemented"; return }

func (b *orcab) Close() { _ = "STUB: not implemented"; return }

func (b *orcab) NewSubConn(addrs []resolver.Address, opts balancer.NewSubConnOptions) (balancer.SubConn, error) {
	_ = "STUB: not implemented"
	return *new(balancer.SubConn), nil
}

func (b *orcab) updateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (b *orcab) UpdateState(state balancer.State) { _ = "STUB: not implemented"; return }

type orcaOOBListener struct {
	subConn  balancer.SubConn
	balancer *orcab
}

func (l *orcaOOBListener) OnLoadReport(r *v3orcapb.OrcaLoadReport) {
	_ = "STUB: not implemented"
	return
}

type oobState struct {
	mu      sync.Mutex
	reports map[balancer.SubConn]*v3orcapb.OrcaLoadReport
}
type orcaPicker struct {
	childPicker balancer.Picker
	oobState    *oobState
}

func (p *orcaPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

func setContextCMR(ctx context.Context, lr *v3orcapb.OrcaLoadReport) {
	_ = "STUB: not implemented"
	return
}

type orcaKey string

var orcaCtxKey = orcaKey("orcaResult")

func contextWithORCAResult(ctx context.Context, result **v3orcapb.OrcaLoadReport) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func orcaResultFromContext(ctx context.Context) **v3orcapb.OrcaLoadReport {
	_ = "STUB: not implemented"
	return nil
}
