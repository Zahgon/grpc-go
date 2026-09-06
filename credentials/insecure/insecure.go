package insecure

import (
	"context"
	"net"

	"google.golang.org/grpc/credentials"
)

func NewCredentials() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

type insecureTC struct{}

func (insecureTC) ClientHandshake(_ context.Context, _ string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (insecureTC) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (insecureTC) Info() credentials.ProtocolInfo {
	_ = "STUB: not implemented"
	return *new(credentials.ProtocolInfo)
}

func (insecureTC) Clone() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (insecureTC) OverrideServerName(string) error { _ = "STUB: not implemented"; return nil }

type info struct {
	credentials.CommonAuthInfo
}

func (info) AuthType() string { _ = "STUB: not implemented"; return "" }

func (info) ValidateAuthority(string) error { _ = "STUB: not implemented"; return nil }

type insecureBundle struct{}

func NewBundle() credentials.Bundle { _ = "STUB: not implemented"; return *new(credentials.Bundle) }

func (insecureBundle) NewWithMode(string) (credentials.Bundle, error) {
	_ = "STUB: not implemented"
	return *new(credentials.Bundle), nil
}

func (insecureBundle) PerRPCCredentials() credentials.PerRPCCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials)
}

func (insecureBundle) TransportCredentials() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}
