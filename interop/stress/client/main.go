package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"

	_ "google.golang.org/grpc/xds/googledirectpath"

	metricspb "google.golang.org/grpc/interop/stress/grpc_testing"
)

const (
	googleDefaultCredsName = "google_default_credentials"
	computeEngineCredsName = "compute_engine_channel_creds"
)

var (
	serverAddresses       = flag.String("server_addresses", "localhost:8080", "a list of server addresses")
	testCases             = flag.String("test_cases", "", "a list of test cases along with the relative weights")
	testDurationSecs      = flag.Int("test_duration_secs", -1, "test duration in seconds")
	numChannelsPerServer  = flag.Int("num_channels_per_server", 1, "Number of channels (i.e connections) to each server")
	numStubsPerChannel    = flag.Int("num_stubs_per_channel", 1, "Number of client stubs per each connection to server")
	metricsPort           = flag.Int("metrics_port", 8081, "The port at which the stress client exposes QPS metrics")
	useTLS                = flag.Bool("use_tls", false, "Connection uses TLS if true, else plain TCP")
	testCA                = flag.Bool("use_test_ca", false, "Whether to replace platform root CAs with test CA as the CA root")
	tlsServerName         = flag.String("server_host_override", "foo.test.google.fr", "The server name use to verify the hostname returned by TLS handshake if it is not empty. Otherwise, --server_host is used.")
	caFile                = flag.String("ca_file", "", "The file containing the CA root cert file")
	customCredentialsType = flag.String("custom_credentials_type", "", "Custom credentials type to use")

	totalNumCalls int64
	logger        = grpclog.Component("stress")
)

type testCaseWithWeight struct {
	name   string
	weight int
}

func parseTestCases(testCaseString string) []testCaseWithWeight {
	_ = "STUB: not implemented"
	return nil
}

type weightedRandomTestSelector struct {
	tests       []testCaseWithWeight
	totalWeight int
}

func newWeightedRandomTestSelector(tests []testCaseWithWeight) *weightedRandomTestSelector {
	_ = "STUB: not implemented"
	return nil
}

func (selector weightedRandomTestSelector) getNextTest() string {
	_ = "STUB: not implemented"
	return ""
}

type gauge struct {
	mutex sync.RWMutex
	val   int64
}

func (g *gauge) set(v int64) { _ = "STUB: not implemented"; return }

func (g *gauge) get() int64 { _ = "STUB: not implemented"; return 0 }

type server struct {
	metricspb.UnimplementedMetricsServiceServer
	mutex sync.RWMutex

	gauges map[string]*gauge
}

func newMetricsServer() *server { _ = "STUB: not implemented"; return nil }

func (s *server) GetAllGauges(_ *metricspb.EmptyMessage, stream metricspb.MetricsService_GetAllGaugesServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) GetGauge(_ context.Context, in *metricspb.GaugeRequest) (*metricspb.GaugeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) createGauge(name string) *gauge { _ = "STUB: not implemented"; return nil }

func startServer(server *server, port int) { _ = "STUB: not implemented"; return }

func performRPCs(gauge *gauge, conn *grpc.ClientConn, selector *weightedRandomTestSelector, stop <-chan bool) {
	_ = "STUB: not implemented"
	return
}

func logParameterInfo(addresses []string, tests []testCaseWithWeight) {
	_ = "STUB: not implemented"
	return
}

func newConn(address string, useTLS, testCA bool, tlsServerName string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	flag.Parse()
	resolver.SetDefaultScheme("dns")
	addresses := strings.Split(*serverAddresses, ",")
	tests := parseTestCases(*testCases)
	logParameterInfo(addresses, tests)
	testSelector := newWeightedRandomTestSelector(tests)
	metricsServer := newMetricsServer()

	var wg sync.WaitGroup
	wg.Add(len(addresses) * *numChannelsPerServer * *numStubsPerChannel)
	stop := make(chan bool)

	for serverIndex, address := range addresses {
		for connIndex := 0; connIndex < *numChannelsPerServer; connIndex++ {
			conn, err := newConn(address, *useTLS, *testCA, *tlsServerName)
			if err != nil {
				logger.Fatalf("Fail to dial: %v", err)
			}
			defer conn.Close()
			for clientIndex := 0; clientIndex < *numStubsPerChannel; clientIndex++ {
				name := fmt.Sprintf("/stress_test/server_%d/channel_%d/stub_%d/qps", serverIndex+1, connIndex+1, clientIndex+1)
				go func() {
					defer wg.Done()
					g := metricsServer.createGauge(name)
					performRPCs(g, conn, testSelector, stop)
				}()
			}

		}
	}
	go startServer(metricsServer, *metricsPort)
	if *testDurationSecs > 0 {
		time.Sleep(time.Duration(*testDurationSecs) * time.Second)
		close(stop)
	}
	wg.Wait()
	fmt.Fprintf(os.Stdout, "Total calls made: %v\n", totalNumCalls)
	logger.Infof(" ===== ALL DONE ===== ")
}
