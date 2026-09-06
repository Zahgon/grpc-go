package alts

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc/credentials"
	core "google.golang.org/grpc/credentials/alts/internal"
	altspb "google.golang.org/grpc/credentials/alts/internal/proto/grpc_gcp"
	"google.golang.org/grpc/grpclog"
)

const (
	hypervisorHandshakerServiceAddress = "dns:///metadata.google.internal.:8080"

	defaultTimeout = 30.0 * time.Second

	protocolVersionMaxMajor = 2
	protocolVersionMaxMinor = 1
	protocolVersionMinMajor = 2
	protocolVersionMinMinor = 1
)

var (
	vmOnGCP       bool
	once          sync.Once
	maxRPCVersion = &altspb.RpcProtocolVersions_Version{
		Major: protocolVersionMaxMajor,
		Minor: protocolVersionMaxMinor,
	}
	minRPCVersion = &altspb.RpcProtocolVersions_Version{
		Major: protocolVersionMinMajor,
		Minor: protocolVersionMinMinor,
	}

	ErrUntrustedPlatform = errors.New("ALTS: untrusted platform. ALTS is only supported on GCP")
	logger               = grpclog.Component("alts")
)

type AuthInfo interface {
	ApplicationProtocol() string

	RecordProtocol() string

	SecurityLevel() altspb.SecurityLevel

	PeerServiceAccount() string

	LocalServiceAccount() string

	PeerRPCVersions() *altspb.RpcProtocolVersions
}

type ClientOptions struct {
	TargetServiceAccounts []string

	HandshakerServiceAddress string
}

func DefaultClientOptions() *ClientOptions { _ = "STUB: not implemented"; return nil }

type ServerOptions struct {
	HandshakerServiceAddress string
}

func DefaultServerOptions() *ServerOptions { _ = "STUB: not implemented"; return nil }

type altsTC struct {
	info             *credentials.ProtocolInfo
	side             core.Side
	accounts         []string
	hsAddress        string
	boundAccessToken string
}

func NewClientCreds(opts *ClientOptions) credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func NewServerCreds(opts *ServerOptions) credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func newALTS(side core.Side, accounts []string, hsAddress string) credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (g *altsTC) ClientHandshake(ctx context.Context, addr string, rawConn net.Conn) (_ net.Conn, _ credentials.AuthInfo, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (g *altsTC) ServerHandshake(rawConn net.Conn) (_ net.Conn, _ credentials.AuthInfo, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (g *altsTC) Info() credentials.ProtocolInfo {
	_ = "STUB: not implemented"
	return *new(credentials.ProtocolInfo)
}

func (g *altsTC) Clone() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (g *altsTC) OverrideServerName(serverNameOverride string) error {
	_ = "STUB: not implemented"
	return nil
}

func compareRPCVersions(v1, v2 *altspb.RpcProtocolVersions_Version) int {
	_ = "STUB: not implemented"
	return 0
}

func checkRPCVersions(local, peer *altspb.RpcProtocolVersions) (bool, *altspb.RpcProtocolVersions_Version) {
	_ = "STUB: not implemented"
	return false, nil
}
