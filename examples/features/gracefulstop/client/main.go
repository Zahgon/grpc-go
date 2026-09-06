package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var addr = flag.String("addr", "localhost:50052", "the address to connect to")

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stream, err := c.ClientStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("Error starting stream: %v", err)
	}

	unaryRequests := 0
	for {
		r, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: "Hello"})
		if err != nil {
			log.Printf("Error calling `UnaryEcho`. Server graceful stop initiated: %v", err)
			break
		}
		unaryRequests++
		time.Sleep(200 * time.Millisecond)
		log.Print(r.Message)
	}
	log.Printf("Successful unary requests made by client: %d", unaryRequests)

	r, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error closing stream: %v", err)
	}
	if fmt.Sprintf("%d", unaryRequests) != r.Message {
		log.Fatalf("Got %s successful unary requests processed from server, want: %d", r.Message, unaryRequests)
	}
	log.Printf("Successful unary requests processed by server and made by client are same.")
}
