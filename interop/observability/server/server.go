package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/gcp/observability"
	"google.golang.org/grpc/interop"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	err := observability.Start(context.Background())
	if err != nil {
		log.Fatalf("observability start failed: %v", err)
	}
	defer observability.End()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	defer server.Stop()
	testgrpc.RegisterTestServiceServer(server, interop.NewTestServer())
	log.Printf("Observability interop server listening on %v", lis.Addr())
	server.Serve(lis)
}
