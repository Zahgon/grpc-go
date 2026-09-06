package stats

import (
	"bytes"
	"runtime"
	"sync"
	"time"
)

type FeatureIndex int

const (
	EnableTraceIndex FeatureIndex = iota
	ReadLatenciesIndex
	ReadKbpsIndex
	ReadMTUIndex
	MaxConcurrentCallsIndex
	ReqSizeBytesIndex
	RespSizeBytesIndex
	ReqPayloadCurveIndex
	RespPayloadCurveIndex
	CompModesIndex
	EnableChannelzIndex
	EnablePreloaderIndex
	ClientReadBufferSize
	ClientWriteBufferSize
	ServerReadBufferSize
	ServerWriteBufferSize
	SleepBetweenRPCs
	RecvBufferPool
	SharedWriteBuffer

	MaxFeatureIndex
)

type Features struct {
	NetworkMode string

	UseBufConn bool

	EnableKeepalive bool

	BenchTime time.Duration

	Connections int

	EnableTrace bool

	Latency time.Duration

	Kbps int

	MTU int

	MaxConcurrentCalls int

	ReqSizeBytes int

	RespSizeBytes int

	ReqPayloadCurve *PayloadCurve

	RespPayloadCurve *PayloadCurve

	ModeCompressor string

	EnableChannelz bool

	EnablePreloader bool

	ClientReadBufferSize int

	ClientWriteBufferSize int

	ServerReadBufferSize int

	ServerWriteBufferSize int

	SleepBetweenRPCs time.Duration

	RecvBufferPool string

	SharedWriteBuffer bool
}

func (f Features) String() string { _ = "STUB: not implemented"; return "" }

func (f Features) SharedFeatures(wantFeatures []bool) string { _ = "STUB: not implemented"; return "" }

func (f Features) PrintableName(wantFeatures []bool) string { _ = "STUB: not implemented"; return "" }

func (f Features) partialString(b *bytes.Buffer, wantFeatures []bool, sep, delim string) {
	_ = "STUB: not implemented"
	return
}

type BenchResults struct {
	GoVersion string

	GrpcVersion string

	RunMode string

	Features Features

	SharedFeatures []bool

	Data RunData
}

type RunData struct {
	TotalOps uint64

	SendOps uint64

	RecvOps uint64

	AllocedBytes float64

	Allocs float64

	ReqT float64

	RespT float64

	Fiftieth time.Duration

	Ninetieth time.Duration

	NinetyNinth time.Duration

	Average time.Duration
}

type durationSlice []time.Duration

func (a durationSlice) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a durationSlice) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (a durationSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type Stats struct {
	mu         sync.Mutex
	numBuckets int
	hw         *histWrapper
	results    []BenchResults
	startMS    runtime.MemStats
	stopMS     runtime.MemStats
}

type histWrapper struct {
	unit      time.Duration
	histogram *Histogram
	durations durationSlice
}

func NewStats(numBuckets int) *Stats { _ = "STUB: not implemented"; return nil }

func (s *Stats) StartRun(mode string, f Features, sf []bool) { _ = "STUB: not implemented"; return }

func (s *Stats) EndRun(count uint64) { _ = "STUB: not implemented"; return }

func (s *Stats) EndUnconstrainedRun(req uint64, resp uint64) { _ = "STUB: not implemented"; return }

func (s *Stats) AddDuration(d time.Duration) { _ = "STUB: not implemented"; return }

func (s *Stats) GetResults() []BenchResults { _ = "STUB: not implemented"; return nil }

func (s *Stats) computeLatencies(result *BenchResults) { _ = "STUB: not implemented"; return }

func (s *Stats) dump(result *BenchResults) { _ = "STUB: not implemented"; return }
