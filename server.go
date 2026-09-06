package grpc

import (
	"context"
	"errors"
	"math"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding"
	estats "google.golang.org/grpc/experimental/stats"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/internal/binarylog"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcsync"
	istats "google.golang.org/grpc/internal/stats"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/mem"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/tap"
)

const (
	defaultServerMaxReceiveMessageSize = 1024 * 1024 * 4
	defaultServerMaxSendMessageSize    = math.MaxInt32

	listenerAddressForServeHTTP = "listenerAddressForServeHTTP"
)

func init() {
	internal.GetServerCredentials = func(srv *Server) credentials.TransportCredentials {
		return srv.opts.creds
	}
	internal.IsRegisteredMethod = func(srv *Server, method string) bool {
		return srv.isRegisteredMethod(method)
	}
	internal.ServerFromContext = serverFromContext
	internal.AddGlobalServerOptions = func(opt ...ServerOption) {
		globalServerOptions = append(globalServerOptions, opt...)
	}
	internal.ClearGlobalServerOptions = func() {
		globalServerOptions = nil
	}
	internal.BinaryLogger = binaryLogger
	internal.JoinServerOptions = newJoinServerOption
	internal.BufferPool = bufferPool
	internal.MetricsRecorderForServer = func(srv *Server) estats.MetricsRecorder {
		return istats.NewMetricsRecorderList(srv.opts.statsHandlers)
	}
	internal.XDSFilterWrapperOption = xdsFilterWrapperOption
}

var statusOK = status.New(codes.OK, "")
var logger = grpclog.Component("core")

type MethodHandler func(srv any, ctx context.Context, dec func(any) error, interceptor UnaryServerInterceptor) (any, error)

type MethodDesc struct {
	MethodName string
	Handler    MethodHandler
}

type ServiceDesc struct {
	ServiceName string

	HandlerType any
	Methods     []MethodDesc
	Streams     []StreamDesc
	Metadata    any
}

type serviceInfo struct {
	serviceImpl any
	streams     map[string]*StreamDesc
	mdata       any
}

type Server struct {
	opts         serverOptions
	statsHandler stats.Handler

	mu  sync.Mutex
	lis map[net.Listener]bool

	conns    map[string]map[transport.ServerTransport]bool
	serve    bool
	drain    bool
	cv       *sync.Cond
	services map[string]*serviceInfo
	events   traceEventLog

	quit               *grpcsync.Event
	done               *grpcsync.Event
	channelzRemoveOnce sync.Once
	serveWG            sync.WaitGroup
	handlersWG         sync.WaitGroup

	channelz *channelz.Server

	serverWorkerChannel      chan func()
	serverWorkerChannelClose func()
}

type serverOptions struct {
	creds                 credentials.TransportCredentials
	codec                 baseCodec
	cp                    Compressor
	dc                    Decompressor
	unaryInt              UnaryServerInterceptor
	streamInt             StreamServerInterceptor
	chainUnaryInts        []UnaryServerInterceptor
	chainStreamInts       []StreamServerInterceptor
	binaryLogger          binarylog.Logger
	inTapHandle           tap.ServerInHandle
	statsHandlers         []stats.Handler
	maxConcurrentStreams  uint32
	maxReceiveMessageSize int
	maxSendMessageSize    int
	unknownStreamDesc     *StreamDesc
	keepaliveParams       keepalive.ServerParameters
	keepalivePolicy       keepalive.EnforcementPolicy
	initialWindowSize     int32
	initialConnWindowSize int32
	writeBufferSize       int
	readBufferSize        int
	sharedWriteBuffer     bool
	connectionTimeout     time.Duration
	maxHeaderListSize     *uint32
	headerTableSize       *uint32
	numServerWorkers      uint32
	bufferPool            mem.BufferPool
	waitForHandlers       bool
	staticWindowSize      bool
	streamWrapper         func(ServerStream) (ServerStream, error)
}

var defaultServerOptions = serverOptions{
	maxConcurrentStreams:  math.MaxUint32,
	maxReceiveMessageSize: defaultServerMaxReceiveMessageSize,
	maxSendMessageSize:    defaultServerMaxSendMessageSize,
	connectionTimeout:     120 * time.Second,
	writeBufferSize:       defaultWriteBufSize,
	sharedWriteBuffer:     true,
	readBufferSize:        defaultReadBufSize,
	bufferPool:            mem.DefaultBufferPool(),
}
var globalServerOptions []ServerOption

type ServerOption interface {
	apply(*serverOptions)
}

type EmptyServerOption struct{}

func (EmptyServerOption) apply(*serverOptions) { _ = "STUB: not implemented"; return }

type funcServerOption struct {
	f func(*serverOptions)
}

func (fdo *funcServerOption) apply(do *serverOptions) { _ = "STUB: not implemented"; return }

func newFuncServerOption(f func(*serverOptions)) *funcServerOption {
	_ = "STUB: not implemented"
	return nil
}

type joinServerOption struct {
	opts []ServerOption
}

func (mdo *joinServerOption) apply(do *serverOptions) { _ = "STUB: not implemented"; return }

func newJoinServerOption(opts ...ServerOption) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func SharedWriteBuffer(val bool) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func WriteBufferSize(s int) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func ReadBufferSize(s int) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func InitialWindowSize(s int32) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func InitialConnWindowSize(s int32) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func StaticStreamWindowSize(s int32) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func StaticConnWindowSize(s int32) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func KeepaliveParams(kp keepalive.ServerParameters) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func KeepaliveEnforcementPolicy(kep keepalive.EnforcementPolicy) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func CustomCodec(codec Codec) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func ForceServerCodec(codec encoding.Codec) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ForceServerCodecV2(codecV2 encoding.CodecV2) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func RPCCompressor(cp Compressor) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func RPCDecompressor(dc Decompressor) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func MaxMsgSize(m int) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func MaxRecvMsgSize(m int) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func MaxSendMsgSize(m int) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func MaxConcurrentStreams(n uint32) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func Creds(c credentials.TransportCredentials) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func UnaryInterceptor(i UnaryServerInterceptor) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ChainUnaryInterceptor(interceptors ...UnaryServerInterceptor) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func StreamInterceptor(i StreamServerInterceptor) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ChainStreamInterceptor(interceptors ...StreamServerInterceptor) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func InTapHandle(h tap.ServerInHandle) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func StatsHandler(h stats.Handler) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func binaryLogger(bl binarylog.Logger) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func UnknownServiceHandler(streamHandler StreamHandler) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ConnectionTimeout(d time.Duration) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

type MaxHeaderListSizeServerOption struct {
	MaxHeaderListSize uint32
}

func (o MaxHeaderListSizeServerOption) apply(so *serverOptions) { _ = "STUB: not implemented"; return }

func MaxHeaderListSize(s uint32) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func HeaderTableSize(s uint32) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func NumStreamWorkers(numServerWorkers uint32) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func WaitForHandlers(w bool) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func bufferPool(bufferPool mem.BufferPool) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func xdsFilterWrapperOption(w func(ServerStream) (ServerStream, error)) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

const serverWorkerResetThreshold = 1 << 16

func (s *Server) serverWorker() { _ = "STUB: not implemented"; return }

func (s *Server) initServerWorkers() { _ = "STUB: not implemented"; return }

func NewServer(opt ...ServerOption) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) printf(format string, a ...any) { _ = "STUB: not implemented"; return }

func (s *Server) errorf(format string, a ...any) { _ = "STUB: not implemented"; return }

type ServiceRegistrar interface {
	RegisterService(desc *ServiceDesc, impl any)
}

func (s *Server) RegisterService(sd *ServiceDesc, ss any) { _ = "STUB: not implemented"; return }

func (s *Server) register(sd *ServiceDesc, ss any) { _ = "STUB: not implemented"; return }

type MethodInfo struct {
	Name string

	IsClientStream bool

	IsServerStream bool
}

type ServiceInfo struct {
	Methods []MethodInfo

	Metadata any
}

func (s *Server) GetServiceInfo() map[string]ServiceInfo { _ = "STUB: not implemented"; return nil }

var ErrServerStopped = errors.New("grpc: the server has been stopped")

type listenSocket struct {
	net.Listener
	channelz *channelz.Socket
}

func (l *listenSocket) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Server) Serve(lis net.Listener) error { _ = "STUB: not implemented"; return nil }

func (s *Server) handleRawConn(lisAddr string, rawConn net.Conn) { _ = "STUB: not implemented"; return }

func (s *Server) newHTTP2Transport(c net.Conn) transport.ServerTransport {
	_ = "STUB: not implemented"
	return *new(transport.ServerTransport)
}

func (s *Server) serveStreams(ctx context.Context, st transport.ServerTransport, rawConn net.Conn) {
	_ = "STUB: not implemented"
	return
}

var _ http.Handler = (*Server)(nil)

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) addConn(addr string, st transport.ServerTransport) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Server) removeConn(addr string, st transport.ServerTransport) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) incrCallsStarted() { _ = "STUB: not implemented"; return }

func (s *Server) incrCallsSucceeded() { _ = "STUB: not implemented"; return }

func (s *Server) incrCallsFailed() { _ = "STUB: not implemented"; return }

func chainUnaryServerInterceptors(s *Server) { _ = "STUB: not implemented"; return }

func chainUnaryInterceptors(interceptors []UnaryServerInterceptor) UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(UnaryServerInterceptor)
}

func getChainUnaryHandler(interceptors []UnaryServerInterceptor, curr int, info *UnaryServerInfo, finalHandler UnaryHandler) UnaryHandler {
	_ = "STUB: not implemented"
	return *new(UnaryHandler)
}

func chainStreamServerInterceptors(s *Server) { _ = "STUB: not implemented"; return }

func chainStreamInterceptors(interceptors []StreamServerInterceptor) StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(StreamServerInterceptor)
}

func getChainStreamHandler(interceptors []StreamServerInterceptor, curr int, info *StreamServerInfo, finalHandler StreamHandler) StreamHandler {
	_ = "STUB: not implemented"
	return *new(StreamHandler)
}

func (s *Server) wrapUnaryHandler(md *MethodDesc) StreamHandler {
	_ = "STUB: not implemented"
	return *new(StreamHandler)
}

func (s *Server) processRPC(ctx context.Context, stream *transport.ServerStream, info *serviceInfo, sd *StreamDesc, trInfo *traceInfo) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleMalformedMethodName(stream *transport.ServerStream, ti *traceInfo) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleStream(t transport.ServerTransport, stream *transport.ServerStream) {
	_ = "STUB: not implemented"
	return
}

type streamKey struct{}

func NewContextWithServerTransportStream(ctx context.Context, stream ServerTransportStream) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type ServerTransportStream interface {
	Method() string
	SetHeader(md metadata.MD) error
	SendHeader(md metadata.MD) error
	SetTrailer(md metadata.MD) error
}

func ServerTransportStreamFromContext(ctx context.Context) ServerTransportStream {
	_ = "STUB: not implemented"
	return *new(ServerTransportStream)
}

func (s *Server) Stop() { _ = "STUB: not implemented"; return }

func (s *Server) GracefulStop() { _ = "STUB: not implemented"; return }

func (s *Server) stop(graceful bool) { _ = "STUB: not implemented"; return }

func (s *Server) closeServerTransportsLocked() { _ = "STUB: not implemented"; return }

func (s *Server) drainAllServerTransportsLocked() { _ = "STUB: not implemented"; return }

func (s *Server) closeListenersLocked() { _ = "STUB: not implemented"; return }

func (s *Server) getCodec(contentSubtype string) baseCodec {
	_ = "STUB: not implemented"
	return *new(baseCodec)
}

type serverKey struct{}

func serverFromContext(ctx context.Context) *Server { _ = "STUB: not implemented"; return nil }

func contextWithServer(ctx context.Context, server *Server) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *Server) isRegisteredMethod(serviceMethod string) bool {
	_ = "STUB: not implemented"
	return false
}

func SetHeader(ctx context.Context, md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func SendHeader(ctx context.Context, md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func SetSendCompressor(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func ClientSupportedCompressors(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetTrailer(ctx context.Context, md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func Method(ctx context.Context) (string, bool) { _ = "STUB: not implemented"; return "", false }

func validateSendCompressor(name string, clientCompressors []string) error {
	_ = "STUB: not implemented"
	return nil
}

type atomicSemaphore struct {
	n    atomic.Int64
	wait chan struct{}
}

func (q *atomicSemaphore) acquire() { _ = "STUB: not implemented"; return }

func (q *atomicSemaphore) release() { _ = "STUB: not implemented"; return }

func newHandlerQuota(n uint32) *atomicSemaphore { _ = "STUB: not implemented"; return nil }
