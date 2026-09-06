package grpc_testing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_grpc_testing_test_proto protoreflect.FileDescriptor

const file_grpc_testing_test_proto_rawDesc = "" +
	"\n" +
	"\x17grpc/testing/test.proto\x12\fgrpc.testing\x1a\x18grpc/testing/empty.proto\x1a\x1bgrpc/testing/messages.proto2\xcb\x05\n" +
	"\vTestService\x125\n" +
	"\tEmptyCall\x12\x13.grpc.testing.Empty\x1a\x13.grpc.testing.Empty\x12F\n" +
	"\tUnaryCall\x12\x1b.grpc.testing.SimpleRequest\x1a\x1c.grpc.testing.SimpleResponse\x12O\n" +
	"\x12CacheableUnaryCall\x12\x1b.grpc.testing.SimpleRequest\x1a\x1c.grpc.testing.SimpleResponse\x12l\n" +
	"\x13StreamingOutputCall\x12(.grpc.testing.StreamingOutputCallRequest\x1a).grpc.testing.StreamingOutputCallResponse0\x01\x12i\n" +
	"\x12StreamingInputCall\x12'.grpc.testing.StreamingInputCallRequest\x1a(.grpc.testing.StreamingInputCallResponse(\x01\x12i\n" +
	"\x0eFullDuplexCall\x12(.grpc.testing.StreamingOutputCallRequest\x1a).grpc.testing.StreamingOutputCallResponse(\x010\x01\x12i\n" +
	"\x0eHalfDuplexCall\x12(.grpc.testing.StreamingOutputCallRequest\x1a).grpc.testing.StreamingOutputCallResponse(\x010\x01\x12=\n" +
	"\x11UnimplementedCall\x12\x13.grpc.testing.Empty\x1a\x13.grpc.testing.Empty2U\n" +
	"\x14UnimplementedService\x12=\n" +
	"\x11UnimplementedCall\x12\x13.grpc.testing.Empty\x1a\x13.grpc.testing.Empty2\x89\x01\n" +
	"\x10ReconnectService\x12;\n" +
	"\x05Start\x12\x1d.grpc.testing.ReconnectParams\x1a\x13.grpc.testing.Empty\x128\n" +
	"\x04Stop\x12\x13.grpc.testing.Empty\x1a\x1b.grpc.testing.ReconnectInfo2\x86\x02\n" +
	"\x18LoadBalancerStatsService\x12c\n" +
	"\x0eGetClientStats\x12&.grpc.testing.LoadBalancerStatsRequest\x1a'.grpc.testing.LoadBalancerStatsResponse\"\x00\x12\x84\x01\n" +
	"\x19GetClientAccumulatedStats\x121.grpc.testing.LoadBalancerAccumulatedStatsRequest\x1a2.grpc.testing.LoadBalancerAccumulatedStatsResponse\"\x002\xcc\x01\n" +
	"\vHookService\x120\n" +
	"\x04Hook\x12\x13.grpc.testing.Empty\x1a\x13.grpc.testing.Empty\x12L\n" +
	"\x0fSetReturnStatus\x12$.grpc.testing.SetReturnStatusRequest\x1a\x13.grpc.testing.Empty\x12=\n" +
	"\x11ClearReturnStatus\x12\x13.grpc.testing.Empty\x1a\x13.grpc.testing.Empty2\xd5\x01\n" +
	"\x16XdsUpdateHealthService\x126\n" +
	"\n" +
	"SetServing\x12\x13.grpc.testing.Empty\x1a\x13.grpc.testing.Empty\x129\n" +
	"\rSetNotServing\x12\x13.grpc.testing.Empty\x1a\x13.grpc.testing.Empty\x12H\n" +
	"\x0fSendHookRequest\x12\x19.grpc.testing.HookRequest\x1a\x1a.grpc.testing.HookResponse2{\n" +
	"\x1fXdsUpdateClientConfigureService\x12X\n" +
	"\tConfigure\x12$.grpc.testing.ClientConfigureRequest\x1a%.grpc.testing.ClientConfigureResponseB\x1d\n" +
	"\x1bio.grpc.testing.integrationb\x06proto3"

var file_grpc_testing_test_proto_goTypes = []any{
	(*Empty)(nil),
	(*SimpleRequest)(nil),
	(*StreamingOutputCallRequest)(nil),
	(*StreamingInputCallRequest)(nil),
	(*ReconnectParams)(nil),
	(*LoadBalancerStatsRequest)(nil),
	(*LoadBalancerAccumulatedStatsRequest)(nil),
	(*SetReturnStatusRequest)(nil),
	(*HookRequest)(nil),
	(*ClientConfigureRequest)(nil),
	(*SimpleResponse)(nil),
	(*StreamingOutputCallResponse)(nil),
	(*StreamingInputCallResponse)(nil),
	(*ReconnectInfo)(nil),
	(*LoadBalancerStatsResponse)(nil),
	(*LoadBalancerAccumulatedStatsResponse)(nil),
	(*HookResponse)(nil),
	(*ClientConfigureResponse)(nil),
}
var file_grpc_testing_test_proto_depIdxs = []int32{
	0,
	1,
	1,
	2,
	3,
	2,
	2,
	0,
	0,
	4,
	0,
	5,
	6,
	0,
	7,
	0,
	0,
	0,
	8,
	9,
	0,
	10,
	10,
	11,
	12,
	11,
	11,
	0,
	0,
	0,
	13,
	14,
	15,
	0,
	0,
	0,
	0,
	0,
	16,
	17,
	20,
	0,
	0,
	0,
	0,
}

func init()                              { file_grpc_testing_test_proto_init() }
func file_grpc_testing_test_proto_init() { _ = "STUB: not implemented"; return }
