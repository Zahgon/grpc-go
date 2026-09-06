package advancedtls

import (
	"crypto/x509"
	"sync"
	"time"
)

const defaultCRLRefreshDuration = 1 * time.Hour
const minCRLRefreshDuration = 1 * time.Minute

type CRLProvider interface {
	CRL(cert *x509.Certificate) (*CRL, error)
}

type StaticCRLProvider struct {
	crls map[string]*CRL
}

func NewStaticCRLProvider(rawCRLs [][]byte) *StaticCRLProvider {
	_ = "STUB: not implemented"
	return nil
}

func (p *StaticCRLProvider) addCRL(crl *CRL) { _ = "STUB: not implemented"; return }

func (p *StaticCRLProvider) CRL(cert *x509.Certificate) (*CRL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FileWatcherOptions struct {
	CRLDirectory               string
	RefreshDuration            time.Duration
	CRLReloadingFailedCallback func(err error)
}

type FileWatcherCRLProvider struct {
	crls map[string]*CRL
	opts FileWatcherOptions
	mu   sync.Mutex
	stop chan struct{}
	done chan struct{}
}

func NewFileWatcherCRLProvider(o FileWatcherOptions) (*FileWatcherCRLProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *FileWatcherOptions) validate() error { _ = "STUB: not implemented"; return nil }

func (p *FileWatcherCRLProvider) run() { _ = "STUB: not implemented"; return }

func (p *FileWatcherCRLProvider) Close() { _ = "STUB: not implemented"; return }

func (p *FileWatcherCRLProvider) scanCRLDirectory() { _ = "STUB: not implemented"; return }

func (p *FileWatcherCRLProvider) CRL(cert *x509.Certificate) (*CRL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
