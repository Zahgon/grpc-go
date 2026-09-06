package telemetry

import (
	"context"

	stats "google.golang.org/grpc/internal/stats"
)

func NewContextWithLabelCallback(ctx context.Context, callback stats.LabelCallback) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
