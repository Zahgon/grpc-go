package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	otelstdouttrace "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	otelpropagation "go.opentelemetry.io/otel/propagation"
	otelmetric "go.opentelemetry.io/otel/sdk/metric"
	otelresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/features/proto/echo"
	oteltracing "google.golang.org/grpc/experimental/opentelemetry"
	"google.golang.org/grpc/stats/opentelemetry"
)

var (
	addr               = flag.String("addr", "localhost:50051", "the server address to connect to")
	prometheusEndpoint = flag.String("prometheus_endpoint", ":9464", "the Prometheus exporter endpoint for metrics")
)

type echoServer struct {
	pb.UnimplementedEchoServer
	addr string
}

func (s *echoServer) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

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
	traceProvider := sdktrace.NewTracerProvider(sdktrace.WithBatcher(traceExporter), sdktrace.WithResource(otelresource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceName("grpc-server"))))

	textMapPropagator := otelpropagation.TraceContext{}
	so := opentelemetry.ServerOption(opentelemetry.Options{
		MetricsOptions: opentelemetry.MetricsOptions{
			MeterProvider: meterProvider,
		},
		TraceOptions: oteltracing.TraceOptions{TracerProvider: traceProvider, TextMapPropagator: textMapPropagator}})

	go http.ListenAndServe(*prometheusEndpoint, promhttp.Handler())

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer(so)
	pb.RegisterEchoServer(s, &echoServer{addr: *addr})

	log.Printf("Serving on %s\n", *addr)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
