package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/credentials/insecure"

	v3orcapb "github.com/cncf/xds/go/xds/data/orca/v3"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")
var test = flag.Bool("test", false, "if set, only 1 RPC is performed before exiting")

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"orca_example":{}}]}`),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		func() {

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			if _, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: "test echo message"}); err != nil {
				log.Fatalf("Error from UnaryEcho call: %v", err)
			}
		}()
		if *test {
			return
		}
	}

}

func init() {
	balancer.Register(orcaLBBuilder{})
}

type orcaLBBuilder struct{}

func (orcaLBBuilder) Name() string { _ = "STUB: not implemented"; return "" }
func (orcaLBBuilder) Build(cc balancer.ClientConn, _ balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

type orcaLB struct {
	cc balancer.ClientConn
	sc balancer.SubConn
}

func (o *orcaLB) UpdateClientConnState(ccs balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *orcaLB) ResolverError(error) { _ = "STUB: not implemented"; return }

func (o *orcaLB) ExitIdle() { _ = "STUB: not implemented"; return }

func (o *orcaLB) UpdateSubConnState(balancer.SubConn, balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (o *orcaLB) Close() { _ = "STUB: not implemented"; return }

type picker struct {
	sc balancer.SubConn
}

func (p *picker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

type orcaLis struct{}

func (orcaLis) OnLoadReport(lr *v3orcapb.OrcaLoadReport) { _ = "STUB: not implemented"; return }
