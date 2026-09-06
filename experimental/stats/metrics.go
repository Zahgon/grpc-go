package stats

import (
	"context"

	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/stats"
)

type customLabelKey struct{}

func NewContextWithCustomLabel(ctx context.Context, label string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func CustomLabelFromContext(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

type MetricsRecorder interface {
	RecordInt64Count(handle *Int64CountHandle, incr int64, labels ...string)

	RecordFloat64Count(handle *Float64CountHandle, incr float64, labels ...string)

	RecordInt64Histo(handle *Int64HistoHandle, incr int64, labels ...string)

	RecordFloat64Histo(handle *Float64HistoHandle, incr float64, labels ...string)

	RecordInt64Gauge(handle *Int64GaugeHandle, incr int64, labels ...string)

	RecordInt64UpDownCount(handle *Int64UpDownCountHandle, incr int64, labels ...string)

	RegisterAsyncReporter(reporter AsyncMetricReporter, descriptors ...AsyncMetric) func()

	internal.EnforceMetricsRecorderEmbedding
}

type AsyncMetricReporter interface {
	Report(AsyncMetricsRecorder) error
}

type AsyncMetricReporterFunc func(AsyncMetricsRecorder) error

func (f AsyncMetricReporterFunc) Report(r AsyncMetricsRecorder) error {
	_ = "STUB: not implemented"
	return nil
}

type AsyncMetricsRecorder interface {
	RecordInt64AsyncGauge(handle *Int64AsyncGaugeHandle, incr int64, labels ...string)
}

type Metrics = stats.MetricSet

type Metric = string

func NewMetrics(metrics ...Metric) *Metrics { _ = "STUB: not implemented"; return nil }

type UnimplementedMetricsRecorder struct {
	internal.EnforceMetricsRecorderEmbedding
}

func (UnimplementedMetricsRecorder) RecordInt64Count(*Int64CountHandle, int64, ...string) {
	_ = "STUB: not implemented"
	return
}

func (UnimplementedMetricsRecorder) RecordFloat64Count(*Float64CountHandle, float64, ...string) {
	_ = "STUB: not implemented"
	return
}

func (UnimplementedMetricsRecorder) RecordInt64Histo(*Int64HistoHandle, int64, ...string) {
	_ = "STUB: not implemented"
	return
}

func (UnimplementedMetricsRecorder) RecordFloat64Histo(*Float64HistoHandle, float64, ...string) {
	_ = "STUB: not implemented"
	return
}

func (UnimplementedMetricsRecorder) RecordInt64Gauge(*Int64GaugeHandle, int64, ...string) {
	_ = "STUB: not implemented"
	return
}

func (UnimplementedMetricsRecorder) RecordInt64UpDownCount(*Int64UpDownCountHandle, int64, ...string) {
	_ = "STUB: not implemented"
	return
}

func (UnimplementedMetricsRecorder) RegisterAsyncReporter(AsyncMetricReporter, ...AsyncMetric) func() {
	_ = "STUB: not implemented"
	return nil
}
