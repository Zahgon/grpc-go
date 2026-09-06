package randomsubsetting

import (
	"encoding/json"
	"math/rand/v2"

	xxhash "github.com/cespare/xxhash/v2"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/balancer/gracefulswitch"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
	iserviceconfig "google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

const Name = "random_subsetting_experimental"

var (
	logger     = grpclog.Component(Name)
	randUint64 = rand.Uint64
)

func prefixLogger(p *subsettingBalancer) *internalgrpclog.PrefixLogger {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	balancer.Register(bb{})
}

type bb struct{}

func (bb) Build(cc balancer.ClientConn, bOpts balancer.BuildOptions) balancer.Balancer {
	_ = "STUB: not implemented"
	return *new(balancer.Balancer)
}

type lbConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`

	SubsetSize  uint32                         `json:"subsetSize,omitempty"`
	ChildPolicy *iserviceconfig.BalancerConfig `json:"childPolicy,omitempty"`
}

func (bb) ParseConfig(s json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	_ = "STUB: not implemented"
	return *new(serviceconfig.LoadBalancingConfig), nil
}

func (bb) Name() string { _ = "STUB: not implemented"; return "" }

type subsettingBalancer struct {
	*gracefulswitch.Balancer

	logger     *internalgrpclog.PrefixLogger
	cfg        *lbConfig
	hashSeed   uint64
	hashDigest *xxhash.Digest
}

func (b *subsettingBalancer) UpdateClientConnState(s balancer.ClientConnState) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *subsettingBalancer) calculateSubset(endpoints []resolver.Endpoint) []resolver.Endpoint {
	_ = "STUB: not implemented"
	return nil
}
