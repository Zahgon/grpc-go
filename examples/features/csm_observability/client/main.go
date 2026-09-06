package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	xdscreds "google.golang.org/grpc/credentials/xds"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/stats/opentelemetry"
	"google.golang.org/grpc/stats/opentelemetry/csm"
	_ "google.golang.org/grpc/xds"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

const defaultName = "world"

var (
	target             = flag.String("target", "xds:///helloworld:50051", "the server address to connect to")
	prometheusEndpoint = flag.String("prometheus_endpoint", ":9464", "the Prometheus exporter endpoint")
	name               = flag.String("name", defaultName, "Name to greet")
)

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

	creds, err := xdscreds.NewClientCredentials(xdscreds.ClientOptions{FallbackCreds: insecure.NewCredentials()})
	if err != nil {
		log.Fatalf("Failed to create xDS credentials: %v", err)
	}
	cc, err := grpc.NewClient(*target, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to start NewClient: %v", err)
	}
	defer cc.Close()
	c := pb.NewGreeterClient(cc)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Fatalf("Could not greet: %v", err)
		}
		fmt.Println(r)
		time.Sleep(time.Second)
		cancel()
	}
}
