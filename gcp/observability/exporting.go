package observability

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/stats/opencensus"

	gcplogging "cloud.google.com/go/logging"
)

var cOptsDisableLogTrace = []option.ClientOption{
	option.WithTelemetryDisabled(),
	option.WithGRPCDialOption(internal.DisableGlobalDialOptions.(func() grpc.DialOption)()),
	option.WithGRPCDialOption(opencensus.DialOption(opencensus.TraceOptions{
		DisableTrace: true,
	})),
}

type loggingExporter interface {
	EmitGcpLoggingEntry(entry gcplogging.Entry)

	Close() error
}

type cloudLoggingExporter struct {
	projectID string
	client    *gcplogging.Client
	logger    *gcplogging.Logger
}

func newCloudLoggingExporter(ctx context.Context, config *config) (loggingExporter, error) {
	_ = "STUB: not implemented"
	return *new(loggingExporter), nil
}

func (cle *cloudLoggingExporter) EmitGcpLoggingEntry(entry gcplogging.Entry) {
	_ = "STUB: not implemented"
	return
}

func (cle *cloudLoggingExporter) Close() error { _ = "STUB: not implemented"; return nil }
