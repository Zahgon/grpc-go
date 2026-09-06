package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var addr = flag.String("addr", "localhost:50052", "the address to connect to")

func unaryCall(c pb.EchoClient, requestID int, message string, want codes.Code) {
	_ = "STUB: not implemented"
	return
}

func streamingCall(c pb.EchoClient, requestID int, message string, want codes.Code) {
	_ = "STUB: not implemented"
	return
}

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	unaryCall(c, 1, "world", codes.OK)

	unaryCall(c, 2, "delay", codes.DeadlineExceeded)

	unaryCall(c, 3, "[propagate me]world", codes.OK)

	unaryCall(c, 4, "[propagate me][propagate me]world", codes.DeadlineExceeded)

	streamingCall(c, 5, "[propagate me]world", codes.OK)

	streamingCall(c, 6, "[propagate me][propagate me]world", codes.DeadlineExceeded)
}
