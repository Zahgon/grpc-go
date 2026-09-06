package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var port = flag.Int("port", 50052, "port number")

type server struct {
	pb.UnimplementedGreeterServer
	mu    sync.Mutex
	count map[string]int
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
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

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{count: make(map[string]int)})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
