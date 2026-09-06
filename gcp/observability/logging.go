package observability

import (
	"context"
	"time"

	gcplogging "cloud.google.com/go/logging"

	binlogpb "google.golang.org/grpc/binarylog/grpc_binarylog_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal"
	iblog "google.golang.org/grpc/internal/binarylog"
)

var lExporter loggingExporter

var newLoggingExporter = newCloudLoggingExporter

var canonicalString = internal.CanonicalString.(func(codes.Code) string)

func translateMetadata(m *binlogpb.Metadata) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func setPeerIfPresent(binlogEntry *binlogpb.GrpcLogEntry, grpcLogEntry *grpcLogEntry) {
	_ = "STUB: not implemented"
	return
}

var loggerTypeToEventLogger = map[binlogpb.GrpcLogEntry_Logger]loggerType{
	binlogpb.GrpcLogEntry_LOGGER_UNKNOWN: loggerUnknown,
	binlogpb.GrpcLogEntry_LOGGER_CLIENT:  loggerClient,
	binlogpb.GrpcLogEntry_LOGGER_SERVER:  loggerServer,
}

type eventType int

const (
	eventTypeUnknown eventType = iota

	eventTypeClientHeader

	eventTypeServerHeader

	eventTypeClientMessage

	eventTypeServerMessage

	eventTypeClientHalfClose

	eventTypeServerTrailer

	eventTypeCancel
)

func (t eventType) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type loggerType int

const (
	loggerUnknown loggerType = iota
	loggerClient
	loggerServer
)

func (t loggerType) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type payload struct {
	Metadata map[string]string `json:"metadata,omitempty"`

	Timeout time.Duration `json:"timeout,omitempty"`

	StatusCode string `json:"statusCode,omitempty"`

	StatusMessage string `json:"statusMessage,omitempty"`

	StatusDetails []byte `json:"statusDetails,omitempty"`

	MessageLength uint32 `json:"messageLength,omitempty"`

	Message []byte `json:"message,omitempty"`
}

type addrType int

const (
	typeUnknown addrType = iota
	ipv4
	ipv6
	unix
)

func (at addrType) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type address struct {
	Type addrType `json:"type,omitempty"`

	Address string `json:"address,omitempty"`

	IPPort uint32 `json:"ipPort,omitempty"`
}

type grpcLogEntry struct {
	CallID string `json:"callId,omitempty"`

	SequenceID uint64 `json:"sequenceId,omitempty"`

	Type eventType `json:"type,omitempty"`

	Logger loggerType `json:"logger,omitempty"`

	Payload payload `json:"payload,omitempty"`

	PayloadTruncated bool `json:"payloadTruncated,omitempty"`

	Peer address `json:"peer,omitempty"`

	Authority string `json:"authority,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	MethodName string `json:"methodName,omitempty"`
}

type methodLoggerBuilder interface {
	Build(iblog.LogEntryConfig) *binlogpb.GrpcLogEntry
}

type binaryMethodLogger struct {
	callID, serviceName, methodName, authority, projectID string

	mlb        methodLoggerBuilder
	exporter   loggingExporter
	clientSide bool
}

func (bml *binaryMethodLogger) buildGCPLoggingEntry(ctx context.Context, c iblog.LogEntryConfig) gcplogging.Entry {
	_ = "STUB: not implemented"
	return *new(gcplogging.Entry)
}

func (bml *binaryMethodLogger) Log(ctx context.Context, c iblog.LogEntryConfig) {
	_ = "STUB: not implemented"
	return
}

type eventConfig struct {
	ServiceMethod map[string]bool
	Services      map[string]bool
	MatchAll      bool

	Exclude      bool
	HeaderBytes  uint64
	MessageBytes uint64
}

type binaryLogger struct {
	EventConfigs []eventConfig
	projectID    string
	exporter     loggingExporter
	clientSide   bool
}

func (bl *binaryLogger) GetMethodLogger(methodName string) iblog.MethodLogger {
	_ = "STUB: not implemented"
	return *new(iblog.MethodLogger)
}

func parseMethod(method string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func registerClientRPCEvents(config *config, exporter loggingExporter) {
	_ = "STUB: not implemented"
	return
}

func registerServerRPCEvents(config *config, exporter loggingExporter) {
	_ = "STUB: not implemented"
	return
}

func startLogging(ctx context.Context, config *config) error { _ = "STUB: not implemented"; return nil }

func stopLogging() { _ = "STUB: not implemented"; return }
