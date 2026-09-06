package main

import (
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	ecpb "google.golang.org/grpc/examples/features/proto/echo"
	hwpb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func callSayHello(c hwpb.GreeterClient, name string) { _ = "STUB: not implemented"; return }

func callUnaryEcho(client ecpb.EchoClient, message string) { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	fmt.Println("--- calling helloworld.Greeter/SayHello ---")

	hwc := hwpb.NewGreeterClient(conn)
	callSayHello(hwc, "multiplex")

	fmt.Println()
	fmt.Println("--- calling routeguide.RouteGuide/GetFeature ---")

	rgc := ecpb.NewEchoClient(conn)
	callUnaryEcho(rgc, "this is examples/multiplex")
}
