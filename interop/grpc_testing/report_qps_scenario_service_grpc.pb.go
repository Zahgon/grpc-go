package grpc_testing

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	ReportQpsScenarioService_ReportScenario_FullMethodName = "/grpc.testing.ReportQpsScenarioService/ReportScenario"
)

type ReportQpsScenarioServiceClient interface {
	ReportScenario(ctx context.Context, in *ScenarioResult, opts ...grpc.CallOption) (*Void, error)
}

type reportQpsScenarioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportQpsScenarioServiceClient(cc grpc.ClientConnInterface) ReportQpsScenarioServiceClient {
	_ = "STUB: not implemented"
	return *new(ReportQpsScenarioServiceClient)
}

func (c *reportQpsScenarioServiceClient) ReportScenario(ctx context.Context, in *ScenarioResult, opts ...grpc.CallOption) (*Void, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ReportQpsScenarioServiceServer interface {
	ReportScenario(context.Context, *ScenarioResult) (*Void, error)
	mustEmbedUnimplementedReportQpsScenarioServiceServer()
}

type UnimplementedReportQpsScenarioServiceServer struct{}

func (UnimplementedReportQpsScenarioServiceServer) ReportScenario(context.Context, *ScenarioResult) (*Void, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedReportQpsScenarioServiceServer) mustEmbedUnimplementedReportQpsScenarioServiceServer() {
	_ = "STUB: not implemented"
	return
}

func (UnimplementedReportQpsScenarioServiceServer) testEmbeddedByValue() {
	_ = "STUB: not implemented"
	return
}

type UnsafeReportQpsScenarioServiceServer interface {
	mustEmbedUnimplementedReportQpsScenarioServiceServer()
}

func RegisterReportQpsScenarioServiceServer(s grpc.ServiceRegistrar, srv ReportQpsScenarioServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _ReportQpsScenarioService_ReportScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var ReportQpsScenarioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.ReportQpsScenarioService",
	HandlerType: (*ReportQpsScenarioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportScenario",
			Handler:    _ReportQpsScenarioService_ReportScenario_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/testing/report_qps_scenario_service.proto",
}
