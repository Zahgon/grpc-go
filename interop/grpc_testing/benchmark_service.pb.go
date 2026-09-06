package grpc_testing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_grpc_testing_benchmark_service_proto protoreflect.FileDescriptor

const file_grpc_testing_benchmark_service_proto_rawDesc = "" +
	"\n" +
	"$grpc/testing/benchmark_service.proto\x12\fgrpc.testing\x1a\x1bgrpc/testing/messages.proto2\xa6\x03\n" +
	"\x10BenchmarkService\x12F\n" +
	"\tUnaryCall\x12\x1b.grpc.testing.SimpleRequest\x1a\x1c.grpc.testing.SimpleResponse\x12N\n" +
	"\rStreamingCall\x12\x1b.grpc.testing.SimpleRequest\x1a\x1c.grpc.testing.SimpleResponse(\x010\x01\x12R\n" +
	"\x13StreamingFromClient\x12\x1b.grpc.testing.SimpleRequest\x1a\x1c.grpc.testing.SimpleResponse(\x01\x12R\n" +
	"\x13StreamingFromServer\x12\x1b.grpc.testing.SimpleRequest\x1a\x1c.grpc.testing.SimpleResponse0\x01\x12R\n" +
	"\x11StreamingBothWays\x12\x1b.grpc.testing.SimpleRequest\x1a\x1c.grpc.testing.SimpleResponse(\x010\x01B*\n" +
	"\x0fio.grpc.testingB\x15BenchmarkServiceProtoP\x01b\x06proto3"

var file_grpc_testing_benchmark_service_proto_goTypes = []any{
	(*SimpleRequest)(nil),
	(*SimpleResponse)(nil),
}
var file_grpc_testing_benchmark_service_proto_depIdxs = []int32{
	0,
	0,
	0,
	0,
	0,
	1,
	1,
	1,
	1,
	1,
	5,
	0,
	0,
	0,
	0,
}

func init()                                           { file_grpc_testing_benchmark_service_proto_init() }
func file_grpc_testing_benchmark_service_proto_init() { _ = "STUB: not implemented"; return }
