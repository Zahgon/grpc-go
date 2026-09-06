//go:build !grpcnotrace

package grpc

import (
	"context"
)

func newTrace(family, title string) traceLog { _ = "STUB: not implemented"; return *new(traceLog) }

func newTraceContext(ctx context.Context, tr traceLog) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func newTraceEventLog(family, title string) traceEventLog {
	_ = "STUB: not implemented"
	return *new(traceEventLog)
}
