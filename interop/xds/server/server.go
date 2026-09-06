package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/admin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/stats/opentelemetry"
	"google.golang.org/grpc/stats/opentelemetry/csm"
	"google.golang.org/grpc/xds"

	xdscreds "google.golang.org/grpc/credentials/xds"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

var (
	port                   = flag.Int("port", 8080, "Listening port for test service")
	maintenancePort        = flag.Int("maintenance_port", 8081, "Listening port for maintenance services like health, reflection, channelz etc when -secure_mode is true. When -secure_mode is false, all these services will be registered on -port")
	serverID               = flag.String("server_id", "go_server", "Server ID included in response")
	secureMode             = flag.Bool("secure_mode", false, "If true, retrieve security configuration from the management server. Else, use insecure credentials.")
	hostNameOverride       = flag.String("host_name_override", "", "If set, use this as the hostname instead of the real hostname")
	enableCSMObservability = flag.Bool("enable_csm_observability", false, "Whether to enable CSM Observability")

	logger = grpclog.Component("interop")
)

const (
	rpcBehaviorMDKey             = "rpc-behavior"
	grpcPreviousRPCAttemptsMDKey = "grpc-previous-rpc-attempts"
	sleepPfx                     = "sleep-"
	keepOpenVal                  = "keep-open"
	errorCodePfx                 = "error-code-"
	succeedOnRetryPfx            = "succeed-on-retry-attempt-"
	hostnamePfx                  = "hostname="
)

func getHostname() string { _ = "STUB: not implemented"; return "" }

type testServiceImpl struct {
	testgrpc.UnimplementedTestServiceServer
	hostname string
	serverID string
}

func (s *testServiceImpl) EmptyCall(ctx context.Context, _ *testpb.Empty) (*testpb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testServiceImpl) UnaryCall(ctx context.Context, in *testpb.SimpleRequest) (*testpb.SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getRPCBehaviorMetadata(ctx context.Context) []string { _ = "STUB: not implemented"; return nil }

func getMetadataValues(ctx context.Context, metadataKey string) []string {
	_ = "STUB: not implemented"
	return nil
}

type xdsUpdateHealthServiceImpl struct {
	testgrpc.UnimplementedXdsUpdateHealthServiceServer
	healthServer *health.Server
}

func (x *xdsUpdateHealthServiceImpl) SetServing(_ context.Context, _ *testpb.Empty) (*testpb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *xdsUpdateHealthServiceImpl) SetNotServing(_ context.Context, _ *testpb.Empty) (*testpb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func xdsServingModeCallback(addr net.Addr, args xds.ServingModeChangeArgs) {
	_ = "STUB: not implemented"
	return
}

func main() {
	flag.Parse()
	if *enableCSMObservability {
		exporter, err := prometheus.New()
		if err != nil {
			logger.Fatalf("Failed to start prometheus exporter: %v", err)
		}
		var addr string
		var ok bool
		if addr, ok = os.LookupEnv("OTEL_EXPORTER_PROMETHEUS_HOST"); !ok {
			addr = ""
		}
		var port string
		if port, ok = os.LookupEnv("OTEL_EXPORTER_PROMETHEUS_PORT"); !ok {
			port = "9464"
		}
		go func() {
			if err := http.ListenAndServe(addr+":"+port, promhttp.Handler()); err != nil {
				logger.Fatalf("error listening: %v", err)
			}
		}()

		provider := metric.NewMeterProvider(
			metric.WithReader(exporter),
		)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		cleanup := csm.EnableObservability(ctx, opentelemetry.Options{MetricsOptions: opentelemetry.MetricsOptions{MeterProvider: provider}})
		defer cleanup()
	}

	if *secureMode && *port == *maintenancePort {
		logger.Fatal("-port and -maintenance_port must be different when -secure_mode is set")
	}

	testService := &testServiceImpl{hostname: getHostname(), serverID: *serverID}
	healthServer := health.NewServer()
	updateHealthService := &xdsUpdateHealthServiceImpl{healthServer: healthServer}

	if !*secureMode {
		addr := fmt.Sprintf(":%d", *port)
		lis, err := net.Listen("tcp4", addr)
		if err != nil {
			logger.Fatalf("net.Listen(%s) failed: %v", addr, err)
		}

		server := grpc.NewServer()
		testgrpc.RegisterTestServiceServer(server, testService)
		healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
		healthgrpc.RegisterHealthServer(server, healthServer)
		testgrpc.RegisterXdsUpdateHealthServiceServer(server, updateHealthService)
		reflection.Register(server)
		cleanup, err := admin.Register(server)
		if err != nil {
			logger.Fatalf("Failed to register admin services: %v", err)
		}
		defer cleanup()
		if err := server.Serve(lis); err != nil {
			logger.Errorf("Serve() failed: %v", err)
		}
		return
	}

	addr := fmt.Sprintf(":%d", *port)
	testLis, err := net.Listen("tcp4", addr)
	if err != nil {
		logger.Fatalf("net.Listen(%s) failed: %v", addr, err)
	}

	creds, err := xdscreds.NewServerCredentials(xdscreds.ServerOptions{FallbackCreds: insecure.NewCredentials()})
	if err != nil {
		logger.Fatalf("Failed to create xDS credentials: %v", err)
	}

	testServer, err := xds.NewGRPCServer(grpc.Creds(creds), xds.ServingModeCallback(xdsServingModeCallback))
	if err != nil {
		logger.Fatal("Failed to create an xDS enabled gRPC server: %v", err)
	}
	testgrpc.RegisterTestServiceServer(testServer, testService)
	go func() {
		if err := testServer.Serve(testLis); err != nil {
			logger.Errorf("test server Serve() failed: %v", err)
		}
	}()
	defer testServer.Stop()

	addr = fmt.Sprintf(":%d", *maintenancePort)
	maintenanceLis, err := net.Listen("tcp4", addr)
	if err != nil {
		logger.Fatalf("net.Listen(%s) failed: %v", addr, err)
	}

	maintenanceServer := grpc.NewServer()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthgrpc.RegisterHealthServer(maintenanceServer, healthServer)
	testgrpc.RegisterXdsUpdateHealthServiceServer(maintenanceServer, updateHealthService)
	reflection.Register(maintenanceServer)
	cleanup, err := admin.Register(maintenanceServer)
	if err != nil {
		logger.Fatalf("Failed to register admin services: %v", err)
	}
	defer cleanup()
	if err := maintenanceServer.Serve(maintenanceLis); err != nil {
		logger.Errorf("maintenance server Serve() failed: %v", err)
	}
}
