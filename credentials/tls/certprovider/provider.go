package certprovider

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"

	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"google.golang.org/grpc/internal"
)

func init() {
	internal.GetCertificateProviderBuilder = getBuilder
}

var (
	errProviderClosed = errors.New("provider instance is closed")

	m = make(map[string]Builder)
)

func Register(b Builder) { _ = "STUB: not implemented"; return }

func getBuilder(name string) Builder { _ = "STUB: not implemented"; return *new(Builder) }

type Builder interface {
	ParseConfig(any) (*BuildableConfig, error)

	Name() string
}

type Provider interface {
	KeyMaterial(ctx context.Context) (*KeyMaterial, error)

	Close()
}

type KeyMaterial struct {
	Certs []tls.Certificate

	Roots *x509.CertPool

	SPIFFEBundleMap map[string]*spiffebundle.Bundle
}

type BuildOptions struct {
	CertName string

	WantRoot bool

	WantIdentity bool
}
