package balancer

import "google.golang.org/grpc/connectivity"

type ConnectivityStateEvaluator struct {
	numReady            uint64
	numConnecting       uint64
	numTransientFailure uint64
	numIdle             uint64
}

func (cse *ConnectivityStateEvaluator) RecordTransition(oldState, newState connectivity.State) connectivity.State {
	_ = "STUB: not implemented"
	return *new(connectivity.State)
}

func (cse *ConnectivityStateEvaluator) CurrentState() connectivity.State {
	_ = "STUB: not implemented"
	return *new(connectivity.State)
}
