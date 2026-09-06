package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var port = flag.Int("port", 50052, "port number")

type failingServer struct {
	pb.UnimplementedEchoServer
	mu sync.Mutex

	reqCounter uint
	reqModulo  uint
}

func (s *failingServer) maybeFailRequest() error { _ = "STUB: not implemented"; return nil }

func (s *failingServer) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	flag.Parse()

	address := fmt.Sprintf(":%v", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("listen on address", address)

	s := grpc.NewServer()

	failingservice := &failingServer{
		reqCounter: 0,
		reqModulo:  4,
	}

	pb.RegisterEchoServer(s, failingservice)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
