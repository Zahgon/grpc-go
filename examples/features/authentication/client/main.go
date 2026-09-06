package main

import (
	"flag"
	"log"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/examples/data"
	ecpb "google.golang.org/grpc/examples/features/proto/echo"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func callUnaryEcho(client ecpb.EchoClient, message string) { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()

	perRPC := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(fetchToken())}
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{

		grpc.WithPerRPCCredentials(perRPC),

		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.NewClient(*addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	rgc := ecpb.NewEchoClient(conn)

	callUnaryEcho(rgc, "hello world")
}

func fetchToken() *oauth2.Token { _ = "STUB: not implemented"; return nil }
