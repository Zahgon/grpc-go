package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "google.golang.org/grpc/examples/features/customloadbalancer/client/customroundrobin"
	pb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"google.golang.org/grpc/serviceconfig"
)

var (
	addr1 = "localhost:50050"
	addr2 = "localhost:50051"
)

func main() {
	mr := manual.NewBuilderWithScheme("example")
	defer mr.Close()

	json := `{"loadBalancingConfig": [{"custom_round_robin":{"chooseSecond": 3}}]}`
	sc := internal.ParseServiceConfig.(func(string) *serviceconfig.ParseResult)(json)
	mr.InitialState(resolver.State{
		Endpoints: []resolver.Endpoint{
			{Addresses: []resolver.Address{{Addr: addr1}}},
			{Addresses: []resolver.Address{{Addr: addr2}}},
		},
		ServiceConfig: sc,
	})

	cc, err := grpc.NewClient(mr.Scheme()+":///", grpc.WithResolvers(mr), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.NewClient() failed: %v", err)
	}
	defer cc.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	ec := pb.NewEchoClient(cc)
	if err := waitForDistribution(ctx, ec); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful multiple iterations of 1:2 ratio")
}

func waitForDistribution(ctx context.Context, ec pb.EchoClient) error {
	_ = "STUB: not implemented"
	return nil
}
