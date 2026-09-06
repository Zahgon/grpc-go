package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var port = flag.Int("port", 50051, "the port to serve on")

const (
	timestampFormat = time.StampNano
	streamingCount  = 10
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) ServerStreamingEcho(in *pb.EchoRequest, stream pb.Echo_ServerStreamingEchoServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) ClientStreamingEcho(stream pb.Echo_ClientStreamingEchoServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) BidirectionalStreamingEcho(stream pb.Echo_BidirectionalStreamingEchoServer) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	s.Serve(lis)
}
