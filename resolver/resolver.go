package resolver

import (
	"context"
	"net"
	"net/url"

	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/experimental/stats"
	"google.golang.org/grpc/serviceconfig"
)

var (
	m = make(map[string]Builder)

	defaultScheme = "passthrough"
)

func Register(b Builder) { _ = "STUB: not implemented"; return }

func Get(scheme string) Builder { _ = "STUB: not implemented"; return *new(Builder) }

func SetDefaultScheme(scheme string) { _ = "STUB: not implemented"; return }

func GetDefaultScheme() string { _ = "STUB: not implemented"; return "" }

type Address struct {
	Addr string

	ServerName string

	Attributes *attributes.Attributes

	BalancerAttributes *attributes.Attributes

	Metadata any
}

func (a Address) Equal(o Address) bool { _ = "STUB: not implemented"; return false }

func (a Address) String() string { _ = "STUB: not implemented"; return "" }

type BuildOptions struct {
	DisableServiceConfig bool

	DialCreds credentials.TransportCredentials

	CredsBundle credentials.Bundle

	Dialer func(context.Context, string) (net.Conn, error)

	Authority string

	MetricsRecorder stats.MetricsRecorder
}

type Endpoint struct {
	Addresses []Address

	Attributes *attributes.Attributes
}

type State struct {
	Addresses []Address

	Endpoints []Endpoint

	ServiceConfig *serviceconfig.ParseResult

	Attributes *attributes.Attributes
}

type ClientConn interface {
	UpdateState(State) error

	ReportError(error)

	NewAddress(addresses []Address)

	ParseServiceConfig(serviceConfigJSON string) *serviceconfig.ParseResult
}

type Target struct {
	URL url.URL
}

func (t Target) Endpoint() string { _ = "STUB: not implemented"; return "" }

func (t Target) String() string { _ = "STUB: not implemented"; return "" }

type Builder interface {
	Build(target Target, cc ClientConn, opts BuildOptions) (Resolver, error)

	Scheme() string
}

type ResolveNowOptions struct{}

type Resolver interface {
	ResolveNow(ResolveNowOptions)

	Close()
}

type AuthorityOverrider interface {
	OverrideAuthority(Target) string
}

func ValidateEndpoints(endpoints []Endpoint) error { _ = "STUB: not implemented"; return nil }
