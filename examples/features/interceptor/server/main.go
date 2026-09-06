package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/status"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var (
	port = flag.Int("port", 50051, "the port to serve on")

	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func logger(format string, a ...any) { _ = "STUB: not implemented"; return }

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(_ context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) BidirectionalStreamingEcho(stream pb.Echo_BidirectionalStreamingEchoServer) error {
	_ = "STUB: not implemented"
	return nil
}

func valid(authorization []string) bool { _ = "STUB: not implemented"; return false }

func unaryInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (w *wrappedStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	_ = "STUB: not implemented"
	return *new(grpc.ServerStream)
}

func streamInterceptor(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(unaryInterceptor), grpc.StreamInterceptor(streamInterceptor))

	pb.RegisterEchoServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
