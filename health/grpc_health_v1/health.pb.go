package grpc_health_v1

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN         HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING         HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING     HealthCheckResponse_ServingStatus = 2
	HealthCheckResponse_SERVICE_UNKNOWN HealthCheckResponse_ServingStatus = 3
)

var (
	HealthCheckResponse_ServingStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "NOT_SERVING",
		3: "SERVICE_UNKNOWN",
	}
	HealthCheckResponse_ServingStatus_value = map[string]int32{
		"UNKNOWN":         0,
		"SERVING":         1,
		"NOT_SERVING":     2,
		"SERVICE_UNKNOWN": 3,
	}
)

func (x HealthCheckResponse_ServingStatus) Enum() *HealthCheckResponse_ServingStatus {
	_ = "STUB: not implemented"
	return nil
}

func (x HealthCheckResponse_ServingStatus) String() string { _ = "STUB: not implemented"; return "" }

func (HealthCheckResponse_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (HealthCheckResponse_ServingStatus) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x HealthCheckResponse_ServingStatus) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Service       string                 `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *HealthCheckRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*HealthCheckRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HealthCheckRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HealthCheckRequest) GetService() string { _ = "STUB: not implemented"; return "" }

type HealthCheckResponse struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Status        HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=grpc.health.v1.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *HealthCheckResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*HealthCheckResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	_ = "STUB: not implemented"
	return *new(HealthCheckResponse_ServingStatus)
}

type HealthListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthListRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *HealthListRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*HealthListRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HealthListRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HealthListRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type HealthListResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Statuses      map[string]*HealthCheckResponse `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthListResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *HealthListResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*HealthListResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HealthListResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HealthListResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HealthListResponse) GetStatuses() map[string]*HealthCheckResponse {
	_ = "STUB: not implemented"
	return nil
}

var File_grpc_health_v1_health_proto protoreflect.FileDescriptor

const file_grpc_health_v1_health_proto_rawDesc = "" +
	"\n" +
	"\x1bgrpc/health/v1/health.proto\x12\x0egrpc.health.v1\".\n" +
	"\x12HealthCheckRequest\x12\x18\n" +
	"\aservice\x18\x01 \x01(\tR\aservice\"\xb1\x01\n" +
	"\x13HealthCheckResponse\x12I\n" +
	"\x06status\x18\x01 \x01(\x0e21.grpc.health.v1.HealthCheckResponse.ServingStatusR\x06status\"O\n" +
	"\rServingStatus\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\v\n" +
	"\aSERVING\x10\x01\x12\x0f\n" +
	"\vNOT_SERVING\x10\x02\x12\x13\n" +
	"\x0fSERVICE_UNKNOWN\x10\x03\"\x13\n" +
	"\x11HealthListRequest\"\xc4\x01\n" +
	"\x12HealthListResponse\x12L\n" +
	"\bstatuses\x18\x01 \x03(\v20.grpc.health.v1.HealthListResponse.StatusesEntryR\bstatuses\x1a`\n" +
	"\rStatusesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x129\n" +
	"\x05value\x18\x02 \x01(\v2#.grpc.health.v1.HealthCheckResponseR\x05value:\x028\x012\xfd\x01\n" +
	"\x06Health\x12P\n" +
	"\x05Check\x12\".grpc.health.v1.HealthCheckRequest\x1a#.grpc.health.v1.HealthCheckResponse\x12M\n" +
	"\x04List\x12!.grpc.health.v1.HealthListRequest\x1a\".grpc.health.v1.HealthListResponse\x12R\n" +
	"\x05Watch\x12\".grpc.health.v1.HealthCheckRequest\x1a#.grpc.health.v1.HealthCheckResponse0\x01Bp\n" +
	"\x11io.grpc.health.v1B\vHealthProtoP\x01Z,google.golang.org/grpc/health/grpc_health_v1\xa2\x02\fGrpcHealthV1\xaa\x02\x0eGrpc.Health.V1b\x06proto3"

var (
	file_grpc_health_v1_health_proto_rawDescOnce sync.Once
	file_grpc_health_v1_health_proto_rawDescData []byte
)

func file_grpc_health_v1_health_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_grpc_health_v1_health_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_health_v1_health_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_health_v1_health_proto_goTypes = []any{
	(HealthCheckResponse_ServingStatus)(0),
	(*HealthCheckRequest)(nil),
	(*HealthCheckResponse)(nil),
	(*HealthListRequest)(nil),
	(*HealthListResponse)(nil),
	nil,
}
var file_grpc_health_v1_health_proto_depIdxs = []int32{
	0,
	5,
	2,
	1,
	3,
	1,
	2,
	4,
	2,
	6,
	3,
	3,
	3,
	0,
}

func init()                                  { file_grpc_health_v1_health_proto_init() }
func file_grpc_health_v1_health_proto_init() { _ = "STUB: not implemented"; return }
