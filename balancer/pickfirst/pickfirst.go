package pickfirst

import (
	"encoding/json"
	"sync"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

func init() {
	balancer.Register(pickfirstBuilder{})
}

const Name = "pick_first"

type enableHealthListenerKeyType struct{}

var logger = grpclog.Component("pick-first-leaf-lb")

const (
	logPrefix = "[pick-first-leaf-lb %p] "

	connectionDelayInterval = 250 * time.Millisecond
)

type ipAddrFamily int

const (
	ipAddrFamilyUnknown ipAddrFamily = iota
	ipAddrFamilyV4
	ipAddrFamilyV6
)

type pickfirstBuilder struct{}

func (pickfirstBuilder) Build(cc balancer.ClientConn, _ balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

func (b pickfirstBuilder) Name() string { _ = "STUB: not implemented"; return "" }

func (pickfirstBuilder) ParseConfig(js json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	_ = "STUB: not implemented"
	return *new(serviceconfig.LoadBalancingConfig), nil
}

func EnableHealthListener(state resolver.State) resolver.State {
	_ = "STUB: not implemented"
	return *new(resolver.State)
}

type pfConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`

	ShuffleAddressList bool `json:"shuffleAddressList"`
}

type scData struct {
	subConn balancer.SubConn
	addr    resolver.Address

	rawConnectivityState connectivity.State

	effectiveState              connectivity.State
	lastErr                     error
	connectionFailedInFirstPass bool
}

func (b *pickfirstBalancer) newSCData(addr resolver.Address) (*scData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type pickfirstBalancer struct {
	logger *internalgrpclog.PrefixLogger
	cc     balancer.ClientConn

	mu sync.Mutex

	state connectivity.State

	subConns              *resolver.AddressMapV2[*scData]
	addressList           addressList
	firstPass             bool
	numTF                 int
	cancelConnectionTimer func()
	healthCheckingEnabled bool
}

func (b *pickfirstBalancer) ResolverError(err error) { _ = "STUB: not implemented"; return }

func (b *pickfirstBalancer) resolverErrorLocked(err error) { _ = "STUB: not implemented"; return }

func (b *pickfirstBalancer) UpdateClientConnState(state balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *pickfirstBalancer) UpdateSubConnState(subConn balancer.SubConn, state balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (b *pickfirstBalancer) Close() { _ = "STUB: not implemented"; return }

func (b *pickfirstBalancer) ExitIdle() { _ = "STUB: not implemented"; return }

func (b *pickfirstBalancer) startFirstPassLocked() { _ = "STUB: not implemented"; return }

func (b *pickfirstBalancer) closeSubConnsLocked() { _ = "STUB: not implemented"; return }

func deDupAddresses(addrs []resolver.Address) []resolver.Address {
	_ = "STUB: not implemented"
	return nil
}

func interleaveAddresses(addrs []resolver.Address) []resolver.Address {
	_ = "STUB: not implemented"
	return nil
}

func addressFamily(address string) ipAddrFamily {
	_ = "STUB: not implemented"
	return *new(ipAddrFamily)
}

func (b *pickfirstBalancer) reconcileSubConnsLocked(newAddrs []resolver.Address) {
	_ = "STUB: not implemented"
	return
}

func (b *pickfirstBalancer) shutdownRemainingLocked(selected *scData) {
	_ = "STUB: not implemented"
	return
}

func (b *pickfirstBalancer) requestConnectionLocked() { _ = "STUB: not implemented"; return }

func (b *pickfirstBalancer) scheduleNextConnectionLocked() { _ = "STUB: not implemented"; return }

func (b *pickfirstBalancer) updateSubConnState(sd *scData, newState balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (b *pickfirstBalancer) endFirstPassIfPossibleLocked(lastErr error) {
	_ = "STUB: not implemented"
	return
}

func (b *pickfirstBalancer) isActiveSCData(sd *scData) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *pickfirstBalancer) updateSubConnHealthState(sd *scData, state balancer.SubConnState) {
	_ = "STUB: not implemented"
	return
}

func (b *pickfirstBalancer) updateBalancerState(newState balancer.State) {
	_ = "STUB: not implemented"
	return
}

func (b *pickfirstBalancer) forceUpdateConcludedStateLocked(newState balancer.State) {
	_ = "STUB: not implemented"
	return
}

type picker struct {
	result balancer.PickResult
	err    error
}

func (p *picker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

type idlePicker struct {
	exitIdle func()
}

func (i *idlePicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

type addressList struct {
	addresses []resolver.Address
	idx       int
}

func (al *addressList) isValid() bool { _ = "STUB: not implemented"; return false }

func (al *addressList) size() int { _ = "STUB: not implemented"; return 0 }

func (al *addressList) increment() bool { _ = "STUB: not implemented"; return false }

func (al *addressList) currentAddress() resolver.Address {
	_ = "STUB: not implemented"
	return *new(resolver.Address)
}

func (al *addressList) reset() { _ = "STUB: not implemented"; return }

func (al *addressList) updateAddrs(addrs []resolver.Address) { _ = "STUB: not implemented"; return }

func (al *addressList) seekTo(needle resolver.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (al *addressList) hasNext() bool { _ = "STUB: not implemented"; return false }

func equalAddressIgnoringBalAttributes(a, b *resolver.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func weightAttribute(e resolver.Endpoint) uint32 { _ = "STUB: not implemented"; return 0 }
