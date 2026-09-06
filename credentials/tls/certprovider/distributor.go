package certprovider

import (
	"context"
	"sync"

	"google.golang.org/grpc/internal/grpcsync"
)

type Distributor struct {
	mu   sync.Mutex
	km   *KeyMaterial
	pErr error

	ready *grpcsync.Event

	closed *grpcsync.Event
}

func NewDistributor() *Distributor { _ = "STUB: not implemented"; return nil }

func (d *Distributor) Set(km *KeyMaterial, err error) { _ = "STUB: not implemented"; return }

func (d *Distributor) KeyMaterial(ctx context.Context) (*KeyMaterial, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Distributor) keyMaterial() (*KeyMaterial, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Distributor) Stop() { _ = "STUB: not implemented"; return }
