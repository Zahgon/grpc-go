package grpc_channelz_v1

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChannelConnectivityState_State int32

const (
	ChannelConnectivityState_UNKNOWN           ChannelConnectivityState_State = 0
	ChannelConnectivityState_IDLE              ChannelConnectivityState_State = 1
	ChannelConnectivityState_CONNECTING        ChannelConnectivityState_State = 2
	ChannelConnectivityState_READY             ChannelConnectivityState_State = 3
	ChannelConnectivityState_TRANSIENT_FAILURE ChannelConnectivityState_State = 4
	ChannelConnectivityState_SHUTDOWN          ChannelConnectivityState_State = 5
)

var (
	ChannelConnectivityState_State_name = map[int32]string{
		0: "UNKNOWN",
		1: "IDLE",
		2: "CONNECTING",
		3: "READY",
		4: "TRANSIENT_FAILURE",
		5: "SHUTDOWN",
	}
	ChannelConnectivityState_State_value = map[string]int32{
		"UNKNOWN":           0,
		"IDLE":              1,
		"CONNECTING":        2,
		"READY":             3,
		"TRANSIENT_FAILURE": 4,
		"SHUTDOWN":          5,
	}
)

func (x ChannelConnectivityState_State) Enum() *ChannelConnectivityState_State {
	_ = "STUB: not implemented"
	return nil
}

func (x ChannelConnectivityState_State) String() string { _ = "STUB: not implemented"; return "" }

func (ChannelConnectivityState_State) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (ChannelConnectivityState_State) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x ChannelConnectivityState_State) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (ChannelConnectivityState_State) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ChannelTraceEvent_Severity int32

const (
	ChannelTraceEvent_CT_UNKNOWN ChannelTraceEvent_Severity = 0
	ChannelTraceEvent_CT_INFO    ChannelTraceEvent_Severity = 1
	ChannelTraceEvent_CT_WARNING ChannelTraceEvent_Severity = 2
	ChannelTraceEvent_CT_ERROR   ChannelTraceEvent_Severity = 3
)

var (
	ChannelTraceEvent_Severity_name = map[int32]string{
		0: "CT_UNKNOWN",
		1: "CT_INFO",
		2: "CT_WARNING",
		3: "CT_ERROR",
	}
	ChannelTraceEvent_Severity_value = map[string]int32{
		"CT_UNKNOWN": 0,
		"CT_INFO":    1,
		"CT_WARNING": 2,
		"CT_ERROR":   3,
	}
)

func (x ChannelTraceEvent_Severity) Enum() *ChannelTraceEvent_Severity {
	_ = "STUB: not implemented"
	return nil
}

func (x ChannelTraceEvent_Severity) String() string { _ = "STUB: not implemented"; return "" }

func (ChannelTraceEvent_Severity) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (ChannelTraceEvent_Severity) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x ChannelTraceEvent_Severity) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (ChannelTraceEvent_Severity) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Channel struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Ref *ChannelRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`

	Data *ChannelData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`

	ChannelRef []*ChannelRef `protobuf:"bytes,3,rep,name=channel_ref,json=channelRef,proto3" json:"channel_ref,omitempty"`

	SubchannelRef []*SubchannelRef `protobuf:"bytes,4,rep,name=subchannel_ref,json=subchannelRef,proto3" json:"subchannel_ref,omitempty"`

	SocketRef     []*SocketRef `protobuf:"bytes,5,rep,name=socket_ref,json=socketRef,proto3" json:"socket_ref,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Channel) Reset() { _ = "STUB: not implemented"; return }

func (x *Channel) String() string { _ = "STUB: not implemented"; return "" }

func (*Channel) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Channel) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Channel) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Channel) GetRef() *ChannelRef { _ = "STUB: not implemented"; return nil }

func (x *Channel) GetData() *ChannelData { _ = "STUB: not implemented"; return nil }

func (x *Channel) GetChannelRef() []*ChannelRef { _ = "STUB: not implemented"; return nil }

func (x *Channel) GetSubchannelRef() []*SubchannelRef { _ = "STUB: not implemented"; return nil }

func (x *Channel) GetSocketRef() []*SocketRef { _ = "STUB: not implemented"; return nil }

type Subchannel struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Ref *SubchannelRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`

	Data *ChannelData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`

	ChannelRef []*ChannelRef `protobuf:"bytes,3,rep,name=channel_ref,json=channelRef,proto3" json:"channel_ref,omitempty"`

	SubchannelRef []*SubchannelRef `protobuf:"bytes,4,rep,name=subchannel_ref,json=subchannelRef,proto3" json:"subchannel_ref,omitempty"`

	SocketRef     []*SocketRef `protobuf:"bytes,5,rep,name=socket_ref,json=socketRef,proto3" json:"socket_ref,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Subchannel) Reset() { _ = "STUB: not implemented"; return }

func (x *Subchannel) String() string { _ = "STUB: not implemented"; return "" }

func (*Subchannel) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Subchannel) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Subchannel) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Subchannel) GetRef() *SubchannelRef { _ = "STUB: not implemented"; return nil }

func (x *Subchannel) GetData() *ChannelData { _ = "STUB: not implemented"; return nil }

func (x *Subchannel) GetChannelRef() []*ChannelRef { _ = "STUB: not implemented"; return nil }

func (x *Subchannel) GetSubchannelRef() []*SubchannelRef { _ = "STUB: not implemented"; return nil }

func (x *Subchannel) GetSocketRef() []*SocketRef { _ = "STUB: not implemented"; return nil }

type ChannelConnectivityState struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	State         ChannelConnectivityState_State `protobuf:"varint,1,opt,name=state,proto3,enum=grpc.channelz.v1.ChannelConnectivityState_State" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChannelConnectivityState) Reset() { _ = "STUB: not implemented"; return }

func (x *ChannelConnectivityState) String() string { _ = "STUB: not implemented"; return "" }

func (*ChannelConnectivityState) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ChannelConnectivityState) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ChannelConnectivityState) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *ChannelConnectivityState) GetState() ChannelConnectivityState_State {
	_ = "STUB: not implemented"
	return *new(ChannelConnectivityState_State)
}

type ChannelData struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	State *ChannelConnectivityState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`

	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`

	Trace *ChannelTrace `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`

	CallsStarted int64 `protobuf:"varint,4,opt,name=calls_started,json=callsStarted,proto3" json:"calls_started,omitempty"`

	CallsSucceeded int64 `protobuf:"varint,5,opt,name=calls_succeeded,json=callsSucceeded,proto3" json:"calls_succeeded,omitempty"`

	CallsFailed int64 `protobuf:"varint,6,opt,name=calls_failed,json=callsFailed,proto3" json:"calls_failed,omitempty"`

	LastCallStartedTimestamp *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_call_started_timestamp,json=lastCallStartedTimestamp,proto3" json:"last_call_started_timestamp,omitempty"`

	MaxConnectionsPerSubchannel uint32 `protobuf:"varint,8,opt,name=max_connections_per_subchannel,json=maxConnectionsPerSubchannel,proto3" json:"max_connections_per_subchannel,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *ChannelData) Reset() { _ = "STUB: not implemented"; return }

func (x *ChannelData) String() string { _ = "STUB: not implemented"; return "" }

func (*ChannelData) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ChannelData) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ChannelData) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ChannelData) GetState() *ChannelConnectivityState { _ = "STUB: not implemented"; return nil }

func (x *ChannelData) GetTarget() string { _ = "STUB: not implemented"; return "" }

func (x *ChannelData) GetTrace() *ChannelTrace { _ = "STUB: not implemented"; return nil }

func (x *ChannelData) GetCallsStarted() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ChannelData) GetCallsSucceeded() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ChannelData) GetCallsFailed() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ChannelData) GetLastCallStartedTimestamp() *timestamppb.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

func (x *ChannelData) GetMaxConnectionsPerSubchannel() uint32 { _ = "STUB: not implemented"; return 0 }

type ChannelTraceEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`

	Severity ChannelTraceEvent_Severity `protobuf:"varint,2,opt,name=severity,proto3,enum=grpc.channelz.v1.ChannelTraceEvent_Severity" json:"severity,omitempty"`

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`

	ChildRef      isChannelTraceEvent_ChildRef `protobuf_oneof:"child_ref"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChannelTraceEvent) Reset() { _ = "STUB: not implemented"; return }

func (x *ChannelTraceEvent) String() string { _ = "STUB: not implemented"; return "" }

func (*ChannelTraceEvent) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ChannelTraceEvent) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ChannelTraceEvent) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ChannelTraceEvent) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (x *ChannelTraceEvent) GetSeverity() ChannelTraceEvent_Severity {
	_ = "STUB: not implemented"
	return *new(ChannelTraceEvent_Severity)
}

func (x *ChannelTraceEvent) GetTimestamp() *timestamppb.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

func (x *ChannelTraceEvent) GetChildRef() isChannelTraceEvent_ChildRef {
	_ = "STUB: not implemented"
	return *new(isChannelTraceEvent_ChildRef)
}

func (x *ChannelTraceEvent) GetChannelRef() *ChannelRef { _ = "STUB: not implemented"; return nil }

func (x *ChannelTraceEvent) GetSubchannelRef() *SubchannelRef {
	_ = "STUB: not implemented"
	return nil
}

type isChannelTraceEvent_ChildRef interface {
	isChannelTraceEvent_ChildRef()
}

type ChannelTraceEvent_ChannelRef struct {
	ChannelRef *ChannelRef `protobuf:"bytes,4,opt,name=channel_ref,json=channelRef,proto3,oneof"`
}

type ChannelTraceEvent_SubchannelRef struct {
	SubchannelRef *SubchannelRef `protobuf:"bytes,5,opt,name=subchannel_ref,json=subchannelRef,proto3,oneof"`
}

func (*ChannelTraceEvent_ChannelRef) isChannelTraceEvent_ChildRef() {
	_ = "STUB: not implemented"
	return
}

func (*ChannelTraceEvent_SubchannelRef) isChannelTraceEvent_ChildRef() {
	_ = "STUB: not implemented"
	return
}

type ChannelTrace struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	NumEventsLogged int64 `protobuf:"varint,1,opt,name=num_events_logged,json=numEventsLogged,proto3" json:"num_events_logged,omitempty"`

	CreationTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"`

	Events        []*ChannelTraceEvent `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChannelTrace) Reset() { _ = "STUB: not implemented"; return }

func (x *ChannelTrace) String() string { _ = "STUB: not implemented"; return "" }

func (*ChannelTrace) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ChannelTrace) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ChannelTrace) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ChannelTrace) GetNumEventsLogged() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ChannelTrace) GetCreationTimestamp() *timestamppb.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

func (x *ChannelTrace) GetEvents() []*ChannelTraceEvent { _ = "STUB: not implemented"; return nil }

type ChannelRef struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ChannelId int64 `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`

	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChannelRef) Reset() { _ = "STUB: not implemented"; return }

func (x *ChannelRef) String() string { _ = "STUB: not implemented"; return "" }

func (*ChannelRef) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ChannelRef) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ChannelRef) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ChannelRef) GetChannelId() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ChannelRef) GetName() string { _ = "STUB: not implemented"; return "" }

type SubchannelRef struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	SubchannelId int64 `protobuf:"varint,7,opt,name=subchannel_id,json=subchannelId,proto3" json:"subchannel_id,omitempty"`

	Name          string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubchannelRef) Reset() { _ = "STUB: not implemented"; return }

func (x *SubchannelRef) String() string { _ = "STUB: not implemented"; return "" }

func (*SubchannelRef) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SubchannelRef) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SubchannelRef) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SubchannelRef) GetSubchannelId() int64 { _ = "STUB: not implemented"; return 0 }

func (x *SubchannelRef) GetName() string { _ = "STUB: not implemented"; return "" }

type SocketRef struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	SocketId int64 `protobuf:"varint,3,opt,name=socket_id,json=socketId,proto3" json:"socket_id,omitempty"`

	Name          string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SocketRef) Reset() { _ = "STUB: not implemented"; return }

func (x *SocketRef) String() string { _ = "STUB: not implemented"; return "" }

func (*SocketRef) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SocketRef) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SocketRef) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SocketRef) GetSocketId() int64 { _ = "STUB: not implemented"; return 0 }

func (x *SocketRef) GetName() string { _ = "STUB: not implemented"; return "" }

type ServerRef struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ServerId int64 `protobuf:"varint,5,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`

	Name          string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerRef) Reset() { _ = "STUB: not implemented"; return }

func (x *ServerRef) String() string { _ = "STUB: not implemented"; return "" }

func (*ServerRef) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServerRef) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServerRef) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ServerRef) GetServerId() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ServerRef) GetName() string { _ = "STUB: not implemented"; return "" }

type Server struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Ref *ServerRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`

	Data *ServerData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`

	ListenSocket  []*SocketRef `protobuf:"bytes,3,rep,name=listen_socket,json=listenSocket,proto3" json:"listen_socket,omitempty"`
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

func (x *Server) GetRef() *ServerRef { _ = "STUB: not implemented"; return nil }

func (x *Server) GetData() *ServerData { _ = "STUB: not implemented"; return nil }

func (x *Server) GetListenSocket() []*SocketRef { _ = "STUB: not implemented"; return nil }

type ServerData struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Trace *ChannelTrace `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`

	CallsStarted int64 `protobuf:"varint,2,opt,name=calls_started,json=callsStarted,proto3" json:"calls_started,omitempty"`

	CallsSucceeded int64 `protobuf:"varint,3,opt,name=calls_succeeded,json=callsSucceeded,proto3" json:"calls_succeeded,omitempty"`

	CallsFailed int64 `protobuf:"varint,4,opt,name=calls_failed,json=callsFailed,proto3" json:"calls_failed,omitempty"`

	LastCallStartedTimestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_call_started_timestamp,json=lastCallStartedTimestamp,proto3" json:"last_call_started_timestamp,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *ServerData) Reset() { _ = "STUB: not implemented"; return }

func (x *ServerData) String() string { _ = "STUB: not implemented"; return "" }

func (*ServerData) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServerData) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServerData) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ServerData) GetTrace() *ChannelTrace { _ = "STUB: not implemented"; return nil }

func (x *ServerData) GetCallsStarted() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ServerData) GetCallsSucceeded() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ServerData) GetCallsFailed() int64 { _ = "STUB: not implemented"; return 0 }

func (x *ServerData) GetLastCallStartedTimestamp() *timestamppb.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

type Socket struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Ref *SocketRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`

	Data *SocketData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`

	Local *Address `protobuf:"bytes,3,opt,name=local,proto3" json:"local,omitempty"`

	Remote *Address `protobuf:"bytes,4,opt,name=remote,proto3" json:"remote,omitempty"`

	Security *Security `protobuf:"bytes,5,opt,name=security,proto3" json:"security,omitempty"`

	RemoteName    string `protobuf:"bytes,6,opt,name=remote_name,json=remoteName,proto3" json:"remote_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Socket) Reset() { _ = "STUB: not implemented"; return }

func (x *Socket) String() string { _ = "STUB: not implemented"; return "" }

func (*Socket) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Socket) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Socket) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Socket) GetRef() *SocketRef { _ = "STUB: not implemented"; return nil }

func (x *Socket) GetData() *SocketData { _ = "STUB: not implemented"; return nil }

func (x *Socket) GetLocal() *Address { _ = "STUB: not implemented"; return nil }

func (x *Socket) GetRemote() *Address { _ = "STUB: not implemented"; return nil }

func (x *Socket) GetSecurity() *Security { _ = "STUB: not implemented"; return nil }

func (x *Socket) GetRemoteName() string { _ = "STUB: not implemented"; return "" }

type SocketData struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	StreamsStarted int64 `protobuf:"varint,1,opt,name=streams_started,json=streamsStarted,proto3" json:"streams_started,omitempty"`

	StreamsSucceeded int64 `protobuf:"varint,2,opt,name=streams_succeeded,json=streamsSucceeded,proto3" json:"streams_succeeded,omitempty"`

	StreamsFailed int64 `protobuf:"varint,3,opt,name=streams_failed,json=streamsFailed,proto3" json:"streams_failed,omitempty"`

	MessagesSent int64 `protobuf:"varint,4,opt,name=messages_sent,json=messagesSent,proto3" json:"messages_sent,omitempty"`

	MessagesReceived int64 `protobuf:"varint,5,opt,name=messages_received,json=messagesReceived,proto3" json:"messages_received,omitempty"`

	KeepAlivesSent int64 `protobuf:"varint,6,opt,name=keep_alives_sent,json=keepAlivesSent,proto3" json:"keep_alives_sent,omitempty"`

	LastLocalStreamCreatedTimestamp *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_local_stream_created_timestamp,json=lastLocalStreamCreatedTimestamp,proto3" json:"last_local_stream_created_timestamp,omitempty"`

	LastRemoteStreamCreatedTimestamp *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=last_remote_stream_created_timestamp,json=lastRemoteStreamCreatedTimestamp,proto3" json:"last_remote_stream_created_timestamp,omitempty"`

	LastMessageSentTimestamp *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=last_message_sent_timestamp,json=lastMessageSentTimestamp,proto3" json:"last_message_sent_timestamp,omitempty"`

	LastMessageReceivedTimestamp *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=last_message_received_timestamp,json=lastMessageReceivedTimestamp,proto3" json:"last_message_received_timestamp,omitempty"`

	LocalFlowControlWindow *wrapperspb.Int64Value `protobuf:"bytes,11,opt,name=local_flow_control_window,json=localFlowControlWindow,proto3" json:"local_flow_control_window,omitempty"`

	RemoteFlowControlWindow *wrapperspb.Int64Value `protobuf:"bytes,12,opt,name=remote_flow_control_window,json=remoteFlowControlWindow,proto3" json:"remote_flow_control_window,omitempty"`

	Option []*SocketOption `protobuf:"bytes,13,rep,name=option,proto3" json:"option,omitempty"`

	ReceivedGoawayError *wrapperspb.UInt32Value `protobuf:"bytes,14,opt,name=received_goaway_error,json=receivedGoawayError,proto3" json:"received_goaway_error,omitempty"`

	PeerMaxConcurrentStreams uint32 `protobuf:"varint,15,opt,name=peer_max_concurrent_streams,json=peerMaxConcurrentStreams,proto3" json:"peer_max_concurrent_streams,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *SocketData) Reset() { _ = "STUB: not implemented"; return }

func (x *SocketData) String() string { _ = "STUB: not implemented"; return "" }

func (*SocketData) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SocketData) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SocketData) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SocketData) GetStreamsStarted() int64 { _ = "STUB: not implemented"; return 0 }

func (x *SocketData) GetStreamsSucceeded() int64 { _ = "STUB: not implemented"; return 0 }

func (x *SocketData) GetStreamsFailed() int64 { _ = "STUB: not implemented"; return 0 }

func (x *SocketData) GetMessagesSent() int64 { _ = "STUB: not implemented"; return 0 }

func (x *SocketData) GetMessagesReceived() int64 { _ = "STUB: not implemented"; return 0 }

func (x *SocketData) GetKeepAlivesSent() int64 { _ = "STUB: not implemented"; return 0 }

func (x *SocketData) GetLastLocalStreamCreatedTimestamp() *timestamppb.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

func (x *SocketData) GetLastRemoteStreamCreatedTimestamp() *timestamppb.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

func (x *SocketData) GetLastMessageSentTimestamp() *timestamppb.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

func (x *SocketData) GetLastMessageReceivedTimestamp() *timestamppb.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

func (x *SocketData) GetLocalFlowControlWindow() *wrapperspb.Int64Value {
	_ = "STUB: not implemented"
	return nil
}

func (x *SocketData) GetRemoteFlowControlWindow() *wrapperspb.Int64Value {
	_ = "STUB: not implemented"
	return nil
}

func (x *SocketData) GetOption() []*SocketOption { _ = "STUB: not implemented"; return nil }

func (x *SocketData) GetReceivedGoawayError() *wrapperspb.UInt32Value {
	_ = "STUB: not implemented"
	return nil
}

func (x *SocketData) GetPeerMaxConcurrentStreams() uint32 { _ = "STUB: not implemented"; return 0 }

type Address struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Address       isAddress_Address `protobuf_oneof:"address"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() { _ = "STUB: not implemented"; return }

func (x *Address) String() string { _ = "STUB: not implemented"; return "" }

func (*Address) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Address) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Address) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Address) GetAddress() isAddress_Address {
	_ = "STUB: not implemented"
	return *new(isAddress_Address)
}

func (x *Address) GetTcpipAddress() *Address_TcpIpAddress { _ = "STUB: not implemented"; return nil }

func (x *Address) GetUdsAddress() *Address_UdsAddress { _ = "STUB: not implemented"; return nil }

func (x *Address) GetOtherAddress() *Address_OtherAddress { _ = "STUB: not implemented"; return nil }

type isAddress_Address interface {
	isAddress_Address()
}

type Address_TcpipAddress struct {
	TcpipAddress *Address_TcpIpAddress `protobuf:"bytes,1,opt,name=tcpip_address,json=tcpipAddress,proto3,oneof"`
}

type Address_UdsAddress_ struct {
	UdsAddress *Address_UdsAddress `protobuf:"bytes,2,opt,name=uds_address,json=udsAddress,proto3,oneof"`
}

type Address_OtherAddress_ struct {
	OtherAddress *Address_OtherAddress `protobuf:"bytes,3,opt,name=other_address,json=otherAddress,proto3,oneof"`
}

func (*Address_TcpipAddress) isAddress_Address() { _ = "STUB: not implemented"; return }

func (*Address_UdsAddress_) isAddress_Address() { _ = "STUB: not implemented"; return }

func (*Address_OtherAddress_) isAddress_Address() { _ = "STUB: not implemented"; return }

type Security struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Model         isSecurity_Model `protobuf_oneof:"model"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Security) Reset() { _ = "STUB: not implemented"; return }

func (x *Security) String() string { _ = "STUB: not implemented"; return "" }

func (*Security) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Security) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Security) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Security) GetModel() isSecurity_Model {
	_ = "STUB: not implemented"
	return *new(isSecurity_Model)
}

func (x *Security) GetTls() *Security_Tls { _ = "STUB: not implemented"; return nil }

func (x *Security) GetOther() *Security_OtherSecurity { _ = "STUB: not implemented"; return nil }

type isSecurity_Model interface {
	isSecurity_Model()
}

type Security_Tls_ struct {
	Tls *Security_Tls `protobuf:"bytes,1,opt,name=tls,proto3,oneof"`
}

type Security_Other struct {
	Other *Security_OtherSecurity `protobuf:"bytes,2,opt,name=other,proto3,oneof"`
}

func (*Security_Tls_) isSecurity_Model() { _ = "STUB: not implemented"; return }

func (*Security_Other) isSecurity_Model() { _ = "STUB: not implemented"; return }

type SocketOption struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`

	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`

	Additional    *anypb.Any `protobuf:"bytes,3,opt,name=additional,proto3" json:"additional,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SocketOption) Reset() { _ = "STUB: not implemented"; return }

func (x *SocketOption) String() string { _ = "STUB: not implemented"; return "" }

func (*SocketOption) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SocketOption) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SocketOption) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SocketOption) GetName() string { _ = "STUB: not implemented"; return "" }

func (x *SocketOption) GetValue() string { _ = "STUB: not implemented"; return "" }

func (x *SocketOption) GetAdditional() *anypb.Any { _ = "STUB: not implemented"; return nil }

type SocketOptionTimeout struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Duration      *durationpb.Duration   `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SocketOptionTimeout) Reset() { _ = "STUB: not implemented"; return }

func (x *SocketOptionTimeout) String() string { _ = "STUB: not implemented"; return "" }

func (*SocketOptionTimeout) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SocketOptionTimeout) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SocketOptionTimeout) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *SocketOptionTimeout) GetDuration() *durationpb.Duration {
	_ = "STUB: not implemented"
	return nil
}

type SocketOptionLinger struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Active bool `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`

	Duration      *durationpb.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SocketOptionLinger) Reset() { _ = "STUB: not implemented"; return }

func (x *SocketOptionLinger) String() string { _ = "STUB: not implemented"; return "" }

func (*SocketOptionLinger) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SocketOptionLinger) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SocketOptionLinger) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SocketOptionLinger) GetActive() bool { _ = "STUB: not implemented"; return false }

func (x *SocketOptionLinger) GetDuration() *durationpb.Duration {
	_ = "STUB: not implemented"
	return nil
}

type SocketOptionTcpInfo struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	TcpiState        uint32                 `protobuf:"varint,1,opt,name=tcpi_state,json=tcpiState,proto3" json:"tcpi_state,omitempty"`
	TcpiCaState      uint32                 `protobuf:"varint,2,opt,name=tcpi_ca_state,json=tcpiCaState,proto3" json:"tcpi_ca_state,omitempty"`
	TcpiRetransmits  uint32                 `protobuf:"varint,3,opt,name=tcpi_retransmits,json=tcpiRetransmits,proto3" json:"tcpi_retransmits,omitempty"`
	TcpiProbes       uint32                 `protobuf:"varint,4,opt,name=tcpi_probes,json=tcpiProbes,proto3" json:"tcpi_probes,omitempty"`
	TcpiBackoff      uint32                 `protobuf:"varint,5,opt,name=tcpi_backoff,json=tcpiBackoff,proto3" json:"tcpi_backoff,omitempty"`
	TcpiOptions      uint32                 `protobuf:"varint,6,opt,name=tcpi_options,json=tcpiOptions,proto3" json:"tcpi_options,omitempty"`
	TcpiSndWscale    uint32                 `protobuf:"varint,7,opt,name=tcpi_snd_wscale,json=tcpiSndWscale,proto3" json:"tcpi_snd_wscale,omitempty"`
	TcpiRcvWscale    uint32                 `protobuf:"varint,8,opt,name=tcpi_rcv_wscale,json=tcpiRcvWscale,proto3" json:"tcpi_rcv_wscale,omitempty"`
	TcpiRto          uint32                 `protobuf:"varint,9,opt,name=tcpi_rto,json=tcpiRto,proto3" json:"tcpi_rto,omitempty"`
	TcpiAto          uint32                 `protobuf:"varint,10,opt,name=tcpi_ato,json=tcpiAto,proto3" json:"tcpi_ato,omitempty"`
	TcpiSndMss       uint32                 `protobuf:"varint,11,opt,name=tcpi_snd_mss,json=tcpiSndMss,proto3" json:"tcpi_snd_mss,omitempty"`
	TcpiRcvMss       uint32                 `protobuf:"varint,12,opt,name=tcpi_rcv_mss,json=tcpiRcvMss,proto3" json:"tcpi_rcv_mss,omitempty"`
	TcpiUnacked      uint32                 `protobuf:"varint,13,opt,name=tcpi_unacked,json=tcpiUnacked,proto3" json:"tcpi_unacked,omitempty"`
	TcpiSacked       uint32                 `protobuf:"varint,14,opt,name=tcpi_sacked,json=tcpiSacked,proto3" json:"tcpi_sacked,omitempty"`
	TcpiLost         uint32                 `protobuf:"varint,15,opt,name=tcpi_lost,json=tcpiLost,proto3" json:"tcpi_lost,omitempty"`
	TcpiRetrans      uint32                 `protobuf:"varint,16,opt,name=tcpi_retrans,json=tcpiRetrans,proto3" json:"tcpi_retrans,omitempty"`
	TcpiFackets      uint32                 `protobuf:"varint,17,opt,name=tcpi_fackets,json=tcpiFackets,proto3" json:"tcpi_fackets,omitempty"`
	TcpiLastDataSent uint32                 `protobuf:"varint,18,opt,name=tcpi_last_data_sent,json=tcpiLastDataSent,proto3" json:"tcpi_last_data_sent,omitempty"`
	TcpiLastAckSent  uint32                 `protobuf:"varint,19,opt,name=tcpi_last_ack_sent,json=tcpiLastAckSent,proto3" json:"tcpi_last_ack_sent,omitempty"`
	TcpiLastDataRecv uint32                 `protobuf:"varint,20,opt,name=tcpi_last_data_recv,json=tcpiLastDataRecv,proto3" json:"tcpi_last_data_recv,omitempty"`
	TcpiLastAckRecv  uint32                 `protobuf:"varint,21,opt,name=tcpi_last_ack_recv,json=tcpiLastAckRecv,proto3" json:"tcpi_last_ack_recv,omitempty"`
	TcpiPmtu         uint32                 `protobuf:"varint,22,opt,name=tcpi_pmtu,json=tcpiPmtu,proto3" json:"tcpi_pmtu,omitempty"`
	TcpiRcvSsthresh  uint32                 `protobuf:"varint,23,opt,name=tcpi_rcv_ssthresh,json=tcpiRcvSsthresh,proto3" json:"tcpi_rcv_ssthresh,omitempty"`
	TcpiRtt          uint32                 `protobuf:"varint,24,opt,name=tcpi_rtt,json=tcpiRtt,proto3" json:"tcpi_rtt,omitempty"`
	TcpiRttvar       uint32                 `protobuf:"varint,25,opt,name=tcpi_rttvar,json=tcpiRttvar,proto3" json:"tcpi_rttvar,omitempty"`
	TcpiSndSsthresh  uint32                 `protobuf:"varint,26,opt,name=tcpi_snd_ssthresh,json=tcpiSndSsthresh,proto3" json:"tcpi_snd_ssthresh,omitempty"`
	TcpiSndCwnd      uint32                 `protobuf:"varint,27,opt,name=tcpi_snd_cwnd,json=tcpiSndCwnd,proto3" json:"tcpi_snd_cwnd,omitempty"`
	TcpiAdvmss       uint32                 `protobuf:"varint,28,opt,name=tcpi_advmss,json=tcpiAdvmss,proto3" json:"tcpi_advmss,omitempty"`
	TcpiReordering   uint32                 `protobuf:"varint,29,opt,name=tcpi_reordering,json=tcpiReordering,proto3" json:"tcpi_reordering,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *SocketOptionTcpInfo) Reset() { _ = "STUB: not implemented"; return }

func (x *SocketOptionTcpInfo) String() string { _ = "STUB: not implemented"; return "" }

func (*SocketOptionTcpInfo) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SocketOptionTcpInfo) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SocketOptionTcpInfo) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *SocketOptionTcpInfo) GetTcpiState() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiCaState() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiRetransmits() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiProbes() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiBackoff() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiOptions() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiSndWscale() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiRcvWscale() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiRto() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiAto() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiSndMss() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiRcvMss() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiUnacked() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiSacked() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiLost() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiRetrans() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiFackets() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiLastDataSent() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiLastAckSent() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiLastDataRecv() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiLastAckRecv() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiPmtu() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiRcvSsthresh() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiRtt() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiRttvar() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiSndSsthresh() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiSndCwnd() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiAdvmss() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *SocketOptionTcpInfo) GetTcpiReordering() uint32 { _ = "STUB: not implemented"; return 0 }

type GetTopChannelsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	StartChannelId int64 `protobuf:"varint,1,opt,name=start_channel_id,json=startChannelId,proto3" json:"start_channel_id,omitempty"`

	MaxResults    int64 `protobuf:"varint,2,opt,name=max_results,json=maxResults,proto3" json:"max_results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTopChannelsRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *GetTopChannelsRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*GetTopChannelsRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetTopChannelsRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetTopChannelsRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *GetTopChannelsRequest) GetStartChannelId() int64 { _ = "STUB: not implemented"; return 0 }

func (x *GetTopChannelsRequest) GetMaxResults() int64 { _ = "STUB: not implemented"; return 0 }

type GetTopChannelsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Channel []*Channel `protobuf:"bytes,1,rep,name=channel,proto3" json:"channel,omitempty"`

	End           bool `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTopChannelsResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *GetTopChannelsResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*GetTopChannelsResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetTopChannelsResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetTopChannelsResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *GetTopChannelsResponse) GetChannel() []*Channel { _ = "STUB: not implemented"; return nil }

func (x *GetTopChannelsResponse) GetEnd() bool { _ = "STUB: not implemented"; return false }

type GetServersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	StartServerId int64 `protobuf:"varint,1,opt,name=start_server_id,json=startServerId,proto3" json:"start_server_id,omitempty"`

	MaxResults    int64 `protobuf:"varint,2,opt,name=max_results,json=maxResults,proto3" json:"max_results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServersRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *GetServersRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*GetServersRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetServersRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetServersRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GetServersRequest) GetStartServerId() int64 { _ = "STUB: not implemented"; return 0 }

func (x *GetServersRequest) GetMaxResults() int64 { _ = "STUB: not implemented"; return 0 }

type GetServersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Server []*Server `protobuf:"bytes,1,rep,name=server,proto3" json:"server,omitempty"`

	End           bool `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServersResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *GetServersResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*GetServersResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetServersResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetServersResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GetServersResponse) GetServer() []*Server { _ = "STUB: not implemented"; return nil }

func (x *GetServersResponse) GetEnd() bool { _ = "STUB: not implemented"; return false }

type GetServerRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ServerId      int64 `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServerRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *GetServerRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*GetServerRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetServerRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetServerRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GetServerRequest) GetServerId() int64 { _ = "STUB: not implemented"; return 0 }

type GetServerResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Server        *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServerResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *GetServerResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*GetServerResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetServerResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetServerResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GetServerResponse) GetServer() *Server { _ = "STUB: not implemented"; return nil }

type GetServerSocketsRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	ServerId int64                  `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`

	StartSocketId int64 `protobuf:"varint,2,opt,name=start_socket_id,json=startSocketId,proto3" json:"start_socket_id,omitempty"`

	MaxResults    int64 `protobuf:"varint,3,opt,name=max_results,json=maxResults,proto3" json:"max_results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServerSocketsRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *GetServerSocketsRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*GetServerSocketsRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetServerSocketsRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetServerSocketsRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *GetServerSocketsRequest) GetServerId() int64 { _ = "STUB: not implemented"; return 0 }

func (x *GetServerSocketsRequest) GetStartSocketId() int64 { _ = "STUB: not implemented"; return 0 }

func (x *GetServerSocketsRequest) GetMaxResults() int64 { _ = "STUB: not implemented"; return 0 }

type GetServerSocketsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	SocketRef []*SocketRef `protobuf:"bytes,1,rep,name=socket_ref,json=socketRef,proto3" json:"socket_ref,omitempty"`

	End           bool `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServerSocketsResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *GetServerSocketsResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*GetServerSocketsResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetServerSocketsResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetServerSocketsResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *GetServerSocketsResponse) GetSocketRef() []*SocketRef {
	_ = "STUB: not implemented"
	return nil
}

func (x *GetServerSocketsResponse) GetEnd() bool { _ = "STUB: not implemented"; return false }

type GetChannelRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ChannelId     int64 `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChannelRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *GetChannelRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*GetChannelRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetChannelRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetChannelRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GetChannelRequest) GetChannelId() int64 { _ = "STUB: not implemented"; return 0 }

type GetChannelResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Channel       *Channel `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChannelResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *GetChannelResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*GetChannelResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetChannelResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetChannelResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GetChannelResponse) GetChannel() *Channel { _ = "STUB: not implemented"; return nil }

type GetSubchannelRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	SubchannelId  int64 `protobuf:"varint,1,opt,name=subchannel_id,json=subchannelId,proto3" json:"subchannel_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSubchannelRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *GetSubchannelRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*GetSubchannelRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetSubchannelRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetSubchannelRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *GetSubchannelRequest) GetSubchannelId() int64 { _ = "STUB: not implemented"; return 0 }

type GetSubchannelResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Subchannel    *Subchannel `protobuf:"bytes,1,opt,name=subchannel,proto3" json:"subchannel,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSubchannelResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *GetSubchannelResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*GetSubchannelResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetSubchannelResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetSubchannelResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *GetSubchannelResponse) GetSubchannel() *Subchannel { _ = "STUB: not implemented"; return nil }

type GetSocketRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	SocketId int64 `protobuf:"varint,1,opt,name=socket_id,json=socketId,proto3" json:"socket_id,omitempty"`

	Summary       bool `protobuf:"varint,2,opt,name=summary,proto3" json:"summary,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSocketRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *GetSocketRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*GetSocketRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetSocketRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetSocketRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GetSocketRequest) GetSocketId() int64 { _ = "STUB: not implemented"; return 0 }

func (x *GetSocketRequest) GetSummary() bool { _ = "STUB: not implemented"; return false }

type GetSocketResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Socket        *Socket `protobuf:"bytes,1,opt,name=socket,proto3" json:"socket,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSocketResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *GetSocketResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*GetSocketResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GetSocketResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GetSocketResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GetSocketResponse) GetSocket() *Socket { _ = "STUB: not implemented"; return nil }

type Address_TcpIpAddress struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	IpAddress []byte `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`

	Port          int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address_TcpIpAddress) Reset() { _ = "STUB: not implemented"; return }

func (x *Address_TcpIpAddress) String() string { _ = "STUB: not implemented"; return "" }

func (*Address_TcpIpAddress) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Address_TcpIpAddress) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Address_TcpIpAddress) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *Address_TcpIpAddress) GetIpAddress() []byte { _ = "STUB: not implemented"; return nil }

func (x *Address_TcpIpAddress) GetPort() int32 { _ = "STUB: not implemented"; return 0 }

type Address_UdsAddress struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address_UdsAddress) Reset() { _ = "STUB: not implemented"; return }

func (x *Address_UdsAddress) String() string { _ = "STUB: not implemented"; return "" }

func (*Address_UdsAddress) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Address_UdsAddress) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Address_UdsAddress) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Address_UdsAddress) GetFilename() string { _ = "STUB: not implemented"; return "" }

type Address_OtherAddress struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`

	Value         *anypb.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address_OtherAddress) Reset() { _ = "STUB: not implemented"; return }

func (x *Address_OtherAddress) String() string { _ = "STUB: not implemented"; return "" }

func (*Address_OtherAddress) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Address_OtherAddress) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Address_OtherAddress) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *Address_OtherAddress) GetName() string { _ = "STUB: not implemented"; return "" }

func (x *Address_OtherAddress) GetValue() *anypb.Any { _ = "STUB: not implemented"; return nil }

type Security_Tls struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	CipherSuite isSecurity_Tls_CipherSuite `protobuf_oneof:"cipher_suite"`

	LocalCertificate []byte `protobuf:"bytes,3,opt,name=local_certificate,json=localCertificate,proto3" json:"local_certificate,omitempty"`

	RemoteCertificate []byte `protobuf:"bytes,4,opt,name=remote_certificate,json=remoteCertificate,proto3" json:"remote_certificate,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Security_Tls) Reset() { _ = "STUB: not implemented"; return }

func (x *Security_Tls) String() string { _ = "STUB: not implemented"; return "" }

func (*Security_Tls) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Security_Tls) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Security_Tls) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Security_Tls) GetCipherSuite() isSecurity_Tls_CipherSuite {
	_ = "STUB: not implemented"
	return *new(isSecurity_Tls_CipherSuite)
}

func (x *Security_Tls) GetStandardName() string { _ = "STUB: not implemented"; return "" }

func (x *Security_Tls) GetOtherName() string { _ = "STUB: not implemented"; return "" }

func (x *Security_Tls) GetLocalCertificate() []byte { _ = "STUB: not implemented"; return nil }

func (x *Security_Tls) GetRemoteCertificate() []byte { _ = "STUB: not implemented"; return nil }

type isSecurity_Tls_CipherSuite interface {
	isSecurity_Tls_CipherSuite()
}

type Security_Tls_StandardName struct {
	StandardName string `protobuf:"bytes,1,opt,name=standard_name,json=standardName,proto3,oneof"`
}

type Security_Tls_OtherName struct {
	OtherName string `protobuf:"bytes,2,opt,name=other_name,json=otherName,proto3,oneof"`
}

func (*Security_Tls_StandardName) isSecurity_Tls_CipherSuite() { _ = "STUB: not implemented"; return }

func (*Security_Tls_OtherName) isSecurity_Tls_CipherSuite() { _ = "STUB: not implemented"; return }

type Security_OtherSecurity struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`

	Value         *anypb.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Security_OtherSecurity) Reset() { _ = "STUB: not implemented"; return }

func (x *Security_OtherSecurity) String() string { _ = "STUB: not implemented"; return "" }

func (*Security_OtherSecurity) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Security_OtherSecurity) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Security_OtherSecurity) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *Security_OtherSecurity) GetName() string { _ = "STUB: not implemented"; return "" }

func (x *Security_OtherSecurity) GetValue() *anypb.Any { _ = "STUB: not implemented"; return nil }

var File_grpc_channelz_v1_channelz_proto protoreflect.FileDescriptor

const file_grpc_channelz_v1_channelz_proto_rawDesc = "" +
	"\n" +
	"\x1fgrpc/channelz/v1/channelz.proto\x12\x10grpc.channelz.v1\x1a\x19google/protobuf/any.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\xaf\x02\n" +
	"\aChannel\x12.\n" +
	"\x03ref\x18\x01 \x01(\v2\x1c.grpc.channelz.v1.ChannelRefR\x03ref\x121\n" +
	"\x04data\x18\x02 \x01(\v2\x1d.grpc.channelz.v1.ChannelDataR\x04data\x12=\n" +
	"\vchannel_ref\x18\x03 \x03(\v2\x1c.grpc.channelz.v1.ChannelRefR\n" +
	"channelRef\x12F\n" +
	"\x0esubchannel_ref\x18\x04 \x03(\v2\x1f.grpc.channelz.v1.SubchannelRefR\rsubchannelRef\x12:\n" +
	"\n" +
	"socket_ref\x18\x05 \x03(\v2\x1b.grpc.channelz.v1.SocketRefR\tsocketRef\"\xb5\x02\n" +
	"\n" +
	"Subchannel\x121\n" +
	"\x03ref\x18\x01 \x01(\v2\x1f.grpc.channelz.v1.SubchannelRefR\x03ref\x121\n" +
	"\x04data\x18\x02 \x01(\v2\x1d.grpc.channelz.v1.ChannelDataR\x04data\x12=\n" +
	"\vchannel_ref\x18\x03 \x03(\v2\x1c.grpc.channelz.v1.ChannelRefR\n" +
	"channelRef\x12F\n" +
	"\x0esubchannel_ref\x18\x04 \x03(\v2\x1f.grpc.channelz.v1.SubchannelRefR\rsubchannelRef\x12:\n" +
	"\n" +
	"socket_ref\x18\x05 \x03(\v2\x1b.grpc.channelz.v1.SocketRefR\tsocketRef\"\xc2\x01\n" +
	"\x18ChannelConnectivityState\x12F\n" +
	"\x05state\x18\x01 \x01(\x0e20.grpc.channelz.v1.ChannelConnectivityState.StateR\x05state\"^\n" +
	"\x05State\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\b\n" +
	"\x04IDLE\x10\x01\x12\x0e\n" +
	"\n" +
	"CONNECTING\x10\x02\x12\t\n" +
	"\x05READY\x10\x03\x12\x15\n" +
	"\x11TRANSIENT_FAILURE\x10\x04\x12\f\n" +
	"\bSHUTDOWN\x10\x05\"\xae\x03\n" +
	"\vChannelData\x12@\n" +
	"\x05state\x18\x01 \x01(\v2*.grpc.channelz.v1.ChannelConnectivityStateR\x05state\x12\x16\n" +
	"\x06target\x18\x02 \x01(\tR\x06target\x124\n" +
	"\x05trace\x18\x03 \x01(\v2\x1e.grpc.channelz.v1.ChannelTraceR\x05trace\x12#\n" +
	"\rcalls_started\x18\x04 \x01(\x03R\fcallsStarted\x12'\n" +
	"\x0fcalls_succeeded\x18\x05 \x01(\x03R\x0ecallsSucceeded\x12!\n" +
	"\fcalls_failed\x18\x06 \x01(\x03R\vcallsFailed\x12Y\n" +
	"\x1blast_call_started_timestamp\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\x18lastCallStartedTimestamp\x12C\n" +
	"\x1emax_connections_per_subchannel\x18\b \x01(\rR\x1bmaxConnectionsPerSubchannel\"\x98\x03\n" +
	"\x11ChannelTraceEvent\x12 \n" +
	"\vdescription\x18\x01 \x01(\tR\vdescription\x12H\n" +
	"\bseverity\x18\x02 \x01(\x0e2,.grpc.channelz.v1.ChannelTraceEvent.SeverityR\bseverity\x128\n" +
	"\ttimestamp\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12?\n" +
	"\vchannel_ref\x18\x04 \x01(\v2\x1c.grpc.channelz.v1.ChannelRefH\x00R\n" +
	"channelRef\x12H\n" +
	"\x0esubchannel_ref\x18\x05 \x01(\v2\x1f.grpc.channelz.v1.SubchannelRefH\x00R\rsubchannelRef\"E\n" +
	"\bSeverity\x12\x0e\n" +
	"\n" +
	"CT_UNKNOWN\x10\x00\x12\v\n" +
	"\aCT_INFO\x10\x01\x12\x0e\n" +
	"\n" +
	"CT_WARNING\x10\x02\x12\f\n" +
	"\bCT_ERROR\x10\x03B\v\n" +
	"\tchild_ref\"\xc2\x01\n" +
	"\fChannelTrace\x12*\n" +
	"\x11num_events_logged\x18\x01 \x01(\x03R\x0fnumEventsLogged\x12I\n" +
	"\x12creation_timestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x11creationTimestamp\x12;\n" +
	"\x06events\x18\x03 \x03(\v2#.grpc.channelz.v1.ChannelTraceEventR\x06events\"c\n" +
	"\n" +
	"ChannelRef\x12\x1d\n" +
	"\n" +
	"channel_id\x18\x01 \x01(\x03R\tchannelId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04nameJ\x04\b\x03\x10\x04J\x04\b\x04\x10\x05J\x04\b\x05\x10\x06J\x04\b\x06\x10\aJ\x04\b\a\x10\bJ\x04\b\b\x10\t\"l\n" +
	"\rSubchannelRef\x12#\n" +
	"\rsubchannel_id\x18\a \x01(\x03R\fsubchannelId\x12\x12\n" +
	"\x04name\x18\b \x01(\tR\x04nameJ\x04\b\x01\x10\x02J\x04\b\x02\x10\x03J\x04\b\x03\x10\x04J\x04\b\x04\x10\x05J\x04\b\x05\x10\x06J\x04\b\x06\x10\a\"`\n" +
	"\tSocketRef\x12\x1b\n" +
	"\tsocket_id\x18\x03 \x01(\x03R\bsocketId\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04nameJ\x04\b\x01\x10\x02J\x04\b\x02\x10\x03J\x04\b\x05\x10\x06J\x04\b\x06\x10\aJ\x04\b\a\x10\bJ\x04\b\b\x10\t\"`\n" +
	"\tServerRef\x12\x1b\n" +
	"\tserver_id\x18\x05 \x01(\x03R\bserverId\x12\x12\n" +
	"\x04name\x18\x06 \x01(\tR\x04nameJ\x04\b\x01\x10\x02J\x04\b\x02\x10\x03J\x04\b\x03\x10\x04J\x04\b\x04\x10\x05J\x04\b\a\x10\bJ\x04\b\b\x10\t\"\xab\x01\n" +
	"\x06Server\x12-\n" +
	"\x03ref\x18\x01 \x01(\v2\x1b.grpc.channelz.v1.ServerRefR\x03ref\x120\n" +
	"\x04data\x18\x02 \x01(\v2\x1c.grpc.channelz.v1.ServerDataR\x04data\x12@\n" +
	"\rlisten_socket\x18\x03 \x03(\v2\x1b.grpc.channelz.v1.SocketRefR\flistenSocket\"\x8e\x02\n" +
	"\n" +
	"ServerData\x124\n" +
	"\x05trace\x18\x01 \x01(\v2\x1e.grpc.channelz.v1.ChannelTraceR\x05trace\x12#\n" +
	"\rcalls_started\x18\x02 \x01(\x03R\fcallsStarted\x12'\n" +
	"\x0fcalls_succeeded\x18\x03 \x01(\x03R\x0ecallsSucceeded\x12!\n" +
	"\fcalls_failed\x18\x04 \x01(\x03R\vcallsFailed\x12Y\n" +
	"\x1blast_call_started_timestamp\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\x18lastCallStartedTimestamp\"\xa6\x02\n" +
	"\x06Socket\x12-\n" +
	"\x03ref\x18\x01 \x01(\v2\x1b.grpc.channelz.v1.SocketRefR\x03ref\x120\n" +
	"\x04data\x18\x02 \x01(\v2\x1c.grpc.channelz.v1.SocketDataR\x04data\x12/\n" +
	"\x05local\x18\x03 \x01(\v2\x19.grpc.channelz.v1.AddressR\x05local\x121\n" +
	"\x06remote\x18\x04 \x01(\v2\x19.grpc.channelz.v1.AddressR\x06remote\x126\n" +
	"\bsecurity\x18\x05 \x01(\v2\x1a.grpc.channelz.v1.SecurityR\bsecurity\x12\x1f\n" +
	"\vremote_name\x18\x06 \x01(\tR\n" +
	"remoteName\"\x94\b\n" +
	"\n" +
	"SocketData\x12'\n" +
	"\x0fstreams_started\x18\x01 \x01(\x03R\x0estreamsStarted\x12+\n" +
	"\x11streams_succeeded\x18\x02 \x01(\x03R\x10streamsSucceeded\x12%\n" +
	"\x0estreams_failed\x18\x03 \x01(\x03R\rstreamsFailed\x12#\n" +
	"\rmessages_sent\x18\x04 \x01(\x03R\fmessagesSent\x12+\n" +
	"\x11messages_received\x18\x05 \x01(\x03R\x10messagesReceived\x12(\n" +
	"\x10keep_alives_sent\x18\x06 \x01(\x03R\x0ekeepAlivesSent\x12h\n" +
	"#last_local_stream_created_timestamp\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\x1flastLocalStreamCreatedTimestamp\x12j\n" +
	"$last_remote_stream_created_timestamp\x18\b \x01(\v2\x1a.google.protobuf.TimestampR lastRemoteStreamCreatedTimestamp\x12Y\n" +
	"\x1blast_message_sent_timestamp\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\x18lastMessageSentTimestamp\x12a\n" +
	"\x1flast_message_received_timestamp\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\x1clastMessageReceivedTimestamp\x12V\n" +
	"\x19local_flow_control_window\x18\v \x01(\v2\x1b.google.protobuf.Int64ValueR\x16localFlowControlWindow\x12X\n" +
	"\x1aremote_flow_control_window\x18\f \x01(\v2\x1b.google.protobuf.Int64ValueR\x17remoteFlowControlWindow\x126\n" +
	"\x06option\x18\r \x03(\v2\x1e.grpc.channelz.v1.SocketOptionR\x06option\x12P\n" +
	"\x15received_goaway_error\x18\x0e \x01(\v2\x1c.google.protobuf.UInt32ValueR\x13receivedGoawayError\x12=\n" +
	"\x1bpeer_max_concurrent_streams\x18\x0f \x01(\rR\x18peerMaxConcurrentStreams\"\xb8\x03\n" +
	"\aAddress\x12M\n" +
	"\rtcpip_address\x18\x01 \x01(\v2&.grpc.channelz.v1.Address.TcpIpAddressH\x00R\ftcpipAddress\x12G\n" +
	"\vuds_address\x18\x02 \x01(\v2$.grpc.channelz.v1.Address.UdsAddressH\x00R\n" +
	"udsAddress\x12M\n" +
	"\rother_address\x18\x03 \x01(\v2&.grpc.channelz.v1.Address.OtherAddressH\x00R\fotherAddress\x1aA\n" +
	"\fTcpIpAddress\x12\x1d\n" +
	"\n" +
	"ip_address\x18\x01 \x01(\fR\tipAddress\x12\x12\n" +
	"\x04port\x18\x02 \x01(\x05R\x04port\x1a(\n" +
	"\n" +
	"UdsAddress\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\x1aN\n" +
	"\fOtherAddress\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12*\n" +
	"\x05value\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\x05valueB\t\n" +
	"\aaddress\"\x96\x03\n" +
	"\bSecurity\x122\n" +
	"\x03tls\x18\x01 \x01(\v2\x1e.grpc.channelz.v1.Security.TlsH\x00R\x03tls\x12@\n" +
	"\x05other\x18\x02 \x01(\v2(.grpc.channelz.v1.Security.OtherSecurityH\x00R\x05other\x1a\xb9\x01\n" +
	"\x03Tls\x12%\n" +
	"\rstandard_name\x18\x01 \x01(\tH\x00R\fstandardName\x12\x1f\n" +
	"\n" +
	"other_name\x18\x02 \x01(\tH\x00R\totherName\x12+\n" +
	"\x11local_certificate\x18\x03 \x01(\fR\x10localCertificate\x12-\n" +
	"\x12remote_certificate\x18\x04 \x01(\fR\x11remoteCertificateB\x0e\n" +
	"\fcipher_suite\x1aO\n" +
	"\rOtherSecurity\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12*\n" +
	"\x05value\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\x05valueB\a\n" +
	"\x05model\"n\n" +
	"\fSocketOption\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\x124\n" +
	"\n" +
	"additional\x18\x03 \x01(\v2\x14.google.protobuf.AnyR\n" +
	"additional\"L\n" +
	"\x13SocketOptionTimeout\x125\n" +
	"\bduration\x18\x01 \x01(\v2\x19.google.protobuf.DurationR\bduration\"c\n" +
	"\x12SocketOptionLinger\x12\x16\n" +
	"\x06active\x18\x01 \x01(\bR\x06active\x125\n" +
	"\bduration\x18\x02 \x01(\v2\x19.google.protobuf.DurationR\bduration\"\xb2\b\n" +
	"\x13SocketOptionTcpInfo\x12\x1d\n" +
	"\n" +
	"tcpi_state\x18\x01 \x01(\rR\ttcpiState\x12\"\n" +
	"\rtcpi_ca_state\x18\x02 \x01(\rR\vtcpiCaState\x12)\n" +
	"\x10tcpi_retransmits\x18\x03 \x01(\rR\x0ftcpiRetransmits\x12\x1f\n" +
	"\vtcpi_probes\x18\x04 \x01(\rR\n" +
	"tcpiProbes\x12!\n" +
	"\ftcpi_backoff\x18\x05 \x01(\rR\vtcpiBackoff\x12!\n" +
	"\ftcpi_options\x18\x06 \x01(\rR\vtcpiOptions\x12&\n" +
	"\x0ftcpi_snd_wscale\x18\a \x01(\rR\rtcpiSndWscale\x12&\n" +
	"\x0ftcpi_rcv_wscale\x18\b \x01(\rR\rtcpiRcvWscale\x12\x19\n" +
	"\btcpi_rto\x18\t \x01(\rR\atcpiRto\x12\x19\n" +
	"\btcpi_ato\x18\n" +
	" \x01(\rR\atcpiAto\x12 \n" +
	"\ftcpi_snd_mss\x18\v \x01(\rR\n" +
	"tcpiSndMss\x12 \n" +
	"\ftcpi_rcv_mss\x18\f \x01(\rR\n" +
	"tcpiRcvMss\x12!\n" +
	"\ftcpi_unacked\x18\r \x01(\rR\vtcpiUnacked\x12\x1f\n" +
	"\vtcpi_sacked\x18\x0e \x01(\rR\n" +
	"tcpiSacked\x12\x1b\n" +
	"\ttcpi_lost\x18\x0f \x01(\rR\btcpiLost\x12!\n" +
	"\ftcpi_retrans\x18\x10 \x01(\rR\vtcpiRetrans\x12!\n" +
	"\ftcpi_fackets\x18\x11 \x01(\rR\vtcpiFackets\x12-\n" +
	"\x13tcpi_last_data_sent\x18\x12 \x01(\rR\x10tcpiLastDataSent\x12+\n" +
	"\x12tcpi_last_ack_sent\x18\x13 \x01(\rR\x0ftcpiLastAckSent\x12-\n" +
	"\x13tcpi_last_data_recv\x18\x14 \x01(\rR\x10tcpiLastDataRecv\x12+\n" +
	"\x12tcpi_last_ack_recv\x18\x15 \x01(\rR\x0ftcpiLastAckRecv\x12\x1b\n" +
	"\ttcpi_pmtu\x18\x16 \x01(\rR\btcpiPmtu\x12*\n" +
	"\x11tcpi_rcv_ssthresh\x18\x17 \x01(\rR\x0ftcpiRcvSsthresh\x12\x19\n" +
	"\btcpi_rtt\x18\x18 \x01(\rR\atcpiRtt\x12\x1f\n" +
	"\vtcpi_rttvar\x18\x19 \x01(\rR\n" +
	"tcpiRttvar\x12*\n" +
	"\x11tcpi_snd_ssthresh\x18\x1a \x01(\rR\x0ftcpiSndSsthresh\x12\"\n" +
	"\rtcpi_snd_cwnd\x18\x1b \x01(\rR\vtcpiSndCwnd\x12\x1f\n" +
	"\vtcpi_advmss\x18\x1c \x01(\rR\n" +
	"tcpiAdvmss\x12'\n" +
	"\x0ftcpi_reordering\x18\x1d \x01(\rR\x0etcpiReordering\"b\n" +
	"\x15GetTopChannelsRequest\x12(\n" +
	"\x10start_channel_id\x18\x01 \x01(\x03R\x0estartChannelId\x12\x1f\n" +
	"\vmax_results\x18\x02 \x01(\x03R\n" +
	"maxResults\"_\n" +
	"\x16GetTopChannelsResponse\x123\n" +
	"\achannel\x18\x01 \x03(\v2\x19.grpc.channelz.v1.ChannelR\achannel\x12\x10\n" +
	"\x03end\x18\x02 \x01(\bR\x03end\"\\\n" +
	"\x11GetServersRequest\x12&\n" +
	"\x0fstart_server_id\x18\x01 \x01(\x03R\rstartServerId\x12\x1f\n" +
	"\vmax_results\x18\x02 \x01(\x03R\n" +
	"maxResults\"X\n" +
	"\x12GetServersResponse\x120\n" +
	"\x06server\x18\x01 \x03(\v2\x18.grpc.channelz.v1.ServerR\x06server\x12\x10\n" +
	"\x03end\x18\x02 \x01(\bR\x03end\"/\n" +
	"\x10GetServerRequest\x12\x1b\n" +
	"\tserver_id\x18\x01 \x01(\x03R\bserverId\"E\n" +
	"\x11GetServerResponse\x120\n" +
	"\x06server\x18\x01 \x01(\v2\x18.grpc.channelz.v1.ServerR\x06server\"\x7f\n" +
	"\x17GetServerSocketsRequest\x12\x1b\n" +
	"\tserver_id\x18\x01 \x01(\x03R\bserverId\x12&\n" +
	"\x0fstart_socket_id\x18\x02 \x01(\x03R\rstartSocketId\x12\x1f\n" +
	"\vmax_results\x18\x03 \x01(\x03R\n" +
	"maxResults\"h\n" +
	"\x18GetServerSocketsResponse\x12:\n" +
	"\n" +
	"socket_ref\x18\x01 \x03(\v2\x1b.grpc.channelz.v1.SocketRefR\tsocketRef\x12\x10\n" +
	"\x03end\x18\x02 \x01(\bR\x03end\"2\n" +
	"\x11GetChannelRequest\x12\x1d\n" +
	"\n" +
	"channel_id\x18\x01 \x01(\x03R\tchannelId\"I\n" +
	"\x12GetChannelResponse\x123\n" +
	"\achannel\x18\x01 \x01(\v2\x19.grpc.channelz.v1.ChannelR\achannel\";\n" +
	"\x14GetSubchannelRequest\x12#\n" +
	"\rsubchannel_id\x18\x01 \x01(\x03R\fsubchannelId\"U\n" +
	"\x15GetSubchannelResponse\x12<\n" +
	"\n" +
	"subchannel\x18\x01 \x01(\v2\x1c.grpc.channelz.v1.SubchannelR\n" +
	"subchannel\"I\n" +
	"\x10GetSocketRequest\x12\x1b\n" +
	"\tsocket_id\x18\x01 \x01(\x03R\bsocketId\x12\x18\n" +
	"\asummary\x18\x02 \x01(\bR\asummary\"E\n" +
	"\x11GetSocketResponse\x120\n" +
	"\x06socket\x18\x01 \x01(\v2\x18.grpc.channelz.v1.SocketR\x06socket2\x9a\x05\n" +
	"\bChannelz\x12c\n" +
	"\x0eGetTopChannels\x12'.grpc.channelz.v1.GetTopChannelsRequest\x1a(.grpc.channelz.v1.GetTopChannelsResponse\x12W\n" +
	"\n" +
	"GetServers\x12#.grpc.channelz.v1.GetServersRequest\x1a$.grpc.channelz.v1.GetServersResponse\x12T\n" +
	"\tGetServer\x12\".grpc.channelz.v1.GetServerRequest\x1a#.grpc.channelz.v1.GetServerResponse\x12i\n" +
	"\x10GetServerSockets\x12).grpc.channelz.v1.GetServerSocketsRequest\x1a*.grpc.channelz.v1.GetServerSocketsResponse\x12W\n" +
	"\n" +
	"GetChannel\x12#.grpc.channelz.v1.GetChannelRequest\x1a$.grpc.channelz.v1.GetChannelResponse\x12`\n" +
	"\rGetSubchannel\x12&.grpc.channelz.v1.GetSubchannelRequest\x1a'.grpc.channelz.v1.GetSubchannelResponse\x12T\n" +
	"\tGetSocket\x12\".grpc.channelz.v1.GetSocketRequest\x1a#.grpc.channelz.v1.GetSocketResponseBX\n" +
	"\x13io.grpc.channelz.v1B\rChannelzProtoP\x01Z0google.golang.org/grpc/channelz/grpc_channelz_v1b\x06proto3"

var (
	file_grpc_channelz_v1_channelz_proto_rawDescOnce sync.Once
	file_grpc_channelz_v1_channelz_proto_rawDescData []byte
)

func file_grpc_channelz_v1_channelz_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_grpc_channelz_v1_channelz_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_grpc_channelz_v1_channelz_proto_msgTypes = make([]protoimpl.MessageInfo, 39)
var file_grpc_channelz_v1_channelz_proto_goTypes = []any{
	(ChannelConnectivityState_State)(0),
	(ChannelTraceEvent_Severity)(0),
	(*Channel)(nil),
	(*Subchannel)(nil),
	(*ChannelConnectivityState)(nil),
	(*ChannelData)(nil),
	(*ChannelTraceEvent)(nil),
	(*ChannelTrace)(nil),
	(*ChannelRef)(nil),
	(*SubchannelRef)(nil),
	(*SocketRef)(nil),
	(*ServerRef)(nil),
	(*Server)(nil),
	(*ServerData)(nil),
	(*Socket)(nil),
	(*SocketData)(nil),
	(*Address)(nil),
	(*Security)(nil),
	(*SocketOption)(nil),
	(*SocketOptionTimeout)(nil),
	(*SocketOptionLinger)(nil),
	(*SocketOptionTcpInfo)(nil),
	(*GetTopChannelsRequest)(nil),
	(*GetTopChannelsResponse)(nil),
	(*GetServersRequest)(nil),
	(*GetServersResponse)(nil),
	(*GetServerRequest)(nil),
	(*GetServerResponse)(nil),
	(*GetServerSocketsRequest)(nil),
	(*GetServerSocketsResponse)(nil),
	(*GetChannelRequest)(nil),
	(*GetChannelResponse)(nil),
	(*GetSubchannelRequest)(nil),
	(*GetSubchannelResponse)(nil),
	(*GetSocketRequest)(nil),
	(*GetSocketResponse)(nil),
	(*Address_TcpIpAddress)(nil),
	(*Address_UdsAddress)(nil),
	(*Address_OtherAddress)(nil),
	(*Security_Tls)(nil),
	(*Security_OtherSecurity)(nil),
	(*timestamppb.Timestamp)(nil),
	(*wrapperspb.Int64Value)(nil),
	(*wrapperspb.UInt32Value)(nil),
	(*anypb.Any)(nil),
	(*durationpb.Duration)(nil),
}
var file_grpc_channelz_v1_channelz_proto_depIdxs = []int32{
	8,
	5,
	8,
	9,
	10,
	9,
	5,
	8,
	9,
	10,
	0,
	4,
	7,
	41,
	1,
	41,
	8,
	9,
	41,
	6,
	11,
	13,
	10,
	7,
	41,
	10,
	15,
	16,
	16,
	17,
	41,
	41,
	41,
	41,
	42,
	42,
	18,
	43,
	36,
	37,
	38,
	39,
	40,
	44,
	45,
	45,
	2,
	12,
	12,
	10,
	2,
	3,
	14,
	44,
	44,
	22,
	24,
	26,
	28,
	30,
	32,
	34,
	23,
	25,
	27,
	29,
	31,
	33,
	35,
	62,
	55,
	55,
	55,
	0,
}

func init()                                      { file_grpc_channelz_v1_channelz_proto_init() }
func file_grpc_channelz_v1_channelz_proto_init() { _ = "STUB: not implemented"; return }
