package csm

import (
	"context"
	"net/url"

	"google.golang.org/grpc"
	"google.golang.org/grpc/stats/opentelemetry"
	otelinternal "google.golang.org/grpc/stats/opentelemetry/internal"
)

func EnableObservability(ctx context.Context, options opentelemetry.Options) func() {
	_ = "STUB: not implemented"
	return nil
}

type perTargetDialOption struct {
	clientSideOTelWithCSM grpc.DialOption
	clientSideOTel        grpc.DialOption
}

func (o *perTargetDialOption) DialOptionForTarget(parsedTarget url.URL) grpc.DialOption {
	_ = "STUB: not implemented"
	return *new(grpc.DialOption)
}

func dialOptionWithCSMPluginOption(options opentelemetry.Options, po otelinternal.PluginOption) grpc.DialOption {
	_ = "STUB: not implemented"
	return *new(grpc.DialOption)
}

func dialOptionSetCSM(options opentelemetry.Options, po otelinternal.PluginOption) grpc.DialOption {
	_ = "STUB: not implemented"
	return *new(grpc.DialOption)
}

func serverOptionWithCSMPluginOption(options opentelemetry.Options, po otelinternal.PluginOption) grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}
