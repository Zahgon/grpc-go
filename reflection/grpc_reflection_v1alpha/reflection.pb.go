package grpc_reflection_v1alpha

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServerReflectionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`

	MessageRequest isServerReflectionRequest_MessageRequest `protobuf_oneof:"message_request"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ServerReflectionRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *ServerReflectionRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*ServerReflectionRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServerReflectionRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServerReflectionRequest) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *ServerReflectionRequest) GetHost() string { _ = "STUB: not implemented"; return "" }

func (x *ServerReflectionRequest) GetMessageRequest() isServerReflectionRequest_MessageRequest {
	_ = "STUB: not implemented"
	return *new(isServerReflectionRequest_MessageRequest)
}

func (x *ServerReflectionRequest) GetFileByFilename() string { _ = "STUB: not implemented"; return "" }

func (x *ServerReflectionRequest) GetFileContainingSymbol() string {
	_ = "STUB: not implemented"
	return ""
}

func (x *ServerReflectionRequest) GetFileContainingExtension() *ExtensionRequest {
	_ = "STUB: not implemented"
	return nil
}

func (x *ServerReflectionRequest) GetAllExtensionNumbersOfType() string {
	_ = "STUB: not implemented"
	return ""
}

func (x *ServerReflectionRequest) GetListServices() string { _ = "STUB: not implemented"; return "" }

type isServerReflectionRequest_MessageRequest interface {
	isServerReflectionRequest_MessageRequest()
}

type ServerReflectionRequest_FileByFilename struct {
	FileByFilename string `protobuf:"bytes,3,opt,name=file_by_filename,json=fileByFilename,proto3,oneof"`
}

type ServerReflectionRequest_FileContainingSymbol struct {
	FileContainingSymbol string `protobuf:"bytes,4,opt,name=file_containing_symbol,json=fileContainingSymbol,proto3,oneof"`
}

type ServerReflectionRequest_FileContainingExtension struct {
	FileContainingExtension *ExtensionRequest `protobuf:"bytes,5,opt,name=file_containing_extension,json=fileContainingExtension,proto3,oneof"`
}

type ServerReflectionRequest_AllExtensionNumbersOfType struct {
	AllExtensionNumbersOfType string `protobuf:"bytes,6,opt,name=all_extension_numbers_of_type,json=allExtensionNumbersOfType,proto3,oneof"`
}

type ServerReflectionRequest_ListServices struct {
	ListServices string `protobuf:"bytes,7,opt,name=list_services,json=listServices,proto3,oneof"`
}

func (*ServerReflectionRequest_FileByFilename) isServerReflectionRequest_MessageRequest() {
	_ = "STUB: not implemented"
	return
}

func (*ServerReflectionRequest_FileContainingSymbol) isServerReflectionRequest_MessageRequest() {
	_ = "STUB: not implemented"
	return
}

func (*ServerReflectionRequest_FileContainingExtension) isServerReflectionRequest_MessageRequest() {
	_ = "STUB: not implemented"
	return
}

func (*ServerReflectionRequest_AllExtensionNumbersOfType) isServerReflectionRequest_MessageRequest() {
	_ = "STUB: not implemented"
	return
}

func (*ServerReflectionRequest_ListServices) isServerReflectionRequest_MessageRequest() {
	_ = "STUB: not implemented"
	return
}

type ExtensionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ContainingType string `protobuf:"bytes,1,opt,name=containing_type,json=containingType,proto3" json:"containing_type,omitempty"`

	ExtensionNumber int32 `protobuf:"varint,2,opt,name=extension_number,json=extensionNumber,proto3" json:"extension_number,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ExtensionRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *ExtensionRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*ExtensionRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ExtensionRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ExtensionRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ExtensionRequest) GetContainingType() string { _ = "STUB: not implemented"; return "" }

func (x *ExtensionRequest) GetExtensionNumber() int32 { _ = "STUB: not implemented"; return 0 }

type ServerReflectionResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ValidHost string `protobuf:"bytes,1,opt,name=valid_host,json=validHost,proto3" json:"valid_host,omitempty"`

	OriginalRequest *ServerReflectionRequest `protobuf:"bytes,2,opt,name=original_request,json=originalRequest,proto3" json:"original_request,omitempty"`

	MessageResponse isServerReflectionResponse_MessageResponse `protobuf_oneof:"message_response"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ServerReflectionResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *ServerReflectionResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*ServerReflectionResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServerReflectionResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServerReflectionResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *ServerReflectionResponse) GetValidHost() string { _ = "STUB: not implemented"; return "" }

func (x *ServerReflectionResponse) GetOriginalRequest() *ServerReflectionRequest {
	_ = "STUB: not implemented"
	return nil
}

func (x *ServerReflectionResponse) GetMessageResponse() isServerReflectionResponse_MessageResponse {
	_ = "STUB: not implemented"
	return *new(isServerReflectionResponse_MessageResponse)
}

func (x *ServerReflectionResponse) GetFileDescriptorResponse() *FileDescriptorResponse {
	_ = "STUB: not implemented"
	return nil
}

func (x *ServerReflectionResponse) GetAllExtensionNumbersResponse() *ExtensionNumberResponse {
	_ = "STUB: not implemented"
	return nil
}

func (x *ServerReflectionResponse) GetListServicesResponse() *ListServiceResponse {
	_ = "STUB: not implemented"
	return nil
}

func (x *ServerReflectionResponse) GetErrorResponse() *ErrorResponse {
	_ = "STUB: not implemented"
	return nil
}

type isServerReflectionResponse_MessageResponse interface {
	isServerReflectionResponse_MessageResponse()
}

type ServerReflectionResponse_FileDescriptorResponse struct {
	FileDescriptorResponse *FileDescriptorResponse `protobuf:"bytes,4,opt,name=file_descriptor_response,json=fileDescriptorResponse,proto3,oneof"`
}

type ServerReflectionResponse_AllExtensionNumbersResponse struct {
	AllExtensionNumbersResponse *ExtensionNumberResponse `protobuf:"bytes,5,opt,name=all_extension_numbers_response,json=allExtensionNumbersResponse,proto3,oneof"`
}

type ServerReflectionResponse_ListServicesResponse struct {
	ListServicesResponse *ListServiceResponse `protobuf:"bytes,6,opt,name=list_services_response,json=listServicesResponse,proto3,oneof"`
}

type ServerReflectionResponse_ErrorResponse struct {
	ErrorResponse *ErrorResponse `protobuf:"bytes,7,opt,name=error_response,json=errorResponse,proto3,oneof"`
}

func (*ServerReflectionResponse_FileDescriptorResponse) isServerReflectionResponse_MessageResponse() {
	_ = "STUB: not implemented"
	return
}

func (*ServerReflectionResponse_AllExtensionNumbersResponse) isServerReflectionResponse_MessageResponse() {
	_ = "STUB: not implemented"
	return
}

func (*ServerReflectionResponse_ListServicesResponse) isServerReflectionResponse_MessageResponse() {
	_ = "STUB: not implemented"
	return
}

func (*ServerReflectionResponse_ErrorResponse) isServerReflectionResponse_MessageResponse() {
	_ = "STUB: not implemented"
	return
}

type FileDescriptorResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	FileDescriptorProto [][]byte `protobuf:"bytes,1,rep,name=file_descriptor_proto,json=fileDescriptorProto,proto3" json:"file_descriptor_proto,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *FileDescriptorResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *FileDescriptorResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*FileDescriptorResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *FileDescriptorResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*FileDescriptorResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *FileDescriptorResponse) GetFileDescriptorProto() [][]byte {
	_ = "STUB: not implemented"
	return nil
}

type ExtensionNumberResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	BaseTypeName string `protobuf:"bytes,1,opt,name=base_type_name,json=baseTypeName,proto3" json:"base_type_name,omitempty"`

	ExtensionNumber []int32 `protobuf:"varint,2,rep,packed,name=extension_number,json=extensionNumber,proto3" json:"extension_number,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ExtensionNumberResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *ExtensionNumberResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*ExtensionNumberResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ExtensionNumberResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ExtensionNumberResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *ExtensionNumberResponse) GetBaseTypeName() string { _ = "STUB: not implemented"; return "" }

func (x *ExtensionNumberResponse) GetExtensionNumber() []int32 {
	_ = "STUB: not implemented"
	return nil
}

type ListServiceResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Service       []*ServiceResponse `protobuf:"bytes,1,rep,name=service,proto3" json:"service,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListServiceResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *ListServiceResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*ListServiceResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ListServiceResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ListServiceResponse) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *ListServiceResponse) GetService() []*ServiceResponse {
	_ = "STUB: not implemented"
	return nil
}

type ServiceResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *ServiceResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*ServiceResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServiceResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ServiceResponse) GetName() string { _ = "STUB: not implemented"; return "" }

type ErrorResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ErrorCode int32 `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`

	ErrorMessage  string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrorResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *ErrorResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*ErrorResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ErrorResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ErrorResponse) GetErrorCode() int32 { _ = "STUB: not implemented"; return 0 }

func (x *ErrorResponse) GetErrorMessage() string { _ = "STUB: not implemented"; return "" }

var File_grpc_reflection_v1alpha_reflection_proto protoreflect.FileDescriptor

const file_grpc_reflection_v1alpha_reflection_proto_rawDesc = "" +
	"\n" +
	"(grpc/reflection/v1alpha/reflection.proto\x12\x17grpc.reflection.v1alpha\"\xf8\x02\n" +
	"\x17ServerReflectionRequest\x12\x12\n" +
	"\x04host\x18\x01 \x01(\tR\x04host\x12*\n" +
	"\x10file_by_filename\x18\x03 \x01(\tH\x00R\x0efileByFilename\x126\n" +
	"\x16file_containing_symbol\x18\x04 \x01(\tH\x00R\x14fileContainingSymbol\x12g\n" +
	"\x19file_containing_extension\x18\x05 \x01(\v2).grpc.reflection.v1alpha.ExtensionRequestH\x00R\x17fileContainingExtension\x12B\n" +
	"\x1dall_extension_numbers_of_type\x18\x06 \x01(\tH\x00R\x19allExtensionNumbersOfType\x12%\n" +
	"\rlist_services\x18\a \x01(\tH\x00R\flistServicesB\x11\n" +
	"\x0fmessage_request\"f\n" +
	"\x10ExtensionRequest\x12'\n" +
	"\x0fcontaining_type\x18\x01 \x01(\tR\x0econtainingType\x12)\n" +
	"\x10extension_number\x18\x02 \x01(\x05R\x0fextensionNumber\"\xc7\x04\n" +
	"\x18ServerReflectionResponse\x12\x1d\n" +
	"\n" +
	"valid_host\x18\x01 \x01(\tR\tvalidHost\x12[\n" +
	"\x10original_request\x18\x02 \x01(\v20.grpc.reflection.v1alpha.ServerReflectionRequestR\x0foriginalRequest\x12k\n" +
	"\x18file_descriptor_response\x18\x04 \x01(\v2/.grpc.reflection.v1alpha.FileDescriptorResponseH\x00R\x16fileDescriptorResponse\x12w\n" +
	"\x1eall_extension_numbers_response\x18\x05 \x01(\v20.grpc.reflection.v1alpha.ExtensionNumberResponseH\x00R\x1ballExtensionNumbersResponse\x12d\n" +
	"\x16list_services_response\x18\x06 \x01(\v2,.grpc.reflection.v1alpha.ListServiceResponseH\x00R\x14listServicesResponse\x12O\n" +
	"\x0eerror_response\x18\a \x01(\v2&.grpc.reflection.v1alpha.ErrorResponseH\x00R\rerrorResponseB\x12\n" +
	"\x10message_response\"L\n" +
	"\x16FileDescriptorResponse\x122\n" +
	"\x15file_descriptor_proto\x18\x01 \x03(\fR\x13fileDescriptorProto\"j\n" +
	"\x17ExtensionNumberResponse\x12$\n" +
	"\x0ebase_type_name\x18\x01 \x01(\tR\fbaseTypeName\x12)\n" +
	"\x10extension_number\x18\x02 \x03(\x05R\x0fextensionNumber\"Y\n" +
	"\x13ListServiceResponse\x12B\n" +
	"\aservice\x18\x01 \x03(\v2(.grpc.reflection.v1alpha.ServiceResponseR\aservice\"%\n" +
	"\x0fServiceResponse\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"S\n" +
	"\rErrorResponse\x12\x1d\n" +
	"\n" +
	"error_code\x18\x01 \x01(\x05R\terrorCode\x12#\n" +
	"\rerror_message\x18\x02 \x01(\tR\ferrorMessage2\x93\x01\n" +
	"\x10ServerReflection\x12\x7f\n" +
	"\x14ServerReflectionInfo\x120.grpc.reflection.v1alpha.ServerReflectionRequest\x1a1.grpc.reflection.v1alpha.ServerReflectionResponse(\x010\x01Bs\n" +
	"\x1aio.grpc.reflection.v1alphaB\x15ServerReflectionProtoP\x01Z9google.golang.org/grpc/reflection/grpc_reflection_v1alpha\xb8\x01\x01b\x06proto3"

var (
	file_grpc_reflection_v1alpha_reflection_proto_rawDescOnce sync.Once
	file_grpc_reflection_v1alpha_reflection_proto_rawDescData []byte
)

func file_grpc_reflection_v1alpha_reflection_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_grpc_reflection_v1alpha_reflection_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_reflection_v1alpha_reflection_proto_goTypes = []any{
	(*ServerReflectionRequest)(nil),
	(*ExtensionRequest)(nil),
	(*ServerReflectionResponse)(nil),
	(*FileDescriptorResponse)(nil),
	(*ExtensionNumberResponse)(nil),
	(*ListServiceResponse)(nil),
	(*ServiceResponse)(nil),
	(*ErrorResponse)(nil),
}
var file_grpc_reflection_v1alpha_reflection_proto_depIdxs = []int32{
	1,
	0,
	3,
	4,
	5,
	7,
	6,
	0,
	2,
	8,
	7,
	7,
	7,
	0,
}

func init()                                               { file_grpc_reflection_v1alpha_reflection_proto_init() }
func file_grpc_reflection_v1alpha_reflection_proto_init() { _ = "STUB: not implemented"; return }
