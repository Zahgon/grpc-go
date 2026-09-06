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

type PayloadType int32

const (
	PayloadType_COMPRESSABLE PayloadType = 0
)

var (
	PayloadType_name = map[int32]string{
		0: "COMPRESSABLE",
	}
	PayloadType_value = map[string]int32{
		"COMPRESSABLE": 0,
	}
)

func (x PayloadType) Enum() *PayloadType { _ = "STUB: not implemented"; return nil }

func (x PayloadType) String() string { _ = "STUB: not implemented"; return "" }

func (PayloadType) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (PayloadType) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x PayloadType) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (PayloadType) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type GrpclbRouteType int32

const (
	GrpclbRouteType_GRPCLB_ROUTE_TYPE_UNKNOWN GrpclbRouteType = 0

	GrpclbRouteType_GRPCLB_ROUTE_TYPE_FALLBACK GrpclbRouteType = 1

	GrpclbRouteType_GRPCLB_ROUTE_TYPE_BACKEND GrpclbRouteType = 2
)

var (
	GrpclbRouteType_name = map[int32]string{
		0: "GRPCLB_ROUTE_TYPE_UNKNOWN",
		1: "GRPCLB_ROUTE_TYPE_FALLBACK",
		2: "GRPCLB_ROUTE_TYPE_BACKEND",
	}
	GrpclbRouteType_value = map[string]int32{
		"GRPCLB_ROUTE_TYPE_UNKNOWN":  0,
		"GRPCLB_ROUTE_TYPE_FALLBACK": 1,
		"GRPCLB_ROUTE_TYPE_BACKEND":  2,
	}
)

func (x GrpclbRouteType) Enum() *GrpclbRouteType { _ = "STUB: not implemented"; return nil }

func (x GrpclbRouteType) String() string { _ = "STUB: not implemented"; return "" }

func (GrpclbRouteType) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (GrpclbRouteType) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x GrpclbRouteType) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (GrpclbRouteType) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type LoadBalancerStatsResponse_MetadataType int32

const (
	LoadBalancerStatsResponse_UNKNOWN  LoadBalancerStatsResponse_MetadataType = 0
	LoadBalancerStatsResponse_INITIAL  LoadBalancerStatsResponse_MetadataType = 1
	LoadBalancerStatsResponse_TRAILING LoadBalancerStatsResponse_MetadataType = 2
)

var (
	LoadBalancerStatsResponse_MetadataType_name = map[int32]string{
		0: "UNKNOWN",
		1: "INITIAL",
		2: "TRAILING",
	}
	LoadBalancerStatsResponse_MetadataType_value = map[string]int32{
		"UNKNOWN":  0,
		"INITIAL":  1,
		"TRAILING": 2,
	}
)

func (x LoadBalancerStatsResponse_MetadataType) Enum() *LoadBalancerStatsResponse_MetadataType {
	_ = "STUB: not implemented"
	return nil
}

func (x LoadBalancerStatsResponse_MetadataType) String() string {
	_ = "STUB: not implemented"
	return ""
}

func (LoadBalancerStatsResponse_MetadataType) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (LoadBalancerStatsResponse_MetadataType) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x LoadBalancerStatsResponse_MetadataType) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (LoadBalancerStatsResponse_MetadataType) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ClientConfigureRequest_RpcType int32

const (
	ClientConfigureRequest_EMPTY_CALL ClientConfigureRequest_RpcType = 0
	ClientConfigureRequest_UNARY_CALL ClientConfigureRequest_RpcType = 1
)

var (
	ClientConfigureRequest_RpcType_name = map[int32]string{
		0: "EMPTY_CALL",
		1: "UNARY_CALL",
	}
	ClientConfigureRequest_RpcType_value = map[string]int32{
		"EMPTY_CALL": 0,
		"UNARY_CALL": 1,
	}
)

func (x ClientConfigureRequest_RpcType) Enum() *ClientConfigureRequest_RpcType {
	_ = "STUB: not implemented"
	return nil
}

func (x ClientConfigureRequest_RpcType) String() string { _ = "STUB: not implemented"; return "" }

func (ClientConfigureRequest_RpcType) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (ClientConfigureRequest_RpcType) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x ClientConfigureRequest_RpcType) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (ClientConfigureRequest_RpcType) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type HookRequest_HookRequestCommand int32

const (
	HookRequest_UNSPECIFIED HookRequest_HookRequestCommand = 0

	HookRequest_START HookRequest_HookRequestCommand = 1

	HookRequest_STOP HookRequest_HookRequestCommand = 2

	HookRequest_RETURN HookRequest_HookRequestCommand = 3
)

var (
	HookRequest_HookRequestCommand_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "START",
		2: "STOP",
		3: "RETURN",
	}
	HookRequest_HookRequestCommand_value = map[string]int32{
		"UNSPECIFIED": 0,
		"START":       1,
		"STOP":        2,
		"RETURN":      3,
	}
)

func (x HookRequest_HookRequestCommand) Enum() *HookRequest_HookRequestCommand {
	_ = "STUB: not implemented"
	return nil
}

func (x HookRequest_HookRequestCommand) String() string { _ = "STUB: not implemented"; return "" }

func (HookRequest_HookRequestCommand) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (HookRequest_HookRequestCommand) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x HookRequest_HookRequestCommand) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (HookRequest_HookRequestCommand) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type BoolValue struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Value         bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BoolValue) Reset() { _ = "STUB: not implemented"; return }

func (x *BoolValue) String() string { _ = "STUB: not implemented"; return "" }

func (*BoolValue) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *BoolValue) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*BoolValue) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *BoolValue) GetValue() bool { _ = "STUB: not implemented"; return false }

type Payload struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Type PayloadType `protobuf:"varint,1,opt,name=type,proto3,enum=grpc.testing.PayloadType" json:"type,omitempty"`

	Body          []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Payload) Reset() { _ = "STUB: not implemented"; return }

func (x *Payload) String() string { _ = "STUB: not implemented"; return "" }

func (*Payload) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Payload) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Payload) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Payload) GetType() PayloadType { _ = "STUB: not implemented"; return *new(PayloadType) }

func (x *Payload) GetBody() []byte { _ = "STUB: not implemented"; return nil }

type EchoStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EchoStatus) Reset() { _ = "STUB: not implemented"; return }

func (x *EchoStatus) String() string { _ = "STUB: not implemented"; return "" }

func (*EchoStatus) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *EchoStatus) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*EchoStatus) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *EchoStatus) GetCode() int32 { _ = "STUB: not implemented"; return 0 }

func (x *EchoStatus) GetMessage() string { _ = "STUB: not implemented"; return "" }

type SimpleRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ResponseType PayloadType `protobuf:"varint,1,opt,name=response_type,json=responseType,proto3,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`

	ResponseSize int32 `protobuf:"varint,2,opt,name=response_size,json=responseSize,proto3" json:"response_size,omitempty"`

	Payload *Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`

	FillUsername bool `protobuf:"varint,4,opt,name=fill_username,json=fillUsername,proto3" json:"fill_username,omitempty"`

	FillOauthScope bool `protobuf:"varint,5,opt,name=fill_oauth_scope,json=fillOauthScope,proto3" json:"fill_oauth_scope,omitempty"`

	ResponseCompressed *BoolValue `protobuf:"bytes,6,opt,name=response_compressed,json=responseCompressed,proto3" json:"response_compressed,omitempty"`

	ResponseStatus *EchoStatus `protobuf:"bytes,7,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`

	ExpectCompressed *BoolValue `protobuf:"bytes,8,opt,name=expect_compressed,json=expectCompressed,proto3" json:"expect_compressed,omitempty"`

	FillServerId bool `protobuf:"varint,9,opt,name=fill_server_id,json=fillServerId,proto3" json:"fill_server_id,omitempty"`

	FillGrpclbRouteType bool `protobuf:"varint,10,opt,name=fill_grpclb_route_type,json=fillGrpclbRouteType,proto3" json:"fill_grpclb_route_type,omitempty"`

	OrcaPerQueryReport *TestOrcaReport `protobuf:"bytes,11,opt,name=orca_per_query_report,json=orcaPerQueryReport,proto3" json:"orca_per_query_report,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *SimpleRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *SimpleRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*SimpleRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SimpleRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SimpleRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SimpleRequest) GetResponseType() PayloadType {
	_ = "STUB: not implemented"
	return *new(PayloadType)
}

func (x *SimpleRequest) GetResponseSize() int32 { _ = "STUB: not implemented"; return 0 }

func (x *SimpleRequest) GetPayload() *Payload { _ = "STUB: not implemented"; return nil }

func (x *SimpleRequest) GetFillUsername() bool { _ = "STUB: not implemented"; return false }

func (x *SimpleRequest) GetFillOauthScope() bool { _ = "STUB: not implemented"; return false }

func (x *SimpleRequest) GetResponseCompressed() *BoolValue { _ = "STUB: not implemented"; return nil }

func (x *SimpleRequest) GetResponseStatus() *EchoStatus { _ = "STUB: not implemented"; return nil }

func (x *SimpleRequest) GetExpectCompressed() *BoolValue { _ = "STUB: not implemented"; return nil }

func (x *SimpleRequest) GetFillServerId() bool { _ = "STUB: not implemented"; return false }

func (x *SimpleRequest) GetFillGrpclbRouteType() bool { _ = "STUB: not implemented"; return false }

func (x *SimpleRequest) GetOrcaPerQueryReport() *TestOrcaReport {
	_ = "STUB: not implemented"
	return nil
}

type SimpleResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Payload *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`

	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`

	OauthScope string `protobuf:"bytes,3,opt,name=oauth_scope,json=oauthScope,proto3" json:"oauth_scope,omitempty"`

	ServerId string `protobuf:"bytes,4,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`

	GrpclbRouteType GrpclbRouteType `protobuf:"varint,5,opt,name=grpclb_route_type,json=grpclbRouteType,proto3,enum=grpc.testing.GrpclbRouteType" json:"grpclb_route_type,omitempty"`

	Hostname      string `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SimpleResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *SimpleResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*SimpleResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SimpleResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SimpleResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SimpleResponse) GetPayload() *Payload { _ = "STUB: not implemented"; return nil }

func (x *SimpleResponse) GetUsername() string { _ = "STUB: not implemented"; return "" }

func (x *SimpleResponse) GetOauthScope() string { _ = "STUB: not implemented"; return "" }

func (x *SimpleResponse) GetServerId() string { _ = "STUB: not implemented"; return "" }

func (x *SimpleResponse) GetGrpclbRouteType() GrpclbRouteType {
	_ = "STUB: not implemented"
	return *new(GrpclbRouteType)
}

func (x *SimpleResponse) GetHostname() string { _ = "STUB: not implemented"; return "" }

type StreamingInputCallRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Payload *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`

	ExpectCompressed *BoolValue `protobuf:"bytes,2,opt,name=expect_compressed,json=expectCompressed,proto3" json:"expect_compressed,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *StreamingInputCallRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *StreamingInputCallRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*StreamingInputCallRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *StreamingInputCallRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*StreamingInputCallRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *StreamingInputCallRequest) GetPayload() *Payload { _ = "STUB: not implemented"; return nil }

func (x *StreamingInputCallRequest) GetExpectCompressed() *BoolValue {
	_ = "STUB: not implemented"
	return nil
}

type StreamingInputCallResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	AggregatedPayloadSize int32 `protobuf:"varint,1,opt,name=aggregated_payload_size,json=aggregatedPayloadSize,proto3" json:"aggregated_payload_size,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *StreamingInputCallResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *StreamingInputCallResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*StreamingInputCallResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *StreamingInputCallResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*StreamingInputCallResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *StreamingInputCallResponse) GetAggregatedPayloadSize() int32 {
	_ = "STUB: not implemented"
	return 0
}

type ResponseParameters struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`

	IntervalUs int32 `protobuf:"varint,2,opt,name=interval_us,json=intervalUs,proto3" json:"interval_us,omitempty"`

	Compressed *BoolValue `protobuf:"bytes,3,opt,name=compressed,proto3" json:"compressed,omitempty"`

	FillPeerSocketAddress *BoolValue `protobuf:"bytes,4,opt,name=fill_peer_socket_address,json=fillPeerSocketAddress,proto3" json:"fill_peer_socket_address,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ResponseParameters) Reset() { _ = "STUB: not implemented"; return }

func (x *ResponseParameters) String() string { _ = "STUB: not implemented"; return "" }

func (*ResponseParameters) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ResponseParameters) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ResponseParameters) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ResponseParameters) GetSize() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ResponseParameters) GetIntervalUs() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ResponseParameters) GetCompressed() *BoolValue { _ = "STUB: not implemented"; return nil }

func (x *ResponseParameters) GetFillPeerSocketAddress() *BoolValue {
	_ = "STUB: not implemented"
	return nil
}

type StreamingOutputCallRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ResponseType PayloadType `protobuf:"varint,1,opt,name=response_type,json=responseType,proto3,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`

	ResponseParameters []*ResponseParameters `protobuf:"bytes,2,rep,name=response_parameters,json=responseParameters,proto3" json:"response_parameters,omitempty"`

	Payload *Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`

	ResponseStatus *EchoStatus `protobuf:"bytes,7,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`

	OrcaOobReport *TestOrcaReport `protobuf:"bytes,8,opt,name=orca_oob_report,json=orcaOobReport,proto3" json:"orca_oob_report,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamingOutputCallRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *StreamingOutputCallRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*StreamingOutputCallRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *StreamingOutputCallRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*StreamingOutputCallRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *StreamingOutputCallRequest) GetResponseType() PayloadType {
	_ = "STUB: not implemented"
	return *new(PayloadType)
}

func (x *StreamingOutputCallRequest) GetResponseParameters() []*ResponseParameters {
	_ = "STUB: not implemented"
	return nil
}

func (x *StreamingOutputCallRequest) GetPayload() *Payload { _ = "STUB: not implemented"; return nil }

func (x *StreamingOutputCallRequest) GetResponseStatus() *EchoStatus {
	_ = "STUB: not implemented"
	return nil
}

func (x *StreamingOutputCallRequest) GetOrcaOobReport() *TestOrcaReport {
	_ = "STUB: not implemented"
	return nil
}

type StreamingOutputCallResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Payload *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`

	PeerSocketAddress string `protobuf:"bytes,2,opt,name=peer_socket_address,json=peerSocketAddress,proto3" json:"peer_socket_address,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *StreamingOutputCallResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *StreamingOutputCallResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*StreamingOutputCallResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *StreamingOutputCallResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*StreamingOutputCallResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *StreamingOutputCallResponse) GetPayload() *Payload { _ = "STUB: not implemented"; return nil }

func (x *StreamingOutputCallResponse) GetPeerSocketAddress() string {
	_ = "STUB: not implemented"
	return ""
}

type ReconnectParams struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	MaxReconnectBackoffMs int32                  `protobuf:"varint,1,opt,name=max_reconnect_backoff_ms,json=maxReconnectBackoffMs,proto3" json:"max_reconnect_backoff_ms,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ReconnectParams) Reset() { _ = "STUB: not implemented"; return }

func (x *ReconnectParams) String() string { _ = "STUB: not implemented"; return "" }

func (*ReconnectParams) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ReconnectParams) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ReconnectParams) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ReconnectParams) GetMaxReconnectBackoffMs() int32 { _ = "STUB: not implemented"; return 0 }

type ReconnectInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Passed        bool                   `protobuf:"varint,1,opt,name=passed,proto3" json:"passed,omitempty"`
	BackoffMs     []int32                `protobuf:"varint,2,rep,packed,name=backoff_ms,json=backoffMs,proto3" json:"backoff_ms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReconnectInfo) Reset() { _ = "STUB: not implemented"; return }

func (x *ReconnectInfo) String() string { _ = "STUB: not implemented"; return "" }

func (*ReconnectInfo) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ReconnectInfo) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ReconnectInfo) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ReconnectInfo) GetPassed() bool { _ = "STUB: not implemented"; return false }

func (x *ReconnectInfo) GetBackoffMs() []int32 { _ = "STUB: not implemented"; return nil }

type LoadBalancerStatsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	NumRpcs int32 `protobuf:"varint,1,opt,name=num_rpcs,json=numRpcs,proto3" json:"num_rpcs,omitempty"`

	TimeoutSec int32 `protobuf:"varint,2,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`

	MetadataKeys  []string `protobuf:"bytes,3,rep,name=metadata_keys,json=metadataKeys,proto3" json:"metadata_keys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoadBalancerStatsRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*LoadBalancerStatsRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalancerStatsRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *LoadBalancerStatsRequest) GetNumRpcs() int32 { _ = "STUB: not implemented"; return 0 }

func (x *LoadBalancerStatsRequest) GetTimeoutSec() int32 { _ = "STUB: not implemented"; return 0 }

func (x *LoadBalancerStatsRequest) GetMetadataKeys() []string {
	_ = "STUB: not implemented"
	return nil
}

type LoadBalancerStatsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	RpcsByPeer map[string]int32 `protobuf:"bytes,1,rep,name=rpcs_by_peer,json=rpcsByPeer,proto3" json:"rpcs_by_peer,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`

	NumFailures  int32                                            `protobuf:"varint,2,opt,name=num_failures,json=numFailures,proto3" json:"num_failures,omitempty"`
	RpcsByMethod map[string]*LoadBalancerStatsResponse_RpcsByPeer `protobuf:"bytes,3,rep,name=rpcs_by_method,json=rpcsByMethod,proto3" json:"rpcs_by_method,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`

	MetadatasByPeer map[string]*LoadBalancerStatsResponse_MetadataByPeer `protobuf:"bytes,4,rep,name=metadatas_by_peer,json=metadatasByPeer,proto3" json:"metadatas_by_peer,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *LoadBalancerStatsResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*LoadBalancerStatsResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalancerStatsResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *LoadBalancerStatsResponse) GetRpcsByPeer() map[string]int32 {
	_ = "STUB: not implemented"
	return nil
}

func (x *LoadBalancerStatsResponse) GetNumFailures() int32 { _ = "STUB: not implemented"; return 0 }

func (x *LoadBalancerStatsResponse) GetRpcsByMethod() map[string]*LoadBalancerStatsResponse_RpcsByPeer {
	_ = "STUB: not implemented"
	return nil
}

func (x *LoadBalancerStatsResponse) GetMetadatasByPeer() map[string]*LoadBalancerStatsResponse_MetadataByPeer {
	_ = "STUB: not implemented"
	return nil
}

type LoadBalancerAccumulatedStatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoadBalancerAccumulatedStatsRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerAccumulatedStatsRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*LoadBalancerAccumulatedStatsRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerAccumulatedStatsRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalancerAccumulatedStatsRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type LoadBalancerAccumulatedStatsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	NumRpcsStartedByMethod map[string]int32 `protobuf:"bytes,1,rep,name=num_rpcs_started_by_method,json=numRpcsStartedByMethod,proto3" json:"num_rpcs_started_by_method,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`

	NumRpcsSucceededByMethod map[string]int32 `protobuf:"bytes,2,rep,name=num_rpcs_succeeded_by_method,json=numRpcsSucceededByMethod,proto3" json:"num_rpcs_succeeded_by_method,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`

	NumRpcsFailedByMethod map[string]int32 `protobuf:"bytes,3,rep,name=num_rpcs_failed_by_method,json=numRpcsFailedByMethod,proto3" json:"num_rpcs_failed_by_method,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`

	StatsPerMethod map[string]*LoadBalancerAccumulatedStatsResponse_MethodStats `protobuf:"bytes,4,rep,name=stats_per_method,json=statsPerMethod,proto3" json:"stats_per_method,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *LoadBalancerAccumulatedStatsResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerAccumulatedStatsResponse) String() string {
	_ = "STUB: not implemented"
	return ""
}

func (*LoadBalancerAccumulatedStatsResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerAccumulatedStatsResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalancerAccumulatedStatsResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *LoadBalancerAccumulatedStatsResponse) GetNumRpcsStartedByMethod() map[string]int32 {
	_ = "STUB: not implemented"
	return nil
}

func (x *LoadBalancerAccumulatedStatsResponse) GetNumRpcsSucceededByMethod() map[string]int32 {
	_ = "STUB: not implemented"
	return nil
}

func (x *LoadBalancerAccumulatedStatsResponse) GetNumRpcsFailedByMethod() map[string]int32 {
	_ = "STUB: not implemented"
	return nil
}

func (x *LoadBalancerAccumulatedStatsResponse) GetStatsPerMethod() map[string]*LoadBalancerAccumulatedStatsResponse_MethodStats {
	_ = "STUB: not implemented"
	return nil
}

type ClientConfigureRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Types []ClientConfigureRequest_RpcType `protobuf:"varint,1,rep,packed,name=types,proto3,enum=grpc.testing.ClientConfigureRequest_RpcType" json:"types,omitempty"`

	Metadata []*ClientConfigureRequest_Metadata `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty"`

	TimeoutSec    int32 `protobuf:"varint,3,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientConfigureRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *ClientConfigureRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*ClientConfigureRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClientConfigureRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClientConfigureRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *ClientConfigureRequest) GetTypes() []ClientConfigureRequest_RpcType {
	_ = "STUB: not implemented"
	return nil
}

func (x *ClientConfigureRequest) GetMetadata() []*ClientConfigureRequest_Metadata {
	_ = "STUB: not implemented"
	return nil
}

func (x *ClientConfigureRequest) GetTimeoutSec() int32 { _ = "STUB: not implemented"; return 0 }

type ClientConfigureResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientConfigureResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *ClientConfigureResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*ClientConfigureResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClientConfigureResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClientConfigureResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MemorySize struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rss           int64                  `protobuf:"varint,1,opt,name=rss,proto3" json:"rss,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemorySize) Reset() { _ = "STUB: not implemented"; return }

func (x *MemorySize) String() string { _ = "STUB: not implemented"; return "" }

func (*MemorySize) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *MemorySize) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*MemorySize) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *MemorySize) GetRss() int64 { _ = "STUB: not implemented"; return 0 }

type TestOrcaReport struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CpuUtilization    float64                `protobuf:"fixed64,1,opt,name=cpu_utilization,json=cpuUtilization,proto3" json:"cpu_utilization,omitempty"`
	MemoryUtilization float64                `protobuf:"fixed64,2,opt,name=memory_utilization,json=memoryUtilization,proto3" json:"memory_utilization,omitempty"`
	RequestCost       map[string]float64     `protobuf:"bytes,3,rep,name=request_cost,json=requestCost,proto3" json:"request_cost,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	Utilization       map[string]float64     `protobuf:"bytes,4,rep,name=utilization,proto3" json:"utilization,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *TestOrcaReport) Reset() { _ = "STUB: not implemented"; return }

func (x *TestOrcaReport) String() string { _ = "STUB: not implemented"; return "" }

func (*TestOrcaReport) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *TestOrcaReport) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*TestOrcaReport) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *TestOrcaReport) GetCpuUtilization() float64 { _ = "STUB: not implemented"; return 0 }

func (x *TestOrcaReport) GetMemoryUtilization() float64 { _ = "STUB: not implemented"; return 0 }

func (x *TestOrcaReport) GetRequestCost() map[string]float64 { _ = "STUB: not implemented"; return nil }

func (x *TestOrcaReport) GetUtilization() map[string]float64 { _ = "STUB: not implemented"; return nil }

type SetReturnStatusRequest struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	GrpcCodeToReturn      int32                  `protobuf:"varint,1,opt,name=grpc_code_to_return,json=grpcCodeToReturn,proto3" json:"grpc_code_to_return,omitempty"`
	GrpcStatusDescription string                 `protobuf:"bytes,2,opt,name=grpc_status_description,json=grpcStatusDescription,proto3" json:"grpc_status_description,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *SetReturnStatusRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *SetReturnStatusRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*SetReturnStatusRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SetReturnStatusRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SetReturnStatusRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *SetReturnStatusRequest) GetGrpcCodeToReturn() int32 { _ = "STUB: not implemented"; return 0 }

func (x *SetReturnStatusRequest) GetGrpcStatusDescription() string {
	_ = "STUB: not implemented"
	return ""
}

type HookRequest struct {
	state                 protoimpl.MessageState         `protogen:"open.v1"`
	Command               HookRequest_HookRequestCommand `protobuf:"varint,1,opt,name=command,proto3,enum=grpc.testing.HookRequest_HookRequestCommand" json:"command,omitempty"`
	GrpcCodeToReturn      int32                          `protobuf:"varint,2,opt,name=grpc_code_to_return,json=grpcCodeToReturn,proto3" json:"grpc_code_to_return,omitempty"`
	GrpcStatusDescription string                         `protobuf:"bytes,3,opt,name=grpc_status_description,json=grpcStatusDescription,proto3" json:"grpc_status_description,omitempty"`

	ServerPort    int32 `protobuf:"varint,4,opt,name=server_port,json=serverPort,proto3" json:"server_port,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HookRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *HookRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*HookRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HookRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HookRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HookRequest) GetCommand() HookRequest_HookRequestCommand {
	_ = "STUB: not implemented"
	return *new(HookRequest_HookRequestCommand)
}

func (x *HookRequest) GetGrpcCodeToReturn() int32 { _ = "STUB: not implemented"; return 0 }

func (x *HookRequest) GetGrpcStatusDescription() string { _ = "STUB: not implemented"; return "" }

func (x *HookRequest) GetServerPort() int32 { _ = "STUB: not implemented"; return 0 }

type HookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HookResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *HookResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*HookResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HookResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HookResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type LoadBalancerStatsResponse_MetadataEntry struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`

	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`

	Type          LoadBalancerStatsResponse_MetadataType `protobuf:"varint,3,opt,name=type,proto3,enum=grpc.testing.LoadBalancerStatsResponse_MetadataType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoadBalancerStatsResponse_MetadataEntry) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsResponse_MetadataEntry) String() string {
	_ = "STUB: not implemented"
	return ""
}

func (*LoadBalancerStatsResponse_MetadataEntry) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsResponse_MetadataEntry) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalancerStatsResponse_MetadataEntry) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *LoadBalancerStatsResponse_MetadataEntry) GetKey() string {
	_ = "STUB: not implemented"
	return ""
}

func (x *LoadBalancerStatsResponse_MetadataEntry) GetValue() string {
	_ = "STUB: not implemented"
	return ""
}

func (x *LoadBalancerStatsResponse_MetadataEntry) GetType() LoadBalancerStatsResponse_MetadataType {
	_ = "STUB: not implemented"
	return *new(LoadBalancerStatsResponse_MetadataType)
}

type LoadBalancerStatsResponse_RpcMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Metadata      []*LoadBalancerStatsResponse_MetadataEntry `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoadBalancerStatsResponse_RpcMetadata) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsResponse_RpcMetadata) String() string {
	_ = "STUB: not implemented"
	return ""
}

func (*LoadBalancerStatsResponse_RpcMetadata) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsResponse_RpcMetadata) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalancerStatsResponse_RpcMetadata) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *LoadBalancerStatsResponse_RpcMetadata) GetMetadata() []*LoadBalancerStatsResponse_MetadataEntry {
	_ = "STUB: not implemented"
	return nil
}

type LoadBalancerStatsResponse_MetadataByPeer struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	RpcMetadata   []*LoadBalancerStatsResponse_RpcMetadata `protobuf:"bytes,1,rep,name=rpc_metadata,json=rpcMetadata,proto3" json:"rpc_metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoadBalancerStatsResponse_MetadataByPeer) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsResponse_MetadataByPeer) String() string {
	_ = "STUB: not implemented"
	return ""
}

func (*LoadBalancerStatsResponse_MetadataByPeer) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsResponse_MetadataByPeer) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalancerStatsResponse_MetadataByPeer) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *LoadBalancerStatsResponse_MetadataByPeer) GetRpcMetadata() []*LoadBalancerStatsResponse_RpcMetadata {
	_ = "STUB: not implemented"
	return nil
}

type LoadBalancerStatsResponse_RpcsByPeer struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	RpcsByPeer    map[string]int32 `protobuf:"bytes,1,rep,name=rpcs_by_peer,json=rpcsByPeer,proto3" json:"rpcs_by_peer,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoadBalancerStatsResponse_RpcsByPeer) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsResponse_RpcsByPeer) String() string {
	_ = "STUB: not implemented"
	return ""
}

func (*LoadBalancerStatsResponse_RpcsByPeer) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadBalancerStatsResponse_RpcsByPeer) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalancerStatsResponse_RpcsByPeer) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *LoadBalancerStatsResponse_RpcsByPeer) GetRpcsByPeer() map[string]int32 {
	_ = "STUB: not implemented"
	return nil
}

type LoadBalancerAccumulatedStatsResponse_MethodStats struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	RpcsStarted int32 `protobuf:"varint,1,opt,name=rpcs_started,json=rpcsStarted,proto3" json:"rpcs_started,omitempty"`

	Result        map[int32]int32 `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoadBalancerAccumulatedStatsResponse_MethodStats) Reset() {
	_ = "STUB: not implemented"
	return
}

func (x *LoadBalancerAccumulatedStatsResponse_MethodStats) String() string {
	_ = "STUB: not implemented"
	return ""
}

func (*LoadBalancerAccumulatedStatsResponse_MethodStats) ProtoMessage() {
	_ = "STUB: not implemented"
	return
}

func (x *LoadBalancerAccumulatedStatsResponse_MethodStats) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadBalancerAccumulatedStatsResponse_MethodStats) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *LoadBalancerAccumulatedStatsResponse_MethodStats) GetRpcsStarted() int32 {
	_ = "STUB: not implemented"
	return 0
}

func (x *LoadBalancerAccumulatedStatsResponse_MethodStats) GetResult() map[int32]int32 {
	_ = "STUB: not implemented"
	return nil
}

type ClientConfigureRequest_Metadata struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Type          ClientConfigureRequest_RpcType `protobuf:"varint,1,opt,name=type,proto3,enum=grpc.testing.ClientConfigureRequest_RpcType" json:"type,omitempty"`
	Key           string                         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                         `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientConfigureRequest_Metadata) Reset() { _ = "STUB: not implemented"; return }

func (x *ClientConfigureRequest_Metadata) String() string { _ = "STUB: not implemented"; return "" }

func (*ClientConfigureRequest_Metadata) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClientConfigureRequest_Metadata) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClientConfigureRequest_Metadata) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *ClientConfigureRequest_Metadata) GetType() ClientConfigureRequest_RpcType {
	_ = "STUB: not implemented"
	return *new(ClientConfigureRequest_RpcType)
}

func (x *ClientConfigureRequest_Metadata) GetKey() string { _ = "STUB: not implemented"; return "" }

func (x *ClientConfigureRequest_Metadata) GetValue() string { _ = "STUB: not implemented"; return "" }

var File_grpc_testing_messages_proto protoreflect.FileDescriptor

const file_grpc_testing_messages_proto_rawDesc = "" +
	"\n" +
	"\x1bgrpc/testing/messages.proto\x12\fgrpc.testing\"!\n" +
	"\tBoolValue\x12\x14\n" +
	"\x05value\x18\x01 \x01(\bR\x05value\"L\n" +
	"\aPayload\x12-\n" +
	"\x04type\x18\x01 \x01(\x0e2\x19.grpc.testing.PayloadTypeR\x04type\x12\x12\n" +
	"\x04body\x18\x02 \x01(\fR\x04body\":\n" +
	"\n" +
	"EchoStatus\x12\x12\n" +
	"\x04code\x18\x01 \x01(\x05R\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"\xf3\x04\n" +
	"\rSimpleRequest\x12>\n" +
	"\rresponse_type\x18\x01 \x01(\x0e2\x19.grpc.testing.PayloadTypeR\fresponseType\x12#\n" +
	"\rresponse_size\x18\x02 \x01(\x05R\fresponseSize\x12/\n" +
	"\apayload\x18\x03 \x01(\v2\x15.grpc.testing.PayloadR\apayload\x12#\n" +
	"\rfill_username\x18\x04 \x01(\bR\ffillUsername\x12(\n" +
	"\x10fill_oauth_scope\x18\x05 \x01(\bR\x0efillOauthScope\x12H\n" +
	"\x13response_compressed\x18\x06 \x01(\v2\x17.grpc.testing.BoolValueR\x12responseCompressed\x12A\n" +
	"\x0fresponse_status\x18\a \x01(\v2\x18.grpc.testing.EchoStatusR\x0eresponseStatus\x12D\n" +
	"\x11expect_compressed\x18\b \x01(\v2\x17.grpc.testing.BoolValueR\x10expectCompressed\x12$\n" +
	"\x0efill_server_id\x18\t \x01(\bR\ffillServerId\x123\n" +
	"\x16fill_grpclb_route_type\x18\n" +
	" \x01(\bR\x13fillGrpclbRouteType\x12O\n" +
	"\x15orca_per_query_report\x18\v \x01(\v2\x1c.grpc.testing.TestOrcaReportR\x12orcaPerQueryReport\"\x82\x02\n" +
	"\x0eSimpleResponse\x12/\n" +
	"\apayload\x18\x01 \x01(\v2\x15.grpc.testing.PayloadR\apayload\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x1f\n" +
	"\voauth_scope\x18\x03 \x01(\tR\n" +
	"oauthScope\x12\x1b\n" +
	"\tserver_id\x18\x04 \x01(\tR\bserverId\x12I\n" +
	"\x11grpclb_route_type\x18\x05 \x01(\x0e2\x1d.grpc.testing.GrpclbRouteTypeR\x0fgrpclbRouteType\x12\x1a\n" +
	"\bhostname\x18\x06 \x01(\tR\bhostname\"\x92\x01\n" +
	"\x19StreamingInputCallRequest\x12/\n" +
	"\apayload\x18\x01 \x01(\v2\x15.grpc.testing.PayloadR\apayload\x12D\n" +
	"\x11expect_compressed\x18\x02 \x01(\v2\x17.grpc.testing.BoolValueR\x10expectCompressed\"T\n" +
	"\x1aStreamingInputCallResponse\x126\n" +
	"\x17aggregated_payload_size\x18\x01 \x01(\x05R\x15aggregatedPayloadSize\"\xd4\x01\n" +
	"\x12ResponseParameters\x12\x12\n" +
	"\x04size\x18\x01 \x01(\x05R\x04size\x12\x1f\n" +
	"\vinterval_us\x18\x02 \x01(\x05R\n" +
	"intervalUs\x127\n" +
	"\n" +
	"compressed\x18\x03 \x01(\v2\x17.grpc.testing.BoolValueR\n" +
	"compressed\x12P\n" +
	"\x18fill_peer_socket_address\x18\x04 \x01(\v2\x17.grpc.testing.BoolValueR\x15fillPeerSocketAddress\"\xe9\x02\n" +
	"\x1aStreamingOutputCallRequest\x12>\n" +
	"\rresponse_type\x18\x01 \x01(\x0e2\x19.grpc.testing.PayloadTypeR\fresponseType\x12Q\n" +
	"\x13response_parameters\x18\x02 \x03(\v2 .grpc.testing.ResponseParametersR\x12responseParameters\x12/\n" +
	"\apayload\x18\x03 \x01(\v2\x15.grpc.testing.PayloadR\apayload\x12A\n" +
	"\x0fresponse_status\x18\a \x01(\v2\x18.grpc.testing.EchoStatusR\x0eresponseStatus\x12D\n" +
	"\x0forca_oob_report\x18\b \x01(\v2\x1c.grpc.testing.TestOrcaReportR\rorcaOobReport\"~\n" +
	"\x1bStreamingOutputCallResponse\x12/\n" +
	"\apayload\x18\x01 \x01(\v2\x15.grpc.testing.PayloadR\apayload\x12.\n" +
	"\x13peer_socket_address\x18\x02 \x01(\tR\x11peerSocketAddress\"J\n" +
	"\x0fReconnectParams\x127\n" +
	"\x18max_reconnect_backoff_ms\x18\x01 \x01(\x05R\x15maxReconnectBackoffMs\"F\n" +
	"\rReconnectInfo\x12\x16\n" +
	"\x06passed\x18\x01 \x01(\bR\x06passed\x12\x1d\n" +
	"\n" +
	"backoff_ms\x18\x02 \x03(\x05R\tbackoffMs\"{\n" +
	"\x18LoadBalancerStatsRequest\x12\x19\n" +
	"\bnum_rpcs\x18\x01 \x01(\x05R\anumRpcs\x12\x1f\n" +
	"\vtimeout_sec\x18\x02 \x01(\x05R\n" +
	"timeoutSec\x12#\n" +
	"\rmetadata_keys\x18\x03 \x03(\tR\fmetadataKeys\"\xd0\t\n" +
	"\x19LoadBalancerStatsResponse\x12Y\n" +
	"\frpcs_by_peer\x18\x01 \x03(\v27.grpc.testing.LoadBalancerStatsResponse.RpcsByPeerEntryR\n" +
	"rpcsByPeer\x12!\n" +
	"\fnum_failures\x18\x02 \x01(\x05R\vnumFailures\x12_\n" +
	"\x0erpcs_by_method\x18\x03 \x03(\v29.grpc.testing.LoadBalancerStatsResponse.RpcsByMethodEntryR\frpcsByMethod\x12h\n" +
	"\x11metadatas_by_peer\x18\x04 \x03(\v2<.grpc.testing.LoadBalancerStatsResponse.MetadatasByPeerEntryR\x0fmetadatasByPeer\x1a\x81\x01\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\x12H\n" +
	"\x04type\x18\x03 \x01(\x0e24.grpc.testing.LoadBalancerStatsResponse.MetadataTypeR\x04type\x1a`\n" +
	"\vRpcMetadata\x12Q\n" +
	"\bmetadata\x18\x01 \x03(\v25.grpc.testing.LoadBalancerStatsResponse.MetadataEntryR\bmetadata\x1ah\n" +
	"\x0eMetadataByPeer\x12V\n" +
	"\frpc_metadata\x18\x01 \x03(\v23.grpc.testing.LoadBalancerStatsResponse.RpcMetadataR\vrpcMetadata\x1a\xb1\x01\n" +
	"\n" +
	"RpcsByPeer\x12d\n" +
	"\frpcs_by_peer\x18\x01 \x03(\v2B.grpc.testing.LoadBalancerStatsResponse.RpcsByPeer.RpcsByPeerEntryR\n" +
	"rpcsByPeer\x1a=\n" +
	"\x0fRpcsByPeerEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x05R\x05value:\x028\x01\x1a=\n" +
	"\x0fRpcsByPeerEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x05R\x05value:\x028\x01\x1as\n" +
	"\x11RpcsByMethodEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12H\n" +
	"\x05value\x18\x02 \x01(\v22.grpc.testing.LoadBalancerStatsResponse.RpcsByPeerR\x05value:\x028\x01\x1az\n" +
	"\x14MetadatasByPeerEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12L\n" +
	"\x05value\x18\x02 \x01(\v26.grpc.testing.LoadBalancerStatsResponse.MetadataByPeerR\x05value:\x028\x01\"6\n" +
	"\fMetadataType\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\v\n" +
	"\aINITIAL\x10\x01\x12\f\n" +
	"\bTRAILING\x10\x02\"%\n" +
	"#LoadBalancerAccumulatedStatsRequest\"\x86\t\n" +
	"$LoadBalancerAccumulatedStatsResponse\x12\x8e\x01\n" +
	"\x1anum_rpcs_started_by_method\x18\x01 \x03(\v2N.grpc.testing.LoadBalancerAccumulatedStatsResponse.NumRpcsStartedByMethodEntryB\x02\x18\x01R\x16numRpcsStartedByMethod\x12\x94\x01\n" +
	"\x1cnum_rpcs_succeeded_by_method\x18\x02 \x03(\v2P.grpc.testing.LoadBalancerAccumulatedStatsResponse.NumRpcsSucceededByMethodEntryB\x02\x18\x01R\x18numRpcsSucceededByMethod\x12\x8b\x01\n" +
	"\x19num_rpcs_failed_by_method\x18\x03 \x03(\v2M.grpc.testing.LoadBalancerAccumulatedStatsResponse.NumRpcsFailedByMethodEntryB\x02\x18\x01R\x15numRpcsFailedByMethod\x12p\n" +
	"\x10stats_per_method\x18\x04 \x03(\v2F.grpc.testing.LoadBalancerAccumulatedStatsResponse.StatsPerMethodEntryR\x0estatsPerMethod\x1aI\n" +
	"\x1bNumRpcsStartedByMethodEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x05R\x05value:\x028\x01\x1aK\n" +
	"\x1dNumRpcsSucceededByMethodEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x05R\x05value:\x028\x01\x1aH\n" +
	"\x1aNumRpcsFailedByMethodEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x05R\x05value:\x028\x01\x1a\xcf\x01\n" +
	"\vMethodStats\x12!\n" +
	"\frpcs_started\x18\x01 \x01(\x05R\vrpcsStarted\x12b\n" +
	"\x06result\x18\x02 \x03(\v2J.grpc.testing.LoadBalancerAccumulatedStatsResponse.MethodStats.ResultEntryR\x06result\x1a9\n" +
	"\vResultEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x05R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x05R\x05value:\x028\x01\x1a\x81\x01\n" +
	"\x13StatsPerMethodEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12T\n" +
	"\x05value\x18\x02 \x01(\v2>.grpc.testing.LoadBalancerAccumulatedStatsResponse.MethodStatsR\x05value:\x028\x01\"\xe9\x02\n" +
	"\x16ClientConfigureRequest\x12B\n" +
	"\x05types\x18\x01 \x03(\x0e2,.grpc.testing.ClientConfigureRequest.RpcTypeR\x05types\x12I\n" +
	"\bmetadata\x18\x02 \x03(\v2-.grpc.testing.ClientConfigureRequest.MetadataR\bmetadata\x12\x1f\n" +
	"\vtimeout_sec\x18\x03 \x01(\x05R\n" +
	"timeoutSec\x1at\n" +
	"\bMetadata\x12@\n" +
	"\x04type\x18\x01 \x01(\x0e2,.grpc.testing.ClientConfigureRequest.RpcTypeR\x04type\x12\x10\n" +
	"\x03key\x18\x02 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x03 \x01(\tR\x05value\")\n" +
	"\aRpcType\x12\x0e\n" +
	"\n" +
	"EMPTY_CALL\x10\x00\x12\x0e\n" +
	"\n" +
	"UNARY_CALL\x10\x01\"\x19\n" +
	"\x17ClientConfigureResponse\"\x1e\n" +
	"\n" +
	"MemorySize\x12\x10\n" +
	"\x03rss\x18\x01 \x01(\x03R\x03rss\"\x8b\x03\n" +
	"\x0eTestOrcaReport\x12'\n" +
	"\x0fcpu_utilization\x18\x01 \x01(\x01R\x0ecpuUtilization\x12-\n" +
	"\x12memory_utilization\x18\x02 \x01(\x01R\x11memoryUtilization\x12P\n" +
	"\frequest_cost\x18\x03 \x03(\v2-.grpc.testing.TestOrcaReport.RequestCostEntryR\vrequestCost\x12O\n" +
	"\vutilization\x18\x04 \x03(\v2-.grpc.testing.TestOrcaReport.UtilizationEntryR\vutilization\x1a>\n" +
	"\x10RequestCostEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x01R\x05value:\x028\x01\x1a>\n" +
	"\x10UtilizationEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x01R\x05value:\x028\x01\"\x7f\n" +
	"\x16SetReturnStatusRequest\x12-\n" +
	"\x13grpc_code_to_return\x18\x01 \x01(\x05R\x10grpcCodeToReturn\x126\n" +
	"\x17grpc_status_description\x18\x02 \x01(\tR\x15grpcStatusDescription\"\xa5\x02\n" +
	"\vHookRequest\x12F\n" +
	"\acommand\x18\x01 \x01(\x0e2,.grpc.testing.HookRequest.HookRequestCommandR\acommand\x12-\n" +
	"\x13grpc_code_to_return\x18\x02 \x01(\x05R\x10grpcCodeToReturn\x126\n" +
	"\x17grpc_status_description\x18\x03 \x01(\tR\x15grpcStatusDescription\x12\x1f\n" +
	"\vserver_port\x18\x04 \x01(\x05R\n" +
	"serverPort\"F\n" +
	"\x12HookRequestCommand\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\t\n" +
	"\x05START\x10\x01\x12\b\n" +
	"\x04STOP\x10\x02\x12\n" +
	"\n" +
	"\x06RETURN\x10\x03\"\x0e\n" +
	"\fHookResponse*\x1f\n" +
	"\vPayloadType\x12\x10\n" +
	"\fCOMPRESSABLE\x10\x00*o\n" +
	"\x0fGrpclbRouteType\x12\x1d\n" +
	"\x19GRPCLB_ROUTE_TYPE_UNKNOWN\x10\x00\x12\x1e\n" +
	"\x1aGRPCLB_ROUTE_TYPE_FALLBACK\x10\x01\x12\x1d\n" +
	"\x19GRPCLB_ROUTE_TYPE_BACKEND\x10\x02B\x1d\n" +
	"\x1bio.grpc.testing.integrationb\x06proto3"

var (
	file_grpc_testing_messages_proto_rawDescOnce sync.Once
	file_grpc_testing_messages_proto_rawDescData []byte
)

func file_grpc_testing_messages_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_grpc_testing_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_grpc_testing_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 40)
var file_grpc_testing_messages_proto_goTypes = []any{
	(PayloadType)(0),
	(GrpclbRouteType)(0),
	(LoadBalancerStatsResponse_MetadataType)(0),
	(ClientConfigureRequest_RpcType)(0),
	(HookRequest_HookRequestCommand)(0),
	(*BoolValue)(nil),
	(*Payload)(nil),
	(*EchoStatus)(nil),
	(*SimpleRequest)(nil),
	(*SimpleResponse)(nil),
	(*StreamingInputCallRequest)(nil),
	(*StreamingInputCallResponse)(nil),
	(*ResponseParameters)(nil),
	(*StreamingOutputCallRequest)(nil),
	(*StreamingOutputCallResponse)(nil),
	(*ReconnectParams)(nil),
	(*ReconnectInfo)(nil),
	(*LoadBalancerStatsRequest)(nil),
	(*LoadBalancerStatsResponse)(nil),
	(*LoadBalancerAccumulatedStatsRequest)(nil),
	(*LoadBalancerAccumulatedStatsResponse)(nil),
	(*ClientConfigureRequest)(nil),
	(*ClientConfigureResponse)(nil),
	(*MemorySize)(nil),
	(*TestOrcaReport)(nil),
	(*SetReturnStatusRequest)(nil),
	(*HookRequest)(nil),
	(*HookResponse)(nil),
	(*LoadBalancerStatsResponse_MetadataEntry)(nil),
	(*LoadBalancerStatsResponse_RpcMetadata)(nil),
	(*LoadBalancerStatsResponse_MetadataByPeer)(nil),
	(*LoadBalancerStatsResponse_RpcsByPeer)(nil),
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	(*LoadBalancerAccumulatedStatsResponse_MethodStats)(nil),
	nil,
	nil,
	(*ClientConfigureRequest_Metadata)(nil),
	nil,
	nil,
}
var file_grpc_testing_messages_proto_depIdxs = []int32{
	0,
	0,
	6,
	5,
	7,
	5,
	24,
	6,
	1,
	6,
	5,
	5,
	5,
	0,
	12,
	6,
	7,
	24,
	6,
	32,
	33,
	34,
	36,
	37,
	38,
	40,
	3,
	42,
	43,
	44,
	4,
	2,
	28,
	29,
	35,
	31,
	30,
	41,
	39,
	3,
	40,
	40,
	40,
	40,
	0,
}

func init()                                  { file_grpc_testing_messages_proto_init() }
func file_grpc_testing_messages_proto_init() { _ = "STUB: not implemented"; return }
