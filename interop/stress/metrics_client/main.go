package main

import (
	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	metricspb "google.golang.org/grpc/interop/stress/grpc_testing"
)

var (
	metricsServerAddress = flag.String("metrics_server_address", "", "The metrics server addresses in the format <hostname>:<port>")
	totalOnly            = flag.Bool("total_only", false, "If true, this prints only the total value of all gauges")

	logger = grpclog.Component("stress")
)

func printMetrics(client metricspb.MetricsServiceClient, totalOnly bool) {
	_ = "STUB: not implemented"
	return
}

func main() {
	flag.Parse()
	if *metricsServerAddress == "" {
		logger.Fatal("-metrics_server_address is unset")
	}

	conn, err := grpc.NewClient(*metricsServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalf("grpc.NewClient(%q) = %v", metricsServerAddress, err)
	}
	defer conn.Close()

	c := metricspb.NewMetricsServiceClient(conn)
	printMetrics(c, *totalOnly)
}
