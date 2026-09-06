package rls

import (
	"container/list"
	"time"

	"google.golang.org/grpc/internal/backoff"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/internal/grpcsync"
)

type cacheKey struct {
	path string

	keys string
}

type cacheEntry struct {
	childPolicyWrappers []*childPolicyWrapper

	headerData string

	expiryTime time.Time

	staleTime time.Time

	earliestEvictTime time.Time

	status error

	backoffState *backoffState

	backoffTime time.Time

	backoffExpiryTime time.Time

	size int64
}

type backoffState struct {
	retries int

	bs backoff.Strategy

	timer *time.Timer
}

type lru struct {
	ll *list.List

	m map[cacheKey]*list.Element
}

func newLRU() *lru { _ = "STUB: not implemented"; return nil }

func (l *lru) addEntry(key cacheKey) { _ = "STUB: not implemented"; return }

func (l *lru) makeRecent(key cacheKey) { _ = "STUB: not implemented"; return }

func (l *lru) removeEntry(key cacheKey) { _ = "STUB: not implemented"; return }

func (l *lru) getLeastRecentlyUsed() cacheKey { _ = "STUB: not implemented"; return *new(cacheKey) }

type dataCache struct {
	maxSize         int64
	currentSize     int64
	keys            *lru
	entries         map[cacheKey]*cacheEntry
	logger          *internalgrpclog.PrefixLogger
	shutdown        *grpcsync.Event
	rlsServerTarget string

	grpcTarget string
	uuid       string
}

func newDataCache(size int64, logger *internalgrpclog.PrefixLogger, grpcTarget string) *dataCache {
	_ = "STUB: not implemented"
	return nil
}

func (dc *dataCache) updateRLSServerTarget(rlsServerTarget string) {
	_ = "STUB: not implemented"
	return
}

func (dc *dataCache) resize(size int64) (backoffCancelled bool) {
	_ = "STUB: not implemented"
	return false
}

func (dc *dataCache) evictExpiredEntries() bool { _ = "STUB: not implemented"; return false }

func (dc *dataCache) resetBackoffState(newBackoffState *backoffState) bool {
	_ = "STUB: not implemented"
	return false
}

func (dc *dataCache) addEntry(key cacheKey, entry *cacheEntry) (backoffCancelled bool, ok bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (dc *dataCache) updateEntrySize(entry *cacheEntry, newSize int64) {
	_ = "STUB: not implemented"
	return
}

func (dc *dataCache) getEntry(key cacheKey) *cacheEntry { _ = "STUB: not implemented"; return nil }

func (dc *dataCache) removeEntryForTesting(key cacheKey) { _ = "STUB: not implemented"; return }

func (dc *dataCache) deleteAndCleanup(key cacheKey, entry *cacheEntry) {
	_ = "STUB: not implemented"
	return
}

func (dc *dataCache) stop() { _ = "STUB: not implemented"; return }
