package observability

import (
	"context"

	"google.golang.org/grpc/grpclog"
)

var logger = grpclog.Component("observability")

func Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func End() { _ = "STUB: not implemented"; return }
