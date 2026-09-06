package grpclog

import "google.golang.org/grpc/grpclog/internal"

type Logger internal.Logger

func SetLogger(l Logger) { _ = "STUB: not implemented"; return }
