package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/authz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/status"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

const (
	unaryEchoWriterRole      = "UNARY_ECHO:W"
	streamEchoReadWriterRole = "STREAM_ECHO:RW"
	authzPolicy              = `
	{
		"name": "authz",
		"allow_rules": [
			{
				"name": "allow_UnaryEcho",
				"request": {
					"paths": ["/grpc.examples.echo.Echo/UnaryEcho"],
					"headers": [
						{
							"key": "UNARY_ECHO:W",
							"values": ["true"]
						}
					]
				}
			},
			{
				"name": "allow_BidirectionalStreamingEcho",
				"request": {
					"paths": ["/grpc.examples.echo.Echo/BidirectionalStreamingEcho"],
					"headers": [
						{
							"key": "STREAM_ECHO:RW",
							"values": ["true"]
						}
					]
				}
			}
		],
		"deny_rules": []
	}
	`
	authzOptStatic      = "static"
	authzOptFileWatcher = "filewatcher"
)

var (
	port     = flag.Int("port", 50051, "the port to serve on")
	authzOpt = flag.String("authz-option", authzOptStatic, "the authz option (static or filewatcher)")

	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
)

func newContextWithRoles(ctx context.Context, username string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

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

func isAuthenticated(authorization []string) (username string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func authUnaryInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func newWrappedStream(ctx context.Context, s grpc.ServerStream) grpc.ServerStream {
	_ = "STUB: not implemented"
	return *new(grpc.ServerStream)
}

func authStreamInterceptor(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	flag.Parse()

	if *authzOpt != authzOptStatic && *authzOpt != authzOptFileWatcher {
		log.Fatalf("Invalid authz option: %s", *authzOpt)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Listening on local port %q: %v", *port, err)
	}

	creds, err := credentials.NewServerTLSFromFile(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("Loading credentials: %v", err)
	}

	var unaryAuthzInterceptor grpc.UnaryServerInterceptor
	var streamAuthzInterceptor grpc.StreamServerInterceptor
	if *authzOpt == authzOptStatic {

		staticInterceptor, err := authz.NewStatic(authzPolicy)
		if err != nil {
			log.Fatalf("Creating a static authz interceptor: %v", err)
		}
		unaryAuthzInterceptor, streamAuthzInterceptor = staticInterceptor.UnaryInterceptor, staticInterceptor.StreamInterceptor
	} else if *authzOpt == authzOptFileWatcher {

		fileWatcherInterceptor, err := authz.NewFileWatcher(data.Path("rbac/policy.json"), 100*time.Millisecond)
		if err != nil {
			log.Fatalf("Creating a file watcher authz interceptor: %v", err)
		}
		unaryAuthzInterceptor, streamAuthzInterceptor = fileWatcherInterceptor.UnaryInterceptor, fileWatcherInterceptor.StreamInterceptor
	}

	unaryInterceptors := grpc.ChainUnaryInterceptor(authUnaryInterceptor, unaryAuthzInterceptor)
	streamInterceptors := grpc.ChainStreamInterceptor(authStreamInterceptor, streamAuthzInterceptor)
	s := grpc.NewServer(grpc.Creds(creds), unaryInterceptors, streamInterceptors)

	pb.RegisterEchoServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Serving Echo service on local port: %v", err)
	}
}
