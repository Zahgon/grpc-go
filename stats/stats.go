package stats

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc/metadata"
)

type RPCStats interface {
	isRPCStats()

	IsClient() bool
}

type Begin struct {
	Client bool

	BeginTime time.Time

	FailFast bool

	IsClientStream bool

	IsServerStream bool

	IsTransparentRetryAttempt bool
}

func (s *Begin) IsClient() bool { _ = "STUB: not implemented"; return false }

func (s *Begin) isRPCStats() { _ = "STUB: not implemented"; return }

type DelayedPickComplete struct{}

func (*DelayedPickComplete) IsClient() bool { _ = "STUB: not implemented"; return false }

func (*DelayedPickComplete) isRPCStats() { _ = "STUB: not implemented"; return }

type PickerUpdated = DelayedPickComplete

type InPayload struct {
	Client bool

	Payload any

	Length int

	CompressedLength int

	WireLength int

	RecvTime time.Time
}

func (s *InPayload) IsClient() bool { _ = "STUB: not implemented"; return false }

func (s *InPayload) isRPCStats() { _ = "STUB: not implemented"; return }

type InHeader struct {
	Client bool

	WireLength int

	Compression string

	Header metadata.MD

	FullMethod string

	RemoteAddr net.Addr

	LocalAddr net.Addr
}

func (s *InHeader) IsClient() bool { _ = "STUB: not implemented"; return false }

func (s *InHeader) isRPCStats() { _ = "STUB: not implemented"; return }

type InTrailer struct {
	Client bool

	WireLength int

	Trailer metadata.MD
}

func (s *InTrailer) IsClient() bool { _ = "STUB: not implemented"; return false }

func (s *InTrailer) isRPCStats() { _ = "STUB: not implemented"; return }

type OutPayload struct {
	Client bool

	Payload any

	Length int

	CompressedLength int

	WireLength int

	SentTime time.Time
}

func (s *OutPayload) IsClient() bool { _ = "STUB: not implemented"; return false }

func (s *OutPayload) isRPCStats() { _ = "STUB: not implemented"; return }

type OutHeader struct {
	Client bool

	Compression string

	Header metadata.MD

	Authority string

	FullMethod string

	RemoteAddr net.Addr

	LocalAddr net.Addr
}

func (s *OutHeader) IsClient() bool { _ = "STUB: not implemented"; return false }

func (s *OutHeader) isRPCStats() { _ = "STUB: not implemented"; return }

type OutTrailer struct {
	Client bool

	WireLength int

	Trailer metadata.MD
}

func (s *OutTrailer) IsClient() bool { _ = "STUB: not implemented"; return false }

func (s *OutTrailer) isRPCStats() { _ = "STUB: not implemented"; return }

type End struct {
	Client bool

	BeginTime time.Time

	EndTime time.Time

	Trailer metadata.MD

	Error error
}

func (s *End) IsClient() bool { _ = "STUB: not implemented"; return false }

func (s *End) isRPCStats() { _ = "STUB: not implemented"; return }

type ConnStats interface {
	isConnStats()

	IsClient() bool
}

type ConnBegin struct {
	Client bool
}

func (s *ConnBegin) IsClient() bool { _ = "STUB: not implemented"; return false }

func (s *ConnBegin) isConnStats() { _ = "STUB: not implemented"; return }

type ConnEnd struct {
	Client bool
}

func (s *ConnEnd) IsClient() bool { _ = "STUB: not implemented"; return false }

func (s *ConnEnd) isConnStats() { _ = "STUB: not implemented"; return }

func SetTags(ctx context.Context, b []byte) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func Tags(ctx context.Context) []byte { _ = "STUB: not implemented"; return nil }

func SetTrace(ctx context.Context, b []byte) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func Trace(ctx context.Context) []byte { _ = "STUB: not implemented"; return nil }
