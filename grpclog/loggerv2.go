package grpclog

import (
	"io"

	"google.golang.org/grpc/grpclog/internal"
)

type LoggerV2 internal.LoggerV2

func SetLoggerV2(l LoggerV2) { _ = "STUB: not implemented"; return }

func NewLoggerV2(infoW, warningW, errorW io.Writer) LoggerV2 {
	_ = "STUB: not implemented"
	return *new(LoggerV2)
}

func NewLoggerV2WithVerbosity(infoW, warningW, errorW io.Writer, v int) LoggerV2 {
	_ = "STUB: not implemented"
	return *new(LoggerV2)
}

func newLoggerV2() LoggerV2 { _ = "STUB: not implemented"; return *new(LoggerV2) }

type DepthLoggerV2 internal.DepthLoggerV2
