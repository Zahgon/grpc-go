package stats

import (
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/stats"
)

func init() {
	internal.SnapshotMetricRegistryForTesting = snapshotMetricsRegistryForTesting
}

var logger = grpclog.Component("metrics-registry")

var DefaultMetrics = stats.NewMetricSet()

type MetricDescriptor struct {
	Name string

	Description string

	Unit string

	Labels []string

	OptionalLabels []string

	Default bool

	Type MetricType

	Bounds []float64
}

type MetricType int

const (
	MetricTypeIntCount MetricType = iota
	MetricTypeFloatCount
	MetricTypeIntHisto
	MetricTypeFloatHisto
	MetricTypeIntGauge
	MetricTypeIntUpDownCount
	MetricTypeIntAsyncGauge
)

type Int64CountHandle MetricDescriptor

func (h *Int64CountHandle) Descriptor() *MetricDescriptor { _ = "STUB: not implemented"; return nil }

func (h *Int64CountHandle) Record(recorder MetricsRecorder, incr int64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

type Int64UpDownCountHandle MetricDescriptor

func (h *Int64UpDownCountHandle) Descriptor() *MetricDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (h *Int64UpDownCountHandle) Record(recorder MetricsRecorder, v int64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

type Float64CountHandle MetricDescriptor

func (h *Float64CountHandle) Descriptor() *MetricDescriptor { _ = "STUB: not implemented"; return nil }

func (h *Float64CountHandle) Record(recorder MetricsRecorder, incr float64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

type Int64HistoHandle MetricDescriptor

func (h *Int64HistoHandle) Descriptor() *MetricDescriptor { _ = "STUB: not implemented"; return nil }

func (h *Int64HistoHandle) Record(recorder MetricsRecorder, incr int64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

type Float64HistoHandle MetricDescriptor

func (h *Float64HistoHandle) Descriptor() *MetricDescriptor { _ = "STUB: not implemented"; return nil }

func (h *Float64HistoHandle) Record(recorder MetricsRecorder, incr float64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

type Int64GaugeHandle MetricDescriptor

func (h *Int64GaugeHandle) Descriptor() *MetricDescriptor { _ = "STUB: not implemented"; return nil }

func (h *Int64GaugeHandle) Record(recorder MetricsRecorder, incr int64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

type AsyncMetric interface {
	isAsync()
	Descriptor() *MetricDescriptor
}

type Int64AsyncGaugeHandle MetricDescriptor

func (h *Int64AsyncGaugeHandle) isAsync() { _ = "STUB: not implemented"; return }

func (h *Int64AsyncGaugeHandle) Descriptor() *MetricDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (h *Int64AsyncGaugeHandle) Record(recorder AsyncMetricsRecorder, value int64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

var registeredMetrics = make(map[string]bool)

var metricsRegistry = make(map[string]*MetricDescriptor)

func DescriptorForMetric(metricName string) *MetricDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func registerMetric(metricName string, def bool) { _ = "STUB: not implemented"; return }

func RegisterInt64Count(descriptor MetricDescriptor) *Int64CountHandle {
	_ = "STUB: not implemented"
	return nil
}

func RegisterFloat64Count(descriptor MetricDescriptor) *Float64CountHandle {
	_ = "STUB: not implemented"
	return nil
}

func RegisterInt64Histo(descriptor MetricDescriptor) *Int64HistoHandle {
	_ = "STUB: not implemented"
	return nil
}

func RegisterFloat64Histo(descriptor MetricDescriptor) *Float64HistoHandle {
	_ = "STUB: not implemented"
	return nil
}

func RegisterInt64Gauge(descriptor MetricDescriptor) *Int64GaugeHandle {
	_ = "STUB: not implemented"
	return nil
}

func RegisterInt64UpDownCount(descriptor MetricDescriptor) *Int64UpDownCountHandle {
	_ = "STUB: not implemented"
	return nil
}

func RegisterInt64AsyncGauge(descriptor MetricDescriptor) *Int64AsyncGaugeHandle {
	_ = "STUB: not implemented"
	return nil
}

func snapshotMetricsRegistryForTesting() func() { _ = "STUB: not implemented"; return nil }
