package main

import (
	"context"
	"crypto/tls"
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
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

var port = flag.Int("port", 50051, "the port to serve on")

func main() {
	flag.Parse()
	fmt.Printf("server starting on port %d...\n", *port)

	cert, err := tls.LoadX509KeyPair(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}
	opts := []grpc.ServerOption{

		grpc.UnaryInterceptor(ensureValidToken),

		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterEchoServer(s, &ecServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type ecServer struct {
	pb.UnimplementedEchoServer
}

func (s *ecServer) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func valid(authorization []string) bool { _ = "STUB: not implemented"; return false }

func ensureValidToken(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
