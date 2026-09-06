package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/grpclb"
	_ "google.golang.org/grpc/xds/googledirectpath"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var (
	customCredentialsType   = flag.String("custom_credentials_type", "", "Client creds to use")
	serverURI               = flag.String("server_uri", "dns:///staging-grpc-directpath-fallback-test.googleapis.com:443", "The server host name")
	induceFallbackCmd       = flag.String("induce_fallback_cmd", "", "Command to induce fallback e.g. by making certain addresses unroutable")
	fallbackDeadlineSeconds = flag.Int("fallback_deadline_seconds", 1, "How long to wait for fallback to happen after induce_fallback_cmd")
	testCase                = flag.String("test_case", "",
		`Configure different test cases. Valid options are:
        fallback_before_startup : LB/backend connections fail before RPC's have been made;
        fallback_after_startup : LB/backend connections fail after RPC's have been made;`)
	infoLog  = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func doRPCAndGetPath(client testgrpc.TestServiceClient, timeout time.Duration) testpb.GrpclbRouteType {
	_ = "STUB: not implemented"
	return *new(testpb.GrpclbRouteType)
}

func dialTCPUserTimeout(ctx context.Context, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func createTestConn() *grpc.ClientConn { _ = "STUB: not implemented"; return nil }

func runCmd(command string) { _ = "STUB: not implemented"; return }

func waitForFallbackAndDoRPCs(client testgrpc.TestServiceClient, fallbackDeadline time.Time) {
	_ = "STUB: not implemented"
	return
}

func doFallbackBeforeStartup() { _ = "STUB: not implemented"; return }

func doFallbackAfterStartup() { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()
	if len(*induceFallbackCmd) == 0 {
		errorLog.Fatalf("--induce_fallback_cmd unset")
	}
	switch *testCase {
	case "fallback_before_startup":
		doFallbackBeforeStartup()
		log.Printf("FallbackBeforeStartup done!\n")
	case "fallback_after_startup":
		doFallbackAfterStartup()
		log.Printf("FallbackAfterStartup done!\n")
	default:
		errorLog.Fatalf("Unsupported test case: %v", *testCase)
	}
}
