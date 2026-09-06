package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var (
	port = flag.Int("port", 50052, "port number")
)

type server struct {
	pb.UnimplementedEchoServer

	unaryRequests atomic.Int32
	streamStart   chan struct{}
}

func (s *server) ClientStreamingEcho(stream pb.Echo_ClientStreamingEchoServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	flag.Parse()

	address := fmt.Sprintf(":%v", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ss := &server{streamStart: make(chan struct{})}
	pb.RegisterEchoServer(s, ss)
	serverStopped := make(chan struct{}, 1)

	go func() {
		<-ss.streamStart
		time.Sleep(1 * time.Second)
		log.Println("Initiating graceful shutdown...")
		timer := time.AfterFunc(10*time.Second, func() {
			log.Println("Server couldn't stop gracefully in time. Doing force stop.")
			s.Stop()

			select {
			case serverStopped <- struct{}{}:
			default:
			}
		})
		defer timer.Stop()
		s.GracefulStop()
		log.Println("Server stopped gracefully.")

		select {
		case serverStopped <- struct{}{}:
		default:
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	<-serverStopped
}
