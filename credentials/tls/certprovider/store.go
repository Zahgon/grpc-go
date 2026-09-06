package certprovider

import (
	"context"
	"sync"
	"sync/atomic"
)

var provStore = &store{
	providers: make(map[storeKey]*wrappedProvider),
}

type storeKey struct {
	name string

	config string

	opts BuildOptions
}

type wrappedProvider struct {
	Provider
	refCount int

	storeKey storeKey
	store    *store
}

type closedProvider struct{}

func (c closedProvider) KeyMaterial(context.Context) (*KeyMaterial, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c closedProvider) Close() { _ = "STUB: not implemented"; return }

type singleCloseWrappedProvider struct {
	provider atomic.Pointer[Provider]
}

type store struct {
	mu        sync.Mutex
	providers map[storeKey]*wrappedProvider
}

func (wp *wrappedProvider) Close() { _ = "STUB: not implemented"; return }

func (w *singleCloseWrappedProvider) Close() { _ = "STUB: not implemented"; return }

func (w *singleCloseWrappedProvider) KeyMaterial(ctx context.Context) (*KeyMaterial, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newSingleCloseWrappedProvider(provider Provider) *singleCloseWrappedProvider {
	_ = "STUB: not implemented"
	return nil
}

type BuildableConfig struct {
	name    string
	config  []byte
	starter func(BuildOptions) Provider
	pStore  *store
}

func NewBuildableConfig(name string, config []byte, starter func(BuildOptions) Provider) *BuildableConfig {
	_ = "STUB: not implemented"
	return nil
}

func (bc *BuildableConfig) Build(opts BuildOptions) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}

func (bc *BuildableConfig) String() string { _ = "STUB: not implemented"; return "" }

func ParseConfig(name string, config any) (*BuildableConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetProvider(name string, config any, opts BuildOptions) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}
