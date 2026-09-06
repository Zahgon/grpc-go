package test

import (
	"testing"

	"google.golang.org/grpc/resolver/manual"
	"google.golang.org/grpc/serviceconfig"
)

func parseServiceConfig(t *testing.T, r *manual.Resolver, sc string) *serviceconfig.ParseResult {
	_ = "STUB: not implemented"
	return nil
}
