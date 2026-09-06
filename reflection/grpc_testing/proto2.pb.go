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

type ToBeExtended struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Foo             *int32                 `protobuf:"varint,1,req,name=foo" json:"foo,omitempty"`
	extensionFields protoimpl.ExtensionFields
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ToBeExtended) Reset() { _ = "STUB: not implemented"; return }

func (x *ToBeExtended) String() string { _ = "STUB: not implemented"; return "" }

func (*ToBeExtended) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ToBeExtended) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ToBeExtended) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ToBeExtended) GetFoo() int32 { _ = "STUB: not implemented"; return 0 }

var File_reflection_grpc_testing_proto2_proto protoreflect.FileDescriptor

const file_reflection_grpc_testing_proto2_proto_rawDesc = "" +
	"\n" +
	"$reflection/grpc_testing/proto2.proto\x12\fgrpc.testing\"&\n" +
	"\fToBeExtended\x12\x10\n" +
	"\x03foo\x18\x01 \x02(\x05R\x03foo*\x04\b\n" +
	"\x10\x1fB0Z.google.golang.org/grpc/reflection/grpc_testing"

var (
	file_reflection_grpc_testing_proto2_proto_rawDescOnce sync.Once
	file_reflection_grpc_testing_proto2_proto_rawDescData []byte
)

func file_reflection_grpc_testing_proto2_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_reflection_grpc_testing_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reflection_grpc_testing_proto2_proto_goTypes = []any{
	(*ToBeExtended)(nil),
}
var file_reflection_grpc_testing_proto2_proto_depIdxs = []int32{
	0,
	0,
	0,
	0,
	0,
}

func init()                                           { file_reflection_grpc_testing_proto2_proto_init() }
func file_reflection_grpc_testing_proto2_proto_init() { _ = "STUB: not implemented"; return }
