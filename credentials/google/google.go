package google

import (
	"context"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/alts"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/grpclog"
)

const defaultCloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

var logger = grpclog.Component("credentials")

type DefaultCredentialsOptions struct {
	PerRPCCreds credentials.PerRPCCredentials

	ALTSPerRPCCreds credentials.PerRPCCredentials
}

func NewDefaultCredentialsWithOptions(opts DefaultCredentialsOptions) credentials.Bundle {
	_ = "STUB: not implemented"
	return *new(credentials.Bundle)
}

func NewDefaultCredentials() credentials.Bundle {
	_ = "STUB: not implemented"
	return *new(credentials.Bundle)
}

func NewComputeEngineCredentials() credentials.Bundle {
	_ = "STUB: not implemented"
	return *new(credentials.Bundle)
}

type creds struct {
	opts DefaultCredentialsOptions

	mode string

	transportCreds credentials.TransportCredentials

	perRPCCreds credentials.PerRPCCredentials
}

func (c *creds) TransportCredentials() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

func (c *creds) PerRPCCredentials() credentials.PerRPCCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials)
}

var (
	newTLS = func() credentials.TransportCredentials {
		return credentials.NewTLS(nil)
	}
	newALTS = func() credentials.TransportCredentials {
		return alts.NewClientCreds(alts.DefaultClientOptions())
	}
	newADC = func(ctx context.Context) (credentials.PerRPCCredentials, error) {
		return oauth.NewApplicationDefault(ctx, defaultCloudPlatformScope)
	}
)

func (c *creds) NewWithMode(mode string) (credentials.Bundle, error) {
	_ = "STUB: not implemented"
	return *new(credentials.Bundle), nil
}

type dualPerRPCCreds struct {
	perRPCCreds     credentials.PerRPCCredentials
	altsPerRPCCreds credentials.PerRPCCredentials
}

func (d *dualPerRPCCreds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dualPerRPCCreds) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }
