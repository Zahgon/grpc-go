//go:build grpcnotrace

package grpc

import (
	"context"
	"fmt"
)

type notrace struct{}

func (notrace) LazyLog(x fmt.Stringer, sensitive bool) { _ = "STUB: not implemented"; return }
func (notrace) LazyPrintf(format string, a ...any)     { _ = "STUB: not implemented"; return }
func (notrace) SetError()                              { _ = "STUB: not implemented"; return }
func (notrace) SetRecycler(f func(any))                { _ = "STUB: not implemented"; return }
func (notrace) SetTraceInfo(traceID, spanID uint64)    { _ = "STUB: not implemented"; return }
func (notrace) SetMaxEvents(m int)                     { _ = "STUB: not implemented"; return }
func (notrace) Finish()                                { _ = "STUB: not implemented"; return }

func newTrace(family, title string) traceLog { _ = "STUB: not implemented"; return *new(traceLog) }

func newTraceContext(ctx context.Context, tr traceLog) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func newTraceEventLog(family, title string) traceEventLog {
	_ = "STUB: not implemented"
	return *new(traceEventLog)
}
