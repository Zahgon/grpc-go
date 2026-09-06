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

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() { _ = "STUB: not implemented"; return }

func (x *Empty) String() string { _ = "STUB: not implemented"; return "" }

func (*Empty) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Empty) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Empty) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

var File_grpc_testing_empty_proto protoreflect.FileDescriptor

const file_grpc_testing_empty_proto_rawDesc = "" +
	"\n" +
	"\x18grpc/testing/empty.proto\x12\fgrpc.testing\"\a\n" +
	"\x05EmptyB*\n" +
	"\x1bio.grpc.testing.integrationB\vEmptyProtosb\x06proto3"

var (
	file_grpc_testing_empty_proto_rawDescOnce sync.Once
	file_grpc_testing_empty_proto_rawDescData []byte
)

func file_grpc_testing_empty_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_grpc_testing_empty_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_testing_empty_proto_goTypes = []any{
	(*Empty)(nil),
}
var file_grpc_testing_empty_proto_depIdxs = []int32{
	0,
	0,
	0,
	0,
	0,
}

func init()                               { file_grpc_testing_empty_proto_init() }
func file_grpc_testing_empty_proto_init() { _ = "STUB: not implemented"; return }
