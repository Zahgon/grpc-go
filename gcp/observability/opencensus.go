package observability

import (
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"

	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/stats/opencensus"
)

var (
	defaultMetricsReportingInterval = time.Second * 30
	defaultViews                    = []*view.View{
		opencensus.ClientStartedRPCsView,
		opencensus.ClientCompletedRPCsView,
		opencensus.ClientRoundtripLatencyView,
		opencensus.ClientSentCompressedMessageBytesPerRPCView,
		opencensus.ClientReceivedCompressedMessageBytesPerRPCView,
		opencensus.ClientAPILatencyView,
		opencensus.ServerStartedRPCsView,
		opencensus.ServerCompletedRPCsView,
		opencensus.ServerSentCompressedMessageBytesPerRPCView,
		opencensus.ServerReceivedCompressedMessageBytesPerRPCView,
		opencensus.ServerLatencyView,
	}
)

func labelsToMonitoringLabels(labels map[string]string) *stackdriver.Labels {
	_ = "STUB: not implemented"
	return nil
}

func labelsToTraceAttributes(labels map[string]string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

type tracingMetricsExporter interface {
	trace.Exporter
	view.Exporter
	Flush()
	Close() error
}

var exporter tracingMetricsExporter

var newExporter = newStackdriverExporter

func newStackdriverExporter(config *config) (tracingMetricsExporter, error) {
	_ = "STUB: not implemented"
	return *new(tracingMetricsExporter), nil
}

func generateUniqueProcessIdentifier() string { _ = "STUB: not implemented"; return "" }

func startOpenCensus(config *config) error { _ = "STUB: not implemented"; return nil }

func stopOpenCensus() { _ = "STUB: not implemented"; return }
