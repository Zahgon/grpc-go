package main

import (
	"flag"
	"sync"
	"time"

	"google.golang.org/grpc/internal/syscall"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var (
	certFile = flag.String("tls_cert_file", "", "The TLS cert file")
	keyFile  = flag.String("tls_key_file", "", "The TLS key file")
)

type benchmarkServer struct {
	port      int
	cores     int
	closeFunc func()

	mu              sync.Mutex
	lastResetTime   time.Time
	rusageLastReset *syscall.Rusage
}

func printServerConfig(config *testpb.ServerConfig) { _ = "STUB: not implemented"; return }

func startBenchmarkServer(config *testpb.ServerConfig, serverPort int) (*benchmarkServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bs *benchmarkServer) getStats(reset bool) *testpb.ServerStats {
	_ = "STUB: not implemented"
	return nil
}
