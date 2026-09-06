package stats

type MetricSet struct {
	metrics map[string]bool
}

func NewMetricSet(metricNames ...string) *MetricSet { _ = "STUB: not implemented"; return nil }

func (m *MetricSet) Metrics() map[string]bool { _ = "STUB: not implemented"; return nil }

func (m *MetricSet) Add(metricNames ...string) *MetricSet { _ = "STUB: not implemented"; return nil }

func (m *MetricSet) Join(metrics *MetricSet) *MetricSet { _ = "STUB: not implemented"; return nil }

func (m *MetricSet) Remove(metricNames ...string) *MetricSet { _ = "STUB: not implemented"; return nil }
