package opentelemetry

import (
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type TraceOptions struct {
	TracerProvider trace.TracerProvider

	TextMapPropagator propagation.TextMapPropagator
}
