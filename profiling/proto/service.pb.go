package proto

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnableRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Enabled       bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnableRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *EnableRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*EnableRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *EnableRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*EnableRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *EnableRequest) GetEnabled() bool { _ = "STUB: not implemented"; return false }

type EnableResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnableResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *EnableResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*EnableResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *EnableResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*EnableResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type GetStreamStatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStreamStatsRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *GetStreamStatsRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*GetStreamStatsRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetStreamStatsRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetStreamStatsRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GetStreamStatsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StreamStats   []*Stat                `protobuf:"bytes,1,rep,name=stream_stats,json=streamStats,proto3" json:"stream_stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStreamStatsResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *GetStreamStatsResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*GetStreamStatsResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetStreamStatsResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetStreamStatsResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *GetStreamStatsResponse) GetStreamStats() []*Stat { _ = "STUB: not implemented"; return nil }

type Timer struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Tags string `protobuf:"bytes,1,opt,name=tags,proto3" json:"tags,omitempty"`

	BeginSec  int64 `protobuf:"varint,2,opt,name=begin_sec,json=beginSec,proto3" json:"begin_sec,omitempty"`
	BeginNsec int32 `protobuf:"varint,3,opt,name=begin_nsec,json=beginNsec,proto3" json:"begin_nsec,omitempty"`

	EndSec  int64 `protobuf:"varint,4,opt,name=end_sec,json=endSec,proto3" json:"end_sec,omitempty"`
	EndNsec int32 `protobuf:"varint,5,opt,name=end_nsec,json=endNsec,proto3" json:"end_nsec,omitempty"`

	GoId          int64 `protobuf:"varint,6,opt,name=go_id,json=goId,proto3" json:"go_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Timer) Reset() { _ = "STUB: not implemented"; return }

func (x *Timer) String() string { _ = "STUB: not implemented"; return "" }

func (*Timer) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Timer) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Timer) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Timer) GetTags() string { _ = "STUB: not implemented"; return "" }

func (x *Timer) GetBeginSec() int64 { _ = "STUB: not implemented"; return 0 }

func (x *Timer) GetBeginNsec() int32 { _ = "STUB: not implemented"; return 0 }

func (x *Timer) GetEndSec() int64 { _ = "STUB: not implemented"; return 0 }

func (x *Timer) GetEndNsec() int32 { _ = "STUB: not implemented"; return 0 }

func (x *Timer) GetGoId() int64 { _ = "STUB: not implemented"; return 0 }

type Stat struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Tags string `protobuf:"bytes,1,opt,name=tags,proto3" json:"tags,omitempty"`

	Timers []*Timer `protobuf:"bytes,2,rep,name=timers,proto3" json:"timers,omitempty"`

	Metadata      []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Stat) Reset() { _ = "STUB: not implemented"; return }

func (x *Stat) String() string { _ = "STUB: not implemented"; return "" }

func (*Stat) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Stat) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Stat) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Stat) GetTags() string { _ = "STUB: not implemented"; return "" }

func (x *Stat) GetTimers() []*Timer { _ = "STUB: not implemented"; return nil }

func (x *Stat) GetMetadata() []byte { _ = "STUB: not implemented"; return nil }

var File_profiling_proto_service_proto protoreflect.FileDescriptor

const file_profiling_proto_service_proto_rawDesc = "" +
	"\n" +
	"\x1dprofiling/proto/service.proto\x12\x19grpc.go.profiling.v1alpha\")\n" +
	"\rEnableRequest\x12\x18\n" +
	"\aenabled\x18\x01 \x01(\bR\aenabled\"\x10\n" +
	"\x0eEnableResponse\"\x17\n" +
	"\x15GetStreamStatsRequest\"\\\n" +
	"\x16GetStreamStatsResponse\x12B\n" +
	"\fstream_stats\x18\x01 \x03(\v2\x1f.grpc.go.profiling.v1alpha.StatR\vstreamStats\"\xa0\x01\n" +
	"\x05Timer\x12\x12\n" +
	"\x04tags\x18\x01 \x01(\tR\x04tags\x12\x1b\n" +
	"\tbegin_sec\x18\x02 \x01(\x03R\bbeginSec\x12\x1d\n" +
	"\n" +
	"begin_nsec\x18\x03 \x01(\x05R\tbeginNsec\x12\x17\n" +
	"\aend_sec\x18\x04 \x01(\x03R\x06endSec\x12\x19\n" +
	"\bend_nsec\x18\x05 \x01(\x05R\aendNsec\x12\x13\n" +
	"\x05go_id\x18\x06 \x01(\x03R\x04goId\"p\n" +
	"\x04Stat\x12\x12\n" +
	"\x04tags\x18\x01 \x01(\tR\x04tags\x128\n" +
	"\x06timers\x18\x02 \x03(\v2 .grpc.go.profiling.v1alpha.TimerR\x06timers\x12\x1a\n" +
	"\bmetadata\x18\x03 \x01(\fR\bmetadata2\xe1\x01\n" +
	"\tProfiling\x12]\n" +
	"\x06Enable\x12(.grpc.go.profiling.v1alpha.EnableRequest\x1a).grpc.go.profiling.v1alpha.EnableResponse\x12u\n" +
	"\x0eGetStreamStats\x120.grpc.go.profiling.v1alpha.GetStreamStatsRequest\x1a1.grpc.go.profiling.v1alpha.GetStreamStatsResponseB(Z&google.golang.org/grpc/profiling/protob\x06proto3"

var (
	file_profiling_proto_service_proto_rawDescOnce sync.Once
	file_profiling_proto_service_proto_rawDescData []byte
)

func file_profiling_proto_service_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_profiling_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_profiling_proto_service_proto_goTypes = []any{
	(*EnableRequest)(nil),
	(*EnableResponse)(nil),
	(*GetStreamStatsRequest)(nil),
	(*GetStreamStatsResponse)(nil),
	(*Timer)(nil),
	(*Stat)(nil),
}
var file_profiling_proto_service_proto_depIdxs = []int32{
	5,
	4,
	0,
	2,
	1,
	3,
	4,
	2,
	2,
	2,
	0,
}

func init()                                    { file_profiling_proto_service_proto_init() }
func file_profiling_proto_service_proto_init() { _ = "STUB: not implemented"; return }
