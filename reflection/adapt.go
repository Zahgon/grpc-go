package reflection

import (
	v1reflectiongrpc "google.golang.org/grpc/reflection/grpc_reflection_v1"
	v1reflectionpb "google.golang.org/grpc/reflection/grpc_reflection_v1"
	v1alphareflectiongrpc "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

func asV1Alpha(svr v1reflectiongrpc.ServerReflectionServer) v1alphareflectiongrpc.ServerReflectionServer {
	_ = "STUB: not implemented"
	return *new(v1alphareflectiongrpc.ServerReflectionServer)
}

type v1AlphaServerImpl struct {
	svr v1reflectiongrpc.ServerReflectionServer
}

func (s v1AlphaServerImpl) ServerReflectionInfo(stream v1alphareflectiongrpc.ServerReflection_ServerReflectionInfoServer) error {
	_ = "STUB: not implemented"
	return nil
}

type v1AlphaServerStreamAdapter struct {
	v1alphareflectiongrpc.ServerReflection_ServerReflectionInfoServer
}

func (s v1AlphaServerStreamAdapter) Send(response *v1reflectionpb.ServerReflectionResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func (s v1AlphaServerStreamAdapter) Recv() (*v1reflectionpb.ServerReflectionRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
