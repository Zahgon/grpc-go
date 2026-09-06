package sts

import (
	"context"
	"crypto/x509"
	"net/http"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	stsRequestTimeout = 5 * time.Second

	minCachedTokenLifetime = 300 * time.Second

	tokenExchangeGrantType    = "urn:ietf:params:oauth:grant-type:token-exchange"
	defaultCloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

var (
	loadSystemCertPool   = x509.SystemCertPool
	makeHTTPDoer         = makeHTTPClient
	readSubjectTokenFrom = os.ReadFile
	readActorTokenFrom   = os.ReadFile
	logger               = grpclog.Component("credentials")
)

type Options struct {
	TokenExchangeServiceURI string

	Resource string

	Audience string

	Scope string

	RequestedTokenType string

	SubjectTokenPath string

	SubjectTokenType string

	ActorTokenPath string

	ActorTokenType string
}

func (o Options) String() string { _ = "STUB: not implemented"; return "" }

func NewCredentials(opts Options) (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

type callCreds struct {
	opts   Options
	client httpDoer

	mu            sync.Mutex
	tokenMetadata map[string]string
	tokenExpiry   time.Time
}

func (c *callCreds) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *callCreds) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }

type httpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func makeHTTPClient(roots *x509.CertPool) httpDoer {
	_ = "STUB: not implemented"
	return *new(httpDoer)
}

func validateOptions(opts Options) error { _ = "STUB: not implemented"; return nil }

func (c *callCreds) cachedMetadata() map[string]string { _ = "STUB: not implemented"; return nil }

func constructRequest(ctx context.Context, opts Options) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sendRequest(client httpDoer, req *http.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tokenInfoFromResponse(respBody []byte) (*tokenInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type requestParameters struct {
	GrantType string `json:"grant_type"`

	Resource string `json:"resource,omitempty"`

	Audience string `json:"audience,omitempty"`

	Scope string `json:"scope,omitempty"`

	RequestedTokenType string `json:"requested_token_type,omitempty"`

	SubjectToken string `json:"subject_token"`

	SubjectTokenType string `json:"subject_token_type"`

	ActorToken string `json:"actor_token,omitempty"`

	ActorTokenType string `json:"actor_token_type,omitempty"`
}

type responseParameters struct {
	AccessToken string `json:"access_token"`

	IssuedTokenType string `json:"issued_token_type"`

	TokenType string `json:"token_type"`

	ExpiresIn int64 `json:"expires_in"`

	Scope string `json:"scope"`

	RefreshToken string `json:"refresh_token"`
}

type tokenInfo struct {
	tokenType  string
	token      string
	expiryTime time.Time
}
