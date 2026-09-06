package google

import (
	"context"
	"sync"
	"time"

	"cloud.google.com/go/auth"
	"cloud.google.com/go/auth/credentials/idtoken"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/google/internal"
	"google.golang.org/grpc/internal/backoff"
)

const (
	preemptiveRefresh = 1 * time.Minute

	metadataTimeout = 60 * time.Second
)

type gcpServiceAccountIdentityCallCreds struct {
	ctx      context.Context
	audience string
	creds    *auth.Credentials
	backoff  backoff.Strategy

	mu                     sync.Mutex
	token                  *auth.Token
	tokenExpiry            time.Time
	preemptiveTokenRefresh time.Time
	fetching               chan struct{}
	nextRetryTime          time.Time
	retryAttempt           int
	lastErr                error
}

func init() {
	internal.BackoffStrategy = backoff.DefaultExponential
	internal.NewIDTokenCredentials = func(opts *idtoken.Options) (*auth.Credentials, error) {
		return idtoken.NewCredentials(opts)
	}
}

func NewServiceAccountIdentityCredentials(ctx context.Context, audience string) (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

func (c *gcpServiceAccountIdentityCallCreds) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *gcpServiceAccountIdentityCallCreds) cachedRequestMetadata(attemptPreemptiveRefresh bool) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *gcpServiceAccountIdentityCallCreds) cachedRequestMetadataLocked(attemptPreemptiveRefresh bool) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *gcpServiceAccountIdentityCallCreds) RequireTransportSecurity() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *gcpServiceAccountIdentityCallCreds) isTokenStaleLocked() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *gcpServiceAccountIdentityCallCreds) isTokenValidLocked() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *gcpServiceAccountIdentityCallCreds) startFetch() { _ = "STUB: not implemented"; return }

func (c *gcpServiceAccountIdentityCallCreds) updateStateLocked(token *auth.Token, err error) {
	_ = "STUB: not implemented"
	return
}
