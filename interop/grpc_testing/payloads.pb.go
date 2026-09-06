package grpc_testing

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ByteBufferParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReqSize       int32                  `protobuf:"varint,1,opt,name=req_size,json=reqSize,proto3" json:"req_size,omitempty"`
	RespSize      int32                  `protobuf:"varint,2,opt,name=resp_size,json=respSize,proto3" json:"resp_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ByteBufferParams) Reset() { _ = "STUB: not implemented"; return }

func (x *ByteBufferParams) String() string { _ = "STUB: not implemented"; return "" }

func (*ByteBufferParams) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ByteBufferParams) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ByteBufferParams) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ByteBufferParams) GetReqSize() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ByteBufferParams) GetRespSize() int32 { _ = "STUB: not implemented"; return 0 }

type SimpleProtoParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReqSize       int32                  `protobuf:"varint,1,opt,name=req_size,json=reqSize,proto3" json:"req_size,omitempty"`
	RespSize      int32                  `protobuf:"varint,2,opt,name=resp_size,json=respSize,proto3" json:"resp_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SimpleProtoParams) Reset() { _ = "STUB: not implemented"; return }

func (x *SimpleProtoParams) String() string { _ = "STUB: not implemented"; return "" }

func (*SimpleProtoParams) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SimpleProtoParams) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SimpleProtoParams) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SimpleProtoParams) GetReqSize() int32 { _ = "STUB: not implemented"; return 0 }

func (x *SimpleProtoParams) GetRespSize() int32 { _ = "STUB: not implemented"; return 0 }

type ComplexProtoParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ComplexProtoParams) Reset() { _ = "STUB: not implemented"; return }

func (x *ComplexProtoParams) String() string { _ = "STUB: not implemented"; return "" }

func (*ComplexProtoParams) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ComplexProtoParams) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ComplexProtoParams) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type PayloadConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Payload       isPayloadConfig_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PayloadConfig) Reset() { _ = "STUB: not implemented"; return }

func (x *PayloadConfig) String() string { _ = "STUB: not implemented"; return "" }

func (*PayloadConfig) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *PayloadConfig) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*PayloadConfig) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *PayloadConfig) GetPayload() isPayloadConfig_Payload {
	_ = "STUB: not implemented"
	return *new(isPayloadConfig_Payload)
}

func (x *PayloadConfig) GetBytebufParams() *ByteBufferParams { _ = "STUB: not implemented"; return nil }

func (x *PayloadConfig) GetSimpleParams() *SimpleProtoParams { _ = "STUB: not implemented"; return nil }

func (x *PayloadConfig) GetComplexParams() *ComplexProtoParams {
	_ = "STUB: not implemented"
	return nil
}

type isPayloadConfig_Payload interface {
	isPayloadConfig_Payload()
}

type PayloadConfig_BytebufParams struct {
	BytebufParams *ByteBufferParams `protobuf:"bytes,1,opt,name=bytebuf_params,json=bytebufParams,proto3,oneof"`
}

type PayloadConfig_SimpleParams struct {
	SimpleParams *SimpleProtoParams `protobuf:"bytes,2,opt,name=simple_params,json=simpleParams,proto3,oneof"`
}

type PayloadConfig_ComplexParams struct {
	ComplexParams *ComplexProtoParams `protobuf:"bytes,3,opt,name=complex_params,json=complexParams,proto3,oneof"`
}

func (*PayloadConfig_BytebufParams) isPayloadConfig_Payload() { _ = "STUB: not implemented"; return }

func (*PayloadConfig_SimpleParams) isPayloadConfig_Payload() { _ = "STUB: not implemented"; return }

func (*PayloadConfig_ComplexParams) isPayloadConfig_Payload() { _ = "STUB: not implemented"; return }

var File_grpc_testing_payloads_proto protoreflect.FileDescriptor

const file_grpc_testing_payloads_proto_rawDesc = "" +
	"\n" +
	"\x1bgrpc/testing/payloads.proto\x12\fgrpc.testing\"J\n" +
	"\x10ByteBufferParams\x12\x19\n" +
	"\breq_size\x18\x01 \x01(\x05R\areqSize\x12\x1b\n" +
	"\tresp_size\x18\x02 \x01(\x05R\brespSize\"K\n" +
	"\x11SimpleProtoParams\x12\x19\n" +
	"\breq_size\x18\x01 \x01(\x05R\areqSize\x12\x1b\n" +
	"\tresp_size\x18\x02 \x01(\x05R\brespSize\"\x14\n" +
	"\x12ComplexProtoParams\"\xf6\x01\n" +
	"\rPayloadConfig\x12G\n" +
	"\x0ebytebuf_params\x18\x01 \x01(\v2\x1e.grpc.testing.ByteBufferParamsH\x00R\rbytebufParams\x12F\n" +
	"\rsimple_params\x18\x02 \x01(\v2\x1f.grpc.testing.SimpleProtoParamsH\x00R\fsimpleParams\x12I\n" +
	"\x0ecomplex_params\x18\x03 \x01(\v2 .grpc.testing.ComplexProtoParamsH\x00R\rcomplexParamsB\t\n" +
	"\apayloadB\"\n" +
	"\x0fio.grpc.testingB\rPayloadsProtoP\x01b\x06proto3"

var (
	file_grpc_testing_payloads_proto_rawDescOnce sync.Once
	file_grpc_testing_payloads_proto_rawDescData []byte
)

func file_grpc_testing_payloads_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_grpc_testing_payloads_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_testing_payloads_proto_goTypes = []any{
	(*ByteBufferParams)(nil),
	(*SimpleProtoParams)(nil),
	(*ComplexProtoParams)(nil),
	(*PayloadConfig)(nil),
}
var file_grpc_testing_payloads_proto_depIdxs = []int32{
	0,
	1,
	2,
	3,
	3,
	3,
	3,
	0,
}

func init()                                  { file_grpc_testing_payloads_proto_init() }
func file_grpc_testing_payloads_proto_init() { _ = "STUB: not implemented"; return }
