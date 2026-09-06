package main

import (
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

const (
	timestampFormat = time.StampNano
	streamingCount  = 10
)

func unaryCallWithMetadata(c pb.EchoClient, message string) { _ = "STUB: not implemented"; return }

func serverStreamingWithMetadata(c pb.EchoClient, message string) {
	_ = "STUB: not implemented"
	return
}

func clientStreamWithMetadata(c pb.EchoClient, message string) { _ = "STUB: not implemented"; return }

func bidirectionalWithMetadata(c pb.EchoClient, message string) { _ = "STUB: not implemented"; return }

const message = "this is examples/metadata"

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	unaryCallWithMetadata(c, message)
	time.Sleep(1 * time.Second)

	serverStreamingWithMetadata(c, message)
	time.Sleep(1 * time.Second)

	clientStreamWithMetadata(c, message)
	time.Sleep(1 * time.Second)

	bidirectionalWithMetadata(c, message)
}
