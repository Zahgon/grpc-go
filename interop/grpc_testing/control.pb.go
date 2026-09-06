package grpc_testing

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientType int32

const (
	ClientType_SYNC_CLIENT     ClientType = 0
	ClientType_ASYNC_CLIENT    ClientType = 1
	ClientType_OTHER_CLIENT    ClientType = 2
	ClientType_CALLBACK_CLIENT ClientType = 3
)

var (
	ClientType_name = map[int32]string{
		0: "SYNC_CLIENT",
		1: "ASYNC_CLIENT",
		2: "OTHER_CLIENT",
		3: "CALLBACK_CLIENT",
	}
	ClientType_value = map[string]int32{
		"SYNC_CLIENT":     0,
		"ASYNC_CLIENT":    1,
		"OTHER_CLIENT":    2,
		"CALLBACK_CLIENT": 3,
	}
)

func (x ClientType) Enum() *ClientType { _ = "STUB: not implemented"; return nil }

func (x ClientType) String() string { _ = "STUB: not implemented"; return "" }

func (ClientType) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (ClientType) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x ClientType) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (ClientType) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type ServerType int32

const (
	ServerType_SYNC_SERVER          ServerType = 0
	ServerType_ASYNC_SERVER         ServerType = 1
	ServerType_ASYNC_GENERIC_SERVER ServerType = 2
	ServerType_OTHER_SERVER         ServerType = 3
	ServerType_CALLBACK_SERVER      ServerType = 4
)

var (
	ServerType_name = map[int32]string{
		0: "SYNC_SERVER",
		1: "ASYNC_SERVER",
		2: "ASYNC_GENERIC_SERVER",
		3: "OTHER_SERVER",
		4: "CALLBACK_SERVER",
	}
	ServerType_value = map[string]int32{
		"SYNC_SERVER":          0,
		"ASYNC_SERVER":         1,
		"ASYNC_GENERIC_SERVER": 2,
		"OTHER_SERVER":         3,
		"CALLBACK_SERVER":      4,
	}
)

func (x ServerType) Enum() *ServerType { _ = "STUB: not implemented"; return nil }

func (x ServerType) String() string { _ = "STUB: not implemented"; return "" }

func (ServerType) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (ServerType) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x ServerType) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (ServerType) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type RpcType int32

const (
	RpcType_UNARY                 RpcType = 0
	RpcType_STREAMING             RpcType = 1
	RpcType_STREAMING_FROM_CLIENT RpcType = 2
	RpcType_STREAMING_FROM_SERVER RpcType = 3
	RpcType_STREAMING_BOTH_WAYS   RpcType = 4
)

var (
	RpcType_name = map[int32]string{
		0: "UNARY",
		1: "STREAMING",
		2: "STREAMING_FROM_CLIENT",
		3: "STREAMING_FROM_SERVER",
		4: "STREAMING_BOTH_WAYS",
	}
	RpcType_value = map[string]int32{
		"UNARY":                 0,
		"STREAMING":             1,
		"STREAMING_FROM_CLIENT": 2,
		"STREAMING_FROM_SERVER": 3,
		"STREAMING_BOTH_WAYS":   4,
	}
)

func (x RpcType) Enum() *RpcType { _ = "STUB: not implemented"; return nil }

func (x RpcType) String() string { _ = "STUB: not implemented"; return "" }

func (RpcType) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (RpcType) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x RpcType) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (RpcType) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type PoissonParams struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	OfferedLoad   float64 `protobuf:"fixed64,1,opt,name=offered_load,json=offeredLoad,proto3" json:"offered_load,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PoissonParams) Reset() { _ = "STUB: not implemented"; return }

func (x *PoissonParams) String() string { _ = "STUB: not implemented"; return "" }

func (*PoissonParams) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *PoissonParams) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*PoissonParams) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *PoissonParams) GetOfferedLoad() float64 { _ = "STUB: not implemented"; return 0 }

type ClosedLoopParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClosedLoopParams) Reset() { _ = "STUB: not implemented"; return }

func (x *ClosedLoopParams) String() string { _ = "STUB: not implemented"; return "" }

func (*ClosedLoopParams) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClosedLoopParams) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClosedLoopParams) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type LoadParams struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Load          isLoadParams_Load `protobuf_oneof:"load"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoadParams) Reset() { _ = "STUB: not implemented"; return }

func (x *LoadParams) String() string { _ = "STUB: not implemented"; return "" }

func (*LoadParams) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *LoadParams) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*LoadParams) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *LoadParams) GetLoad() isLoadParams_Load {
	_ = "STUB: not implemented"
	return *new(isLoadParams_Load)
}

func (x *LoadParams) GetClosedLoop() *ClosedLoopParams { _ = "STUB: not implemented"; return nil }

func (x *LoadParams) GetPoisson() *PoissonParams { _ = "STUB: not implemented"; return nil }

type isLoadParams_Load interface {
	isLoadParams_Load()
}

type LoadParams_ClosedLoop struct {
	ClosedLoop *ClosedLoopParams `protobuf:"bytes,1,opt,name=closed_loop,json=closedLoop,proto3,oneof"`
}

type LoadParams_Poisson struct {
	Poisson *PoissonParams `protobuf:"bytes,2,opt,name=poisson,proto3,oneof"`
}

func (*LoadParams_ClosedLoop) isLoadParams_Load() { _ = "STUB: not implemented"; return }

func (*LoadParams_Poisson) isLoadParams_Load() { _ = "STUB: not implemented"; return }

type SecurityParams struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	UseTestCa          bool                   `protobuf:"varint,1,opt,name=use_test_ca,json=useTestCa,proto3" json:"use_test_ca,omitempty"`
	ServerHostOverride string                 `protobuf:"bytes,2,opt,name=server_host_override,json=serverHostOverride,proto3" json:"server_host_override,omitempty"`
	CredType           string                 `protobuf:"bytes,3,opt,name=cred_type,json=credType,proto3" json:"cred_type,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *SecurityParams) Reset() { _ = "STUB: not implemented"; return }

func (x *SecurityParams) String() string { _ = "STUB: not implemented"; return "" }

func (*SecurityParams) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SecurityParams) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SecurityParams) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SecurityParams) GetUseTestCa() bool { _ = "STUB: not implemented"; return false }

func (x *SecurityParams) GetServerHostOverride() string { _ = "STUB: not implemented"; return "" }

func (x *SecurityParams) GetCredType() string { _ = "STUB: not implemented"; return "" }

type ChannelArg struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`

	Value         isChannelArg_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChannelArg) Reset() { _ = "STUB: not implemented"; return }

func (x *ChannelArg) String() string { _ = "STUB: not implemented"; return "" }

func (*ChannelArg) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ChannelArg) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ChannelArg) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ChannelArg) GetName() string { _ = "STUB: not implemented"; return "" }

func (x *ChannelArg) GetValue() isChannelArg_Value {
	_ = "STUB: not implemented"
	return *new(isChannelArg_Value)
}

func (x *ChannelArg) GetStrValue() string { _ = "STUB: not implemented"; return "" }

func (x *ChannelArg) GetIntValue() int32 { _ = "STUB: not implemented"; return 0 }

type isChannelArg_Value interface {
	isChannelArg_Value()
}

type ChannelArg_StrValue struct {
	StrValue string `protobuf:"bytes,2,opt,name=str_value,json=strValue,proto3,oneof"`
}

type ChannelArg_IntValue struct {
	IntValue int32 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof"`
}

func (*ChannelArg_StrValue) isChannelArg_Value() { _ = "STUB: not implemented"; return }

func (*ChannelArg_IntValue) isChannelArg_Value() { _ = "STUB: not implemented"; return }

type ClientConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ServerTargets  []string        `protobuf:"bytes,1,rep,name=server_targets,json=serverTargets,proto3" json:"server_targets,omitempty"`
	ClientType     ClientType      `protobuf:"varint,2,opt,name=client_type,json=clientType,proto3,enum=grpc.testing.ClientType" json:"client_type,omitempty"`
	SecurityParams *SecurityParams `protobuf:"bytes,3,opt,name=security_params,json=securityParams,proto3" json:"security_params,omitempty"`

	OutstandingRpcsPerChannel int32 `protobuf:"varint,4,opt,name=outstanding_rpcs_per_channel,json=outstandingRpcsPerChannel,proto3" json:"outstanding_rpcs_per_channel,omitempty"`

	ClientChannels int32 `protobuf:"varint,5,opt,name=client_channels,json=clientChannels,proto3" json:"client_channels,omitempty"`

	AsyncClientThreads int32   `protobuf:"varint,7,opt,name=async_client_threads,json=asyncClientThreads,proto3" json:"async_client_threads,omitempty"`
	RpcType            RpcType `protobuf:"varint,8,opt,name=rpc_type,json=rpcType,proto3,enum=grpc.testing.RpcType" json:"rpc_type,omitempty"`

	LoadParams      *LoadParams      `protobuf:"bytes,10,opt,name=load_params,json=loadParams,proto3" json:"load_params,omitempty"`
	PayloadConfig   *PayloadConfig   `protobuf:"bytes,11,opt,name=payload_config,json=payloadConfig,proto3" json:"payload_config,omitempty"`
	HistogramParams *HistogramParams `protobuf:"bytes,12,opt,name=histogram_params,json=histogramParams,proto3" json:"histogram_params,omitempty"`

	CoreList  []int32 `protobuf:"varint,13,rep,packed,name=core_list,json=coreList,proto3" json:"core_list,omitempty"`
	CoreLimit int32   `protobuf:"varint,14,opt,name=core_limit,json=coreLimit,proto3" json:"core_limit,omitempty"`

	OtherClientApi string        `protobuf:"bytes,15,opt,name=other_client_api,json=otherClientApi,proto3" json:"other_client_api,omitempty"`
	ChannelArgs    []*ChannelArg `protobuf:"bytes,16,rep,name=channel_args,json=channelArgs,proto3" json:"channel_args,omitempty"`

	ThreadsPerCq int32 `protobuf:"varint,17,opt,name=threads_per_cq,json=threadsPerCq,proto3" json:"threads_per_cq,omitempty"`

	MessagesPerStream int32 `protobuf:"varint,18,opt,name=messages_per_stream,json=messagesPerStream,proto3" json:"messages_per_stream,omitempty"`

	UseCoalesceApi bool `protobuf:"varint,19,opt,name=use_coalesce_api,json=useCoalesceApi,proto3" json:"use_coalesce_api,omitempty"`

	MedianLatencyCollectionIntervalMillis int32 `protobuf:"varint,20,opt,name=median_latency_collection_interval_millis,json=medianLatencyCollectionIntervalMillis,proto3" json:"median_latency_collection_interval_millis,omitempty"`

	ClientProcesses int32 `protobuf:"varint,21,opt,name=client_processes,json=clientProcesses,proto3" json:"client_processes,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ClientConfig) Reset() { _ = "STUB: not implemented"; return }

func (x *ClientConfig) String() string { _ = "STUB: not implemented"; return "" }

func (*ClientConfig) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClientConfig) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ClientConfig) GetServerTargets() []string { _ = "STUB: not implemented"; return nil }

func (x *ClientConfig) GetClientType() ClientType {
	_ = "STUB: not implemented"
	return *new(ClientType)
}

func (x *ClientConfig) GetSecurityParams() *SecurityParams { _ = "STUB: not implemented"; return nil }

func (x *ClientConfig) GetOutstandingRpcsPerChannel() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ClientConfig) GetClientChannels() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ClientConfig) GetAsyncClientThreads() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ClientConfig) GetRpcType() RpcType { _ = "STUB: not implemented"; return *new(RpcType) }

func (x *ClientConfig) GetLoadParams() *LoadParams { _ = "STUB: not implemented"; return nil }

func (x *ClientConfig) GetPayloadConfig() *PayloadConfig { _ = "STUB: not implemented"; return nil }

func (x *ClientConfig) GetHistogramParams() *HistogramParams { _ = "STUB: not implemented"; return nil }

func (x *ClientConfig) GetCoreList() []int32 { _ = "STUB: not implemented"; return nil }

func (x *ClientConfig) GetCoreLimit() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ClientConfig) GetOtherClientApi() string { _ = "STUB: not implemented"; return "" }

func (x *ClientConfig) GetChannelArgs() []*ChannelArg { _ = "STUB: not implemented"; return nil }

func (x *ClientConfig) GetThreadsPerCq() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ClientConfig) GetMessagesPerStream() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ClientConfig) GetUseCoalesceApi() bool { _ = "STUB: not implemented"; return false }

func (x *ClientConfig) GetMedianLatencyCollectionIntervalMillis() int32 {
	_ = "STUB: not implemented"
	return 0
}

func (x *ClientConfig) GetClientProcesses() int32 { _ = "STUB: not implemented"; return 0 }

type ClientStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Stats         *ClientStats           `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientStatus) Reset() { _ = "STUB: not implemented"; return }

func (x *ClientStatus) String() string { _ = "STUB: not implemented"; return "" }

func (*ClientStatus) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClientStatus) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClientStatus) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ClientStatus) GetStats() *ClientStats { _ = "STUB: not implemented"; return nil }

type Mark struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Reset_        bool `protobuf:"varint,1,opt,name=reset,proto3" json:"reset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Mark) Reset() { _ = "STUB: not implemented"; return }

func (x *Mark) String() string { _ = "STUB: not implemented"; return "" }

func (*Mark) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Mark) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Mark) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Mark) GetReset_() bool { _ = "STUB: not implemented"; return false }

type ClientArgs struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Argtype       isClientArgs_Argtype `protobuf_oneof:"argtype"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientArgs) Reset() { _ = "STUB: not implemented"; return }

func (x *ClientArgs) String() string { _ = "STUB: not implemented"; return "" }

func (*ClientArgs) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClientArgs) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClientArgs) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ClientArgs) GetArgtype() isClientArgs_Argtype {
	_ = "STUB: not implemented"
	return *new(isClientArgs_Argtype)
}

func (x *ClientArgs) GetSetup() *ClientConfig { _ = "STUB: not implemented"; return nil }

func (x *ClientArgs) GetMark() *Mark { _ = "STUB: not implemented"; return nil }

type isClientArgs_Argtype interface {
	isClientArgs_Argtype()
}

type ClientArgs_Setup struct {
	Setup *ClientConfig `protobuf:"bytes,1,opt,name=setup,proto3,oneof"`
}

type ClientArgs_Mark struct {
	Mark *Mark `protobuf:"bytes,2,opt,name=mark,proto3,oneof"`
}

func (*ClientArgs_Setup) isClientArgs_Argtype() { _ = "STUB: not implemented"; return }

func (*ClientArgs_Mark) isClientArgs_Argtype() { _ = "STUB: not implemented"; return }

type ServerConfig struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ServerType     ServerType             `protobuf:"varint,1,opt,name=server_type,json=serverType,proto3,enum=grpc.testing.ServerType" json:"server_type,omitempty"`
	SecurityParams *SecurityParams        `protobuf:"bytes,2,opt,name=security_params,json=securityParams,proto3" json:"security_params,omitempty"`

	Port int32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`

	AsyncServerThreads int32 `protobuf:"varint,7,opt,name=async_server_threads,json=asyncServerThreads,proto3" json:"async_server_threads,omitempty"`

	CoreLimit int32 `protobuf:"varint,8,opt,name=core_limit,json=coreLimit,proto3" json:"core_limit,omitempty"`

	PayloadConfig *PayloadConfig `protobuf:"bytes,9,opt,name=payload_config,json=payloadConfig,proto3" json:"payload_config,omitempty"`

	CoreList []int32 `protobuf:"varint,10,rep,packed,name=core_list,json=coreList,proto3" json:"core_list,omitempty"`

	OtherServerApi string `protobuf:"bytes,11,opt,name=other_server_api,json=otherServerApi,proto3" json:"other_server_api,omitempty"`

	ThreadsPerCq int32 `protobuf:"varint,12,opt,name=threads_per_cq,json=threadsPerCq,proto3" json:"threads_per_cq,omitempty"`

	ResourceQuotaSize int32         `protobuf:"varint,1001,opt,name=resource_quota_size,json=resourceQuotaSize,proto3" json:"resource_quota_size,omitempty"`
	ChannelArgs       []*ChannelArg `protobuf:"bytes,1002,rep,name=channel_args,json=channelArgs,proto3" json:"channel_args,omitempty"`

	ServerProcesses int32 `protobuf:"varint,21,opt,name=server_processes,json=serverProcesses,proto3" json:"server_processes,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ServerConfig) Reset() { _ = "STUB: not implemented"; return }

func (x *ServerConfig) String() string { _ = "STUB: not implemented"; return "" }

func (*ServerConfig) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServerConfig) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ServerConfig) GetServerType() ServerType {
	_ = "STUB: not implemented"
	return *new(ServerType)
}

func (x *ServerConfig) GetSecurityParams() *SecurityParams { _ = "STUB: not implemented"; return nil }

func (x *ServerConfig) GetPort() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ServerConfig) GetAsyncServerThreads() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ServerConfig) GetCoreLimit() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ServerConfig) GetPayloadConfig() *PayloadConfig { _ = "STUB: not implemented"; return nil }

func (x *ServerConfig) GetCoreList() []int32 { _ = "STUB: not implemented"; return nil }

func (x *ServerConfig) GetOtherServerApi() string { _ = "STUB: not implemented"; return "" }

func (x *ServerConfig) GetThreadsPerCq() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ServerConfig) GetResourceQuotaSize() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ServerConfig) GetChannelArgs() []*ChannelArg { _ = "STUB: not implemented"; return nil }

func (x *ServerConfig) GetServerProcesses() int32 { _ = "STUB: not implemented"; return 0 }

type ServerArgs struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Argtype       isServerArgs_Argtype `protobuf_oneof:"argtype"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerArgs) Reset() { _ = "STUB: not implemented"; return }

func (x *ServerArgs) String() string { _ = "STUB: not implemented"; return "" }

func (*ServerArgs) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServerArgs) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServerArgs) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ServerArgs) GetArgtype() isServerArgs_Argtype {
	_ = "STUB: not implemented"
	return *new(isServerArgs_Argtype)
}

func (x *ServerArgs) GetSetup() *ServerConfig { _ = "STUB: not implemented"; return nil }

func (x *ServerArgs) GetMark() *Mark { _ = "STUB: not implemented"; return nil }

type isServerArgs_Argtype interface {
	isServerArgs_Argtype()
}

type ServerArgs_Setup struct {
	Setup *ServerConfig `protobuf:"bytes,1,opt,name=setup,proto3,oneof"`
}

type ServerArgs_Mark struct {
	Mark *Mark `protobuf:"bytes,2,opt,name=mark,proto3,oneof"`
}

func (*ServerArgs_Setup) isServerArgs_Argtype() { _ = "STUB: not implemented"; return }

func (*ServerArgs_Mark) isServerArgs_Argtype() { _ = "STUB: not implemented"; return }

type ServerStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Stats *ServerStats           `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`

	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`

	Cores         int32 `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerStatus) Reset() { _ = "STUB: not implemented"; return }

func (x *ServerStatus) String() string { _ = "STUB: not implemented"; return "" }

func (*ServerStatus) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServerStatus) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServerStatus) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ServerStatus) GetStats() *ServerStats { _ = "STUB: not implemented"; return nil }

func (x *ServerStatus) GetPort() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ServerStatus) GetCores() int32 { _ = "STUB: not implemented"; return 0 }

type CoreRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CoreRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *CoreRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*CoreRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *CoreRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*CoreRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type CoreResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Cores         int32 `protobuf:"varint,1,opt,name=cores,proto3" json:"cores,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CoreResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *CoreResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*CoreResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *CoreResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*CoreResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *CoreResponse) GetCores() int32 { _ = "STUB: not implemented"; return 0 }

type Void struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Void) Reset() { _ = "STUB: not implemented"; return }

func (x *Void) String() string { _ = "STUB: not implemented"; return "" }

func (*Void) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Void) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Void) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type Scenario struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`

	ClientConfig *ClientConfig `protobuf:"bytes,2,opt,name=client_config,json=clientConfig,proto3" json:"client_config,omitempty"`

	NumClients int32 `protobuf:"varint,3,opt,name=num_clients,json=numClients,proto3" json:"num_clients,omitempty"`

	ServerConfig *ServerConfig `protobuf:"bytes,4,opt,name=server_config,json=serverConfig,proto3" json:"server_config,omitempty"`

	NumServers int32 `protobuf:"varint,5,opt,name=num_servers,json=numServers,proto3" json:"num_servers,omitempty"`

	WarmupSeconds int32 `protobuf:"varint,6,opt,name=warmup_seconds,json=warmupSeconds,proto3" json:"warmup_seconds,omitempty"`

	BenchmarkSeconds int32 `protobuf:"varint,7,opt,name=benchmark_seconds,json=benchmarkSeconds,proto3" json:"benchmark_seconds,omitempty"`

	SpawnLocalWorkerCount int32 `protobuf:"varint,8,opt,name=spawn_local_worker_count,json=spawnLocalWorkerCount,proto3" json:"spawn_local_worker_count,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *Scenario) Reset() { _ = "STUB: not implemented"; return }

func (x *Scenario) String() string { _ = "STUB: not implemented"; return "" }

func (*Scenario) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Scenario) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Scenario) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Scenario) GetName() string { _ = "STUB: not implemented"; return "" }

func (x *Scenario) GetClientConfig() *ClientConfig { _ = "STUB: not implemented"; return nil }

func (x *Scenario) GetNumClients() int32 { _ = "STUB: not implemented"; return 0 }

func (x *Scenario) GetServerConfig() *ServerConfig { _ = "STUB: not implemented"; return nil }

func (x *Scenario) GetNumServers() int32 { _ = "STUB: not implemented"; return 0 }

func (x *Scenario) GetWarmupSeconds() int32 { _ = "STUB: not implemented"; return 0 }

func (x *Scenario) GetBenchmarkSeconds() int32 { _ = "STUB: not implemented"; return 0 }

func (x *Scenario) GetSpawnLocalWorkerCount() int32 { _ = "STUB: not implemented"; return 0 }

type Scenarios struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Scenarios     []*Scenario            `protobuf:"bytes,1,rep,name=scenarios,proto3" json:"scenarios,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scenarios) Reset() { _ = "STUB: not implemented"; return }

func (x *Scenarios) String() string { _ = "STUB: not implemented"; return "" }

func (*Scenarios) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Scenarios) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Scenarios) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Scenarios) GetScenarios() []*Scenario { _ = "STUB: not implemented"; return nil }

type ScenarioResultSummary struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Qps float64 `protobuf:"fixed64,1,opt,name=qps,proto3" json:"qps,omitempty"`

	QpsPerServerCore float64 `protobuf:"fixed64,2,opt,name=qps_per_server_core,json=qpsPerServerCore,proto3" json:"qps_per_server_core,omitempty"`

	ServerSystemTime float64 `protobuf:"fixed64,3,opt,name=server_system_time,json=serverSystemTime,proto3" json:"server_system_time,omitempty"`

	ServerUserTime float64 `protobuf:"fixed64,4,opt,name=server_user_time,json=serverUserTime,proto3" json:"server_user_time,omitempty"`

	ClientSystemTime float64 `protobuf:"fixed64,5,opt,name=client_system_time,json=clientSystemTime,proto3" json:"client_system_time,omitempty"`

	ClientUserTime float64 `protobuf:"fixed64,6,opt,name=client_user_time,json=clientUserTime,proto3" json:"client_user_time,omitempty"`

	Latency_50  float64 `protobuf:"fixed64,7,opt,name=latency_50,json=latency50,proto3" json:"latency_50,omitempty"`
	Latency_90  float64 `protobuf:"fixed64,8,opt,name=latency_90,json=latency90,proto3" json:"latency_90,omitempty"`
	Latency_95  float64 `protobuf:"fixed64,9,opt,name=latency_95,json=latency95,proto3" json:"latency_95,omitempty"`
	Latency_99  float64 `protobuf:"fixed64,10,opt,name=latency_99,json=latency99,proto3" json:"latency_99,omitempty"`
	Latency_999 float64 `protobuf:"fixed64,11,opt,name=latency_999,json=latency999,proto3" json:"latency_999,omitempty"`

	ServerCpuUsage float64 `protobuf:"fixed64,12,opt,name=server_cpu_usage,json=serverCpuUsage,proto3" json:"server_cpu_usage,omitempty"`

	SuccessfulRequestsPerSecond float64 `protobuf:"fixed64,13,opt,name=successful_requests_per_second,json=successfulRequestsPerSecond,proto3" json:"successful_requests_per_second,omitempty"`
	FailedRequestsPerSecond     float64 `protobuf:"fixed64,14,opt,name=failed_requests_per_second,json=failedRequestsPerSecond,proto3" json:"failed_requests_per_second,omitempty"`

	ClientPollsPerRequest float64 `protobuf:"fixed64,15,opt,name=client_polls_per_request,json=clientPollsPerRequest,proto3" json:"client_polls_per_request,omitempty"`
	ServerPollsPerRequest float64 `protobuf:"fixed64,16,opt,name=server_polls_per_request,json=serverPollsPerRequest,proto3" json:"server_polls_per_request,omitempty"`

	ServerQueriesPerCpuSec float64 `protobuf:"fixed64,17,opt,name=server_queries_per_cpu_sec,json=serverQueriesPerCpuSec,proto3" json:"server_queries_per_cpu_sec,omitempty"`
	ClientQueriesPerCpuSec float64 `protobuf:"fixed64,18,opt,name=client_queries_per_cpu_sec,json=clientQueriesPerCpuSec,proto3" json:"client_queries_per_cpu_sec,omitempty"`

	StartTime     *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScenarioResultSummary) Reset() { _ = "STUB: not implemented"; return }

func (x *ScenarioResultSummary) String() string { _ = "STUB: not implemented"; return "" }

func (*ScenarioResultSummary) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ScenarioResultSummary) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ScenarioResultSummary) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *ScenarioResultSummary) GetQps() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetQpsPerServerCore() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetServerSystemTime() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetServerUserTime() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetClientSystemTime() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetClientUserTime() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetLatency_50() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetLatency_90() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetLatency_95() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetLatency_99() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetLatency_999() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetServerCpuUsage() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ScenarioResultSummary) GetSuccessfulRequestsPerSecond() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (x *ScenarioResultSummary) GetFailedRequestsPerSecond() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (x *ScenarioResultSummary) GetClientPollsPerRequest() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (x *ScenarioResultSummary) GetServerPollsPerRequest() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (x *ScenarioResultSummary) GetServerQueriesPerCpuSec() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (x *ScenarioResultSummary) GetClientQueriesPerCpuSec() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (x *ScenarioResultSummary) GetStartTime() *timestamppb.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

func (x *ScenarioResultSummary) GetEndTime() *timestamppb.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

type ScenarioResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Scenario *Scenario `protobuf:"bytes,1,opt,name=scenario,proto3" json:"scenario,omitempty"`

	Latencies *HistogramData `protobuf:"bytes,2,opt,name=latencies,proto3" json:"latencies,omitempty"`

	ClientStats []*ClientStats `protobuf:"bytes,3,rep,name=client_stats,json=clientStats,proto3" json:"client_stats,omitempty"`

	ServerStats []*ServerStats `protobuf:"bytes,4,rep,name=server_stats,json=serverStats,proto3" json:"server_stats,omitempty"`

	ServerCores []int32 `protobuf:"varint,5,rep,packed,name=server_cores,json=serverCores,proto3" json:"server_cores,omitempty"`

	Summary *ScenarioResultSummary `protobuf:"bytes,6,opt,name=summary,proto3" json:"summary,omitempty"`

	ClientSuccess []bool `protobuf:"varint,7,rep,packed,name=client_success,json=clientSuccess,proto3" json:"client_success,omitempty"`
	ServerSuccess []bool `protobuf:"varint,8,rep,packed,name=server_success,json=serverSuccess,proto3" json:"server_success,omitempty"`

	RequestResults []*RequestResultCount `protobuf:"bytes,9,rep,name=request_results,json=requestResults,proto3" json:"request_results,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ScenarioResult) Reset() { _ = "STUB: not implemented"; return }

func (x *ScenarioResult) String() string { _ = "STUB: not implemented"; return "" }

func (*ScenarioResult) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ScenarioResult) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ScenarioResult) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ScenarioResult) GetScenario() *Scenario { _ = "STUB: not implemented"; return nil }

func (x *ScenarioResult) GetLatencies() *HistogramData { _ = "STUB: not implemented"; return nil }

func (x *ScenarioResult) GetClientStats() []*ClientStats { _ = "STUB: not implemented"; return nil }

func (x *ScenarioResult) GetServerStats() []*ServerStats { _ = "STUB: not implemented"; return nil }

func (x *ScenarioResult) GetServerCores() []int32 { _ = "STUB: not implemented"; return nil }

func (x *ScenarioResult) GetSummary() *ScenarioResultSummary { _ = "STUB: not implemented"; return nil }

func (x *ScenarioResult) GetClientSuccess() []bool { _ = "STUB: not implemented"; return nil }

func (x *ScenarioResult) GetServerSuccess() []bool { _ = "STUB: not implemented"; return nil }

func (x *ScenarioResult) GetRequestResults() []*RequestResultCount {
	_ = "STUB: not implemented"
	return nil
}

var File_grpc_testing_control_proto protoreflect.FileDescriptor

const file_grpc_testing_control_proto_rawDesc = "" +
	"\n" +
	"\x1agrpc/testing/control.proto\x12\fgrpc.testing\x1a\x1bgrpc/testing/payloads.proto\x1a\x18grpc/testing/stats.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"2\n" +
	"\rPoissonParams\x12!\n" +
	"\foffered_load\x18\x01 \x01(\x01R\vofferedLoad\"\x12\n" +
	"\x10ClosedLoopParams\"\x90\x01\n" +
	"\n" +
	"LoadParams\x12A\n" +
	"\vclosed_loop\x18\x01 \x01(\v2\x1e.grpc.testing.ClosedLoopParamsH\x00R\n" +
	"closedLoop\x127\n" +
	"\apoisson\x18\x02 \x01(\v2\x1b.grpc.testing.PoissonParamsH\x00R\apoissonB\x06\n" +
	"\x04load\"\x7f\n" +
	"\x0eSecurityParams\x12\x1e\n" +
	"\vuse_test_ca\x18\x01 \x01(\bR\tuseTestCa\x120\n" +
	"\x14server_host_override\x18\x02 \x01(\tR\x12serverHostOverride\x12\x1b\n" +
	"\tcred_type\x18\x03 \x01(\tR\bcredType\"g\n" +
	"\n" +
	"ChannelArg\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1d\n" +
	"\tstr_value\x18\x02 \x01(\tH\x00R\bstrValue\x12\x1d\n" +
	"\tint_value\x18\x03 \x01(\x05H\x00R\bintValueB\a\n" +
	"\x05value\"\xf6\a\n" +
	"\fClientConfig\x12%\n" +
	"\x0eserver_targets\x18\x01 \x03(\tR\rserverTargets\x129\n" +
	"\vclient_type\x18\x02 \x01(\x0e2\x18.grpc.testing.ClientTypeR\n" +
	"clientType\x12E\n" +
	"\x0fsecurity_params\x18\x03 \x01(\v2\x1c.grpc.testing.SecurityParamsR\x0esecurityParams\x12?\n" +
	"\x1coutstanding_rpcs_per_channel\x18\x04 \x01(\x05R\x19outstandingRpcsPerChannel\x12'\n" +
	"\x0fclient_channels\x18\x05 \x01(\x05R\x0eclientChannels\x120\n" +
	"\x14async_client_threads\x18\a \x01(\x05R\x12asyncClientThreads\x120\n" +
	"\brpc_type\x18\b \x01(\x0e2\x15.grpc.testing.RpcTypeR\arpcType\x129\n" +
	"\vload_params\x18\n" +
	" \x01(\v2\x18.grpc.testing.LoadParamsR\n" +
	"loadParams\x12B\n" +
	"\x0epayload_config\x18\v \x01(\v2\x1b.grpc.testing.PayloadConfigR\rpayloadConfig\x12H\n" +
	"\x10histogram_params\x18\f \x01(\v2\x1d.grpc.testing.HistogramParamsR\x0fhistogramParams\x12\x1b\n" +
	"\tcore_list\x18\r \x03(\x05R\bcoreList\x12\x1d\n" +
	"\n" +
	"core_limit\x18\x0e \x01(\x05R\tcoreLimit\x12(\n" +
	"\x10other_client_api\x18\x0f \x01(\tR\x0eotherClientApi\x12;\n" +
	"\fchannel_args\x18\x10 \x03(\v2\x18.grpc.testing.ChannelArgR\vchannelArgs\x12$\n" +
	"\x0ethreads_per_cq\x18\x11 \x01(\x05R\fthreadsPerCq\x12.\n" +
	"\x13messages_per_stream\x18\x12 \x01(\x05R\x11messagesPerStream\x12(\n" +
	"\x10use_coalesce_api\x18\x13 \x01(\bR\x0euseCoalesceApi\x12X\n" +
	")median_latency_collection_interval_millis\x18\x14 \x01(\x05R%medianLatencyCollectionIntervalMillis\x12)\n" +
	"\x10client_processes\x18\x15 \x01(\x05R\x0fclientProcesses\"?\n" +
	"\fClientStatus\x12/\n" +
	"\x05stats\x18\x01 \x01(\v2\x19.grpc.testing.ClientStatsR\x05stats\"\x1c\n" +
	"\x04Mark\x12\x14\n" +
	"\x05reset\x18\x01 \x01(\bR\x05reset\"u\n" +
	"\n" +
	"ClientArgs\x122\n" +
	"\x05setup\x18\x01 \x01(\v2\x1a.grpc.testing.ClientConfigH\x00R\x05setup\x12(\n" +
	"\x04mark\x18\x02 \x01(\v2\x12.grpc.testing.MarkH\x00R\x04markB\t\n" +
	"\aargtype\"\xc0\x04\n" +
	"\fServerConfig\x129\n" +
	"\vserver_type\x18\x01 \x01(\x0e2\x18.grpc.testing.ServerTypeR\n" +
	"serverType\x12E\n" +
	"\x0fsecurity_params\x18\x02 \x01(\v2\x1c.grpc.testing.SecurityParamsR\x0esecurityParams\x12\x12\n" +
	"\x04port\x18\x04 \x01(\x05R\x04port\x120\n" +
	"\x14async_server_threads\x18\a \x01(\x05R\x12asyncServerThreads\x12\x1d\n" +
	"\n" +
	"core_limit\x18\b \x01(\x05R\tcoreLimit\x12B\n" +
	"\x0epayload_config\x18\t \x01(\v2\x1b.grpc.testing.PayloadConfigR\rpayloadConfig\x12\x1b\n" +
	"\tcore_list\x18\n" +
	" \x03(\x05R\bcoreList\x12(\n" +
	"\x10other_server_api\x18\v \x01(\tR\x0eotherServerApi\x12$\n" +
	"\x0ethreads_per_cq\x18\f \x01(\x05R\fthreadsPerCq\x12/\n" +
	"\x13resource_quota_size\x18\xe9\a \x01(\x05R\x11resourceQuotaSize\x12<\n" +
	"\fchannel_args\x18\xea\a \x03(\v2\x18.grpc.testing.ChannelArgR\vchannelArgs\x12)\n" +
	"\x10server_processes\x18\x15 \x01(\x05R\x0fserverProcesses\"u\n" +
	"\n" +
	"ServerArgs\x122\n" +
	"\x05setup\x18\x01 \x01(\v2\x1a.grpc.testing.ServerConfigH\x00R\x05setup\x12(\n" +
	"\x04mark\x18\x02 \x01(\v2\x12.grpc.testing.MarkH\x00R\x04markB\t\n" +
	"\aargtype\"i\n" +
	"\fServerStatus\x12/\n" +
	"\x05stats\x18\x01 \x01(\v2\x19.grpc.testing.ServerStatsR\x05stats\x12\x12\n" +
	"\x04port\x18\x02 \x01(\x05R\x04port\x12\x14\n" +
	"\x05cores\x18\x03 \x01(\x05R\x05cores\"\r\n" +
	"\vCoreRequest\"$\n" +
	"\fCoreResponse\x12\x14\n" +
	"\x05cores\x18\x01 \x01(\x05R\x05cores\"\x06\n" +
	"\x04Void\"\xef\x02\n" +
	"\bScenario\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12?\n" +
	"\rclient_config\x18\x02 \x01(\v2\x1a.grpc.testing.ClientConfigR\fclientConfig\x12\x1f\n" +
	"\vnum_clients\x18\x03 \x01(\x05R\n" +
	"numClients\x12?\n" +
	"\rserver_config\x18\x04 \x01(\v2\x1a.grpc.testing.ServerConfigR\fserverConfig\x12\x1f\n" +
	"\vnum_servers\x18\x05 \x01(\x05R\n" +
	"numServers\x12%\n" +
	"\x0ewarmup_seconds\x18\x06 \x01(\x05R\rwarmupSeconds\x12+\n" +
	"\x11benchmark_seconds\x18\a \x01(\x05R\x10benchmarkSeconds\x127\n" +
	"\x18spawn_local_worker_count\x18\b \x01(\x05R\x15spawnLocalWorkerCount\"A\n" +
	"\tScenarios\x124\n" +
	"\tscenarios\x18\x01 \x03(\v2\x16.grpc.testing.ScenarioR\tscenarios\"\xad\a\n" +
	"\x15ScenarioResultSummary\x12\x10\n" +
	"\x03qps\x18\x01 \x01(\x01R\x03qps\x12-\n" +
	"\x13qps_per_server_core\x18\x02 \x01(\x01R\x10qpsPerServerCore\x12,\n" +
	"\x12server_system_time\x18\x03 \x01(\x01R\x10serverSystemTime\x12(\n" +
	"\x10server_user_time\x18\x04 \x01(\x01R\x0eserverUserTime\x12,\n" +
	"\x12client_system_time\x18\x05 \x01(\x01R\x10clientSystemTime\x12(\n" +
	"\x10client_user_time\x18\x06 \x01(\x01R\x0eclientUserTime\x12\x1d\n" +
	"\n" +
	"latency_50\x18\a \x01(\x01R\tlatency50\x12\x1d\n" +
	"\n" +
	"latency_90\x18\b \x01(\x01R\tlatency90\x12\x1d\n" +
	"\n" +
	"latency_95\x18\t \x01(\x01R\tlatency95\x12\x1d\n" +
	"\n" +
	"latency_99\x18\n" +
	" \x01(\x01R\tlatency99\x12\x1f\n" +
	"\vlatency_999\x18\v \x01(\x01R\n" +
	"latency999\x12(\n" +
	"\x10server_cpu_usage\x18\f \x01(\x01R\x0eserverCpuUsage\x12C\n" +
	"\x1esuccessful_requests_per_second\x18\r \x01(\x01R\x1bsuccessfulRequestsPerSecond\x12;\n" +
	"\x1afailed_requests_per_second\x18\x0e \x01(\x01R\x17failedRequestsPerSecond\x127\n" +
	"\x18client_polls_per_request\x18\x0f \x01(\x01R\x15clientPollsPerRequest\x127\n" +
	"\x18server_polls_per_request\x18\x10 \x01(\x01R\x15serverPollsPerRequest\x12:\n" +
	"\x1aserver_queries_per_cpu_sec\x18\x11 \x01(\x01R\x16serverQueriesPerCpuSec\x12:\n" +
	"\x1aclient_queries_per_cpu_sec\x18\x12 \x01(\x01R\x16clientQueriesPerCpuSec\x129\n" +
	"\n" +
	"start_time\x18\x13 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x125\n" +
	"\bend_time\x18\x14 \x01(\v2\x1a.google.protobuf.TimestampR\aendTime\"\xf6\x03\n" +
	"\x0eScenarioResult\x122\n" +
	"\bscenario\x18\x01 \x01(\v2\x16.grpc.testing.ScenarioR\bscenario\x129\n" +
	"\tlatencies\x18\x02 \x01(\v2\x1b.grpc.testing.HistogramDataR\tlatencies\x12<\n" +
	"\fclient_stats\x18\x03 \x03(\v2\x19.grpc.testing.ClientStatsR\vclientStats\x12<\n" +
	"\fserver_stats\x18\x04 \x03(\v2\x19.grpc.testing.ServerStatsR\vserverStats\x12!\n" +
	"\fserver_cores\x18\x05 \x03(\x05R\vserverCores\x12=\n" +
	"\asummary\x18\x06 \x01(\v2#.grpc.testing.ScenarioResultSummaryR\asummary\x12%\n" +
	"\x0eclient_success\x18\a \x03(\bR\rclientSuccess\x12%\n" +
	"\x0eserver_success\x18\b \x03(\bR\rserverSuccess\x12I\n" +
	"\x0frequest_results\x18\t \x03(\v2 .grpc.testing.RequestResultCountR\x0erequestResults*V\n" +
	"\n" +
	"ClientType\x12\x0f\n" +
	"\vSYNC_CLIENT\x10\x00\x12\x10\n" +
	"\fASYNC_CLIENT\x10\x01\x12\x10\n" +
	"\fOTHER_CLIENT\x10\x02\x12\x13\n" +
	"\x0fCALLBACK_CLIENT\x10\x03*p\n" +
	"\n" +
	"ServerType\x12\x0f\n" +
	"\vSYNC_SERVER\x10\x00\x12\x10\n" +
	"\fASYNC_SERVER\x10\x01\x12\x18\n" +
	"\x14ASYNC_GENERIC_SERVER\x10\x02\x12\x10\n" +
	"\fOTHER_SERVER\x10\x03\x12\x13\n" +
	"\x0fCALLBACK_SERVER\x10\x04*r\n" +
	"\aRpcType\x12\t\n" +
	"\x05UNARY\x10\x00\x12\r\n" +
	"\tSTREAMING\x10\x01\x12\x19\n" +
	"\x15STREAMING_FROM_CLIENT\x10\x02\x12\x19\n" +
	"\x15STREAMING_FROM_SERVER\x10\x03\x12\x17\n" +
	"\x13STREAMING_BOTH_WAYS\x10\x04B!\n" +
	"\x0fio.grpc.testingB\fControlProtoP\x01b\x06proto3"

var (
	file_grpc_testing_control_proto_rawDescOnce sync.Once
	file_grpc_testing_control_proto_rawDescData []byte
)

func file_grpc_testing_control_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_grpc_testing_control_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_grpc_testing_control_proto_msgTypes = make([]protoimpl.MessageInfo, 19)
var file_grpc_testing_control_proto_goTypes = []any{
	(ClientType)(0),
	(ServerType)(0),
	(RpcType)(0),
	(*PoissonParams)(nil),
	(*ClosedLoopParams)(nil),
	(*LoadParams)(nil),
	(*SecurityParams)(nil),
	(*ChannelArg)(nil),
	(*ClientConfig)(nil),
	(*ClientStatus)(nil),
	(*Mark)(nil),
	(*ClientArgs)(nil),
	(*ServerConfig)(nil),
	(*ServerArgs)(nil),
	(*ServerStatus)(nil),
	(*CoreRequest)(nil),
	(*CoreResponse)(nil),
	(*Void)(nil),
	(*Scenario)(nil),
	(*Scenarios)(nil),
	(*ScenarioResultSummary)(nil),
	(*ScenarioResult)(nil),
	(*PayloadConfig)(nil),
	(*HistogramParams)(nil),
	(*ClientStats)(nil),
	(*ServerStats)(nil),
	(*timestamppb.Timestamp)(nil),
	(*HistogramData)(nil),
	(*RequestResultCount)(nil),
}
var file_grpc_testing_control_proto_depIdxs = []int32{
	4,
	3,
	0,
	6,
	2,
	5,
	22,
	23,
	7,
	24,
	8,
	10,
	1,
	6,
	22,
	7,
	12,
	10,
	25,
	8,
	12,
	18,
	26,
	26,
	18,
	27,
	24,
	25,
	20,
	28,
	30,
	30,
	30,
	30,
	0,
}

func init()                                 { file_grpc_testing_control_proto_init() }
func file_grpc_testing_control_proto_init() { _ = "STUB: not implemented"; return }
