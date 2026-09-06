package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	echogrpc "google.golang.org/grpc/examples/features/proto/echo"
	echopb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/examples/features/stats_monitoring/statshandler"
)

var port = flag.Int("port", 50051, "the port to serve on")

type server struct {
	echogrpc.UnimplementedEchoServer
}

func (s *server) UnaryEcho(_ context.Context, req *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", *port, err)
	}
	log.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer(grpc.StatsHandler(statshandler.New()))
	echogrpc.RegisterEchoServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
