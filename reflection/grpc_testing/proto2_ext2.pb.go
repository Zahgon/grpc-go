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

type AnotherExtension struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Whatchamacallit *int32                 `protobuf:"varint,1,opt,name=whatchamacallit" json:"whatchamacallit,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AnotherExtension) Reset() { _ = "STUB: not implemented"; return }

func (x *AnotherExtension) String() string { _ = "STUB: not implemented"; return "" }

func (*AnotherExtension) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *AnotherExtension) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*AnotherExtension) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *AnotherExtension) GetWhatchamacallit() int32 { _ = "STUB: not implemented"; return 0 }

var file_reflection_grpc_testing_proto2_ext2_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*ToBeExtended)(nil),
		ExtensionType: (*string)(nil),
		Field:         23,
		Name:          "grpc.testing.frob",
		Tag:           "bytes,23,opt,name=frob",
		Filename:      "reflection/grpc_testing/proto2_ext2.proto",
	},
	{
		ExtendedType:  (*ToBeExtended)(nil),
		ExtensionType: (*AnotherExtension)(nil),
		Field:         29,
		Name:          "grpc.testing.nitz",
		Tag:           "bytes,29,opt,name=nitz",
		Filename:      "reflection/grpc_testing/proto2_ext2.proto",
	},
}

var (
	E_Frob = &file_reflection_grpc_testing_proto2_ext2_proto_extTypes[0]

	E_Nitz = &file_reflection_grpc_testing_proto2_ext2_proto_extTypes[1]
)

var File_reflection_grpc_testing_proto2_ext2_proto protoreflect.FileDescriptor

const file_reflection_grpc_testing_proto2_ext2_proto_rawDesc = "" +
	"\n" +
	")reflection/grpc_testing/proto2_ext2.proto\x12\fgrpc.testing\x1a$reflection/grpc_testing/proto2.proto\"<\n" +
	"\x10AnotherExtension\x12(\n" +
	"\x0fwhatchamacallit\x18\x01 \x01(\x05R\x0fwhatchamacallit:.\n" +
	"\x04frob\x12\x1a.grpc.testing.ToBeExtended\x18\x17 \x01(\tR\x04frob:N\n" +
	"\x04nitz\x12\x1a.grpc.testing.ToBeExtended\x18\x1d \x01(\v2\x1e.grpc.testing.AnotherExtensionR\x04nitzB0Z.google.golang.org/grpc/reflection/grpc_testing"

var (
	file_reflection_grpc_testing_proto2_ext2_proto_rawDescOnce sync.Once
	file_reflection_grpc_testing_proto2_ext2_proto_rawDescData []byte
)

func file_reflection_grpc_testing_proto2_ext2_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_reflection_grpc_testing_proto2_ext2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reflection_grpc_testing_proto2_ext2_proto_goTypes = []any{
	(*AnotherExtension)(nil),
	(*ToBeExtended)(nil),
}
var file_reflection_grpc_testing_proto2_ext2_proto_depIdxs = []int32{
	1,
	1,
	0,
	3,
	3,
	2,
	0,
	0,
}

func init()                                                { file_reflection_grpc_testing_proto2_ext2_proto_init() }
func file_reflection_grpc_testing_proto2_ext2_proto_init() { _ = "STUB: not implemented"; return }
