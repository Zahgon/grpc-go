package grpc_testing

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	TestService_EmptyCall_FullMethodName           = "/grpc.testing.TestService/EmptyCall"
	TestService_UnaryCall_FullMethodName           = "/grpc.testing.TestService/UnaryCall"
	TestService_CacheableUnaryCall_FullMethodName  = "/grpc.testing.TestService/CacheableUnaryCall"
	TestService_StreamingOutputCall_FullMethodName = "/grpc.testing.TestService/StreamingOutputCall"
	TestService_StreamingInputCall_FullMethodName  = "/grpc.testing.TestService/StreamingInputCall"
	TestService_FullDuplexCall_FullMethodName      = "/grpc.testing.TestService/FullDuplexCall"
	TestService_HalfDuplexCall_FullMethodName      = "/grpc.testing.TestService/HalfDuplexCall"
	TestService_UnimplementedCall_FullMethodName   = "/grpc.testing.TestService/UnimplementedCall"
)

type TestServiceClient interface {
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)

	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)

	CacheableUnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)

	StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamingOutputCallResponse], error)

	StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[StreamingInputCallRequest, StreamingInputCallResponse], error)

	FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamingOutputCallRequest, StreamingOutputCallResponse], error)

	HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamingOutputCallRequest, StreamingOutputCallResponse], error)

	UnimplementedCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	_ = "STUB: not implemented"
	return *new(TestServiceClient)
}

func (c *testServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *testServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *testServiceClient) CacheableUnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *testServiceClient) StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamingOutputCallResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestService_StreamingOutputCallClient = grpc.ServerStreamingClient[StreamingOutputCallResponse]

func (c *testServiceClient) StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[StreamingInputCallRequest, StreamingInputCallResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestService_StreamingInputCallClient = grpc.ClientStreamingClient[StreamingInputCallRequest, StreamingInputCallResponse]

func (c *testServiceClient) FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamingOutputCallRequest, StreamingOutputCallResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestService_FullDuplexCallClient = grpc.BidiStreamingClient[StreamingOutputCallRequest, StreamingOutputCallResponse]

func (c *testServiceClient) HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamingOutputCallRequest, StreamingOutputCallResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestService_HalfDuplexCallClient = grpc.BidiStreamingClient[StreamingOutputCallRequest, StreamingOutputCallResponse]

func (c *testServiceClient) UnimplementedCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestServiceServer interface {
	EmptyCall(context.Context, *Empty) (*Empty, error)

	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)

	CacheableUnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)

	StreamingOutputCall(*StreamingOutputCallRequest, grpc.ServerStreamingServer[StreamingOutputCallResponse]) error

	StreamingInputCall(grpc.ClientStreamingServer[StreamingInputCallRequest, StreamingInputCallResponse]) error

	FullDuplexCall(grpc.BidiStreamingServer[StreamingOutputCallRequest, StreamingOutputCallResponse]) error

	HalfDuplexCall(grpc.BidiStreamingServer[StreamingOutputCallRequest, StreamingOutputCallResponse]) error

	UnimplementedCall(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedTestServiceServer()
}

type UnimplementedTestServiceServer struct{}

func (UnimplementedTestServiceServer) EmptyCall(context.Context, *Empty) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedTestServiceServer) UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedTestServiceServer) CacheableUnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedTestServiceServer) StreamingOutputCall(*StreamingOutputCallRequest, grpc.ServerStreamingServer[StreamingOutputCallResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedTestServiceServer) StreamingInputCall(grpc.ClientStreamingServer[StreamingInputCallRequest, StreamingInputCallResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedTestServiceServer) FullDuplexCall(grpc.BidiStreamingServer[StreamingOutputCallRequest, StreamingOutputCallResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedTestServiceServer) HalfDuplexCall(grpc.BidiStreamingServer[StreamingOutputCallRequest, StreamingOutputCallResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedTestServiceServer) UnimplementedCall(context.Context, *Empty) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedTestServiceServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _TestService_EmptyCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _TestService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _TestService_CacheableUnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _TestService_StreamingOutputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type TestService_StreamingOutputCallServer = grpc.ServerStreamingServer[StreamingOutputCallResponse]

func _TestService_StreamingInputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type TestService_StreamingInputCallServer = grpc.ClientStreamingServer[StreamingInputCallRequest, StreamingInputCallResponse]

func _TestService_FullDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type TestService_FullDuplexCallServer = grpc.BidiStreamingServer[StreamingOutputCallRequest, StreamingOutputCallResponse]

func _TestService_HalfDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type TestService_HalfDuplexCallServer = grpc.BidiStreamingServer[StreamingOutputCallRequest, StreamingOutputCallResponse]

func _TestService_UnimplementedCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyCall",
			Handler:    _TestService_EmptyCall_Handler,
		},
		{
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		},
		{
			MethodName: "CacheableUnaryCall",
			Handler:    _TestService_CacheableUnaryCall_Handler,
		},
		{
			MethodName: "UnimplementedCall",
			Handler:    _TestService_UnimplementedCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingOutputCall",
			Handler:       _TestService_StreamingOutputCall_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamingInputCall",
			Handler:       _TestService_StreamingInputCall_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FullDuplexCall",
			Handler:       _TestService_FullDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "HalfDuplexCall",
			Handler:       _TestService_HalfDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/testing/test.proto",
}

const (
	UnimplementedService_UnimplementedCall_FullMethodName = "/grpc.testing.UnimplementedService/UnimplementedCall"
)

type UnimplementedServiceClient interface {
	UnimplementedCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type unimplementedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUnimplementedServiceClient(cc grpc.ClientConnInterface) UnimplementedServiceClient {
	_ = "STUB: not implemented"
	return *new(UnimplementedServiceClient)
}

func (c *unimplementedServiceClient) UnimplementedCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type UnimplementedServiceServer interface {
	UnimplementedCall(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedUnimplementedServiceServer()
}

type UnimplementedUnimplementedServiceServer struct{}

func (UnimplementedUnimplementedServiceServer) UnimplementedCall(context.Context, *Empty) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedUnimplementedServiceServer) mustEmbedUnimplementedUnimplementedServiceServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedUnimplementedServiceServer) testEmbeddedByValue() {
	_ = "STUB: not implemented"
	return
}

type UnsafeUnimplementedServiceServer interface {
	mustEmbedUnimplementedUnimplementedServiceServer()
}

func RegisterUnimplementedServiceServer(s grpc.ServiceRegistrar, srv UnimplementedServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _UnimplementedService_UnimplementedCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var UnimplementedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.UnimplementedService",
	HandlerType: (*UnimplementedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnimplementedCall",
			Handler:    _UnimplementedService_UnimplementedCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/testing/test.proto",
}

const (
	ReconnectService_Start_FullMethodName = "/grpc.testing.ReconnectService/Start"
	ReconnectService_Stop_FullMethodName  = "/grpc.testing.ReconnectService/Stop"
)

type ReconnectServiceClient interface {
	Start(ctx context.Context, in *ReconnectParams, opts ...grpc.CallOption) (*Empty, error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReconnectInfo, error)
}

type reconnectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReconnectServiceClient(cc grpc.ClientConnInterface) ReconnectServiceClient {
	_ = "STUB: not implemented"
	return *new(ReconnectServiceClient)
}

func (c *reconnectServiceClient) Start(ctx context.Context, in *ReconnectParams, opts ...grpc.CallOption) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *reconnectServiceClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReconnectInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ReconnectServiceServer interface {
	Start(context.Context, *ReconnectParams) (*Empty, error)
	Stop(context.Context, *Empty) (*ReconnectInfo, error)
	mustEmbedUnimplementedReconnectServiceServer()
}

type UnimplementedReconnectServiceServer struct{}

func (UnimplementedReconnectServiceServer) Start(context.Context, *ReconnectParams) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedReconnectServiceServer) Stop(context.Context, *Empty) (*ReconnectInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedReconnectServiceServer) mustEmbedUnimplementedReconnectServiceServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedReconnectServiceServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeReconnectServiceServer interface {
	mustEmbedUnimplementedReconnectServiceServer()
}

func RegisterReconnectServiceServer(s grpc.ServiceRegistrar, srv ReconnectServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _ReconnectService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _ReconnectService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var ReconnectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.ReconnectService",
	HandlerType: (*ReconnectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _ReconnectService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ReconnectService_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/testing/test.proto",
}

const (
	LoadBalancerStatsService_GetClientStats_FullMethodName            = "/grpc.testing.LoadBalancerStatsService/GetClientStats"
	LoadBalancerStatsService_GetClientAccumulatedStats_FullMethodName = "/grpc.testing.LoadBalancerStatsService/GetClientAccumulatedStats"
)

type LoadBalancerStatsServiceClient interface {
	GetClientStats(ctx context.Context, in *LoadBalancerStatsRequest, opts ...grpc.CallOption) (*LoadBalancerStatsResponse, error)

	GetClientAccumulatedStats(ctx context.Context, in *LoadBalancerAccumulatedStatsRequest, opts ...grpc.CallOption) (*LoadBalancerAccumulatedStatsResponse, error)
}

type loadBalancerStatsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadBalancerStatsServiceClient(cc grpc.ClientConnInterface) LoadBalancerStatsServiceClient {
	_ = "STUB: not implemented"
	return *new(LoadBalancerStatsServiceClient)
}

func (c *loadBalancerStatsServiceClient) GetClientStats(ctx context.Context, in *LoadBalancerStatsRequest, opts ...grpc.CallOption) (*LoadBalancerStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *loadBalancerStatsServiceClient) GetClientAccumulatedStats(ctx context.Context, in *LoadBalancerAccumulatedStatsRequest, opts ...grpc.CallOption) (*LoadBalancerAccumulatedStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type LoadBalancerStatsServiceServer interface {
	GetClientStats(context.Context, *LoadBalancerStatsRequest) (*LoadBalancerStatsResponse, error)

	GetClientAccumulatedStats(context.Context, *LoadBalancerAccumulatedStatsRequest) (*LoadBalancerAccumulatedStatsResponse, error)
	mustEmbedUnimplementedLoadBalancerStatsServiceServer()
}

type UnimplementedLoadBalancerStatsServiceServer struct{}

func (UnimplementedLoadBalancerStatsServiceServer) GetClientStats(context.Context, *LoadBalancerStatsRequest) (*LoadBalancerStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedLoadBalancerStatsServiceServer) GetClientAccumulatedStats(context.Context, *LoadBalancerAccumulatedStatsRequest) (*LoadBalancerAccumulatedStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedLoadBalancerStatsServiceServer) mustEmbedUnimplementedLoadBalancerStatsServiceServer() {
	_ = "STUB: not implemented"
	return
}

func (UnimplementedLoadBalancerStatsServiceServer) testEmbeddedByValue() {
	_ = "STUB: not implemented"
	return
}

type UnsafeLoadBalancerStatsServiceServer interface {
	mustEmbedUnimplementedLoadBalancerStatsServiceServer()
}

func RegisterLoadBalancerStatsServiceServer(s grpc.ServiceRegistrar, srv LoadBalancerStatsServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _LoadBalancerStatsService_GetClientStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _LoadBalancerStatsService_GetClientAccumulatedStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var LoadBalancerStatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.LoadBalancerStatsService",
	HandlerType: (*LoadBalancerStatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClientStats",
			Handler:    _LoadBalancerStatsService_GetClientStats_Handler,
		},
		{
			MethodName: "GetClientAccumulatedStats",
			Handler:    _LoadBalancerStatsService_GetClientAccumulatedStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/testing/test.proto",
}

const (
	HookService_Hook_FullMethodName              = "/grpc.testing.HookService/Hook"
	HookService_SetReturnStatus_FullMethodName   = "/grpc.testing.HookService/SetReturnStatus"
	HookService_ClearReturnStatus_FullMethodName = "/grpc.testing.HookService/ClearReturnStatus"
)

type HookServiceClient interface {
	Hook(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)

	SetReturnStatus(ctx context.Context, in *SetReturnStatusRequest, opts ...grpc.CallOption) (*Empty, error)

	ClearReturnStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type hookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHookServiceClient(cc grpc.ClientConnInterface) HookServiceClient {
	_ = "STUB: not implemented"
	return *new(HookServiceClient)
}

func (c *hookServiceClient) Hook(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *hookServiceClient) SetReturnStatus(ctx context.Context, in *SetReturnStatusRequest, opts ...grpc.CallOption) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *hookServiceClient) ClearReturnStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type HookServiceServer interface {
	Hook(context.Context, *Empty) (*Empty, error)

	SetReturnStatus(context.Context, *SetReturnStatusRequest) (*Empty, error)

	ClearReturnStatus(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedHookServiceServer()
}

type UnimplementedHookServiceServer struct{}

func (UnimplementedHookServiceServer) Hook(context.Context, *Empty) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedHookServiceServer) SetReturnStatus(context.Context, *SetReturnStatusRequest) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedHookServiceServer) ClearReturnStatus(context.Context, *Empty) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedHookServiceServer) mustEmbedUnimplementedHookServiceServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedHookServiceServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeHookServiceServer interface {
	mustEmbedUnimplementedHookServiceServer()
}

func RegisterHookServiceServer(s grpc.ServiceRegistrar, srv HookServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _HookService_Hook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _HookService_SetReturnStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _HookService_ClearReturnStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var HookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.HookService",
	HandlerType: (*HookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hook",
			Handler:    _HookService_Hook_Handler,
		},
		{
			MethodName: "SetReturnStatus",
			Handler:    _HookService_SetReturnStatus_Handler,
		},
		{
			MethodName: "ClearReturnStatus",
			Handler:    _HookService_ClearReturnStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/testing/test.proto",
}

const (
	XdsUpdateHealthService_SetServing_FullMethodName      = "/grpc.testing.XdsUpdateHealthService/SetServing"
	XdsUpdateHealthService_SetNotServing_FullMethodName   = "/grpc.testing.XdsUpdateHealthService/SetNotServing"
	XdsUpdateHealthService_SendHookRequest_FullMethodName = "/grpc.testing.XdsUpdateHealthService/SendHookRequest"
)

type XdsUpdateHealthServiceClient interface {
	SetServing(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	SetNotServing(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	SendHookRequest(ctx context.Context, in *HookRequest, opts ...grpc.CallOption) (*HookResponse, error)
}

type xdsUpdateHealthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewXdsUpdateHealthServiceClient(cc grpc.ClientConnInterface) XdsUpdateHealthServiceClient {
	_ = "STUB: not implemented"
	return *new(XdsUpdateHealthServiceClient)
}

func (c *xdsUpdateHealthServiceClient) SetServing(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *xdsUpdateHealthServiceClient) SetNotServing(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *xdsUpdateHealthServiceClient) SendHookRequest(ctx context.Context, in *HookRequest, opts ...grpc.CallOption) (*HookResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type XdsUpdateHealthServiceServer interface {
	SetServing(context.Context, *Empty) (*Empty, error)
	SetNotServing(context.Context, *Empty) (*Empty, error)
	SendHookRequest(context.Context, *HookRequest) (*HookResponse, error)
	mustEmbedUnimplementedXdsUpdateHealthServiceServer()
}

type UnimplementedXdsUpdateHealthServiceServer struct{}

func (UnimplementedXdsUpdateHealthServiceServer) SetServing(context.Context, *Empty) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedXdsUpdateHealthServiceServer) SetNotServing(context.Context, *Empty) (*Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedXdsUpdateHealthServiceServer) SendHookRequest(context.Context, *HookRequest) (*HookResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedXdsUpdateHealthServiceServer) mustEmbedUnimplementedXdsUpdateHealthServiceServer() {
	_ = "STUB: not implemented"
	return
}

func (UnimplementedXdsUpdateHealthServiceServer) testEmbeddedByValue() {
	_ = "STUB: not implemented"
	return
}

type UnsafeXdsUpdateHealthServiceServer interface {
	mustEmbedUnimplementedXdsUpdateHealthServiceServer()
}

func RegisterXdsUpdateHealthServiceServer(s grpc.ServiceRegistrar, srv XdsUpdateHealthServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _XdsUpdateHealthService_SetServing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _XdsUpdateHealthService_SetNotServing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _XdsUpdateHealthService_SendHookRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var XdsUpdateHealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.XdsUpdateHealthService",
	HandlerType: (*XdsUpdateHealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetServing",
			Handler:    _XdsUpdateHealthService_SetServing_Handler,
		},
		{
			MethodName: "SetNotServing",
			Handler:    _XdsUpdateHealthService_SetNotServing_Handler,
		},
		{
			MethodName: "SendHookRequest",
			Handler:    _XdsUpdateHealthService_SendHookRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/testing/test.proto",
}

const (
	XdsUpdateClientConfigureService_Configure_FullMethodName = "/grpc.testing.XdsUpdateClientConfigureService/Configure"
)

type XdsUpdateClientConfigureServiceClient interface {
	Configure(ctx context.Context, in *ClientConfigureRequest, opts ...grpc.CallOption) (*ClientConfigureResponse, error)
}

type xdsUpdateClientConfigureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewXdsUpdateClientConfigureServiceClient(cc grpc.ClientConnInterface) XdsUpdateClientConfigureServiceClient {
	_ = "STUB: not implemented"
	return *new(XdsUpdateClientConfigureServiceClient)
}

func (c *xdsUpdateClientConfigureServiceClient) Configure(ctx context.Context, in *ClientConfigureRequest, opts ...grpc.CallOption) (*ClientConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type XdsUpdateClientConfigureServiceServer interface {
	Configure(context.Context, *ClientConfigureRequest) (*ClientConfigureResponse, error)
	mustEmbedUnimplementedXdsUpdateClientConfigureServiceServer()
}

type UnimplementedXdsUpdateClientConfigureServiceServer struct{}

func (UnimplementedXdsUpdateClientConfigureServiceServer) Configure(context.Context, *ClientConfigureRequest) (*ClientConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedXdsUpdateClientConfigureServiceServer) mustEmbedUnimplementedXdsUpdateClientConfigureServiceServer() {
	_ = "STUB: not implemented"
	return
}

func (UnimplementedXdsUpdateClientConfigureServiceServer) testEmbeddedByValue() {
	_ = "STUB: not implemented"
	return
}

type UnsafeXdsUpdateClientConfigureServiceServer interface {
	mustEmbedUnimplementedXdsUpdateClientConfigureServiceServer()
}

func RegisterXdsUpdateClientConfigureServiceServer(s grpc.ServiceRegistrar, srv XdsUpdateClientConfigureServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _XdsUpdateClientConfigureService_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var XdsUpdateClientConfigureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.XdsUpdateClientConfigureService",
	HandlerType: (*XdsUpdateClientConfigureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _XdsUpdateClientConfigureService_Configure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/testing/test.proto",
}
