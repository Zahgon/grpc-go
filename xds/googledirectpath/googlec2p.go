package googledirectpath

import (
	rand "math/rand/v2"
	"sync"
	"time"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/googlecloud"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/internal/xds/xdsclient"
	"google.golang.org/grpc/resolver"

	_ "google.golang.org/grpc/xds"
)

const (
	c2pScheme    = "google-c2p"
	c2pAuthority = "traffic-director-c2p.xds.googleapis.com"

	defaultUniverseDomain   = "googleapis.com"
	zoneURL                 = "http://metadata.google.internal/computeMetadata/v1/instance/zone"
	ipv6URL                 = "http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ipv6s"
	ipv6CapableMetadataName = "TRAFFICDIRECTOR_DIRECTPATH_C2P_IPV6_CAPABLE"
	httpReqTimeout          = 10 * time.Second

	logPrefix        = "[google-c2p-resolver]"
	dnsName, xdsName = "dns", "xds"
)

var (
	logger           = internalgrpclog.NewPrefixLogger(grpclog.Component("directpath"), logPrefix)
	universeDomainMu sync.Mutex
	universeDomain   = ""

	onGCE         = googlecloud.OnGCE
	randInt       = rand.Int
	xdsClientPool = xdsclient.DefaultPool
)

func init() {
	resolver.Register(c2pResolverBuilder{})
}

func SetUniverseDomain(domain string) error { _ = "STUB: not implemented"; return nil }

func getXdsServerURI() string { _ = "STUB: not implemented"; return "" }

type c2pResolverWrapper struct {
	resolver.Resolver
	cancel func()
}

func (r *c2pResolverWrapper) Close() { _ = "STUB: not implemented"; return }

type c2pResolverBuilder struct{}

func (c2pResolverBuilder) Build(t resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Resolver), nil
}

func (b c2pResolverBuilder) Scheme() string { _ = "STUB: not implemented"; return "" }

func newNodeConfig(zone string, ipv6Capable bool, forceXDS bool) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func newAuthoritiesConfig(serverCfg map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func newXdsServerConfig(uri string) map[string]any { _ = "STUB: not implemented"; return nil }
