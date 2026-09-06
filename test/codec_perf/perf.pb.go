package codec_perf

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Buffer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Body          []byte                 `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Buffer) Reset() { _ = "STUB: not implemented"; return }

func (x *Buffer) String() string { _ = "STUB: not implemented"; return "" }

func (*Buffer) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Buffer) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Buffer) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Buffer) GetBody() []byte { _ = "STUB: not implemented"; return nil }

var File_test_codec_perf_perf_proto protoreflect.FileDescriptor

const file_test_codec_perf_perf_proto_rawDesc = "" +
	"\n" +
	"\x1atest/codec_perf/perf.proto\x12\n" +
	"codec.perf\"\x1c\n" +
	"\x06Buffer\x12\x12\n" +
	"\x04body\x18\x01 \x01(\fR\x04bodyB(Z&google.golang.org/grpc/test/codec_perfb\x06proto3"

var (
	file_test_codec_perf_perf_proto_rawDescOnce sync.Once
	file_test_codec_perf_perf_proto_rawDescData []byte
)

func file_test_codec_perf_perf_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_test_codec_perf_perf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test_codec_perf_perf_proto_goTypes = []any{
	(*Buffer)(nil),
}
var file_test_codec_perf_perf_proto_depIdxs = []int32{
	0,
	0,
	0,
	0,
	0,
}

func init()                                 { file_test_codec_perf_perf_proto_init() }
func file_test_codec_perf_perf_proto_init() { _ = "STUB: not implemented"; return }
