package xds

import (
	"context"
	"net"

	"google.golang.org/grpc/credentials"
)

type ClientOptions struct {
	FallbackCreds credentials.TransportCredentials
}

func NewClientCredentials(opts ClientOptions) (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials), nil
}

type ServerOptions struct {
	FallbackCreds credentials.TransportCredentials
}

func NewServerCredentials(opts ServerOptions) (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials), nil
}

type credsImpl struct {
	isClient bool
	fallback credentials.TransportCredentials
}

func (c *credsImpl) ClientHandshake(ctx context.Context, authority string, rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c *credsImpl) ServerHandshake(rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c *credsImpl) Info() credentials.ProtocolInfo {
	_ = "STUB: not implemented"
	return *new(credentials.ProtocolInfo)
}

func (c *credsImpl) Clone() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (c *credsImpl) OverrideServerName(_ string) error { _ = "STUB: not implemented"; return nil }

func (c *credsImpl) UsesXDS() bool { _ = "STUB: not implemented"; return false }
