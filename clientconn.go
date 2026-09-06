package grpc

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/pickfirst"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	expstats "google.golang.org/grpc/experimental/stats"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/internal/idle"
	iresolver "google.golang.org/grpc/internal/resolver"
	istats "google.golang.org/grpc/internal/stats"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"

	_ "google.golang.org/grpc/balancer/roundrobin"
	_ "google.golang.org/grpc/internal/resolver/passthrough"
	_ "google.golang.org/grpc/internal/resolver/unix"
	_ "google.golang.org/grpc/resolver/dns"
)

const (
	minConnectTimeout = 20 * time.Second
)

var (
	ErrClientConnClosing = status.Error(codes.Canceled, "grpc: the client connection is closing")

	errConnDrain = errors.New("grpc: the connection is drained")

	errConnClosing = errors.New("grpc: the connection is closing")

	errConnIdling = errors.New("grpc: the connection is closing due to channel idleness")

	invalidDefaultServiceConfigErrPrefix = "grpc: the provided default service config is invalid"

	PickFirstBalancerName = pickfirst.Name
)

var (
	errNoTransportSecurity = errors.New("grpc: no transport security set (use grpc.WithTransportCredentials(insecure.NewCredentials()) explicitly or set credentials)")

	errTransportCredsAndBundle = errors.New("grpc: credentials.Bundle may not be used with individual TransportCredentials")

	errNoTransportCredsInBundle = errors.New("grpc: credentials.Bundle must return non-nil transport credentials")

	errTransportCredentialsMissing = errors.New("grpc: the credentials require transport level security (use grpc.WithTransportCredentials() to set)")
)

var (
	disconnectionsMetric = expstats.RegisterInt64Count(expstats.MetricDescriptor{
		Name:           "grpc.subchannel.disconnections",
		Description:    "EXPERIMENTAL. Number of times the selected subchannel becomes disconnected.",
		Unit:           "{disconnection}",
		Labels:         []string{"grpc.target"},
		OptionalLabels: []string{"grpc.lb.backend_service", "grpc.lb.locality", "grpc.disconnect_error"},
		Default:        false,
	})
	connectionAttemptsSucceededMetric = expstats.RegisterInt64Count(expstats.MetricDescriptor{
		Name:           "grpc.subchannel.connection_attempts_succeeded",
		Description:    "EXPERIMENTAL. Number of successful connection attempts.",
		Unit:           "{attempt}",
		Labels:         []string{"grpc.target"},
		OptionalLabels: []string{"grpc.lb.backend_service", "grpc.lb.locality"},
		Default:        false,
	})
	connectionAttemptsFailedMetric = expstats.RegisterInt64Count(expstats.MetricDescriptor{
		Name:           "grpc.subchannel.connection_attempts_failed",
		Description:    "EXPERIMENTAL. Number of failed connection attempts.",
		Unit:           "{attempt}",
		Labels:         []string{"grpc.target"},
		OptionalLabels: []string{"grpc.lb.backend_service", "grpc.lb.locality"},
		Default:        false,
	})
	openConnectionsMetric = expstats.RegisterInt64UpDownCount(expstats.MetricDescriptor{
		Name:           "grpc.subchannel.open_connections",
		Description:    "EXPERIMENTAL. Number of open connections.",
		Unit:           "{attempt}",
		Labels:         []string{"grpc.target"},
		OptionalLabels: []string{"grpc.lb.backend_service", "grpc.security_level", "grpc.lb.locality"},
		Default:        false,
	})
)

const (
	defaultClientMaxReceiveMessageSize = 1024 * 1024 * 4
	defaultClientMaxSendMessageSize    = math.MaxInt32

	defaultWriteBufSize = 32 * 1024
	defaultReadBufSize  = 32 * 1024
)

type defaultConfigSelector struct {
	sc *ServiceConfig
}

func (dcs *defaultConfigSelector) SelectConfig(rpcInfo iresolver.RPCInfo) (*iresolver.RPCConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewClient(target string, opts ...DialOption) (conn *ClientConn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Dial(target string, opts ...DialOption) (*ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DialContext(ctx context.Context, target string, opts ...DialOption) (conn *ClientConn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cc *ClientConn) addTraceEvent(msg string) { _ = "STUB: not implemented"; return }

type idler ClientConn

func (i *idler) EnterIdleMode() { _ = "STUB: not implemented"; return }

func (i *idler) ExitIdleMode() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) exitIdleMode() error { _ = "STUB: not implemented"; return nil }

func (cc *ClientConn) initIdleStateLocked() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) enterIdleMode() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) validateTransportCredentials() error { _ = "STUB: not implemented"; return nil }

func (cc *ClientConn) channelzRegistration(target string) { _ = "STUB: not implemented"; return }

func chainUnaryClientInterceptors(cc *ClientConn) { _ = "STUB: not implemented"; return }

func getChainUnaryInvoker(interceptors []UnaryClientInterceptor, curr int, finalInvoker UnaryInvoker) UnaryInvoker {
	_ = "STUB: not implemented"
	return *new(UnaryInvoker)
}

func chainStreamClientInterceptors(cc *ClientConn) { _ = "STUB: not implemented"; return }

func getChainStreamer(interceptors []StreamClientInterceptor, curr int, finalStreamer Streamer) Streamer {
	_ = "STUB: not implemented"
	return *new(Streamer)
}

func newConnectivityStateManager(ctx context.Context, channel *channelz.Channel) *connectivityStateManager {
	_ = "STUB: not implemented"
	return nil
}

type connectivityStateManager struct {
	mu         sync.Mutex
	state      connectivity.State
	notifyChan chan struct{}
	channelz   *channelz.Channel
	pubSub     *grpcsync.PubSub
}

func (csm *connectivityStateManager) updateState(state connectivity.State) {
	_ = "STUB: not implemented"
	return
}

func (csm *connectivityStateManager) getState() connectivity.State {
	_ = "STUB: not implemented"
	return *new(connectivity.State)
}

func (csm *connectivityStateManager) getNotifyChan() <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

type ClientConnInterface interface {
	Invoke(ctx context.Context, method string, args any, reply any, opts ...CallOption) error

	NewStream(ctx context.Context, desc *StreamDesc, method string, opts ...CallOption) (ClientStream, error)
}

var _ ClientConnInterface = (*ClientConn)(nil)

type ClientConn struct {
	ctx    context.Context
	cancel context.CancelFunc

	target              string
	parsedTarget        resolver.Target
	authority           string
	dopts               dialOptions
	channelz            *channelz.Channel
	resolverBuilder     resolver.Builder
	idlenessMgr         *idle.Manager
	metricsRecorderList *istats.MetricsRecorderList
	statsHandler        stats.Handler

	csMgr              *connectivityStateManager
	pickerWrapper      *pickerWrapper
	safeConfigSelector iresolver.SafeConfigSelector
	retryThrottler     atomic.Value

	mu              sync.RWMutex
	resolverWrapper *ccResolverWrapper
	balancerWrapper *ccBalancerWrapper
	sc              *ServiceConfig
	conns           map[*addrConn]struct{}
	keepaliveParams keepalive.ClientParameters

	firstResolveEvent *grpcsync.Event

	lceMu               sync.Mutex
	lastConnectionError error
}

func (cc *ClientConn) WaitForStateChange(ctx context.Context, sourceState connectivity.State) bool {
	_ = "STUB: not implemented"
	return false
}

func (cc *ClientConn) GetState() connectivity.State {
	_ = "STUB: not implemented"
	return *new(connectivity.State)
}

func (cc *ClientConn) Connect() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) waitForResolvedAddrs(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

var emptyServiceConfig *ServiceConfig

func init() {
	cfg := parseServiceConfig("{}", defaultMaxCallAttempts)
	if cfg.Err != nil {
		panic(fmt.Sprintf("impossible error parsing empty service config: %v", cfg.Err))
	}
	emptyServiceConfig = cfg.Config.(*ServiceConfig)

	internal.SubscribeToConnectivityStateChanges = func(cc *ClientConn, s grpcsync.Subscriber) func() {
		return cc.csMgr.pubSub.Subscribe(s)
	}
	internal.EnterIdleModeForTesting = func(cc *ClientConn) {
		cc.idlenessMgr.EnterIdleModeForTesting()
	}
	internal.ExitIdleModeForTesting = func(cc *ClientConn) {
		cc.idlenessMgr.ExitIdleMode()
	}
}

func (cc *ClientConn) maybeApplyDefaultServiceConfig() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) updateResolverStateAndUnlock(s resolver.State, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *ClientConn) applyFailingLBLocked(sc *serviceconfig.ParseResult) {
	_ = "STUB: not implemented"
	return
}

func copyAddresses(in []resolver.Address) []resolver.Address { _ = "STUB: not implemented"; return nil }

func (cc *ClientConn) newAddrConnLocked(addrs []resolver.Address, opts balancer.NewSubConnOptions) (*addrConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cc *ClientConn) removeAddrConn(ac *addrConn, err error) { _ = "STUB: not implemented"; return }

func (cc *ClientConn) Target() string { _ = "STUB: not implemented"; return "" }

func (cc *ClientConn) CanonicalTarget() string { _ = "STUB: not implemented"; return "" }

func (cc *ClientConn) incrCallsStarted() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) incrCallsSucceeded() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) incrCallsFailed() { _ = "STUB: not implemented"; return }

func (ac *addrConn) connect() { _ = "STUB: not implemented"; return }

func equalAddressIgnoringBalAttributes(a, b *resolver.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func equalAddressesIgnoringBalAttributes(a, b []resolver.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (ac *addrConn) updateAddrs(addrs []resolver.Address) { _ = "STUB: not implemented"; return }

func (cc *ClientConn) getServerName(addr resolver.Address) string {
	_ = "STUB: not implemented"
	return ""
}

func getMethodConfig(sc *ServiceConfig, method string) MethodConfig {
	_ = "STUB: not implemented"
	return *new(MethodConfig)
}

func (cc *ClientConn) GetMethodConfig(method string) MethodConfig {
	_ = "STUB: not implemented"
	return *new(MethodConfig)
}

func (cc *ClientConn) healthCheckConfig() *healthCheckConfig { _ = "STUB: not implemented"; return nil }

func (cc *ClientConn) applyServiceConfigAndBalancer(sc *ServiceConfig, configSelector iresolver.ConfigSelector) {
	_ = "STUB: not implemented"
	return
}

func (cc *ClientConn) resolveNow(o resolver.ResolveNowOptions) { _ = "STUB: not implemented"; return }

func (cc *ClientConn) resolveNowLocked(o resolver.ResolveNowOptions) {
	_ = "STUB: not implemented"
	return
}

func (cc *ClientConn) ResetConnectBackoff() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) Close() error { _ = "STUB: not implemented"; return nil }

type addrConn struct {
	ctx    context.Context
	cancel context.CancelFunc

	cc     *ClientConn
	dopts  dialOptions
	acbw   *acBalancerWrapper
	scopts balancer.NewSubConnOptions

	transport transport.ClientTransport

	mu      sync.Mutex
	curAddr resolver.Address
	addrs   []resolver.Address

	state connectivity.State

	backoffIdx   int
	resetBackoff chan struct{}

	channelz *channelz.SubChannel

	localityLabel        string
	backendServiceLabel  string
	disconnectErrorLabel string
}

func (ac *addrConn) updateConnectivityState(s connectivity.State, lastErr error) {
	_ = "STUB: not implemented"
	return
}

func (ac *addrConn) adjustParams(r transport.GoAwayReason) { _ = "STUB: not implemented"; return }

func (ac *addrConn) resetTransportAndUnlock() { _ = "STUB: not implemented"; return }

func (ac *addrConn) updateTelemetryLabelsLocked() { _ = "STUB: not implemented"; return }

type securityLevelKey struct{}

func (ac *addrConn) securityLevelLocked() string { _ = "STUB: not implemented"; return "" }

func (ac *addrConn) tryAllAddrs(ctx context.Context, addrs []resolver.Address, connectDeadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (ac *addrConn) createTransport(ctx context.Context, addr resolver.Address, copts transport.ConnectOptions, connectDeadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func disconnectErrorString(info transport.GoAwayInfo) string { _ = "STUB: not implemented"; return "" }

func (ac *addrConn) startHealthCheck(ctx context.Context) { _ = "STUB: not implemented"; return }

func (ac *addrConn) resetConnectBackoff() { _ = "STUB: not implemented"; return }

func (ac *addrConn) getReadyTransport() transport.ClientTransport {
	_ = "STUB: not implemented"
	return *new(transport.ClientTransport)
}

func (ac *addrConn) tearDown(err error) { _ = "STUB: not implemented"; return }

type retryThrottler struct {
	max    float64
	thresh float64
	ratio  float64

	mu     sync.Mutex
	tokens float64
}

func (rt *retryThrottler) throttle() bool { _ = "STUB: not implemented"; return false }

func (rt *retryThrottler) successfulRPC() { _ = "STUB: not implemented"; return }

func (ac *addrConn) incrCallsStarted() { _ = "STUB: not implemented"; return }

func (ac *addrConn) incrCallsSucceeded() { _ = "STUB: not implemented"; return }

func (ac *addrConn) incrCallsFailed() { _ = "STUB: not implemented"; return }

var ErrClientConnTimeout = errors.New("grpc: timed out when dialing")

func (cc *ClientConn) getResolver(scheme string) resolver.Builder {
	_ = "STUB: not implemented"
	return *new(resolver.Builder)
}

func (cc *ClientConn) updateConnectionError(err error) { _ = "STUB: not implemented"; return }

func (cc *ClientConn) connectionError() error { _ = "STUB: not implemented"; return nil }

func (cc *ClientConn) initParsedTargetAndResolverBuilder() error {
	_ = "STUB: not implemented"
	return nil
}

func parseTarget(target string) (resolver.Target, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Target), nil
}

func encodeAuthority(authority string) string { _ = "STUB: not implemented"; return "" }

func (cc *ClientConn) initAuthority() error { _ = "STUB: not implemented"; return nil }
