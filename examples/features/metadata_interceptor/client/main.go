package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func callUnaryEcho(ctx context.Context, client pb.EchoClient) { _ = "STUB: not implemented"; return }

func callBidiStreamingEcho(ctx context.Context, client pb.EchoClient) {
	_ = "STUB: not implemented"
	return
}

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.NewClient(%q): %v", *addr, err)
	}
	defer conn.Close()

	ec := pb.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	callUnaryEcho(ctx, ec)

	callBidiStreamingEcho(ctx, ec)
}
