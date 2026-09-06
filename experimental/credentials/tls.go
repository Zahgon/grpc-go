package credentials

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"

	"google.golang.org/grpc/credentials"
)

type tlsCreds struct {
	config *tls.Config
}

func (c tlsCreds) Info() credentials.ProtocolInfo {
	_ = "STUB: not implemented"
	return *new(credentials.ProtocolInfo)
}

func (c *tlsCreds) ClientHandshake(ctx context.Context, authority string, rawConn net.Conn) (_ net.Conn, _ credentials.AuthInfo, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c *tlsCreds) ServerHandshake(rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c *tlsCreds) Clone() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (c *tlsCreds) OverrideServerName(serverNameOverride string) error {
	_ = "STUB: not implemented"
	return nil
}

var tls12ForbiddenCipherSuites = map[uint16]struct{}{
	tls.TLS_RSA_WITH_AES_128_CBC_SHA:         {},
	tls.TLS_RSA_WITH_AES_256_CBC_SHA:         {},
	tls.TLS_RSA_WITH_AES_128_GCM_SHA256:      {},
	tls.TLS_RSA_WITH_AES_256_GCM_SHA384:      {},
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA: {},
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA: {},
	tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:   {},
	tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:   {},
}

func NewTLSWithALPNDisabled(c *tls.Config) credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func applyDefaults(c *tls.Config) *tls.Config { _ = "STUB: not implemented"; return nil }

func NewClientTLSFromCertWithALPNDisabled(cp *x509.CertPool, serverNameOverride string) credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func NewClientTLSFromFileWithALPNDisabled(certFile, serverNameOverride string) (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials), nil
}

func NewServerTLSFromCertWithALPNDisabled(cert *tls.Certificate) credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func NewServerTLSFromFileWithALPNDisabled(certFile, keyFile string) (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials), nil
}

func cloneTLSConfig(cfg *tls.Config) *tls.Config { _ = "STUB: not implemented"; return nil }

func appendH2ToNextProtos(ps []string) []string { _ = "STUB: not implemented"; return nil }
