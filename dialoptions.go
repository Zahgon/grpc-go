package grpc

import (
	"context"
	"net"
	"net/url"
	"time"

	"google.golang.org/grpc/channelz"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal"
	internalbackoff "google.golang.org/grpc/internal/backoff"
	"google.golang.org/grpc/internal/binarylog"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/mem"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/stats"
)

const (
	defaultMaxCallAttempts = 5
)

func init() {
	internal.AddGlobalDialOptions = func(opt ...DialOption) {
		globalDialOptions = append(globalDialOptions, opt...)
	}
	internal.ClearGlobalDialOptions = func() {
		globalDialOptions = nil
	}
	internal.AddGlobalPerTargetDialOptions = func(opt any) {
		if ptdo, ok := opt.(perTargetDialOption); ok {
			globalPerTargetDialOptions = append(globalPerTargetDialOptions, ptdo)
		}
	}
	internal.ClearGlobalPerTargetDialOptions = func() {
		globalPerTargetDialOptions = nil
	}
	internal.WithBinaryLogger = withBinaryLogger
	internal.JoinDialOptions = newJoinDialOption
	internal.DisableGlobalDialOptions = newDisableGlobalDialOptions
	internal.WithBufferPool = withBufferPool
}

type dialOptions struct {
	unaryInt  UnaryClientInterceptor
	streamInt StreamClientInterceptor

	chainUnaryInts  []UnaryClientInterceptor
	chainStreamInts []StreamClientInterceptor

	compressorV0                Compressor
	dc                          Decompressor
	bs                          internalbackoff.Strategy
	block                       bool
	returnLastError             bool
	timeout                     time.Duration
	authority                   string
	binaryLogger                binarylog.Logger
	copts                       transport.ConnectOptions
	callOptions                 []CallOption
	channelzParent              channelz.Identifier
	disableServiceConfig        bool
	disableRetry                bool
	disableHealthCheck          bool
	minConnectTimeout           func() time.Duration
	defaultServiceConfig        *ServiceConfig
	defaultServiceConfigRawJSON *string
	resolvers                   []resolver.Builder
	idleTimeout                 time.Duration
	defaultScheme               string
	maxCallAttempts             int
	enableLocalDNSResolution    bool
	useProxy                    bool
}

type DialOption interface {
	apply(*dialOptions)
}

var globalDialOptions []DialOption

type perTargetDialOption interface {
	DialOptionForTarget(parsedTarget url.URL) DialOption
}

var globalPerTargetDialOptions []perTargetDialOption

type EmptyDialOption struct{}

func (EmptyDialOption) apply(*dialOptions) { _ = "STUB: not implemented"; return }

type disableGlobalDialOptions struct{}

func (disableGlobalDialOptions) apply(*dialOptions) { _ = "STUB: not implemented"; return }

func newDisableGlobalDialOptions() DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

type funcDialOption struct {
	f func(*dialOptions)
}

func (fdo *funcDialOption) apply(do *dialOptions) { _ = "STUB: not implemented"; return }

func newFuncDialOption(f func(*dialOptions)) *funcDialOption { _ = "STUB: not implemented"; return nil }

type joinDialOption struct {
	opts []DialOption
}

func (jdo *joinDialOption) apply(do *dialOptions) { _ = "STUB: not implemented"; return }

func newJoinDialOption(opts ...DialOption) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithSharedWriteBuffer(val bool) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithWriteBufferSize(s int) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithReadBufferSize(s int) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithInitialWindowSize(s int32) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithInitialConnWindowSize(s int32) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithStaticStreamWindowSize(s int32) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithStaticConnWindowSize(s int32) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithMaxMsgSize(s int) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithDefaultCallOptions(cos ...CallOption) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithCodec(c Codec) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithCompressor(cp Compressor) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithDecompressor(dc Decompressor) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithConnectParams(p ConnectParams) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithBackoffMaxDelay(md time.Duration) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithBackoffConfig(b BackoffConfig) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func withBackoff(bs internalbackoff.Strategy) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithBlock() DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithReturnConnectionError() DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithInsecure() DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithNoProxy() DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithLocalDNSResolution() DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithTransportCredentials(creds credentials.TransportCredentials) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithPerRPCCredentials(creds credentials.PerRPCCredentials) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithCredentialsBundle(b credentials.Bundle) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithTimeout(d time.Duration) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithContextDialer(f func(context.Context, string) (net.Conn, error)) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithDialer(f func(string, time.Duration) (net.Conn, error)) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithStatsHandler(h stats.Handler) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func withBinaryLogger(bl binarylog.Logger) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func FailOnNonTempDialError(f bool) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithUserAgent(s string) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithKeepaliveParams(kp keepalive.ClientParameters) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithUnaryInterceptor(f UnaryClientInterceptor) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithChainUnaryInterceptor(interceptors ...UnaryClientInterceptor) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithStreamInterceptor(f StreamClientInterceptor) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithChainStreamInterceptor(interceptors ...StreamClientInterceptor) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithAuthority(a string) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithChannelzParentID(c channelz.Identifier) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithDisableServiceConfig() DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithDefaultServiceConfig(s string) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithDisableRetry() DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

type MaxHeaderListSizeDialOption struct {
	MaxHeaderListSize uint32
}

func (o MaxHeaderListSizeDialOption) apply(do *dialOptions) { _ = "STUB: not implemented"; return }

func WithMaxHeaderListSize(s uint32) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithDisableHealthCheck() DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func defaultDialOptions() dialOptions { _ = "STUB: not implemented"; return *new(dialOptions) }

func withMinConnectDeadline(f func() time.Duration) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func withDefaultScheme(s string) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func WithResolvers(rs ...resolver.Builder) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithIdleTimeout(d time.Duration) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func WithMaxCallAttempts(n int) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

func withBufferPool(bufferPool mem.BufferPool) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}
