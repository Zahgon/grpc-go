package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/examples/features/authz/token"
	ecpb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/status"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func callUnaryEcho(ctx context.Context, client ecpb.EchoClient, message string, opts ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

func callBidiStreamingEcho(ctx context.Context, client ecpb.EchoClient, opts ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

func newCredentialsCallOption(t token.Token) grpc.CallOption {
	_ = "STUB: not implemented"
	return *new(grpc.CallOption)
}

func main() {
	flag.Parse()

	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("grpc.NewClient(%q): %v", *addr, err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := ecpb.NewEchoClient(conn)

	authorisedUserTokenCallOption := newCredentialsCallOption(token.Token{Username: "super-user", Secret: "super-secret"})
	if err := callUnaryEcho(ctx, client, "hello world", authorisedUserTokenCallOption); err != nil {
		log.Fatalf("Unary RPC by authorized user failed: %v", err)
	}
	if err := callBidiStreamingEcho(ctx, client, authorisedUserTokenCallOption); err != nil {
		log.Fatalf("Bidirectional RPC by authorized user failed: %v", err)
	}

	unauthorisedUserTokenCallOption := newCredentialsCallOption(token.Token{Username: "bad-actor", Secret: "super-secret"})
	if err := callUnaryEcho(ctx, client, "hello world", unauthorisedUserTokenCallOption); err != nil {
		switch c := status.Code(err); c {
		case codes.PermissionDenied:
			log.Printf("Unary RPC by unauthorized user failed as expected: %v", err)
		default:
			log.Fatalf("Unary RPC by unauthorized user failed unexpectedly: %v, %v", c, err)
		}
	}
	if err := callBidiStreamingEcho(ctx, client, unauthorisedUserTokenCallOption); err != nil {
		switch c := status.Code(err); c {
		case codes.PermissionDenied:
			log.Printf("Bidirectional RPC by unauthorized user failed as expected: %v", err)
		default:
			log.Fatalf("Bidirectional RPC by unauthorized user failed unexpectedly: %v", err)
		}
	}
}
