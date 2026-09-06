package opentelemetry

import (
	"context"

	otelpropagation "go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

const grpcTraceBinHeaderKey = "grpc-trace-bin"

type GRPCTraceBinPropagator struct{}

func (GRPCTraceBinPropagator) Inject(ctx context.Context, carrier otelpropagation.TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

func (GRPCTraceBinPropagator) Extract(ctx context.Context, carrier otelpropagation.TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (GRPCTraceBinPropagator) Fields() []string { _ = "STUB: not implemented"; return nil }

func toBinary(sc oteltrace.SpanContext) []byte { _ = "STUB: not implemented"; return nil }

func fromBinary(b []byte) (oteltrace.SpanContext, bool) {
	_ = "STUB: not implemented"
	return *new(oteltrace.SpanContext), false
}
