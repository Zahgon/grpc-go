package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
	ecpb "google.golang.org/grpc/examples/features/proto/echo"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func callUnaryEcho(client ecpb.EchoClient, message string) { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()

	altsTC := alts.NewClientCreds(alts.DefaultClientOptions())

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(altsTC))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	rgc := ecpb.NewEchoClient(conn)
	callUnaryEcho(rgc, "hello world")
}
