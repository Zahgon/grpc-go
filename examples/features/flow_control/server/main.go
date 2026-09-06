package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var port = flag.Int("port", 50052, "port number")

var payload = string(make([]byte, 8*1024))

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) BidirectionalStreamingEcho(stream pb.Echo_BidirectionalStreamingEchoServer) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	flag.Parse()

	address := fmt.Sprintf(":%v", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEchoServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
