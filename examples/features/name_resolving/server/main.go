package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

const addr = "localhost:50051"

type ecServer struct {
	pb.UnimplementedEchoServer
	addr string
}

func (s *ecServer) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &ecServer{addr: addr})
	log.Printf("serving on %s\n", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
