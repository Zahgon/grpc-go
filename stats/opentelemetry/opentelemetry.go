package opentelemetry

import (
	"context"
	"sync/atomic"
	"time"

	otelmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	experimental "google.golang.org/grpc/experimental/opentelemetry"
	estats "google.golang.org/grpc/experimental/stats"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/stats"
	otelinternal "google.golang.org/grpc/stats/opentelemetry/internal"
)

func init() {
	otelinternal.SetPluginOption = func(o *Options, po otelinternal.PluginOption) {
		o.MetricsOptions.pluginOption = po

		if (o.TraceOptions.TextMapPropagator == nil) != (o.TraceOptions.TracerProvider == nil) {
			logger.Warning("Tracing will not be recorded because traceOptions are not set properly: one of TextMapPropagator or TracerProvider is missing")
		}
	}
}

var (
	logger          = grpclog.Component("otel-plugin")
	canonicalString = internal.CanonicalString.(func(codes.Code) string)
	joinDialOptions = internal.JoinDialOptions.(func(...grpc.DialOption) grpc.DialOption)
)

type Options struct {
	MetricsOptions MetricsOptions

	TraceOptions experimental.TraceOptions
}

func (o *Options) isMetricsEnabled() bool { _ = "STUB: not implemented"; return false }

func (o *Options) isTracingEnabled() bool { _ = "STUB: not implemented"; return false }

type MetricsOptions struct {
	MeterProvider otelmetric.MeterProvider

	Metrics *stats.MetricSet

	MethodAttributeFilter func(string) bool

	OptionalLabels []string

	pluginOption otelinternal.PluginOption
}

func DialOption(o Options) grpc.DialOption { _ = "STUB: not implemented"; return *new(grpc.DialOption) }

var joinServerOptions = internal.JoinServerOptions.(func(...grpc.ServerOption) grpc.ServerOption)

func ServerOption(o Options) grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}

type callInfo struct {
	target string

	method string

	nameResolutionEventAdded atomic.Bool

	previousRPCAttempts atomic.Uint32
}

type callInfoKey struct{}

func getCallInfo(ctx context.Context) *callInfo { _ = "STUB: not implemented"; return nil }

type rpcInfo struct {
	ai *attemptInfo
}

type clientRPCInfoKey struct{}
type serverRPCInfoKey struct{}

func clientRPCInfo(ctx context.Context) *rpcInfo { _ = "STUB: not implemented"; return nil }

func serverRPCInfo(ctx context.Context) *rpcInfo { _ = "STUB: not implemented"; return nil }

func getOrCreateClientRPCInfo(ctx context.Context, info *stats.RPCTagInfo) (context.Context, *rpcInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func getOrCreateServerRPCInfo(ctx context.Context, info *stats.RPCTagInfo, options MetricsOptions) (context.Context, *rpcInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func removeLeadingSlash(mn string) string { _ = "STUB: not implemented"; return "" }

type attemptInfo struct {
	sentCompressedBytes int64

	recvCompressedBytes int64

	startTime time.Time
	method    string

	pluginOptionLabels map[string]string
	xdsLabels          map[string]string

	traceSpan trace.Span

	countSentMsg uint32
	countRecvMsg uint32
}

type clientMetrics struct {
	attemptStarted otelmetric.Int64Counter

	attemptDuration otelmetric.Float64Histogram

	attemptSentTotalCompressedMessageSize otelmetric.Int64Histogram

	attemptRcvdTotalCompressedMessageSize otelmetric.Int64Histogram

	callDuration otelmetric.Float64Histogram
}

type serverMetrics struct {
	callStarted otelmetric.Int64Counter

	callSentTotalCompressedMessageSize otelmetric.Int64Histogram

	callRcvdTotalCompressedMessageSize otelmetric.Int64Histogram

	callDuration otelmetric.Float64Histogram
}

func createInt64Counter(setOfMetrics map[string]bool, metricName string, meter otelmetric.Meter, options ...otelmetric.Int64CounterOption) otelmetric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(otelmetric.Int64Counter)
}

func createInt64UpDownCounter(setOfMetrics map[string]bool, metricName string, meter otelmetric.Meter, options ...otelmetric.Int64UpDownCounterOption) otelmetric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(otelmetric.Int64UpDownCounter)
}

func createFloat64Counter(setOfMetrics map[string]bool, metricName string, meter otelmetric.Meter, options ...otelmetric.Float64CounterOption) otelmetric.Float64Counter {
	_ = "STUB: not implemented"
	return *new(otelmetric.Float64Counter)
}

func createInt64Histogram(setOfMetrics map[string]bool, metricName string, meter otelmetric.Meter, options ...otelmetric.Int64HistogramOption) otelmetric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(otelmetric.Int64Histogram)
}

func createFloat64Histogram(setOfMetrics map[string]bool, metricName string, meter otelmetric.Meter, options ...otelmetric.Float64HistogramOption) otelmetric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(otelmetric.Float64Histogram)
}

func createInt64Gauge(setOfMetrics map[string]bool, metricName string, meter otelmetric.Meter, options ...otelmetric.Int64GaugeOption) otelmetric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(otelmetric.Int64Gauge)
}

func createInt64ObservableGauge(setOfMetrics map[string]bool, metricName string, meter otelmetric.Meter, options ...otelmetric.Int64ObservableGaugeOption) otelmetric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(otelmetric.Int64ObservableGauge)
}

func optionFromLabels(labelKeys []string, optionalLabelKeys []string, optionalLabels []string, labelVals ...string) otelmetric.MeasurementOption {
	_ = "STUB: not implemented"
	return *new(otelmetric.MeasurementOption)
}

type registryMetrics struct {
	internal.EnforceMetricsRecorderEmbedding
	intCounts       map[*estats.MetricDescriptor]otelmetric.Int64Counter
	floatCounts     map[*estats.MetricDescriptor]otelmetric.Float64Counter
	intHistos       map[*estats.MetricDescriptor]otelmetric.Int64Histogram
	floatHistos     map[*estats.MetricDescriptor]otelmetric.Float64Histogram
	intGauges       map[*estats.MetricDescriptor]otelmetric.Int64Gauge
	intUpDownCounts map[*estats.MetricDescriptor]otelmetric.Int64UpDownCounter

	intObservableGauges map[*estats.MetricDescriptor]otelmetric.Int64ObservableGauge

	meter          otelmetric.Meter
	optionalLabels []string
}

func (rm *registryMetrics) registerMetrics(metrics *stats.MetricSet, meter otelmetric.Meter) {
	_ = "STUB: not implemented"
	return
}

func (rm *registryMetrics) RecordInt64Count(handle *estats.Int64CountHandle, incr int64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

func (rm *registryMetrics) RecordInt64UpDownCount(handle *estats.Int64UpDownCountHandle, incr int64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

func (rm *registryMetrics) RecordFloat64Count(handle *estats.Float64CountHandle, incr float64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

func (rm *registryMetrics) RecordInt64Histo(handle *estats.Int64HistoHandle, incr int64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

func (rm *registryMetrics) RecordFloat64Histo(handle *estats.Float64HistoHandle, incr float64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

func (rm *registryMetrics) RecordInt64Gauge(handle *estats.Int64GaugeHandle, incr int64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

func (rm *registryMetrics) RegisterAsyncReporter(reporter estats.AsyncMetricReporter, metrics ...estats.AsyncMetric) func() {
	_ = "STUB: not implemented"
	return nil
}

var (
	DefaultLatencyBounds = []float64{0, 0.00001, 0.00005, 0.0001, 0.0003, 0.0006, 0.0008, 0.001, 0.002, 0.003, 0.004, 0.005, 0.006, 0.008, 0.01, 0.013, 0.016, 0.02, 0.025, 0.03, 0.04, 0.05, 0.065, 0.08, 0.1, 0.13, 0.16, 0.2, 0.25, 0.3, 0.4, 0.5, 0.65, 0.8, 1, 2, 5, 10, 20, 50, 100}

	DefaultSizeBounds = []float64{0, 1024, 2048, 4096, 16384, 65536, 262144, 1048576, 4194304, 16777216, 67108864, 268435456, 1073741824, 4294967296}

	defaultPerCallMetrics = stats.NewMetricSet(ClientAttemptStartedMetricName, ClientAttemptDurationMetricName, ClientAttemptSentCompressedTotalMessageSizeMetricName, ClientAttemptRcvdCompressedTotalMessageSizeMetricName, ClientCallDurationMetricName, ServerCallStartedMetricName, ServerCallSentCompressedTotalMessageSizeMetricName, ServerCallRcvdCompressedTotalMessageSizeMetricName, ServerCallDurationMetricName)
)

func DefaultMetrics() *stats.MetricSet { _ = "STUB: not implemented"; return nil }

type observerAdapter struct {
	observableMap  map[*estats.MetricDescriptor]otelmetric.Observable
	optionalLabels []string
	delegate       otelmetric.Observer
}

func (a *observerAdapter) RecordInt64AsyncGauge(handle *estats.Int64AsyncGaugeHandle, val int64, labels ...string) {
	_ = "STUB: not implemented"
	return
}
