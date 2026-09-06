package grpc_testing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_grpc_testing_report_qps_scenario_service_proto protoreflect.FileDescriptor

const file_grpc_testing_report_qps_scenario_service_proto_rawDesc = "" +
	"\n" +
	".grpc/testing/report_qps_scenario_service.proto\x12\fgrpc.testing\x1a\x1agrpc/testing/control.proto2^\n" +
	"\x18ReportQpsScenarioService\x12B\n" +
	"\x0eReportScenario\x12\x1c.grpc.testing.ScenarioResult\x1a\x12.grpc.testing.VoidB2\n" +
	"\x0fio.grpc.testingB\x1dReportQpsScenarioServiceProtoP\x01b\x06proto3"

var file_grpc_testing_report_qps_scenario_service_proto_goTypes = []any{
	(*ScenarioResult)(nil),
	(*Void)(nil),
}
var file_grpc_testing_report_qps_scenario_service_proto_depIdxs = []int32{
	0,
	1,
	1,
	0,
	0,
	0,
	0,
}

func init()                                                     { file_grpc_testing_report_qps_scenario_service_proto_init() }
func file_grpc_testing_report_qps_scenario_service_proto_init() { _ = "STUB: not implemented"; return }
