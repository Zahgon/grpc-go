package binarylog

import (
	binlogpb "google.golang.org/grpc/binarylog/grpc_binarylog_v1"
)

func SetSink(s Sink) { _ = "STUB: not implemented"; return }

type Sink interface {
	Write(*binlogpb.GrpcLogEntry) error

	Close() error
}

func NewTempFileSink() (Sink, error) { _ = "STUB: not implemented"; return *new(Sink), nil }
