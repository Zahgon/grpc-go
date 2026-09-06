package jwt

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/backoff"
)

const preemptiveRefreshThreshold = time.Minute

type jwtTokenFileCallCreds struct {
	fileReader      *jwtFileReader
	backoffStrategy backoff.Strategy

	mu               sync.Mutex
	cachedAuthHeader string
	cachedExpiry     time.Time
	cachedError      error
	retryAttempt     int
	nextRetryTime    time.Time
	pendingRefresh   bool
}

func NewTokenFileCallCredentials(tokenFilePath string) (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

func (c *jwtTokenFileCallCreds) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *jwtTokenFileCallCreds) RequireTransportSecurity() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *jwtTokenFileCallCreds) isTokenValidLocked() bool { _ = "STUB: not implemented"; return false }

func (c *jwtTokenFileCallCreds) refreshToken() { _ = "STUB: not implemented"; return }

func (c *jwtTokenFileCallCreds) updateCacheLocked(token string, expiry time.Time, err error) {
	_ = "STUB: not implemented"
	return
}
