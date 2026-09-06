package orca

import (
	"sync/atomic"

	v3orcapb "github.com/cncf/xds/go/xds/data/orca/v3"
)

type ServerMetrics struct {
	CPUUtilization float64
	MemUtilization float64
	AppUtilization float64
	QPS            float64
	EPS            float64

	Utilization  map[string]float64
	RequestCost  map[string]float64
	NamedMetrics map[string]float64
}

func (sm *ServerMetrics) toLoadReportProto() *v3orcapb.OrcaLoadReport {
	_ = "STUB: not implemented"
	return nil
}

func (sm *ServerMetrics) merge(o *ServerMetrics) { _ = "STUB: not implemented"; return }

func mergeMap(a, b map[string]float64) { _ = "STUB: not implemented"; return }

type ServerMetricsRecorder interface {
	ServerMetricsProvider

	SetCPUUtilization(float64)

	DeleteCPUUtilization()

	SetMemoryUtilization(float64)

	DeleteMemoryUtilization()

	SetApplicationUtilization(float64)

	DeleteApplicationUtilization()

	SetQPS(float64)

	DeleteQPS()

	SetEPS(float64)

	DeleteEPS()

	SetNamedUtilization(name string, val float64)

	DeleteNamedUtilization(name string)
}

type serverMetricsRecorder struct {
	state atomic.Pointer[ServerMetrics]
}

func NewServerMetricsRecorder() ServerMetricsRecorder {
	_ = "STUB: not implemented"
	return *new(ServerMetricsRecorder)
}

func newServerMetricsRecorder() *serverMetricsRecorder { _ = "STUB: not implemented"; return nil }

func (s *serverMetricsRecorder) ServerMetrics() *ServerMetrics {
	_ = "STUB: not implemented"
	return nil
}

func copyMap(m map[string]float64) map[string]float64 { _ = "STUB: not implemented"; return nil }

func copyServerMetrics(sm *ServerMetrics) *ServerMetrics { _ = "STUB: not implemented"; return nil }

func (s *serverMetricsRecorder) SetCPUUtilization(val float64) { _ = "STUB: not implemented"; return }

func (s *serverMetricsRecorder) DeleteCPUUtilization() { _ = "STUB: not implemented"; return }

func (s *serverMetricsRecorder) SetMemoryUtilization(val float64) {
	_ = "STUB: not implemented"
	return
}

func (s *serverMetricsRecorder) DeleteMemoryUtilization() { _ = "STUB: not implemented"; return }

func (s *serverMetricsRecorder) SetApplicationUtilization(val float64) {
	_ = "STUB: not implemented"
	return
}

func (s *serverMetricsRecorder) DeleteApplicationUtilization() { _ = "STUB: not implemented"; return }

func (s *serverMetricsRecorder) SetQPS(val float64) { _ = "STUB: not implemented"; return }

func (s *serverMetricsRecorder) DeleteQPS() { _ = "STUB: not implemented"; return }

func (s *serverMetricsRecorder) SetEPS(val float64) { _ = "STUB: not implemented"; return }

func (s *serverMetricsRecorder) DeleteEPS() { _ = "STUB: not implemented"; return }

func (s *serverMetricsRecorder) SetNamedUtilization(name string, val float64) {
	_ = "STUB: not implemented"
	return
}

func (s *serverMetricsRecorder) DeleteNamedUtilization(name string) {
	_ = "STUB: not implemented"
	return
}

func (s *serverMetricsRecorder) SetRequestCost(name string, val float64) {
	_ = "STUB: not implemented"
	return
}

func (s *serverMetricsRecorder) DeleteRequestCost(name string) { _ = "STUB: not implemented"; return }

func (s *serverMetricsRecorder) SetNamedMetric(name string, val float64) {
	_ = "STUB: not implemented"
	return
}

func (s *serverMetricsRecorder) DeleteNamedMetric(name string) { _ = "STUB: not implemented"; return }
