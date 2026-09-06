package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	xdscreds "google.golang.org/grpc/credentials/xds"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/stats/opentelemetry"
	"google.golang.org/grpc/stats/opentelemetry/csm"
	"google.golang.org/grpc/xds"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

var (
	port               = flag.String("port", "50051", "the server address to connect to")
	prometheusEndpoint = flag.String("prometheus_endpoint", ":9464", "the Prometheus exporter endpoint")
)

type server struct {
	pb.UnimplementedGreeterServer
	addr string
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	flag.Parse()
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("Failed to start prometheus exporter: %v", err)
	}
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	go http.ListenAndServe(*prometheusEndpoint, promhttp.Handler())

	cleanup := csm.EnableObservability(context.Background(), opentelemetry.Options{MetricsOptions: opentelemetry.MetricsOptions{MeterProvider: provider}})
	defer cleanup()

	lis, err := net.Listen("tcp4", "0.0.0.0:"+*port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	creds, err := xdscreds.NewServerCredentials(xdscreds.ServerOptions{FallbackCreds: insecure.NewCredentials()})
	if err != nil {
		log.Fatalf("Failed to create xDS credentials: %v", err)
	}
	s, err := xds.NewGRPCServer(grpc.Creds(creds))
	if err != nil {
		log.Fatalf("Failed to start xDS Server: %v", err)
	}
	pb.RegisterGreeterServer(s, &server{addr: ":" + *port})

	log.Printf("Serving on %s\n", *port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
