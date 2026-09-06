package grpc_binarylog_v1

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

type GrpcLogEntry_EventType int32

const (
	GrpcLogEntry_EVENT_TYPE_UNKNOWN GrpcLogEntry_EventType = 0

	GrpcLogEntry_EVENT_TYPE_CLIENT_HEADER GrpcLogEntry_EventType = 1

	GrpcLogEntry_EVENT_TYPE_SERVER_HEADER GrpcLogEntry_EventType = 2

	GrpcLogEntry_EVENT_TYPE_CLIENT_MESSAGE GrpcLogEntry_EventType = 3

	GrpcLogEntry_EVENT_TYPE_SERVER_MESSAGE GrpcLogEntry_EventType = 4

	GrpcLogEntry_EVENT_TYPE_CLIENT_HALF_CLOSE GrpcLogEntry_EventType = 5

	GrpcLogEntry_EVENT_TYPE_SERVER_TRAILER GrpcLogEntry_EventType = 6

	GrpcLogEntry_EVENT_TYPE_CANCEL GrpcLogEntry_EventType = 7
)

var (
	GrpcLogEntry_EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNKNOWN",
		1: "EVENT_TYPE_CLIENT_HEADER",
		2: "EVENT_TYPE_SERVER_HEADER",
		3: "EVENT_TYPE_CLIENT_MESSAGE",
		4: "EVENT_TYPE_SERVER_MESSAGE",
		5: "EVENT_TYPE_CLIENT_HALF_CLOSE",
		6: "EVENT_TYPE_SERVER_TRAILER",
		7: "EVENT_TYPE_CANCEL",
	}
	GrpcLogEntry_EventType_value = map[string]int32{
		"EVENT_TYPE_UNKNOWN":           0,
		"EVENT_TYPE_CLIENT_HEADER":     1,
		"EVENT_TYPE_SERVER_HEADER":     2,
		"EVENT_TYPE_CLIENT_MESSAGE":    3,
		"EVENT_TYPE_SERVER_MESSAGE":    4,
		"EVENT_TYPE_CLIENT_HALF_CLOSE": 5,
		"EVENT_TYPE_SERVER_TRAILER":    6,
		"EVENT_TYPE_CANCEL":            7,
	}
)

func (x GrpcLogEntry_EventType) Enum() *GrpcLogEntry_EventType {
	_ = "STUB: not implemented"
	return nil
}

func (x GrpcLogEntry_EventType) String() string { _ = "STUB: not implemented"; return "" }

func (GrpcLogEntry_EventType) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (GrpcLogEntry_EventType) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x GrpcLogEntry_EventType) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (GrpcLogEntry_EventType) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GrpcLogEntry_Logger int32

const (
	GrpcLogEntry_LOGGER_UNKNOWN GrpcLogEntry_Logger = 0
	GrpcLogEntry_LOGGER_CLIENT  GrpcLogEntry_Logger = 1
	GrpcLogEntry_LOGGER_SERVER  GrpcLogEntry_Logger = 2
)

var (
	GrpcLogEntry_Logger_name = map[int32]string{
		0: "LOGGER_UNKNOWN",
		1: "LOGGER_CLIENT",
		2: "LOGGER_SERVER",
	}
	GrpcLogEntry_Logger_value = map[string]int32{
		"LOGGER_UNKNOWN": 0,
		"LOGGER_CLIENT":  1,
		"LOGGER_SERVER":  2,
	}
)

func (x GrpcLogEntry_Logger) Enum() *GrpcLogEntry_Logger { _ = "STUB: not implemented"; return nil }

func (x GrpcLogEntry_Logger) String() string { _ = "STUB: not implemented"; return "" }

func (GrpcLogEntry_Logger) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (GrpcLogEntry_Logger) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x GrpcLogEntry_Logger) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (GrpcLogEntry_Logger) EnumDescriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Address_Type int32

const (
	Address_TYPE_UNKNOWN Address_Type = 0

	Address_TYPE_IPV4 Address_Type = 1

	Address_TYPE_IPV6 Address_Type = 2

	Address_TYPE_UNIX Address_Type = 3
)

var (
	Address_Type_name = map[int32]string{
		0: "TYPE_UNKNOWN",
		1: "TYPE_IPV4",
		2: "TYPE_IPV6",
		3: "TYPE_UNIX",
	}
	Address_Type_value = map[string]int32{
		"TYPE_UNKNOWN": 0,
		"TYPE_IPV4":    1,
		"TYPE_IPV6":    2,
		"TYPE_UNIX":    3,
	}
)

func (x Address_Type) Enum() *Address_Type { _ = "STUB: not implemented"; return nil }

func (x Address_Type) String() string { _ = "STUB: not implemented"; return "" }

func (Address_Type) Descriptor() protoreflect.EnumDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumDescriptor)
}

func (Address_Type) Type() protoreflect.EnumType {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumType)
}

func (x Address_Type) Number() protoreflect.EnumNumber {
	_ = "STUB: not implemented"
	return *new(protoreflect.EnumNumber)
}

func (Address_Type) EnumDescriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type GrpcLogEntry struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`

	CallId uint64 `protobuf:"varint,2,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`

	SequenceIdWithinCall uint64                 `protobuf:"varint,3,opt,name=sequence_id_within_call,json=sequenceIdWithinCall,proto3" json:"sequence_id_within_call,omitempty"`
	Type                 GrpcLogEntry_EventType `protobuf:"varint,4,opt,name=type,proto3,enum=grpc.binarylog.v1.GrpcLogEntry_EventType" json:"type,omitempty"`
	Logger               GrpcLogEntry_Logger    `protobuf:"varint,5,opt,name=logger,proto3,enum=grpc.binarylog.v1.GrpcLogEntry_Logger" json:"logger,omitempty"`

	Payload isGrpcLogEntry_Payload `protobuf_oneof:"payload"`

	PayloadTruncated bool `protobuf:"varint,10,opt,name=payload_truncated,json=payloadTruncated,proto3" json:"payload_truncated,omitempty"`

	Peer          *Address `protobuf:"bytes,11,opt,name=peer,proto3" json:"peer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrpcLogEntry) Reset() { _ = "STUB: not implemented"; return }

func (x *GrpcLogEntry) String() string { _ = "STUB: not implemented"; return "" }

func (*GrpcLogEntry) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *GrpcLogEntry) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*GrpcLogEntry) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *GrpcLogEntry) GetTimestamp() *timestamppb.Timestamp { _ = "STUB: not implemented"; return nil }

func (x *GrpcLogEntry) GetCallId() uint64 { _ = "STUB: not implemented"; return 0 }

func (x *GrpcLogEntry) GetSequenceIdWithinCall() uint64 { _ = "STUB: not implemented"; return 0 }

func (x *GrpcLogEntry) GetType() GrpcLogEntry_EventType {
	_ = "STUB: not implemented"
	return *new(GrpcLogEntry_EventType)
}

func (x *GrpcLogEntry) GetLogger() GrpcLogEntry_Logger {
	_ = "STUB: not implemented"
	return *new(GrpcLogEntry_Logger)
}

func (x *GrpcLogEntry) GetPayload() isGrpcLogEntry_Payload {
	_ = "STUB: not implemented"
	return *new(isGrpcLogEntry_Payload)
}

func (x *GrpcLogEntry) GetClientHeader() *ClientHeader { _ = "STUB: not implemented"; return nil }

func (x *GrpcLogEntry) GetServerHeader() *ServerHeader { _ = "STUB: not implemented"; return nil }

func (x *GrpcLogEntry) GetMessage() *Message { _ = "STUB: not implemented"; return nil }

func (x *GrpcLogEntry) GetTrailer() *Trailer { _ = "STUB: not implemented"; return nil }

func (x *GrpcLogEntry) GetPayloadTruncated() bool { _ = "STUB: not implemented"; return false }

func (x *GrpcLogEntry) GetPeer() *Address { _ = "STUB: not implemented"; return nil }

type isGrpcLogEntry_Payload interface {
	isGrpcLogEntry_Payload()
}

type GrpcLogEntry_ClientHeader struct {
	ClientHeader *ClientHeader `protobuf:"bytes,6,opt,name=client_header,json=clientHeader,proto3,oneof"`
}

type GrpcLogEntry_ServerHeader struct {
	ServerHeader *ServerHeader `protobuf:"bytes,7,opt,name=server_header,json=serverHeader,proto3,oneof"`
}

type GrpcLogEntry_Message struct {
	Message *Message `protobuf:"bytes,8,opt,name=message,proto3,oneof"`
}

type GrpcLogEntry_Trailer struct {
	Trailer *Trailer `protobuf:"bytes,9,opt,name=trailer,proto3,oneof"`
}

func (*GrpcLogEntry_ClientHeader) isGrpcLogEntry_Payload() { _ = "STUB: not implemented"; return }

func (*GrpcLogEntry_ServerHeader) isGrpcLogEntry_Payload() { _ = "STUB: not implemented"; return }

func (*GrpcLogEntry_Message) isGrpcLogEntry_Payload() { _ = "STUB: not implemented"; return }

func (*GrpcLogEntry_Trailer) isGrpcLogEntry_Payload() { _ = "STUB: not implemented"; return }

type ClientHeader struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`

	MethodName string `protobuf:"bytes,2,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`

	Authority string `protobuf:"bytes,3,opt,name=authority,proto3" json:"authority,omitempty"`

	Timeout       *durationpb.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientHeader) Reset() { _ = "STUB: not implemented"; return }

func (x *ClientHeader) String() string { _ = "STUB: not implemented"; return "" }

func (*ClientHeader) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClientHeader) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClientHeader) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ClientHeader) GetMetadata() *Metadata { _ = "STUB: not implemented"; return nil }

func (x *ClientHeader) GetMethodName() string { _ = "STUB: not implemented"; return "" }

func (x *ClientHeader) GetAuthority() string { _ = "STUB: not implemented"; return "" }

func (x *ClientHeader) GetTimeout() *durationpb.Duration { _ = "STUB: not implemented"; return nil }

type ServerHeader struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Metadata      *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerHeader) Reset() { _ = "STUB: not implemented"; return }

func (x *ServerHeader) String() string { _ = "STUB: not implemented"; return "" }

func (*ServerHeader) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServerHeader) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServerHeader) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ServerHeader) GetMetadata() *Metadata { _ = "STUB: not implemented"; return nil }

type Trailer struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`

	StatusCode uint32 `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`

	StatusMessage string `protobuf:"bytes,3,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`

	StatusDetails []byte `protobuf:"bytes,4,opt,name=status_details,json=statusDetails,proto3" json:"status_details,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Trailer) Reset() { _ = "STUB: not implemented"; return }

func (x *Trailer) String() string { _ = "STUB: not implemented"; return "" }

func (*Trailer) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Trailer) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Trailer) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Trailer) GetMetadata() *Metadata { _ = "STUB: not implemented"; return nil }

func (x *Trailer) GetStatusCode() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *Trailer) GetStatusMessage() string { _ = "STUB: not implemented"; return "" }

func (x *Trailer) GetStatusDetails() []byte { _ = "STUB: not implemented"; return nil }

type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Length uint32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`

	Data          []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() { _ = "STUB: not implemented"; return }

func (x *Message) String() string { _ = "STUB: not implemented"; return "" }

func (*Message) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Message) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Message) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Message) GetLength() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *Message) GetData() []byte { _ = "STUB: not implemented"; return nil }

type Metadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Entry         []*MetadataEntry       `protobuf:"bytes,1,rep,name=entry,proto3" json:"entry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Metadata) Reset() { _ = "STUB: not implemented"; return }

func (x *Metadata) String() string { _ = "STUB: not implemented"; return "" }

func (*Metadata) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Metadata) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Metadata) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Metadata) GetEntry() []*MetadataEntry { _ = "STUB: not implemented"; return nil }

type MetadataEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetadataEntry) Reset() { _ = "STUB: not implemented"; return }

func (x *MetadataEntry) String() string { _ = "STUB: not implemented"; return "" }

func (*MetadataEntry) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *MetadataEntry) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*MetadataEntry) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *MetadataEntry) GetKey() string { _ = "STUB: not implemented"; return "" }

func (x *MetadataEntry) GetValue() []byte { _ = "STUB: not implemented"; return nil }

type Address struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Type    Address_Type           `protobuf:"varint,1,opt,name=type,proto3,enum=grpc.binarylog.v1.Address_Type" json:"type,omitempty"`
	Address string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`

	IpPort        uint32 `protobuf:"varint,3,opt,name=ip_port,json=ipPort,proto3" json:"ip_port,omitempty"`
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

func (x *Address) GetType() Address_Type { _ = "STUB: not implemented"; return *new(Address_Type) }

func (x *Address) GetAddress() string { _ = "STUB: not implemented"; return "" }

func (x *Address) GetIpPort() uint32 { _ = "STUB: not implemented"; return 0 }

var File_grpc_binlog_v1_binarylog_proto protoreflect.FileDescriptor

const file_grpc_binlog_v1_binarylog_proto_rawDesc = "" +
	"\n" +
	"\x1egrpc/binlog/v1/binarylog.proto\x12\x11grpc.binarylog.v1\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xbb\a\n" +
	"\fGrpcLogEntry\x128\n" +
	"\ttimestamp\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12\x17\n" +
	"\acall_id\x18\x02 \x01(\x04R\x06callId\x125\n" +
	"\x17sequence_id_within_call\x18\x03 \x01(\x04R\x14sequenceIdWithinCall\x12=\n" +
	"\x04type\x18\x04 \x01(\x0e2).grpc.binarylog.v1.GrpcLogEntry.EventTypeR\x04type\x12>\n" +
	"\x06logger\x18\x05 \x01(\x0e2&.grpc.binarylog.v1.GrpcLogEntry.LoggerR\x06logger\x12F\n" +
	"\rclient_header\x18\x06 \x01(\v2\x1f.grpc.binarylog.v1.ClientHeaderH\x00R\fclientHeader\x12F\n" +
	"\rserver_header\x18\a \x01(\v2\x1f.grpc.binarylog.v1.ServerHeaderH\x00R\fserverHeader\x126\n" +
	"\amessage\x18\b \x01(\v2\x1a.grpc.binarylog.v1.MessageH\x00R\amessage\x126\n" +
	"\atrailer\x18\t \x01(\v2\x1a.grpc.binarylog.v1.TrailerH\x00R\atrailer\x12+\n" +
	"\x11payload_truncated\x18\n" +
	" \x01(\bR\x10payloadTruncated\x12.\n" +
	"\x04peer\x18\v \x01(\v2\x1a.grpc.binarylog.v1.AddressR\x04peer\"\xf5\x01\n" +
	"\tEventType\x12\x16\n" +
	"\x12EVENT_TYPE_UNKNOWN\x10\x00\x12\x1c\n" +
	"\x18EVENT_TYPE_CLIENT_HEADER\x10\x01\x12\x1c\n" +
	"\x18EVENT_TYPE_SERVER_HEADER\x10\x02\x12\x1d\n" +
	"\x19EVENT_TYPE_CLIENT_MESSAGE\x10\x03\x12\x1d\n" +
	"\x19EVENT_TYPE_SERVER_MESSAGE\x10\x04\x12 \n" +
	"\x1cEVENT_TYPE_CLIENT_HALF_CLOSE\x10\x05\x12\x1d\n" +
	"\x19EVENT_TYPE_SERVER_TRAILER\x10\x06\x12\x15\n" +
	"\x11EVENT_TYPE_CANCEL\x10\a\"B\n" +
	"\x06Logger\x12\x12\n" +
	"\x0eLOGGER_UNKNOWN\x10\x00\x12\x11\n" +
	"\rLOGGER_CLIENT\x10\x01\x12\x11\n" +
	"\rLOGGER_SERVER\x10\x02B\t\n" +
	"\apayload\"\xbb\x01\n" +
	"\fClientHeader\x127\n" +
	"\bmetadata\x18\x01 \x01(\v2\x1b.grpc.binarylog.v1.MetadataR\bmetadata\x12\x1f\n" +
	"\vmethod_name\x18\x02 \x01(\tR\n" +
	"methodName\x12\x1c\n" +
	"\tauthority\x18\x03 \x01(\tR\tauthority\x123\n" +
	"\atimeout\x18\x04 \x01(\v2\x19.google.protobuf.DurationR\atimeout\"G\n" +
	"\fServerHeader\x127\n" +
	"\bmetadata\x18\x01 \x01(\v2\x1b.grpc.binarylog.v1.MetadataR\bmetadata\"\xb1\x01\n" +
	"\aTrailer\x127\n" +
	"\bmetadata\x18\x01 \x01(\v2\x1b.grpc.binarylog.v1.MetadataR\bmetadata\x12\x1f\n" +
	"\vstatus_code\x18\x02 \x01(\rR\n" +
	"statusCode\x12%\n" +
	"\x0estatus_message\x18\x03 \x01(\tR\rstatusMessage\x12%\n" +
	"\x0estatus_details\x18\x04 \x01(\fR\rstatusDetails\"5\n" +
	"\aMessage\x12\x16\n" +
	"\x06length\x18\x01 \x01(\rR\x06length\x12\x12\n" +
	"\x04data\x18\x02 \x01(\fR\x04data\"B\n" +
	"\bMetadata\x126\n" +
	"\x05entry\x18\x01 \x03(\v2 .grpc.binarylog.v1.MetadataEntryR\x05entry\"7\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\fR\x05value\"\xb8\x01\n" +
	"\aAddress\x123\n" +
	"\x04type\x18\x01 \x01(\x0e2\x1f.grpc.binarylog.v1.Address.TypeR\x04type\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\tR\aaddress\x12\x17\n" +
	"\aip_port\x18\x03 \x01(\rR\x06ipPort\"E\n" +
	"\x04Type\x12\x10\n" +
	"\fTYPE_UNKNOWN\x10\x00\x12\r\n" +
	"\tTYPE_IPV4\x10\x01\x12\r\n" +
	"\tTYPE_IPV6\x10\x02\x12\r\n" +
	"\tTYPE_UNIX\x10\x03B\\\n" +
	"\x14io.grpc.binarylog.v1B\x0eBinaryLogProtoP\x01Z2google.golang.org/grpc/binarylog/grpc_binarylog_v1b\x06proto3"

var (
	file_grpc_binlog_v1_binarylog_proto_rawDescOnce sync.Once
	file_grpc_binlog_v1_binarylog_proto_rawDescData []byte
)

func file_grpc_binlog_v1_binarylog_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_grpc_binlog_v1_binarylog_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_grpc_binlog_v1_binarylog_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_binlog_v1_binarylog_proto_goTypes = []any{
	(GrpcLogEntry_EventType)(0),
	(GrpcLogEntry_Logger)(0),
	(Address_Type)(0),
	(*GrpcLogEntry)(nil),
	(*ClientHeader)(nil),
	(*ServerHeader)(nil),
	(*Trailer)(nil),
	(*Message)(nil),
	(*Metadata)(nil),
	(*MetadataEntry)(nil),
	(*Address)(nil),
	(*timestamppb.Timestamp)(nil),
	(*durationpb.Duration)(nil),
}
var file_grpc_binlog_v1_binarylog_proto_depIdxs = []int32{
	11,
	0,
	1,
	4,
	5,
	7,
	6,
	10,
	8,
	12,
	8,
	8,
	9,
	2,
	14,
	14,
	14,
	14,
	0,
}

func init()                                     { file_grpc_binlog_v1_binarylog_proto_init() }
func file_grpc_binlog_v1_binarylog_proto_init() { _ = "STUB: not implemented"; return }
