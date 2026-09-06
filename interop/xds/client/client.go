package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/admin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/xds"
	"google.golang.org/grpc/grpclog"
	_ "google.golang.org/grpc/interop/xds"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/stats/opentelemetry"
	"google.golang.org/grpc/stats/opentelemetry/csm"
	_ "google.golang.org/grpc/xds"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

func init() {
	rpcCfgs.Store([]*rpcConfig{{typ: unaryCall}})
}

type statsWatcherKey struct {
	startID int32
	endID   int32
}

type rpcInfo struct {
	typ      string
	hostname string
}

type statsWatcher struct {
	rpcsByPeer    map[string]int32
	rpcsByType    map[string]map[string]int32
	numFailures   int32
	remainingRPCs int32
	chanHosts     chan *rpcInfo
}

func (watcher *statsWatcher) buildResp() *testpb.LoadBalancerStatsResponse {
	_ = "STUB: not implemented"
	return nil
}

type accumulatedStats struct {
	mu                       sync.Mutex
	numRPCsStartedByMethod   map[string]int32
	numRPCsSucceededByMethod map[string]int32
	numRPCsFailedByMethod    map[string]int32
	rpcStatusByMethod        map[string]map[int32]int32
}

func convertRPCName(in string) string { _ = "STUB: not implemented"; return "" }

func copyStatsMap(originalMap map[string]int32) map[string]int32 {
	_ = "STUB: not implemented"
	return nil
}

func copyStatsIntMap(originalMap map[int32]int32) map[int32]int32 {
	_ = "STUB: not implemented"
	return nil
}

func (as *accumulatedStats) makeStatsMap() map[string]*testpb.LoadBalancerAccumulatedStatsResponse_MethodStats {
	_ = "STUB: not implemented"
	return nil
}

func (as *accumulatedStats) buildResp() *testpb.LoadBalancerAccumulatedStatsResponse {
	_ = "STUB: not implemented"
	return nil
}

func (as *accumulatedStats) startRPC(rpcType string) { _ = "STUB: not implemented"; return }

func (as *accumulatedStats) finishRPC(rpcType string, err error) { _ = "STUB: not implemented"; return }

var (
	failOnFailedRPC        = flag.Bool("fail_on_failed_rpc", false, "Fail client if any RPCs fail after first success")
	numChannels            = flag.Int("num_channels", 1, "Num of channels")
	printResponse          = flag.Bool("print_response", false, "Write RPC response to stdout")
	qps                    = flag.Int("qps", 1, "QPS per channel, for each type of RPC")
	rpc                    = flag.String("rpc", "UnaryCall", "Types of RPCs to make, ',' separated string. RPCs can be EmptyCall or UnaryCall. Deprecated: Use Configure RPC to XdsUpdateClientConfigureServiceServer instead.")
	rpcMetadata            = flag.String("metadata", "", "The metadata to send with RPC, in format EmptyCall:key1:value1,UnaryCall:key2:value2. Deprecated: Use Configure RPC to XdsUpdateClientConfigureServiceServer instead.")
	rpcTimeout             = flag.Duration("rpc_timeout", 20*time.Second, "Per RPC timeout")
	server                 = flag.String("server", "localhost:8080", "Address of server to connect to")
	statsPort              = flag.Int("stats_port", 8081, "Port to expose peer distribution stats service")
	secureMode             = flag.Bool("secure_mode", false, "If true, retrieve security configuration from the management server. Else, use insecure credentials.")
	enableCSMObservability = flag.Bool("enable_csm_observability", false, "Whether to enable CSM Observability")
	requestPayloadSize     = flag.Int("request_payload_size", 0, "Ask the server to respond with SimpleResponse.payload.body of the given length (may not be implemented on the server).")
	responsePayloadSize    = flag.Int("response_payload_size", 0, "Ask the server to respond with SimpleResponse.payload.body of the given length (may not be implemented on the server).")

	rpcCfgs atomic.Value

	mu               sync.Mutex
	currentRequestID int32
	watchers         = make(map[statsWatcherKey]*statsWatcher)

	accStats = accumulatedStats{
		numRPCsStartedByMethod:   make(map[string]int32),
		numRPCsSucceededByMethod: make(map[string]int32),
		numRPCsFailedByMethod:    make(map[string]int32),
		rpcStatusByMethod:        make(map[string]map[int32]int32),
	}

	rpcSucceeded uint32

	logger = grpclog.Component("interop")
)

type statsService struct {
	testgrpc.UnimplementedLoadBalancerStatsServiceServer
}

func hasRPCSucceeded() bool { _ = "STUB: not implemented"; return false }

func setRPCSucceeded() { _ = "STUB: not implemented"; return }

func (s *statsService) GetClientStats(ctx context.Context, in *testpb.LoadBalancerStatsRequest) (*testpb.LoadBalancerStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *statsService) GetClientAccumulatedStats(context.Context, *testpb.LoadBalancerAccumulatedStatsRequest) (*testpb.LoadBalancerAccumulatedStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type configureService struct {
	testgrpc.UnimplementedXdsUpdateClientConfigureServiceServer
}

func (s *configureService) Configure(_ context.Context, in *testpb.ClientConfigureRequest) (*testpb.ClientConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	unaryCall string = "UnaryCall"
	emptyCall string = "EmptyCall"
)

func parseRPCTypes(rpcStr string) []string { _ = "STUB: not implemented"; return nil }

type rpcConfig struct {
	typ     string
	md      metadata.MD
	timeout int32
}

func parseRPCMetadata(rpcMetadataStr string, rpcs []string) []*rpcConfig {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	flag.Parse()
	if *enableCSMObservability {
		exporter, err := prometheus.New()
		if err != nil {
			logger.Fatalf("Failed to start prometheus exporter: %v", err)
		}
		provider := metric.NewMeterProvider(
			metric.WithReader(exporter),
		)
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

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		cleanup := csm.EnableObservability(ctx, opentelemetry.Options{MetricsOptions: opentelemetry.MetricsOptions{MeterProvider: provider}})
		defer cleanup()
	}

	rpcCfgs.Store(parseRPCMetadata(*rpcMetadata, parseRPCTypes(*rpc)))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *statsPort))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	defer s.Stop()
	testgrpc.RegisterLoadBalancerStatsServiceServer(s, &statsService{})
	testgrpc.RegisterXdsUpdateClientConfigureServiceServer(s, &configureService{})
	reflection.Register(s)
	cleanup, err := admin.Register(s)
	if err != nil {
		logger.Fatalf("Failed to register admin: %v", err)
	}
	defer cleanup()
	go s.Serve(lis)

	creds := insecure.NewCredentials()
	if *secureMode {
		var err error
		creds, err = xds.NewClientCredentials(xds.ClientOptions{FallbackCreds: insecure.NewCredentials()})
		if err != nil {
			logger.Fatalf("Failed to create xDS credentials: %v", err)
		}
	}

	clients := make([]testgrpc.TestServiceClient, *numChannels)
	for i := 0; i < *numChannels; i++ {
		conn, err := grpc.NewClient(*server, grpc.WithTransportCredentials(creds))
		if err != nil {
			logger.Fatalf("grpc.NewClient(%q) = %v", *server, err)
		}
		defer conn.Close()
		clients[i] = testgrpc.NewTestServiceClient(conn)
	}
	ticker := time.NewTicker(time.Second / time.Duration(*qps**numChannels))
	defer ticker.Stop()
	sendRPCs(clients, ticker)
}

func makeOneRPC(c testgrpc.TestServiceClient, cfg *rpcConfig) (*peer.Peer, *rpcInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func sendRPCs(clients []testgrpc.TestServiceClient, ticker *time.Ticker) {
	_ = "STUB: not implemented"
	return
}
