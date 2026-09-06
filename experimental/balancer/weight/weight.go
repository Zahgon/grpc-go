package weight

import "google.golang.org/grpc/resolver"

type attributeKey struct{}

type EndpointInfo struct {
	Weight uint32
}

func (a EndpointInfo) Equal(o any) bool { _ = "STUB: not implemented"; return false }

func Set(endpoint resolver.Endpoint, epInfo EndpointInfo) resolver.Endpoint {
	_ = "STUB: not implemented"
	return *new(resolver.Endpoint)
}

func FromEndpoint(endpoint resolver.Endpoint) EndpointInfo {
	_ = "STUB: not implemented"
	return *new(EndpointInfo)
}
