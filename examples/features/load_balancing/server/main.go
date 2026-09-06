package main

import (
	"context"
	"sync"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var (
	addrs = []string{":50051", ":50052"}
)

type ecServer struct {
	pb.UnimplementedEchoServer
	addr string
}

func (s *ecServer) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func startServer(addr string) { _ = "STUB: not implemented"; return }

func main() {
	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			startServer(addr)
		}(addr)
	}
	wg.Wait()
}
