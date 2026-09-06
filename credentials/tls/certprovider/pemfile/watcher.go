package pemfile

import (
	"context"
	"time"

	"google.golang.org/grpc/credentials/tls/certprovider"
	"google.golang.org/grpc/grpclog"
)

const defaultCertRefreshDuration = 1 * time.Hour

var (
	newDistributor = func() distributor { return certprovider.NewDistributor() }

	logger = grpclog.Component("pemfile")
)

type Options struct {
	CertFile string

	KeyFile string

	RootFile string

	SPIFFEBundleMapFile string

	RefreshDuration time.Duration
}

func (o Options) canonical() []byte { _ = "STUB: not implemented"; return nil }

func (o Options) validate() error { _ = "STUB: not implemented"; return nil }

func NewProvider(o Options) (certprovider.Provider, error) {
	_ = "STUB: not implemented"
	return *new(certprovider.Provider), nil
}

func newProvider(o Options) certprovider.Provider {
	_ = "STUB: not implemented"
	return *new(certprovider.Provider)
}

type watcher struct {
	identityDistributor         distributor
	rootDistributor             distributor
	opts                        Options
	certFileContents            []byte
	keyFileContents             []byte
	rootFileContents            []byte
	spiffeBundleMapFileContents []byte
	cancel                      context.CancelFunc
}

type distributor interface {
	KeyMaterial(ctx context.Context) (*certprovider.KeyMaterial, error)
	Set(km *certprovider.KeyMaterial, err error)
	Stop()
}

func (w *watcher) updateIdentityDistributor() { _ = "STUB: not implemented"; return }

func (w *watcher) updateRootDistributor() { _ = "STUB: not implemented"; return }

func (w *watcher) maybeUpdateSPIFFEBundleMap() { _ = "STUB: not implemented"; return }

func (w *watcher) maybeUpdateRootFile() { _ = "STUB: not implemented"; return }

func (w *watcher) run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (w *watcher) KeyMaterial(ctx context.Context) (*certprovider.KeyMaterial, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) Close() { _ = "STUB: not implemented"; return }
