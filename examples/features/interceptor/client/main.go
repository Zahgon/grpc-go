package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	ecpb "google.golang.org/grpc/examples/features/proto/echo"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

const fallbackToken = "some-secret-token"

func logger(format string, a ...any) { _ = "STUB: not implemented"; return }

func unaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (w *wrappedStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream)
}

func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientStream), nil
}

func callUnaryEcho(client ecpb.EchoClient, message string) { _ = "STUB: not implemented"; return }

func callBidiStreamingEcho(client ecpb.EchoClient) { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()

	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(creds), grpc.WithUnaryInterceptor(unaryInterceptor), grpc.WithStreamInterceptor(streamInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	rgc := ecpb.NewEchoClient(conn)
	callUnaryEcho(rgc, "hello world")
	callBidiStreamingEcho(rgc)
}
