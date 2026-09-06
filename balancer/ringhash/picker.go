package ringhash

import (
	"google.golang.org/grpc/balancer"
)

type picker struct {
	ring *ring

	endpointStates map[string]endpointState

	requestHashHeader string

	hasEndpointInConnectingState bool

	randUint64 func() uint64
}

func (p *picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

func (p *picker) endpointState(e *ringEntry) endpointState {
	_ = "STUB: not implemented"
	return *new(endpointState)
}
