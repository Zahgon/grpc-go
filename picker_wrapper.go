package grpc

import (
	"context"
	"sync/atomic"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/internal/transport"
)

type pickerGeneration struct {
	picker balancer.Picker

	blockingCh chan struct{}
}

type pickerWrapper struct {
	pickerGen atomic.Pointer[pickerGeneration]
}

func newPickerWrapper() *pickerWrapper { _ = "STUB: not implemented"; return nil }

func (pw *pickerWrapper) updatePicker(p balancer.Picker) { _ = "STUB: not implemented"; return }

func doneChannelzWrapper(acbw *acBalancerWrapper, result *balancer.PickResult) {
	_ = "STUB: not implemented"
	return
}

type pick struct {
	transport transport.ClientTransport
	result    balancer.PickResult
	blocked   bool
}

func (pw *pickerWrapper) pick(ctx context.Context, failfast bool, info balancer.PickInfo) (pick, error) {
	_ = "STUB: not implemented"
	return *new(pick), nil
}

func (pw *pickerWrapper) close() { _ = "STUB: not implemented"; return }

func (pw *pickerWrapper) reset() { _ = "STUB: not implemented"; return }

type dropError struct {
	error
}
