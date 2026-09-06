package oauth

import (
	"context"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/grpc/credentials"
)

type TokenSource struct {
	oauth2.TokenSource
}

func (ts TokenSource) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ts TokenSource) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }

func removeServiceNameFromJWTURI(uri string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type jwtAccess struct {
	jsonKey []byte
}

func NewJWTAccessFromFile(keyFile string) (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

func NewJWTAccessFromKey(jsonKey []byte) (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

func (j jwtAccess) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jwtAccess) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }

type oauthAccess struct {
	token oauth2.Token
}

func NewOauthAccess(token *oauth2.Token) credentials.PerRPCCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials)
}

func (oa oauthAccess) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (oa oauthAccess) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }

func NewComputeEngine() credentials.PerRPCCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials)
}

type serviceAccount struct {
	mu     sync.Mutex
	config *jwt.Config
	t      *oauth2.Token
}

func (s *serviceAccount) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serviceAccount) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }

func NewServiceAccountFromKey(jsonKey []byte, scope ...string) (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

func NewServiceAccountFromFile(keyFile string, scope ...string) (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

func NewApplicationDefault(ctx context.Context, scope ...string) (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}
