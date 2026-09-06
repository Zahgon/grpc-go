package helloworld

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *HelloRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*HelloRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HelloRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HelloRequest) GetName() string { _ = "STUB: not implemented"; return "" }

type HelloReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloReply) Reset() { _ = "STUB: not implemented"; return }

func (x *HelloReply) String() string { _ = "STUB: not implemented"; return "" }

func (*HelloReply) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HelloReply) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HelloReply) GetMessage() string { _ = "STUB: not implemented"; return "" }

var File_examples_helloworld_helloworld_helloworld_proto protoreflect.FileDescriptor

const file_examples_helloworld_helloworld_helloworld_proto_rawDesc = "" +
	"\n" +
	"/examples/helloworld/helloworld/helloworld.proto\x12\n" +
	"helloworld\"\"\n" +
	"\fHelloRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"&\n" +
	"\n" +
	"HelloReply\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2I\n" +
	"\aGreeter\x12>\n" +
	"\bSayHello\x12\x18.helloworld.HelloRequest\x1a\x16.helloworld.HelloReply\"\x00Bg\n" +
	"\x1bio.grpc.examples.helloworldB\x0fHelloWorldProtoP\x01Z5google.golang.org/grpc/examples/helloworld/helloworldb\x06proto3"

var (
	file_examples_helloworld_helloworld_helloworld_proto_rawDescOnce sync.Once
	file_examples_helloworld_helloworld_helloworld_proto_rawDescData []byte
)

func file_examples_helloworld_helloworld_helloworld_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_examples_helloworld_helloworld_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_examples_helloworld_helloworld_helloworld_proto_goTypes = []any{
	(*HelloRequest)(nil),
	(*HelloReply)(nil),
}
var file_examples_helloworld_helloworld_helloworld_proto_depIdxs = []int32{
	0,
	1,
	1,
	0,
	0,
	0,
	0,
}

func init()                                                      { file_examples_helloworld_helloworld_helloworld_proto_init() }
func file_examples_helloworld_helloworld_helloworld_proto_init() { _ = "STUB: not implemented"; return }
