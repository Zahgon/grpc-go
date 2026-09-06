package grpc

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/internal/binarylog"
	iresolver "google.golang.org/grpc/internal/resolver"
	"google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/mem"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
)

var metadataFromOutgoingContextRaw = internal.FromOutgoingContextRaw.(func(context.Context) (metadata.MD, [][]string, bool))

type StreamHandler func(srv any, stream ServerStream) error

type StreamDesc struct {
	StreamName string
	Handler    StreamHandler

	ServerStreams bool
	ClientStreams bool
}

type Stream interface {
	Context() context.Context

	SendMsg(m any) error

	RecvMsg(m any) error
}

type ClientStream interface {
	Header() (metadata.MD, error)

	Trailer() metadata.MD

	CloseSend() error

	Context() context.Context

	SendMsg(m any) error

	RecvMsg(m any) error
}

type clientStreamWrapper struct {
	ClientStream
	desc *StreamDesc
}

func (w *clientStreamWrapper) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (w *clientStreamWrapper) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }

func defaultStreamInterceptor(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

func (cc *ClientConn) NewStream(ctx context.Context, desc *StreamDesc, method string, opts ...CallOption) (ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

func NewClientStream(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, opts ...CallOption) (ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

var emptyMethodConfig = serviceconfig.MethodConfig{}

func endOfClientStream(cc *ClientConn, err error, opts ...CallOption) {
	_ = "STUB: not implemented"
	return
}

type clientInterceptor interface {
	NewStream(ctx context.Context, ri iresolver.RPCInfo, newStream func(ctx context.Context, opts ...CallOption) (ClientStream, error), opts ...CallOption) (ClientStream, error)
	Close()
}

func newClientStream(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, opts ...CallOption) (_ ClientStream, err error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

func newClientStreamWithParams(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, mc *serviceconfig.MethodConfig, onCommit func(), nameResolutionDelayed bool, opts ...CallOption) (_ ClientStream, err error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

func (cs *clientStream) newAttemptLocked(isTransparent bool) (*csAttempt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *csAttempt) getTransport() error { _ = "STUB: not implemented"; return nil }

func (a *csAttempt) newStream() error { _ = "STUB: not implemented"; return nil }

type clientStream struct {
	callHdr  *transport.CallHdr
	opts     []CallOption
	callInfo *callInfo
	cc       *ClientConn
	desc     *StreamDesc

	codec        baseCodec
	compressorV0 Compressor
	compressorV1 encoding.Compressor

	cancel context.CancelFunc

	methodConfig *MethodConfig

	ctx context.Context

	retryThrottler *retryThrottler

	binlogs []binarylog.MethodLogger

	mu                      sync.Mutex
	numRetries              int
	numRetriesSincePushback int

	attempt *csAttempt

	onCommit         func()
	replayBuffer     []replayOp
	replayBufferSize int

	sentLast         bool
	receivedFirstMsg bool

	serverHeaderBinlogged bool

	nameResolutionDelay bool

	firstAttempt bool
	finished     bool
	committed    bool
}

type replayOp struct {
	op      func(a *csAttempt) error
	cleanup func()
}

type csAttempt struct {
	ctx             context.Context
	cs              *clientStream
	transport       transport.ClientTransport
	transportStream *transport.ClientStream
	parser          parser
	pickResult      balancer.PickResult

	decompressorV0 Decompressor
	decompressorV1 encoding.Compressor

	mu sync.Mutex

	trInfo *traceInfo

	statsHandler stats.Handler
	beginTime    time.Time

	decompressorSet       bool
	allowTransparentRetry bool
	drop                  bool

	finished bool
}

func (cs *clientStream) commitAttemptLocked() { _ = "STUB: not implemented"; return }

func (cs *clientStream) commitAttempt() { _ = "STUB: not implemented"; return }

func (a *csAttempt) shouldRetry(err error) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs *clientStream) retryLocked(attempt *csAttempt, lastErr error) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *clientStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (cs *clientStream) withRetry(op func(a *csAttempt) error, onSuccess func()) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *clientStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (cs *clientStream) Trailer() metadata.MD { _ = "STUB: not implemented"; return *new(metadata.MD) }

func (cs *clientStream) replayBufferLocked(attempt *csAttempt) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *clientStream) bufferForRetryLocked(sz int, op func(a *csAttempt) error, cleanup func()) {
	_ = "STUB: not implemented"
	return
}

func (cs *clientStream) SendMsg(m any) (err error) { _ = "STUB: not implemented"; return nil }

func (cs *clientStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (cs *clientStream) CloseSend() error { _ = "STUB: not implemented"; return nil }

func (cs *clientStream) finish(err error) { _ = "STUB: not implemented"; return }

func (a *csAttempt) sendMsg(m any, hdr []byte, payld mem.BufferSlice, dataLength, payloadLength int) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *csAttempt) recvMsg(m any, payInfo *payloadInfo) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (a *csAttempt) finish(err error) { _ = "STUB: not implemented"; return }

func newNonRetryClientStream(ctx context.Context, desc *StreamDesc, method string, t transport.ClientTransport, ac *addrConn, opts ...CallOption) (_ ClientStream, err error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

type addrConnStream struct {
	transportStream  *transport.ClientStream
	ac               *addrConn
	callHdr          *transport.CallHdr
	cancel           context.CancelFunc
	opts             []CallOption
	callInfo         *callInfo
	transport        transport.ClientTransport
	ctx              context.Context
	desc             *StreamDesc
	codec            baseCodec
	sendCompressorV0 Compressor
	sendCompressorV1 encoding.Compressor
	decompressorV0   Decompressor
	decompressorV1   encoding.Compressor

	mu     sync.Mutex
	parser parser

	sentLast         bool
	receivedFirstMsg bool
	decompressorSet  bool

	finished bool
}

func (as *addrConnStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (as *addrConnStream) Trailer() metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

func (as *addrConnStream) CloseSend() error { _ = "STUB: not implemented"; return nil }

func (as *addrConnStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (as *addrConnStream) SendMsg(m any) (err error) { _ = "STUB: not implemented"; return nil }

func (as *addrConnStream) RecvMsg(m any) (err error) { _ = "STUB: not implemented"; return nil }

func (as *addrConnStream) finish(err error) { _ = "STUB: not implemented"; return }

type ServerStream interface {
	SetHeader(metadata.MD) error

	SendHeader(metadata.MD) error

	SetTrailer(metadata.MD)

	Context() context.Context

	SendMsg(m any) error

	RecvMsg(m any) error
}

type serverStream struct {
	ctx   context.Context
	s     *transport.ServerStream
	p     parser
	codec baseCodec
	desc  *StreamDesc

	compressorV0   Compressor
	compressorV1   encoding.Compressor
	decompressorV0 Decompressor
	decompressorV1 encoding.Compressor

	sendCompressorName string

	maxReceiveMessageSize int
	maxSendMessageSize    int

	mu     sync.Mutex
	trInfo *traceInfo

	statsHandler stats.Handler

	binlogs []binarylog.MethodLogger

	recvFirstMsg bool

	serverHeaderBinlogged bool
}

func (ss *serverStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (ss *serverStream) SetHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (ss *serverStream) SendHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (ss *serverStream) SetTrailer(md metadata.MD) { _ = "STUB: not implemented"; return }

func (ss *serverStream) SendMsg(m any) (err error) { _ = "STUB: not implemented"; return nil }

func (ss *serverStream) RecvMsg(m any) (err error) { _ = "STUB: not implemented"; return nil }

func MethodFromServerStream(stream ServerStream) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func prepareMsg(m any, codec baseCodec, cp Compressor, comp encoding.Compressor, pool mem.BufferPool) (hdr []byte, data, payload mem.BufferSlice, pf payloadFormat, err error) {
	_ = "STUB: not implemented"
	return nil, *new(mem.BufferSlice), *new(mem.BufferSlice), *new(payloadFormat), nil
}
