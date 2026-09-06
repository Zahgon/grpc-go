package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark/flags"
	"google.golang.org/grpc/benchmark/latency"
	"google.golang.org/grpc/benchmark/stats"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/mem"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var (
	workloads = flags.StringWithAllowedValues("workloads", workloadsAll,
		fmt.Sprintf("Workloads to execute - One of: %v", strings.Join(allWorkloads, ", ")), allWorkloads)
	traceMode = flags.StringWithAllowedValues("trace", toggleModeOff,
		fmt.Sprintf("Trace mode - One of: %v", strings.Join(allToggleModes, ", ")), allToggleModes)
	preloaderMode = flags.StringWithAllowedValues("preloader", toggleModeOff,
		fmt.Sprintf("Preloader mode - One of: %v, preloader works only in streaming and unconstrained modes and will be ignored in unary mode",
			strings.Join(allToggleModes, ", ")), allToggleModes)
	channelzOn = flags.StringWithAllowedValues("channelz", toggleModeOff,
		fmt.Sprintf("Channelz mode - One of: %v", strings.Join(allToggleModes, ", ")), allToggleModes)
	compressorMode = flags.StringWithAllowedValues("compression", compModeOff,
		fmt.Sprintf("Compression mode - One of: %v", strings.Join(allCompModes, ", ")), allCompModes)
	networkMode = flags.StringWithAllowedValues("networkMode", networkModeNone,
		"Network mode includes LAN, WAN, Local and Longhaul", allNetworkModes)
	readLatency           = flags.DurationSlice("latency", defaultReadLatency, "Simulated one-way network latency - may be a comma-separated list")
	readKbps              = flags.IntSlice("kbps", defaultReadKbps, "Simulated network throughput (in kbps) - may be a comma-separated list")
	readMTU               = flags.IntSlice("mtu", defaultReadMTU, "Simulated network MTU (Maximum Transmission Unit) - may be a comma-separated list")
	maxConcurrentCalls    = flags.IntSlice("maxConcurrentCalls", defaultMaxConcurrentCalls, "Number of concurrent RPCs during benchmarks")
	readReqSizeBytes      = flags.IntSlice("reqSizeBytes", nil, "Request size in bytes - may be a comma-separated list")
	readRespSizeBytes     = flags.IntSlice("respSizeBytes", nil, "Response size in bytes - may be a comma-separated list")
	reqPayloadCurveFiles  = flags.StringSlice("reqPayloadCurveFiles", nil, "comma-separated list of CSV files describing the shape a random distribution of request payload sizes")
	respPayloadCurveFiles = flags.StringSlice("respPayloadCurveFiles", nil, "comma-separated list of CSV files describing the shape a random distribution of response payload sizes")
	benchTime             = flag.Duration("benchtime", time.Second, "Configures the amount of time to run each benchmark")
	memProfile            = flag.String("memProfile", "", "Enables memory profiling output to the filename provided.")
	memProfileRate        = flag.Int("memProfileRate", 512*1024, "Configures the memory profiling rate. \n"+
		"memProfile should be set before setting profile rate. To include every allocated block in the profile, "+
		"set MemProfileRate to 1. To turn off profiling entirely, set MemProfileRate to 0. 512 * 1024 by default.")
	cpuProfile          = flag.String("cpuProfile", "", "Enables CPU profiling output to the filename provided")
	benchmarkResultFile = flag.String("resultFile", "", "Save the benchmark result into a binary file")
	useBufconn          = flag.Bool("bufconn", false, "Use in-memory connection instead of system network I/O")
	enableKeepalive     = flag.Bool("enable_keepalive", false, "Enable client keepalive. \n"+
		"Keepalive.Time is set to 10s, Keepalive.Timeout is set to 1s, Keepalive.PermitWithoutStream is set to true.")
	clientReadBufferSize  = flags.IntSlice("clientReadBufferSize", []int{-1}, "Configures the client read buffer size in bytes. If negative, use the default - may be a comma-separated list")
	clientWriteBufferSize = flags.IntSlice("clientWriteBufferSize", []int{-1}, "Configures the client write buffer size in bytes. If negative, use the default - may be a comma-separated list")
	serverReadBufferSize  = flags.IntSlice("serverReadBufferSize", []int{-1}, "Configures the server read buffer size in bytes. If negative, use the default - may be a comma-separated list")
	serverWriteBufferSize = flags.IntSlice("serverWriteBufferSize", []int{-1}, "Configures the server write buffer size in bytes. If negative, use the default - may be a comma-separated list")
	sleepBetweenRPCs      = flags.DurationSlice("sleepBetweenRPCs", []time.Duration{0}, "Configures the maximum amount of time the client should sleep between consecutive RPCs - may be a comma-separated list")
	connections           = flag.Int("connections", 1, "The number of connections. Each connection will handle maxConcurrentCalls RPC streams")
	recvBufferPool        = flags.StringWithAllowedValues("recvBufferPool", recvBufferPoolSimple, "Configures the shared receive buffer pool. One of: nil, simple, all", allRecvBufferPools)
	sharedWriteBuffer     = flags.StringWithAllowedValues("sharedWriteBuffer", toggleModeOn,
		fmt.Sprintf("Configures both client and server to share write buffer - One of: %v", strings.Join(allToggleModes, ", ")), allToggleModes)

	logger = grpclog.Component("benchmark")
)

const (
	workloadsUnary         = "unary"
	workloadsStreaming     = "streaming"
	workloadsUnconstrained = "unconstrained"
	workloadsAll           = "all"

	compModeOff  = "off"
	compModeGzip = "gzip"
	compModeNop  = "nop"
	compModeAll  = "all"

	toggleModeOff  = "off"
	toggleModeOn   = "on"
	toggleModeBoth = "both"

	networkModeNone  = "none"
	networkModeLocal = "Local"
	networkModeLAN   = "LAN"
	networkModeWAN   = "WAN"
	networkLongHaul  = "Longhaul"

	recvBufferPoolNil    = "nil"
	recvBufferPoolSimple = "simple"
	recvBufferPoolAll    = "all"

	numStatsBuckets = 10
	warmupCallCount = 10
	warmuptime      = time.Second
)

var useNopBufferPool atomic.Bool

type swappableBufferPool struct {
	mem.BufferPool
}

func (p swappableBufferPool) Get(length int) *[]byte { _ = "STUB: not implemented"; return nil }

func (p swappableBufferPool) Put(i *[]byte) { _ = "STUB: not implemented"; return }

func init() {
	internal.SetDefaultBufferPool.(func(mem.BufferPool))(swappableBufferPool{mem.DefaultBufferPool()})
}

var (
	allWorkloads              = []string{workloadsUnary, workloadsStreaming, workloadsUnconstrained, workloadsAll}
	allCompModes              = []string{compModeOff, compModeGzip, compModeNop, compModeAll}
	allToggleModes            = []string{toggleModeOff, toggleModeOn, toggleModeBoth}
	allNetworkModes           = []string{networkModeNone, networkModeLocal, networkModeLAN, networkModeWAN, networkLongHaul}
	allRecvBufferPools        = []string{recvBufferPoolNil, recvBufferPoolSimple, recvBufferPoolAll}
	defaultReadLatency        = []time.Duration{0, 40 * time.Millisecond}
	defaultReadKbps           = []int{0, 10240}
	defaultReadMTU            = []int{0}
	defaultMaxConcurrentCalls = []int{1, 8, 64, 512}
	defaultReqSizeBytes       = []int{1, 1024, 1024 * 1024}
	defaultRespSizeBytes      = []int{1, 1024, 1024 * 1024}
	networks                  = map[string]latency.Network{
		networkModeLocal: latency.Local,
		networkModeLAN:   latency.LAN,
		networkModeWAN:   latency.WAN,
		networkLongHaul:  latency.Longhaul,
	}
	keepaliveTime    = 10 * time.Second
	keepaliveTimeout = 1 * time.Second

	keepaliveMinTime = 8 * time.Second
)

type runModes struct {
	unary, streaming, unconstrained bool
}

func runModesFromWorkloads(workload string) runModes {
	_ = "STUB: not implemented"
	return *new(runModes)
}

type startFunc func(mode string, bf stats.Features)
type stopFunc func(count uint64)
type ucStopFunc func(req uint64, resp uint64)
type rpcCallFunc func(cn, pos int)
type rpcSendFunc func(cn, pos int)
type rpcRecvFunc func(cn, pos int)
type rpcCleanupFunc func()

func unaryBenchmark(start startFunc, stop stopFunc, bf stats.Features, s *stats.Stats) {
	_ = "STUB: not implemented"
	return
}

func streamBenchmark(start startFunc, stop stopFunc, bf stats.Features, s *stats.Stats) {
	_ = "STUB: not implemented"
	return
}

func unconstrainedStreamBenchmark(start startFunc, stop ucStopFunc, bf stats.Features) {
	_ = "STUB: not implemented"
	return
}

func makeClients(bf stats.Features) ([]testgrpc.BenchmarkServiceClient, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeFuncUnary(bf stats.Features) (rpcCallFunc, rpcCleanupFunc) {
	_ = "STUB: not implemented"
	return *new(rpcCallFunc), *new(rpcCleanupFunc)
}

func makeFuncStream(bf stats.Features) (rpcCallFunc, rpcCleanupFunc) {
	_ = "STUB: not implemented"
	return *new(rpcCallFunc), *new(rpcCleanupFunc)
}

func makeFuncUnconstrainedStreamPreloaded(bf stats.Features) (rpcSendFunc, rpcRecvFunc, rpcCleanupFunc) {
	_ = "STUB: not implemented"
	return *new(rpcSendFunc), *new(rpcRecvFunc), *new(rpcCleanupFunc)
}

func makeFuncUnconstrainedStream(bf stats.Features) (rpcSendFunc, rpcRecvFunc, rpcCleanupFunc) {
	_ = "STUB: not implemented"
	return *new(rpcSendFunc), *new(rpcRecvFunc), *new(rpcCleanupFunc)
}

func setupStream(bf stats.Features, unconstrained bool) ([][]testgrpc.BenchmarkService_StreamingCallClient, *testpb.SimpleRequest, rpcCleanupFunc) {
	_ = "STUB: not implemented"
	return nil, nil, *new(rpcCleanupFunc)
}

func prepareMessages(streams [][]testgrpc.BenchmarkService_StreamingCallClient, req *testpb.SimpleRequest) [][]*grpc.PreparedMsg {
	_ = "STUB: not implemented"
	return nil
}

func unaryCaller(client testgrpc.BenchmarkServiceClient, reqSize, respSize int) {
	_ = "STUB: not implemented"
	return
}

func streamCaller(stream testgrpc.BenchmarkService_StreamingCallClient, req any) {
	_ = "STUB: not implemented"
	return
}

func runBenchmark(caller rpcCallFunc, start startFunc, stop stopFunc, bf stats.Features, s *stats.Stats, mode string) {
	_ = "STUB: not implemented"
	return
}

type benchOpts struct {
	rModes              runModes
	benchTime           time.Duration
	memProfileRate      int
	memProfile          string
	cpuProfile          string
	networkMode         string
	benchmarkResultFile string
	useBufconn          bool
	enableKeepalive     bool
	connections         int
	features            *featureOpts
}

type featureOpts struct {
	enableTrace           []bool
	readLatencies         []time.Duration
	readKbps              []int
	readMTU               []int
	maxConcurrentCalls    []int
	reqSizeBytes          []int
	respSizeBytes         []int
	reqPayloadCurves      []*stats.PayloadCurve
	respPayloadCurves     []*stats.PayloadCurve
	compModes             []string
	enableChannelz        []bool
	enablePreloader       []bool
	clientReadBufferSize  []int
	clientWriteBufferSize []int
	serverReadBufferSize  []int
	serverWriteBufferSize []int
	sleepBetweenRPCs      []time.Duration
	recvBufferPools       []string
	sharedWriteBuffer     []bool
}

func makeFeaturesNum(b *benchOpts) []int { _ = "STUB: not implemented"; return nil }

func sharedFeatures(featuresNum []int) []bool { _ = "STUB: not implemented"; return nil }

func (b *benchOpts) generateFeatures(featuresNum []int) []stats.Features {
	_ = "STUB: not implemented"
	return nil
}

func addOne(features []int, featuresMaxPosition []int) { _ = "STUB: not implemented"; return }

func processFlags() *benchOpts { _ = "STUB: not implemented"; return nil }

func setToggleMode(val string) []bool { _ = "STUB: not implemented"; return nil }

func setCompressorMode(val string) []string { _ = "STUB: not implemented"; return nil }

func setRecvBufferPool(val string) []string { _ = "STUB: not implemented"; return nil }

func main() {
	opts := processFlags()
	before(opts)

	s := stats.NewStats(numStatsBuckets)
	featuresNum := makeFeaturesNum(opts)
	sf := sharedFeatures(featuresNum)

	var (
		start  = func(mode string, bf stats.Features) { s.StartRun(mode, bf, sf) }
		stop   = func(count uint64) { s.EndRun(count) }
		ucStop = func(req uint64, resp uint64) { s.EndUnconstrainedRun(req, resp) }
	)

	for _, bf := range opts.generateFeatures(featuresNum) {
		grpc.EnableTracing = bf.EnableTrace
		if bf.EnableChannelz {
			channelz.TurnOn()
		}
		if opts.rModes.unary {
			unaryBenchmark(start, stop, bf, s)
		}
		if opts.rModes.streaming {
			streamBenchmark(start, stop, bf, s)
		}
		if opts.rModes.unconstrained {
			unconstrainedStreamBenchmark(start, ucStop, bf)
		}
	}
	after(opts, s.GetResults())
}

func before(opts *benchOpts) { _ = "STUB: not implemented"; return }

func after(opts *benchOpts, data []stats.BenchResults) { _ = "STUB: not implemented"; return }

type nopCompressor struct{}

func (nopCompressor) Do(w io.Writer, p []byte) error { _ = "STUB: not implemented"; return nil }

func (nopCompressor) Type() string { _ = "STUB: not implemented"; return "" }

type nopDecompressor struct{}

func (nopDecompressor) Do(r io.Reader) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (nopDecompressor) Type() string                   { _ = "STUB: not implemented"; return "" }
