package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/internal/grpcsync"
)

var addr = flag.String("addr", "localhost:50052", "the address to connect to")

var payload = string(make([]byte, 8*1024))

func main() {
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	stream, err := c.BidirectionalStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}
	log.Printf("New stream began.")

	stopSending := grpcsync.NewEvent()
	sentOne := make(chan struct{})
	go func() {
		i := 0
		for !stopSending.HasFired() {
			i++
			if err := stream.Send(&pb.EchoRequest{Message: payload}); err != nil {
				log.Fatalf("Error sending data: %v", err)
			}
			sentOne <- struct{}{}
		}
		log.Printf("Sent %v messages.", i)
		stream.CloseSend()
	}()

	for !stopSending.HasFired() {
		after := time.NewTimer(time.Second)
		select {
		case <-sentOne:
			after.Stop()
		case <-after.C:
			log.Printf("Sending is blocked.")
			stopSending.Fire()
			<-sentOne
		}
	}

	time.Sleep(2 * time.Second)

	for i := 0; true; i++ {
		if _, err := stream.Recv(); err != nil {
			log.Printf("Read %v messages.", i)
			if err == io.EOF {
				log.Printf("Stream ended successfully.")
				return
			}
			log.Fatalf("Error receiving data: %v", err)
		}
	}
}
