package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/features/proto/echo"
	_ "google.golang.org/grpc/health"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

const serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": ""
	}
}`

func callUnaryEcho(c pb.EchoClient) { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()

	r := manual.NewBuilderWithScheme("whatever")
	r.InitialState(resolver.State{
		Addresses: []resolver.Address{
			{Addr: "localhost:50051"},
			{Addr: "localhost:50052"},
		},
	})

	address := fmt.Sprintf("%s:///unused", r.Scheme())

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(r),

		grpc.WithDefaultServiceConfig(serviceConfig),
	}

	conn, err := grpc.NewClient(address, options...)
	if err != nil {
		log.Fatalf("grpc.NewClient(%q): %v", address, err)
	}
	defer conn.Close()

	echoClient := pb.NewEchoClient(conn)

	for {
		callUnaryEcho(echoClient)
		time.Sleep(time.Second)
	}
}
