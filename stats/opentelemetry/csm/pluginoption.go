package csm

import (
	"context"
	"net/url"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats/opentelemetry/internal"
	"google.golang.org/protobuf/types/known/structpb"

	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

var logger = grpclog.Component("csm-observability-plugin")

type pluginOption struct {
	localLabels map[string]string

	metadataExchangeLabelsEncoded string
}

func newPluginOption(ctx context.Context) internal.PluginOption {
	_ = "STUB: not implemented"
	return *new(internal.PluginOption)
}

func (cpo *pluginOption) GetMetadata() metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

func (cpo *pluginOption) GetLabels(md metadata.MD) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func getFromMetadata(metadataKey string, metadata map[string]*structpb.Value) string {
	_ = "STUB: not implemented"
	return ""
}

func getFromResource(resourceKey attribute.Key, set *attribute.Set) string {
	_ = "STUB: not implemented"
	return ""
}

func getEnv(name string) string { _ = "STUB: not implemented"; return "" }

var (
	getAttrSetFromResourceDetector = func(ctx context.Context) *attribute.Set {
		r, err := resource.New(ctx, resource.WithFromEnv(), resource.WithDetectors(gcp.NewDetector()))
		if err != nil {
			logger.Warningf("error reading OpenTelemetry resource: %v", err)
		}
		if r != nil {

			return r.Set()
		}
		return nil
	}
)

func constructMetadataFromEnv(ctx context.Context) (map[string]string, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func initializeLocalAndMetadataLabels(labels map[string]string) (map[string]string, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

const metadataExchangeKey = "x-envoy-peer-metadata"

func determineTargetCSM(parsedTarget *url.URL) bool { _ = "STUB: not implemented"; return false }
