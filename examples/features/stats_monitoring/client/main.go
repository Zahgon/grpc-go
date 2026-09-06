package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	echogrpc "google.golang.org/grpc/examples/features/proto/echo"
	echopb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/examples/features/stats_monitoring/statshandler"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	flag.Parse()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(statshandler.New()),
	}
	conn, err := grpc.NewClient(*addr, opts...)
	if err != nil {
		log.Fatalf("failed to connect to server %q: %v", *addr, err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := echogrpc.NewEchoClient(conn)

	resp, err := c.UnaryEcho(ctx, &echopb.EchoRequest{Message: "stats handler demo"})
	if err != nil {
		log.Fatalf("unexpected error from UnaryEcho: %v", err)
	}
	log.Printf("RPC response: %s", resp.Message)
}
