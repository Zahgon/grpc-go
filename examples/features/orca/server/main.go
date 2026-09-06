package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/orca"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var port = flag.Int("port", 50051, "the port to serve on")

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("Server listening at %v\n", lis.Addr())

	s := grpc.NewServer(orca.CallMetricsServerOption(nil))
	pb.RegisterEchoServer(s, &server{})

	smr := orca.NewServerMetricsRecorder()
	opts := orca.ServiceOptions{
		MinReportingInterval:  3 * time.Second,
		ServerMetricsProvider: smr,
	}
	internal.ORCAAllowAnyMinReportingInterval.(func(so *orca.ServiceOptions))(&opts)
	if err := orca.Register(s, opts); err != nil {
		log.Fatalf("Failed to register ORCA service: %v", err)
	}

	go func() {
		for {
			smr.SetCPUUtilization(.5)
			time.Sleep(2 * time.Second)
			smr.SetCPUUtilization(.9)
			time.Sleep(2 * time.Second)
		}
	}()

	s.Serve(lis)
}
