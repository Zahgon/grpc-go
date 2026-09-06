package advancedtls

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/tls/certprovider"
)

type CertificateChains [][]*x509.Certificate

type HandshakeVerificationInfo struct {
	ServerName string

	RawCerts [][]byte

	VerifiedChains CertificateChains

	Leaf *x509.Certificate

	ConnectionState tls.ConnectionState
}

type PostHandshakeVerificationResults struct{}

type PostHandshakeVerificationFunc func(params *HandshakeVerificationInfo) (*PostHandshakeVerificationResults, error)

type ConnectionInfo struct {
	RawConn net.Conn

	RawCerts [][]byte
}

type RootCertificates struct {
	TrustCerts *x509.CertPool
}

type RootCertificateOptions struct {
	RootCertificates *x509.CertPool

	GetRootCertificates func(params *ConnectionInfo) (*RootCertificates, error)

	RootProvider certprovider.Provider
}

func (o RootCertificateOptions) nonNilFieldCount() int { _ = "STUB: not implemented"; return 0 }

type IdentityCertificateOptions struct {
	Certificates []tls.Certificate

	GetIdentityCertificatesForClient func(*tls.CertificateRequestInfo) (*tls.Certificate, error)

	GetIdentityCertificatesForServer func(*tls.ClientHelloInfo) ([]*tls.Certificate, error)

	IdentityProvider certprovider.Provider
}

func (o IdentityCertificateOptions) nonNilFieldCount() int { _ = "STUB: not implemented"; return 0 }

type VerificationType int

const (
	CertAndHostVerification VerificationType = iota

	CertVerification

	SkipVerification
)

type Options struct {
	IdentityOptions IdentityCertificateOptions

	AdditionalPeerVerification PostHandshakeVerificationFunc

	RootOptions RootCertificateOptions

	RequireClientCert bool

	VerificationType VerificationType

	RevocationOptions *RevocationOptions

	MinTLSVersion uint16

	MaxTLSVersion uint16

	CipherSuites []uint16

	CurvePreferences []tls.CurveID

	serverNameOverride string

	SkipServerAuthEKU bool
}

func (o *Options) clientConfig() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (o *Options) serverConfig() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

type advancedTLSCreds struct {
	config              *tls.Config
	verifyFunc          PostHandshakeVerificationFunc
	getRootCertificates func(params *ConnectionInfo) (*RootCertificates, error)
	isClient            bool
	revocationOptions   *RevocationOptions
	verificationType    VerificationType
	skipServerAuthEKU   bool
}

func (c advancedTLSCreds) Info() credentials.ProtocolInfo {
	_ = "STUB: not implemented"
	return *new(credentials.ProtocolInfo)
}

func (c *advancedTLSCreds) ClientHandshake(ctx context.Context, authority string, rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c *advancedTLSCreds) ServerHandshake(rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

func (c *advancedTLSCreds) Clone() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (c *advancedTLSCreds) OverrideServerName(serverNameOverride string) error {
	_ = "STUB: not implemented"
	return nil
}

type verifyPeerCertificateFunc func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error
type verifyConnectionFunc func(cs tls.ConnectionState) error

func buildVerifyFunc(c *advancedTLSCreds,
	serverName string,
	rawConn net.Conn,
	peerVerifiedChains *CertificateChains) (verifyPeerCertificateFunc, verifyConnectionFunc) {
	_ = "STUB: not implemented"
	return *new(verifyPeerCertificateFunc), *new(verifyConnectionFunc)
}

func NewClientCreds(o *Options) (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials), nil
}

func NewServerCreds(o *Options) (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials), nil
}
