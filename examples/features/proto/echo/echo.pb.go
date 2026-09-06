package echo

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EchoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EchoRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *EchoRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*EchoRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*EchoRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *EchoRequest) GetMessage() string { _ = "STUB: not implemented"; return "" }

type EchoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EchoResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *EchoResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*EchoResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*EchoResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *EchoResponse) GetMessage() string { _ = "STUB: not implemented"; return "" }

var File_examples_features_proto_echo_echo_proto protoreflect.FileDescriptor

const file_examples_features_proto_echo_echo_proto_rawDesc = "" +
	"\n" +
	"'examples/features/proto/echo/echo.proto\x12\x12grpc.examples.echo\"'\n" +
	"\vEchoRequest\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"(\n" +
	"\fEchoResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2\xfb\x02\n" +
	"\x04Echo\x12P\n" +
	"\tUnaryEcho\x12\x1f.grpc.examples.echo.EchoRequest\x1a .grpc.examples.echo.EchoResponse\"\x00\x12\\\n" +
	"\x13ServerStreamingEcho\x12\x1f.grpc.examples.echo.EchoRequest\x1a .grpc.examples.echo.EchoResponse\"\x000\x01\x12\\\n" +
	"\x13ClientStreamingEcho\x12\x1f.grpc.examples.echo.EchoRequest\x1a .grpc.examples.echo.EchoResponse\"\x00(\x01\x12e\n" +
	"\x1aBidirectionalStreamingEcho\x12\x1f.grpc.examples.echo.EchoRequest\x1a .grpc.examples.echo.EchoResponse\"\x00(\x010\x01B5Z3google.golang.org/grpc/examples/features/proto/echob\x06proto3"

var (
	file_examples_features_proto_echo_echo_proto_rawDescOnce sync.Once
	file_examples_features_proto_echo_echo_proto_rawDescData []byte
)

func file_examples_features_proto_echo_echo_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_examples_features_proto_echo_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_examples_features_proto_echo_echo_proto_goTypes = []any{
	(*EchoRequest)(nil),
	(*EchoResponse)(nil),
}
var file_examples_features_proto_echo_echo_proto_depIdxs = []int32{
	0,
	0,
	0,
	0,
	1,
	1,
	1,
	1,
	4,
	0,
	0,
	0,
	0,
}

func init()                                              { file_examples_features_proto_echo_echo_proto_init() }
func file_examples_features_proto_echo_echo_proto_init() { _ = "STUB: not implemented"; return }
