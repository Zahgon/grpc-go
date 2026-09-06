package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	ecpb "google.golang.org/grpc/examples/features/proto/echo"
	hwpb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var port = flag.Int("port", 50051, "the port to serve on")

type hwServer struct {
	hwpb.UnimplementedGreeterServer
}

func (s *hwServer) SayHello(_ context.Context, in *hwpb.HelloRequest) (*hwpb.HelloReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ecServer struct {
	ecpb.UnimplementedEchoServer
}

func (s *ecServer) UnaryEcho(_ context.Context, req *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer()

	hwpb.RegisterGreeterServer(s, &hwServer{})

	ecpb.RegisterEchoServer(s, &ecServer{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
