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

type Extension struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Whatzit       *int32                 `protobuf:"varint,1,opt,name=whatzit" json:"whatzit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Extension) Reset() { _ = "STUB: not implemented"; return }

func (x *Extension) String() string { _ = "STUB: not implemented"; return "" }

func (*Extension) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Extension) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Extension) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Extension) GetWhatzit() int32 { _ = "STUB: not implemented"; return 0 }

var file_reflection_grpc_testing_proto2_ext_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*ToBeExtended)(nil),
		ExtensionType: (*int32)(nil),
		Field:         13,
		Name:          "grpc.testing.foo",
		Tag:           "varint,13,opt,name=foo",
		Filename:      "reflection/grpc_testing/proto2_ext.proto",
	},
	{
		ExtendedType:  (*ToBeExtended)(nil),
		ExtensionType: (*Extension)(nil),
		Field:         17,
		Name:          "grpc.testing.bar",
		Tag:           "bytes,17,opt,name=bar",
		Filename:      "reflection/grpc_testing/proto2_ext.proto",
	},
	{
		ExtendedType:  (*ToBeExtended)(nil),
		ExtensionType: (*SearchRequest)(nil),
		Field:         19,
		Name:          "grpc.testing.baz",
		Tag:           "bytes,19,opt,name=baz",
		Filename:      "reflection/grpc_testing/proto2_ext.proto",
	},
}

var (
	E_Foo = &file_reflection_grpc_testing_proto2_ext_proto_extTypes[0]

	E_Bar = &file_reflection_grpc_testing_proto2_ext_proto_extTypes[1]

	E_Baz = &file_reflection_grpc_testing_proto2_ext_proto_extTypes[2]
)

var File_reflection_grpc_testing_proto2_ext_proto protoreflect.FileDescriptor

const file_reflection_grpc_testing_proto2_ext_proto_rawDesc = "" +
	"\n" +
	"(reflection/grpc_testing/proto2_ext.proto\x12\fgrpc.testing\x1a$reflection/grpc_testing/proto2.proto\x1a\"reflection/grpc_testing/test.proto\"%\n" +
	"\tExtension\x12\x18\n" +
	"\awhatzit\x18\x01 \x01(\x05R\awhatzit:,\n" +
	"\x03foo\x12\x1a.grpc.testing.ToBeExtended\x18\r \x01(\x05R\x03foo:E\n" +
	"\x03bar\x12\x1a.grpc.testing.ToBeExtended\x18\x11 \x01(\v2\x17.grpc.testing.ExtensionR\x03bar:I\n" +
	"\x03baz\x12\x1a.grpc.testing.ToBeExtended\x18\x13 \x01(\v2\x1b.grpc.testing.SearchRequestR\x03bazB0Z.google.golang.org/grpc/reflection/grpc_testing"

var (
	file_reflection_grpc_testing_proto2_ext_proto_rawDescOnce sync.Once
	file_reflection_grpc_testing_proto2_ext_proto_rawDescData []byte
)

func file_reflection_grpc_testing_proto2_ext_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_reflection_grpc_testing_proto2_ext_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reflection_grpc_testing_proto2_ext_proto_goTypes = []any{
	(*Extension)(nil),
	(*ToBeExtended)(nil),
	(*SearchRequest)(nil),
}
var file_reflection_grpc_testing_proto2_ext_proto_depIdxs = []int32{
	1,
	1,
	1,
	0,
	2,
	5,
	5,
	3,
	0,
	0,
}

func init()                                               { file_reflection_grpc_testing_proto2_ext_proto_init() }
func file_reflection_grpc_testing_proto2_ext_proto_init() { _ = "STUB: not implemented"; return }
