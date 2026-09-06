package hostname

import "google.golang.org/grpc/resolver"

type hostnameKey struct{}

func Set(endpoint resolver.Endpoint, hostname string) resolver.Endpoint {
	_ = "STUB: not implemented"
	return *new(resolver.Endpoint)
}

func FromEndpoint(endpoint resolver.Endpoint) string { _ = "STUB: not implemented"; return "" }
