package main

import (
	"context"
	"flag"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
	"google.golang.org/grpc/grpclog"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var (
	hsAddr     = flag.String("alts_handshaker_service_address", "", "ALTS handshaker gRPC service address")
	serverAddr = flag.String("server_address", ":8080", "The port on which the server is listening")

	logger = grpclog.Component("interop")
)

func main() {
	flag.Parse()

	opts := alts.DefaultClientOptions()
	if *hsAddr != "" {
		opts.HandshakerServiceAddress = *hsAddr
	}
	altsTC := alts.NewClientCreds(opts)
	conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(altsTC))
	if err != nil {
		logger.Fatalf("grpc.NewClient(%q) = %v", *serverAddr, err)
	}
	defer conn.Close()
	grpcClient := testgrpc.NewTestServiceClient(conn)

	ctx := context.Background()
	request := &testpb.Empty{}

	if _, err := grpcClient.EmptyCall(ctx, request, grpc.WaitForReady(true)); err != nil {
		logger.Fatalf("grpc Client: EmptyCall(_, %v) failed: %v", request, err)
	}
	logger.Info("grpc Client: empty call succeeded")

	time.Sleep(1 * time.Second)
}
