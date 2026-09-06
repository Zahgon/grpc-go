package ringhash

import (
	"google.golang.org/grpc/resolver"
)

type hashKeyType string

const hashKeyKey = hashKeyType("grpc.resolver.ringhash.hash_key")

func SetHashKey(endpoint resolver.Endpoint, hashKey string) resolver.Endpoint {
	_ = "STUB: not implemented"
	return *new(resolver.Endpoint)
}

func HashKey(endpoint resolver.Endpoint) string { _ = "STUB: not implemented"; return "" }
