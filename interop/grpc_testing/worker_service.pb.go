package grpc_testing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_grpc_testing_worker_service_proto protoreflect.FileDescriptor

const file_grpc_testing_worker_service_proto_rawDesc = "" +
	"\n" +
	"!grpc/testing/worker_service.proto\x12\fgrpc.testing\x1a\x1agrpc/testing/control.proto2\x97\x02\n" +
	"\rWorkerService\x12E\n" +
	"\tRunServer\x12\x18.grpc.testing.ServerArgs\x1a\x1a.grpc.testing.ServerStatus(\x010\x01\x12E\n" +
	"\tRunClient\x12\x18.grpc.testing.ClientArgs\x1a\x1a.grpc.testing.ClientStatus(\x010\x01\x12B\n" +
	"\tCoreCount\x12\x19.grpc.testing.CoreRequest\x1a\x1a.grpc.testing.CoreResponse\x124\n" +
	"\n" +
	"QuitWorker\x12\x12.grpc.testing.Void\x1a\x12.grpc.testing.VoidB'\n" +
	"\x0fio.grpc.testingB\x12WorkerServiceProtoP\x01b\x06proto3"

var file_grpc_testing_worker_service_proto_goTypes = []any{
	(*ServerArgs)(nil),
	(*ClientArgs)(nil),
	(*CoreRequest)(nil),
	(*Void)(nil),
	(*ServerStatus)(nil),
	(*ClientStatus)(nil),
	(*CoreResponse)(nil),
}
var file_grpc_testing_worker_service_proto_depIdxs = []int32{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	3,
	4,
	0,
	0,
	0,
	0,
}

func init()                                        { file_grpc_testing_worker_service_proto_init() }
func file_grpc_testing_worker_service_proto_init() { _ = "STUB: not implemented"; return }
