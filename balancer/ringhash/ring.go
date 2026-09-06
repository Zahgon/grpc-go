package ringhash

import (
	"google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/resolver"
)

type ring struct {
	items []*ringEntry
}

type endpointInfo struct {
	hashKey        string
	scaledWeight   float64
	originalWeight uint32
}

type ringEntry struct {
	idx     int
	hash    uint64
	hashKey string
	weight  uint32
}

func newRing(endpoints *resolver.EndpointMap[*endpointState], minRingSize, maxRingSize uint64, logger *grpclog.PrefixLogger) *ring {
	_ = "STUB: not implemented"
	return nil
}

func normalizeWeights(endpoints *resolver.EndpointMap[*endpointState]) ([]endpointInfo, float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (r *ring) pick(h uint64) *ringEntry { _ = "STUB: not implemented"; return nil }

func (r *ring) next(e *ringEntry) *ringEntry { _ = "STUB: not implemented"; return nil }
