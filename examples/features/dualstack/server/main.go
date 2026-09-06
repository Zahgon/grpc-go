package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	hwpb "google.golang.org/grpc/examples/helloworld/helloworld"
)

type greeterServer struct {
	hwpb.UnimplementedGreeterServer
	addressType string
	address     string
	port        uint32
}

func (s *greeterServer) SayHello(_ context.Context, req *hwpb.HelloRequest) (*hwpb.HelloReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	servers := []*greeterServer{
		{
			addressType: "both IPv4 and IPv6",
			address:     "[::]",
			port:        50051,
		},
		{
			addressType: "IPv4 only",
			address:     "127.0.0.1",
			port:        50052,
		},
		{
			addressType: "IPv6 only",
			address:     "[::1]",
			port:        50053,
		},
	}

	var wg sync.WaitGroup
	for _, server := range servers {
		bindAddr := fmt.Sprintf("%s:%d", server.address, server.port)
		lis, err := net.Listen("tcp", bindAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		hwpb.RegisterGreeterServer(s, server)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.Serve(lis); err != nil {
				log.Panicf("failed to serve: %v", err)
			}
		}()
		log.Printf("serving on %s\n", bindAddr)
	}
	wg.Wait()
}
