package credentials

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/url"

	"google.golang.org/grpc/grpclog"
)

const alpnFailureHelpMessage = "If you upgraded from a grpc-go version earlier than 1.67, your TLS connections may have stopped working due to ALPN enforcement. For more details, see: https://github.com/grpc/grpc-go/issues/434"

var logger = grpclog.Component("credentials")

type TLSInfo struct {
	State tls.ConnectionState
	CommonAuthInfo

	SPIFFEID *url.URL
}

func (t TLSInfo) AuthType() string { _ = "STUB: not implemented"; return "" }

func (t TLSInfo) ValidateAuthority(authority string) error { _ = "STUB: not implemented"; return nil }

func cipherSuiteLookup(cipherSuiteID uint16) string { _ = "STUB: not implemented"; return "" }

func (t TLSInfo) GetSecurityValue() ChannelzSecurityValue {
	_ = "STUB: not implemented"
	return *new(ChannelzSecurityValue)
}

type tlsCreds struct {
	config *tls.Config
}

func (c tlsCreds) Info() ProtocolInfo { _ = "STUB: not implemented"; return *new(ProtocolInfo) }

func (c *tlsCreds) ClientHandshake(ctx context.Context, authority string, rawConn net.Conn) (_ net.Conn, _ AuthInfo, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(AuthInfo), nil
}

func (c *tlsCreds) ServerHandshake(rawConn net.Conn) (net.Conn, AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(AuthInfo), nil
}

func (c *tlsCreds) Clone() TransportCredentials {
	_ = "STUB: not implemented"
	return *new(TransportCredentials)
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

func NewTLS(c *tls.Config) TransportCredentials {
	_ = "STUB: not implemented"
	return *new(TransportCredentials)
}

func applyDefaults(c *tls.Config) *tls.Config { _ = "STUB: not implemented"; return nil }

func NewClientTLSFromCert(cp *x509.CertPool, serverNameOverride string) TransportCredentials {
	_ = "STUB: not implemented"
	return *new(TransportCredentials)
}

func NewClientTLSFromFile(certFile, serverNameOverride string) (TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(TransportCredentials), nil
}

func NewServerTLSFromCert(cert *tls.Certificate) TransportCredentials {
	_ = "STUB: not implemented"
	return *new(TransportCredentials)
}

func NewServerTLSFromFile(certFile, keyFile string) (TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(TransportCredentials), nil
}

type TLSChannelzSecurityValue struct {
	ChannelzSecurityValue
	StandardName      string
	LocalCertificate  []byte
	RemoteCertificate []byte
}
