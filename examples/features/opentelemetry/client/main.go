package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	otelstdouttrace "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	otelpropagation "go.opentelemetry.io/otel/propagation"
	otelmetric "go.opentelemetry.io/otel/sdk/metric"
	otelresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/features/proto/echo"
	oteltracing "google.golang.org/grpc/experimental/opentelemetry"
	"google.golang.org/grpc/stats/opentelemetry"
)

var (
	addr               = flag.String("addr", "localhost:50051", "the server address to connect to")
	prometheusEndpoint = flag.String("prometheus_endpoint", ":9465", "the Prometheus exporter endpoint for metrics")
)

func main() {
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("Failed to start prometheus exporter: %v", err)
	}

	meterProvider := otelmetric.NewMeterProvider(otelmetric.WithReader(exporter))

	traceExporter, err := otelstdouttrace.New(otelstdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("Failed to create stdouttrace exporter: %v", err)
	}
	traceProvider := sdktrace.NewTracerProvider(sdktrace.WithBatcher(traceExporter), sdktrace.WithResource(otelresource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceName("grpc-client"))))

	textMapPropagator := otelpropagation.TraceContext{}
	do := opentelemetry.DialOption(opentelemetry.Options{
		MetricsOptions: opentelemetry.MetricsOptions{
			MeterProvider: meterProvider,

			Metrics: opentelemetry.DefaultMetrics().Add(
				"grpc.subchannel.connection_attempts_succeeded",
				"grpc.subchannel.connection_attempts_failed",
			),
		},
		TraceOptions: oteltracing.TraceOptions{TracerProvider: traceProvider, TextMapPropagator: textMapPropagator},
	})

	go http.ListenAndServe(*prometheusEndpoint, promhttp.Handler())

	cc, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), do)
	if err != nil {
		log.Fatalf("grpc.NewClient() failed: %v", err)
	}
	defer cc.Close()
	c := echo.NewEchoClient(cc)
	ctx := context.Background()

	for {
		r, err := c.UnaryEcho(ctx, &echo.EchoRequest{Message: "this is examples/opentelemetry"})
		if err != nil {
			log.Fatalf("UnaryEcho failed: %v", err)
		}
		fmt.Println(r)
		time.Sleep(time.Second)
	}
}
