package weightedaggregator

import (
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/internal/wrr"
)

type weightedPickerState struct {
	weight uint32
	state  balancer.State

	stateToAggregate connectivity.State
}

func (s *weightedPickerState) String() string { _ = "STUB: not implemented"; return "" }

type Aggregator struct {
	cc     balancer.ClientConn
	logger *grpclog.PrefixLogger
	newWRR func() wrr.WRR

	csEvltr *balancer.ConnectivityStateEvaluator

	mu sync.Mutex

	started bool

	idToPickerState map[string]*weightedPickerState

	pauseUpdateState bool

	needUpdateStateOnResume bool
}

func New(cc balancer.ClientConn, logger *grpclog.PrefixLogger, newWRR func() wrr.WRR) *Aggregator {
	_ = "STUB: not implemented"
	return nil
}

func (wbsa *Aggregator) Start() { _ = "STUB: not implemented"; return }

func (wbsa *Aggregator) Stop() { _ = "STUB: not implemented"; return }

func (wbsa *Aggregator) Add(id string, weight uint32) { _ = "STUB: not implemented"; return }

func (wbsa *Aggregator) Remove(id string) { _ = "STUB: not implemented"; return }

func (wbsa *Aggregator) UpdateWeight(id string, newWeight uint32) {
	_ = "STUB: not implemented"
	return
}

func (wbsa *Aggregator) PauseStateUpdates() { _ = "STUB: not implemented"; return }

func (wbsa *Aggregator) ResumeStateUpdates() { _ = "STUB: not implemented"; return }

func (wbsa *Aggregator) NeedUpdateStateOnResume() { _ = "STUB: not implemented"; return }

func (wbsa *Aggregator) UpdateState(id string, newState balancer.State) {
	_ = "STUB: not implemented"
	return
}

func (wbsa *Aggregator) clearStates() { _ = "STUB: not implemented"; return }

func (wbsa *Aggregator) buildAndUpdateLocked() { _ = "STUB: not implemented"; return }

func (wbsa *Aggregator) build() balancer.State {
	_ = "STUB: not implemented"
	return *new(balancer.State)
}

type weightedPickerGroup struct {
	w wrr.WRR
}

func newWeightedPickerGroup(readyWeightedPickers []weightedPickerState, newWRR func() wrr.WRR) *weightedPickerGroup {
	_ = "STUB: not implemented"
	return nil
}

func (pg *weightedPickerGroup) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}
