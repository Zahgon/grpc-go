package rls

import (
	"errors"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/rls/internal/keys"
	estats "google.golang.org/grpc/experimental/stats"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
	rlspb "google.golang.org/grpc/internal/proto/grpc_lookup_v1"
)

var (
	errRLSThrottled = errors.New("RLS call throttled at client side")

	computeDataCacheEntrySize = dcEntrySize
)

type exitIdler interface {
	ExitIdleOne(id string)
}

type rlsPicker struct {
	kbm keys.BuilderMap

	origEndpoint string

	lb *rlsBalancer

	rlsServerTarget string
	grpcTarget      string
	metricsRecorder estats.MetricsRecorder
	defaultPolicy   *childPolicyWrapper
	ctrlCh          *controlChannel
	maxAge          time.Duration
	staleAge        time.Duration
	bg              exitIdler
	logger          *internalgrpclog.PrefixLogger
}

func isFullMethodNameValid(name string) bool { _ = "STUB: not implemented"; return false }

func (p *rlsPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

func errToPickResult(err error) string { _ = "STUB: not implemented"; return "" }

func (p *rlsPicker) delegateToChildPoliciesLocked(dcEntry *cacheEntry, info balancer.PickInfo) (balancer.PickResult, func(), error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil, nil
}

func (p *rlsPicker) useDefaultPickIfPossible(info balancer.PickInfo, errOnNoDefault error) (balancer.PickResult, func(), error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil, nil
}

func (p *rlsPicker) sendRouteLookupRequestLocked(cacheKey cacheKey, bs *backoffState, reqKeys map[string]string, reason rlspb.RouteLookupRequest_Reason, staleHeaders string) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *rlsPicker) handleRouteLookupResponse(cacheKey cacheKey, targets []string, headerData string, err error) {
	_ = "STUB: not implemented"
	return
}

func (p *rlsPicker) setChildPolicyWrappersInCacheEntry(dcEntry *cacheEntry, newTargets []string) {
	_ = "STUB: not implemented"
	return
}

func dcEntrySize(key cacheKey, entry *cacheEntry) int64 { _ = "STUB: not implemented"; return 0 }
