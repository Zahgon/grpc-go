package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var port = flag.Int("port", 50051, "the port to serve on")

var errMissingMetadata = status.Errorf(codes.InvalidArgument, "no incoming metadata in rpc context")

type server struct {
	pb.UnimplementedEchoServer
}

func unaryInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *wrappedStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func streamInterceptor(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
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
		log.Fatalf("net.Listen() failed: %v", err)
	}
	fmt.Printf("Server listening at %v\n", lis.Addr())

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor), grpc.StreamInterceptor(streamInterceptor))
	pb.RegisterEchoServer(s, &server{})
	s.Serve(lis)
}
