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
	exampleServiceName = "resolver.example.grpc.io"

	backendAddr = "localhost:50051"
)

func callUnaryEcho(c ecpb.EchoClient, message string) { _ = "STUB: not implemented"; return }

func makeRPCs(cc *grpc.ClientConn, n int) { _ = "STUB: not implemented"; return }

func main() {
	passthroughConn, err := grpc.NewClient(
		fmt.Sprintf("passthrough:///%s", backendAddr),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer passthroughConn.Close()

	fmt.Printf("--- calling helloworld.Greeter/SayHello to \"passthrough:///%s\"\n", backendAddr)
	makeRPCs(passthroughConn, 10)

	fmt.Println()

	exampleConn, err := grpc.NewClient(
		fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer exampleConn.Close()

	fmt.Printf("--- calling helloworld.Greeter/SayHello to \"%s:///%s\"\n", exampleScheme, exampleServiceName)
	makeRPCs(exampleConn, 10)
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
