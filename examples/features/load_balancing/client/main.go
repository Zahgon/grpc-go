package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	ecpb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/resolver"
)

const (
	exampleScheme      = "example"
	exampleServiceName = "lb.example.grpc.io"
)

var addrs = []string{"localhost:50051", "localhost:50052"}

func callUnaryEcho(c ecpb.EchoClient, message string) { _ = "STUB: not implemented"; return }

func makeRPCs(cc *grpc.ClientConn, n int) { _ = "STUB: not implemented"; return }

func main() {

	pickfirstConn, err := grpc.NewClient(
		fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer pickfirstConn.Close()

	fmt.Println("--- calling helloworld.Greeter/SayHello with pick_first ---")
	makeRPCs(pickfirstConn, 10)

	fmt.Println()

	roundrobinConn, err := grpc.NewClient(
		fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer roundrobinConn.Close()

	fmt.Println("--- calling helloworld.Greeter/SayHello with round_robin ---")
	makeRPCs(roundrobinConn, 10)
}

type exampleResolverBuilder struct{}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Resolver), nil
}

func (*exampleResolverBuilder) Scheme() string { _ = "STUB: not implemented"; return "" }

type exampleResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *exampleResolver) start() { _ = "STUB: not implemented"; return }

func (*exampleResolver) ResolveNow(resolver.ResolveNowOptions) { _ = "STUB: not implemented"; return }
func (*exampleResolver) Close()                                { _ = "STUB: not implemented"; return }

func init() {
	resolver.Register(&exampleResolverBuilder{})
}
