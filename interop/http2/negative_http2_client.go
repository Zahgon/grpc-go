package main

import (
	"context"
	"flag"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var (
	serverHost = flag.String("server_host", "localhost", "The server host name")
	serverPort = flag.Int("server_port", 8080, "The server port number")
	testCase   = flag.String("test_case", "goaway",
		`Configure different test cases. Valid options are:
        goaway : client sends two requests, the server will send a goaway in between;
        rst_after_header : server will send rst_stream after it sends headers;
        rst_during_data : server will send rst_stream while sending data;
        rst_after_data : server will send rst_stream after sending data;
        ping : server will send pings between each http2 frame;
        max_streams : server will ensure that the max_concurrent_streams limit is upheld;`)
	largeReqSize  = 271828
	largeRespSize = 314159

	logger = grpclog.Component("interop")
)

func largeSimpleRequest() *testpb.SimpleRequest { _ = "STUB: not implemented"; return nil }

func goaway(ctx context.Context, tc testgrpc.TestServiceClient) { _ = "STUB: not implemented"; return }

func rstAfterHeader(tc testgrpc.TestServiceClient) { _ = "STUB: not implemented"; return }

func rstDuringData(tc testgrpc.TestServiceClient) { _ = "STUB: not implemented"; return }

func rstAfterData(tc testgrpc.TestServiceClient) { _ = "STUB: not implemented"; return }

func ping(ctx context.Context, tc testgrpc.TestServiceClient) { _ = "STUB: not implemented"; return }

func maxStreams(ctx context.Context, tc testgrpc.TestServiceClient) {
	_ = "STUB: not implemented"
	return
}

func main() {
	flag.Parse()
	serverAddr := net.JoinHostPort(*serverHost, strconv.Itoa(*serverPort))
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		logger.Fatalf("grpc.NewClient(%q) = %v", serverAddr, err)
	}
	defer conn.Close()
	tc := testgrpc.NewTestServiceClient(conn)
	ctx := context.Background()
	switch *testCase {
	case "goaway":
		goaway(ctx, tc)
		logger.Infoln("goaway done")
	case "rst_after_header":
		rstAfterHeader(tc)
		logger.Infoln("rst_after_header done")
	case "rst_during_data":
		rstDuringData(tc)
		logger.Infoln("rst_during_data done")
	case "rst_after_data":
		rstAfterData(tc)
		logger.Infoln("rst_after_data done")
	case "ping":
		ping(ctx, tc)
		logger.Infoln("ping done")
	case "max_streams":
		maxStreams(ctx, tc)
		logger.Infoln("max_streams done")
	default:
		logger.Fatal("Unsupported test case: ", *testCase)
	}
}
