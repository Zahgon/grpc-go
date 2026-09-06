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

type GaugeResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`

	Value         isGaugeResponse_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GaugeResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *GaugeResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*GaugeResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GaugeResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GaugeResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GaugeResponse) GetName() string { _ = "STUB: not implemented"; return "" }

func (x *GaugeResponse) GetValue() isGaugeResponse_Value {
	_ = "STUB: not implemented"
	return *new(isGaugeResponse_Value)
}

func (x *GaugeResponse) GetLongValue() int64 { _ = "STUB: not implemented"; return 0 }

func (x *GaugeResponse) GetDoubleValue() float64 { _ = "STUB: not implemented"; return 0 }

func (x *GaugeResponse) GetStringValue() string { _ = "STUB: not implemented"; return "" }

type isGaugeResponse_Value interface {
	isGaugeResponse_Value()
}

type GaugeResponse_LongValue struct {
	LongValue int64 `protobuf:"varint,2,opt,name=long_value,json=longValue,proto3,oneof"`
}

type GaugeResponse_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type GaugeResponse_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*GaugeResponse_LongValue) isGaugeResponse_Value() { _ = "STUB: not implemented"; return }

func (*GaugeResponse_DoubleValue) isGaugeResponse_Value() { _ = "STUB: not implemented"; return }

func (*GaugeResponse_StringValue) isGaugeResponse_Value() { _ = "STUB: not implemented"; return }

type GaugeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GaugeRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *GaugeRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*GaugeRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GaugeRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GaugeRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GaugeRequest) GetName() string { _ = "STUB: not implemented"; return "" }

type EmptyMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyMessage) Reset() { _ = "STUB: not implemented"; return }

func (x *EmptyMessage) String() string { _ = "STUB: not implemented"; return "" }

func (*EmptyMessage) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*EmptyMessage) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

var File_interop_stress_grpc_testing_metrics_proto protoreflect.FileDescriptor

const file_interop_stress_grpc_testing_metrics_proto_rawDesc = "" +
	"\n" +
	")interop/stress/grpc_testing/metrics.proto\x12\fgrpc.testing\"\x97\x01\n" +
	"\rGaugeResponse\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1f\n" +
	"\n" +
	"long_value\x18\x02 \x01(\x03H\x00R\tlongValue\x12#\n" +
	"\fdouble_value\x18\x03 \x01(\x01H\x00R\vdoubleValue\x12#\n" +
	"\fstring_value\x18\x04 \x01(\tH\x00R\vstringValueB\a\n" +
	"\x05value\"\"\n" +
	"\fGaugeRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"\x0e\n" +
	"\fEmptyMessage2\xa0\x01\n" +
	"\x0eMetricsService\x12I\n" +
	"\fGetAllGauges\x12\x1a.grpc.testing.EmptyMessage\x1a\x1b.grpc.testing.GaugeResponse0\x01\x12C\n" +
	"\bGetGauge\x12\x1a.grpc.testing.GaugeRequest\x1a\x1b.grpc.testing.GaugeResponseB4Z2google.golang.org/grpc/interop/stress/grpc_testingb\x06proto3"

var (
	file_interop_stress_grpc_testing_metrics_proto_rawDescOnce sync.Once
	file_interop_stress_grpc_testing_metrics_proto_rawDescData []byte
)

func file_interop_stress_grpc_testing_metrics_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_interop_stress_grpc_testing_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_interop_stress_grpc_testing_metrics_proto_goTypes = []any{
	(*GaugeResponse)(nil),
	(*GaugeRequest)(nil),
	(*EmptyMessage)(nil),
}
var file_interop_stress_grpc_testing_metrics_proto_depIdxs = []int32{
	2,
	1,
	0,
	0,
	2,
	0,
	0,
	0,
	0,
}

func init()                                                { file_interop_stress_grpc_testing_metrics_proto_init() }
func file_interop_stress_grpc_testing_metrics_proto_init() { _ = "STUB: not implemented"; return }
