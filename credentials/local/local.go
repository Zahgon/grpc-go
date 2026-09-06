package local

import (
	"context"
	"net"

	"google.golang.org/grpc/credentials"
)

type info struct {
	credentials.CommonAuthInfo
}

func (info) AuthType() string { _ = "STUB: not implemented"; return "" }

func (info) ValidateAuthority(string) error { _ = "STUB: not implemented"; return nil }

type localTC struct {
	info credentials.ProtocolInfo
}

func (c *localTC) Info() credentials.ProtocolInfo {
	_ = "STUB: not implemented"
	return *new(credentials.ProtocolInfo)
}

func getSecurityLevel(network, addr string) (credentials.SecurityLevel, error) {
	_ = "STUB: not implemented"
	return *new(credentials.SecurityLevel), nil
}

func (*localTC) ClientHandshake(_ context.Context, _ string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (*localTC) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func NewCredentials() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (c *localTC) Clone() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (c *localTC) OverrideServerName(serverNameOverride string) error {
	_ = "STUB: not implemented"
	return nil
}
