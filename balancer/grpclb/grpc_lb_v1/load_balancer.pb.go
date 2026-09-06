package grpc_lb_v1

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoadBalanceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	LoadBalanceRequestType isLoadBalanceRequest_LoadBalanceRequestType `protobuf_oneof:"load_balance_request_type"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *LoadBalanceRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadBalanceRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*LoadBalanceRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadBalanceRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalanceRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *LoadBalanceRequest) GetLoadBalanceRequestType() isLoadBalanceRequest_LoadBalanceRequestType {
	_ = "STUB: not implemented"
	return *new(isLoadBalanceRequest_LoadBalanceRequestType)
}

func (x *LoadBalanceRequest) GetInitialRequest() *InitialLoadBalanceRequest {
	_ = "STUB: not implemented"
	return nil
}

func (x *LoadBalanceRequest) GetClientStats() *ClientStats { _ = "STUB: not implemented"; return nil }

type isLoadBalanceRequest_LoadBalanceRequestType interface {
	isLoadBalanceRequest_LoadBalanceRequestType()
}

type LoadBalanceRequest_InitialRequest struct {
	InitialRequest *InitialLoadBalanceRequest `protobuf:"bytes,1,opt,name=initial_request,json=initialRequest,proto3,oneof"`
}

type LoadBalanceRequest_ClientStats struct {
	ClientStats *ClientStats `protobuf:"bytes,2,opt,name=client_stats,json=clientStats,proto3,oneof"`
}

func (*LoadBalanceRequest_InitialRequest) isLoadBalanceRequest_LoadBalanceRequestType() {
	_ = "STUB: not implemented"
	return
}

func (*LoadBalanceRequest_ClientStats) isLoadBalanceRequest_LoadBalanceRequestType() {
	_ = "STUB: not implemented"
	return
}

type InitialLoadBalanceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitialLoadBalanceRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *InitialLoadBalanceRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*InitialLoadBalanceRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *InitialLoadBalanceRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*InitialLoadBalanceRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *InitialLoadBalanceRequest) GetName() string { _ = "STUB: not implemented"; return "" }

type ClientStatsPerToken struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	LoadBalanceToken string `protobuf:"bytes,1,opt,name=load_balance_token,json=loadBalanceToken,proto3" json:"load_balance_token,omitempty"`

	NumCalls      int64 `protobuf:"varint,2,opt,name=num_calls,json=numCalls,proto3" json:"num_calls,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientStatsPerToken) Reset() { _ = "STUB: not implemented"; return }

func (x *ClientStatsPerToken) String() string { _ = "STUB: not implemented"; return "" }

func (*ClientStatsPerToken) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClientStatsPerToken) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClientStatsPerToken) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *ClientStatsPerToken) GetLoadBalanceToken() string { _ = "STUB: not implemented"; return "" }

func (x *ClientStatsPerToken) GetNumCalls() int64 { _ = "STUB: not implemented"; return 0 }

type ClientStats struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`

	NumCallsStarted int64 `protobuf:"varint,2,opt,name=num_calls_started,json=numCallsStarted,proto3" json:"num_calls_started,omitempty"`

	NumCallsFinished int64 `protobuf:"varint,3,opt,name=num_calls_finished,json=numCallsFinished,proto3" json:"num_calls_finished,omitempty"`

	NumCallsFinishedWithClientFailedToSend int64 `protobuf:"varint,6,opt,name=num_calls_finished_with_client_failed_to_send,json=numCallsFinishedWithClientFailedToSend,proto3" json:"num_calls_finished_with_client_failed_to_send,omitempty"`

	NumCallsFinishedKnownReceived int64 `protobuf:"varint,7,opt,name=num_calls_finished_known_received,json=numCallsFinishedKnownReceived,proto3" json:"num_calls_finished_known_received,omitempty"`

	CallsFinishedWithDrop []*ClientStatsPerToken `protobuf:"bytes,8,rep,name=calls_finished_with_drop,json=callsFinishedWithDrop,proto3" json:"calls_finished_with_drop,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ClientStats) Reset() { _ = "STUB: not implemented"; return }

func (x *ClientStats) String() string { _ = "STUB: not implemented"; return "" }

func (*ClientStats) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClientStats) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClientStats) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ClientStats) GetTimestamp() *timestamppb.Timestamp { _ = "STUB: not implemented"; return nil }

func (x *ClientStats) GetNumCallsStarted() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ClientStats) GetNumCallsFinished() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ClientStats) GetNumCallsFinishedWithClientFailedToSend() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (x *ClientStats) GetNumCallsFinishedKnownReceived() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ClientStats) GetCallsFinishedWithDrop() []*ClientStatsPerToken {
	_ = "STUB: not implemented"
	return nil
}

type LoadBalanceResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	LoadBalanceResponseType isLoadBalanceResponse_LoadBalanceResponseType `protobuf_oneof:"load_balance_response_type"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *LoadBalanceResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadBalanceResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*LoadBalanceResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadBalanceResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalanceResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *LoadBalanceResponse) GetLoadBalanceResponseType() isLoadBalanceResponse_LoadBalanceResponseType {
	_ = "STUB: not implemented"
	return *new(isLoadBalanceResponse_LoadBalanceResponseType)
}

func (x *LoadBalanceResponse) GetInitialResponse() *InitialLoadBalanceResponse {
	_ = "STUB: not implemented"
	return nil
}

func (x *LoadBalanceResponse) GetServerList() *ServerList { _ = "STUB: not implemented"; return nil }

func (x *LoadBalanceResponse) GetFallbackResponse() *FallbackResponse {
	_ = "STUB: not implemented"
	return nil
}

type isLoadBalanceResponse_LoadBalanceResponseType interface {
	isLoadBalanceResponse_LoadBalanceResponseType()
}

type LoadBalanceResponse_InitialResponse struct {
	InitialResponse *InitialLoadBalanceResponse `protobuf:"bytes,1,opt,name=initial_response,json=initialResponse,proto3,oneof"`
}

type LoadBalanceResponse_ServerList struct {
	ServerList *ServerList `protobuf:"bytes,2,opt,name=server_list,json=serverList,proto3,oneof"`
}

type LoadBalanceResponse_FallbackResponse struct {
	FallbackResponse *FallbackResponse `protobuf:"bytes,3,opt,name=fallback_response,json=fallbackResponse,proto3,oneof"`
}

func (*LoadBalanceResponse_InitialResponse) isLoadBalanceResponse_LoadBalanceResponseType() {
	_ = "STUB: not implemented"
	return
}

func (*LoadBalanceResponse_ServerList) isLoadBalanceResponse_LoadBalanceResponseType() {
	_ = "STUB: not implemented"
	return
}

func (*LoadBalanceResponse_FallbackResponse) isLoadBalanceResponse_LoadBalanceResponseType() {
	_ = "STUB: not implemented"
	return
}

type FallbackResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FallbackResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *FallbackResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*FallbackResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *FallbackResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*FallbackResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type InitialLoadBalanceResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ClientStatsReportInterval *durationpb.Duration `protobuf:"bytes,2,opt,name=client_stats_report_interval,json=clientStatsReportInterval,proto3" json:"client_stats_report_interval,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *InitialLoadBalanceResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *InitialLoadBalanceResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*InitialLoadBalanceResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *InitialLoadBalanceResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*InitialLoadBalanceResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *InitialLoadBalanceResponse) GetClientStatsReportInterval() *durationpb.Duration {
	_ = "STUB: not implemented"
	return nil
}

type ServerList struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Servers       []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerList) Reset() { _ = "STUB: not implemented"; return }

func (x *ServerList) String() string { _ = "STUB: not implemented"; return "" }

func (*ServerList) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServerList) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServerList) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ServerList) GetServers() []*Server { _ = "STUB: not implemented"; return nil }

type Server struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	IpAddress []byte `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`

	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`

	LoadBalanceToken string `protobuf:"bytes,3,opt,name=load_balance_token,json=loadBalanceToken,proto3" json:"load_balance_token,omitempty"`

	Drop          bool `protobuf:"varint,4,opt,name=drop,proto3" json:"drop,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() { _ = "STUB: not implemented"; return }

func (x *Server) String() string { _ = "STUB: not implemented"; return "" }

func (*Server) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Server) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Server) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Server) GetIpAddress() []byte { _ = "STUB: not implemented"; return nil }

func (x *Server) GetPort() int32 { _ = "STUB: not implemented"; return 0 }

func (x *Server) GetLoadBalanceToken() string { _ = "STUB: not implemented"; return "" }

func (x *Server) GetDrop() bool { _ = "STUB: not implemented"; return false }

var File_grpc_lb_v1_load_balancer_proto protoreflect.FileDescriptor

const file_grpc_lb_v1_load_balancer_proto_rawDesc = "" +
	"\n" +
	"\x1egrpc/lb/v1/load_balancer.proto\x12\n" +
	"grpc.lb.v1\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc1\x01\n" +
	"\x12LoadBalanceRequest\x12P\n" +
	"\x0finitial_request\x18\x01 \x01(\v2%.grpc.lb.v1.InitialLoadBalanceRequestH\x00R\x0einitialRequest\x12<\n" +
	"\fclient_stats\x18\x02 \x01(\v2\x17.grpc.lb.v1.ClientStatsH\x00R\vclientStatsB\x1b\n" +
	"\x19load_balance_request_type\"/\n" +
	"\x19InitialLoadBalanceRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"`\n" +
	"\x13ClientStatsPerToken\x12,\n" +
	"\x12load_balance_token\x18\x01 \x01(\tR\x10loadBalanceToken\x12\x1b\n" +
	"\tnum_calls\x18\x02 \x01(\x03R\bnumCalls\"\xb0\x03\n" +
	"\vClientStats\x128\n" +
	"\ttimestamp\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12*\n" +
	"\x11num_calls_started\x18\x02 \x01(\x03R\x0fnumCallsStarted\x12,\n" +
	"\x12num_calls_finished\x18\x03 \x01(\x03R\x10numCallsFinished\x12]\n" +
	"-num_calls_finished_with_client_failed_to_send\x18\x06 \x01(\x03R&numCallsFinishedWithClientFailedToSend\x12H\n" +
	"!num_calls_finished_known_received\x18\a \x01(\x03R\x1dnumCallsFinishedKnownReceived\x12X\n" +
	"\x18calls_finished_with_drop\x18\b \x03(\v2\x1f.grpc.lb.v1.ClientStatsPerTokenR\x15callsFinishedWithDropJ\x04\b\x04\x10\x05J\x04\b\x05\x10\x06\"\x90\x02\n" +
	"\x13LoadBalanceResponse\x12S\n" +
	"\x10initial_response\x18\x01 \x01(\v2&.grpc.lb.v1.InitialLoadBalanceResponseH\x00R\x0finitialResponse\x129\n" +
	"\vserver_list\x18\x02 \x01(\v2\x16.grpc.lb.v1.ServerListH\x00R\n" +
	"serverList\x12K\n" +
	"\x11fallback_response\x18\x03 \x01(\v2\x1c.grpc.lb.v1.FallbackResponseH\x00R\x10fallbackResponseB\x1c\n" +
	"\x1aload_balance_response_type\"\x12\n" +
	"\x10FallbackResponse\"~\n" +
	"\x1aInitialLoadBalanceResponse\x12Z\n" +
	"\x1cclient_stats_report_interval\x18\x02 \x01(\v2\x19.google.protobuf.DurationR\x19clientStatsReportIntervalJ\x04\b\x01\x10\x02\"@\n" +
	"\n" +
	"ServerList\x12,\n" +
	"\aservers\x18\x01 \x03(\v2\x12.grpc.lb.v1.ServerR\aserversJ\x04\b\x03\x10\x04\"\x83\x01\n" +
	"\x06Server\x12\x1d\n" +
	"\n" +
	"ip_address\x18\x01 \x01(\fR\tipAddress\x12\x12\n" +
	"\x04port\x18\x02 \x01(\x05R\x04port\x12,\n" +
	"\x12load_balance_token\x18\x03 \x01(\tR\x10loadBalanceToken\x12\x12\n" +
	"\x04drop\x18\x04 \x01(\bR\x04dropJ\x04\b\x05\x10\x062b\n" +
	"\fLoadBalancer\x12R\n" +
	"\vBalanceLoad\x12\x1e.grpc.lb.v1.LoadBalanceRequest\x1a\x1f.grpc.lb.v1.LoadBalanceResponse(\x010\x01BW\n" +
	"\rio.grpc.lb.v1B\x11LoadBalancerProtoP\x01Z1google.golang.org/grpc/balancer/grpclb/grpc_lb_v1b\x06proto3"

var (
	file_grpc_lb_v1_load_balancer_proto_rawDescOnce sync.Once
	file_grpc_lb_v1_load_balancer_proto_rawDescData []byte
)

func file_grpc_lb_v1_load_balancer_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_grpc_lb_v1_load_balancer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_grpc_lb_v1_load_balancer_proto_goTypes = []any{
	(*LoadBalanceRequest)(nil),
	(*InitialLoadBalanceRequest)(nil),
	(*ClientStatsPerToken)(nil),
	(*ClientStats)(nil),
	(*LoadBalanceResponse)(nil),
	(*FallbackResponse)(nil),
	(*InitialLoadBalanceResponse)(nil),
	(*ServerList)(nil),
	(*Server)(nil),
	(*timestamppb.Timestamp)(nil),
	(*durationpb.Duration)(nil),
}
var file_grpc_lb_v1_load_balancer_proto_depIdxs = []int32{
	1,
	3,
	9,
	2,
	6,
	7,
	5,
	10,
	8,
	0,
	4,
	10,
	9,
	9,
	9,
	0,
}

func init()                                     { file_grpc_lb_v1_load_balancer_proto_init() }
func file_grpc_lb_v1_load_balancer_proto_init() { _ = "STUB: not implemented"; return }
