package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var (
	ports = []string{":10001", ":10002", ":10003"}
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type slowServer struct {
	pb.UnimplementedGreeterServer
}

func (s *slowServer) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()
	service.RegisterChannelzServiceToServer(s)
	go s.Serve(lis)
	defer s.Stop()

	for i := 0; i < 3; i++ {
		lis, err := net.Listen("tcp", ports[i])
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		defer lis.Close()
		s := grpc.NewServer()
		if i == 2 {
			pb.RegisterGreeterServer(s, &slowServer{})
		} else {
			pb.RegisterGreeterServer(s, &server{})
		}
		go s.Serve(lis)
	}

	select {}
}
