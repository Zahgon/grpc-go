package core

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bucket struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Start         float64                `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	Count         uint64                 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bucket) Reset() { _ = "STUB: not implemented"; return }

func (x *Bucket) String() string { _ = "STUB: not implemented"; return "" }

func (*Bucket) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Bucket) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Bucket) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Bucket) GetStart() float64 { _ = "STUB: not implemented"; return 0 }

func (x *Bucket) GetCount() uint64 { _ = "STUB: not implemented"; return 0 }

type Histogram struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Buckets       []*Bucket              `protobuf:"bytes,1,rep,name=buckets,proto3" json:"buckets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Histogram) Reset() { _ = "STUB: not implemented"; return }

func (x *Histogram) String() string { _ = "STUB: not implemented"; return "" }

func (*Histogram) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Histogram) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Histogram) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Histogram) GetBuckets() []*Bucket { _ = "STUB: not implemented"; return nil }

type Metric struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`

	Value         isMetric_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Metric) Reset() { _ = "STUB: not implemented"; return }

func (x *Metric) String() string { _ = "STUB: not implemented"; return "" }

func (*Metric) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Metric) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Metric) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Metric) GetName() string { _ = "STUB: not implemented"; return "" }

func (x *Metric) GetValue() isMetric_Value { _ = "STUB: not implemented"; return *new(isMetric_Value) }

func (x *Metric) GetCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (x *Metric) GetHistogram() *Histogram { _ = "STUB: not implemented"; return nil }

type isMetric_Value interface {
	isMetric_Value()
}

type Metric_Count struct {
	Count uint64 `protobuf:"varint,10,opt,name=count,proto3,oneof"`
}

type Metric_Histogram struct {
	Histogram *Histogram `protobuf:"bytes,11,opt,name=histogram,proto3,oneof"`
}

func (*Metric_Count) isMetric_Value() { _ = "STUB: not implemented"; return }

func (*Metric_Histogram) isMetric_Value() { _ = "STUB: not implemented"; return }

type Stats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metrics       []*Metric              `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Stats) Reset() { _ = "STUB: not implemented"; return }

func (x *Stats) String() string { _ = "STUB: not implemented"; return "" }

func (*Stats) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Stats) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Stats) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Stats) GetMetrics() []*Metric { _ = "STUB: not implemented"; return nil }

var File_grpc_core_stats_proto protoreflect.FileDescriptor

const file_grpc_core_stats_proto_rawDesc = "" +
	"\n" +
	"\x15grpc/core/stats.proto\x12\tgrpc.core\"4\n" +
	"\x06Bucket\x12\x14\n" +
	"\x05start\x18\x01 \x01(\x01R\x05start\x12\x14\n" +
	"\x05count\x18\x02 \x01(\x04R\x05count\"8\n" +
	"\tHistogram\x12+\n" +
	"\abuckets\x18\x01 \x03(\v2\x11.grpc.core.BucketR\abuckets\"s\n" +
	"\x06Metric\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x16\n" +
	"\x05count\x18\n" +
	" \x01(\x04H\x00R\x05count\x124\n" +
	"\thistogram\x18\v \x01(\v2\x14.grpc.core.HistogramH\x00R\thistogramB\a\n" +
	"\x05value\"4\n" +
	"\x05Stats\x12+\n" +
	"\ametrics\x18\x01 \x03(\v2\x11.grpc.core.MetricR\ametricsb\x06proto3"

var (
	file_grpc_core_stats_proto_rawDescOnce sync.Once
	file_grpc_core_stats_proto_rawDescData []byte
)

func file_grpc_core_stats_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_grpc_core_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_core_stats_proto_goTypes = []any{
	(*Bucket)(nil),
	(*Histogram)(nil),
	(*Metric)(nil),
	(*Stats)(nil),
}
var file_grpc_core_stats_proto_depIdxs = []int32{
	0,
	1,
	2,
	3,
	3,
	3,
	3,
	0,
}

func init()                            { file_grpc_core_stats_proto_init() }
func file_grpc_core_stats_proto_init() { _ = "STUB: not implemented"; return }
