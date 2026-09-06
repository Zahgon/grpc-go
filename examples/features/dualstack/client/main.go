package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	hwpb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/resolver"
)

const (
	port1 = 50051
	port2 = 50052
	port3 = 50053
)

func init() {
	resolver.Register(&exampleResolver{})
}

type exampleResolver struct{}

func (*exampleResolver) Close() { _ = "STUB: not implemented"; return }

func (*exampleResolver) ResolveNow(resolver.ResolveNowOptions) { _ = "STUB: not implemented"; return }

func (*exampleResolver) Build(_ resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Resolver), nil
}

func (*exampleResolver) Scheme() string { _ = "STUB: not implemented"; return "" }

func main() {

	log.Print("**** Use default DNS resolver ****")
	target := fmt.Sprintf("localhost:%d", port1)
	cc, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer cc.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client := hwpb.NewGreeterClient(cc)

	for i := 0; i < 5; i++ {
		resp, err := client.SayHello(ctx, &hwpb.HelloRequest{
			Name: fmt.Sprintf("request:%d", i),
		})
		if err != nil {
			log.Panicf("RPC failed: %v", err)
		}
		log.Print("Greeting:", resp.GetMessage())
	}
	cc.Close()

	log.Print("**** Change to use example name resolver ****")
	dOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	}
	cc, err = grpc.NewClient("example:///ignored", dOpts...)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	client = hwpb.NewGreeterClient(cc)

	if err := waitForDistribution(ctx, client); err != nil {
		log.Panic(err)
	}
	log.Print("Successful multiple iterations of 1:1:1 ratio")
}

func waitForDistribution(ctx context.Context, client hwpb.GreeterClient) error {
	_ = "STUB: not implemented"
	return nil
}
