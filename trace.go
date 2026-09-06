package grpc

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var EnableTracing bool

func methodFamily(m string) string { _ = "STUB: not implemented"; return "" }

type traceEventLog interface {
	Printf(format string, a ...any)
	Errorf(format string, a ...any)
	Finish()
}

type traceLog interface {
	LazyLog(x fmt.Stringer, sensitive bool)
	LazyPrintf(format string, a ...any)
	SetError()
	SetRecycler(f func(any))
	SetTraceInfo(traceID, spanID uint64)
	SetMaxEvents(m int)
	Finish()
}

type traceInfo struct {
	tr        traceLog
	firstLine firstLine
}

type firstLine struct {
	mu         sync.Mutex
	client     bool
	remoteAddr net.Addr
	deadline   time.Duration
}

func (f *firstLine) SetRemoteAddr(addr net.Addr) { _ = "STUB: not implemented"; return }

func (f *firstLine) String() string { _ = "STUB: not implemented"; return "" }

const truncateSize = 100

func truncate(x string, l int) string { _ = "STUB: not implemented"; return "" }

type payload struct {
	sent bool
	msg  any
}

func (p payload) String() string { _ = "STUB: not implemented"; return "" }

type fmtStringer struct {
	format string
	a      []any
}

func (f *fmtStringer) String() string { _ = "STUB: not implemented"; return "" }

type stringer string

func (s stringer) String() string { _ = "STUB: not implemented"; return "" }
