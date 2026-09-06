//go:build linux
// +build linux

package main

import (
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	ecpb "google.golang.org/grpc/examples/features/proto/echo"
)

var (
	addr = flag.String("addr", "abstract-unix-socket", "The unix abstract socket address")
)

func callUnaryEcho(c ecpb.EchoClient, message string) { _ = "STUB: not implemented"; return }

func makeRPCs(cc *grpc.ClientConn, n int) { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()
	sockAddr := fmt.Sprintf("unix-abstract:%v", *addr)
	cc, err := grpc.NewClient(sockAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.NewClient(%q) failed: %v", sockAddr, err)
	}
	defer cc.Close()

	fmt.Printf("--- calling echo.Echo/UnaryEcho to %s\n", sockAddr)
	makeRPCs(cc, 10)
	fmt.Println()
}
