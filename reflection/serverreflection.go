package reflection

import (
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	v1reflectiongrpc "google.golang.org/grpc/reflection/grpc_reflection_v1"
	v1alphareflectiongrpc "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type GRPCServer interface {
	grpc.ServiceRegistrar
	ServiceInfoProvider
}

var _ GRPCServer = (*grpc.Server)(nil)

func Register(s GRPCServer) { _ = "STUB: not implemented"; return }

func RegisterV1(s GRPCServer) { _ = "STUB: not implemented"; return }

type ServiceInfoProvider interface {
	GetServiceInfo() map[string]grpc.ServiceInfo
}

type ExtensionResolver interface {
	protoregistry.ExtensionTypeResolver
	RangeExtensionsByMessage(message protoreflect.FullName, f func(protoreflect.ExtensionType) bool)
}

type ServerOptions struct {
	Services ServiceInfoProvider

	DescriptorResolver protodesc.Resolver

	ExtensionResolver ExtensionResolver
}

func NewServer(opts ServerOptions) v1alphareflectiongrpc.ServerReflectionServer {
	_ = "STUB: not implemented"
	return *new(v1alphareflectiongrpc.ServerReflectionServer)
}

func NewServerV1(opts ServerOptions) v1reflectiongrpc.ServerReflectionServer {
	_ = "STUB: not implemented"
	return *new(v1reflectiongrpc.ServerReflectionServer)
}
