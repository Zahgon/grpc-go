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

type SearchResponse struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Results       []*SearchResponse_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *SearchResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*SearchResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SearchResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SearchResponse) GetResults() []*SearchResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *SearchRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*SearchRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SearchRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SearchRequest) GetQuery() string { _ = "STUB: not implemented"; return "" }

type SearchResponse_Result struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Snippets      []string               `protobuf:"bytes,3,rep,name=snippets,proto3" json:"snippets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchResponse_Result) Reset() { _ = "STUB: not implemented"; return }

func (x *SearchResponse_Result) String() string { _ = "STUB: not implemented"; return "" }

func (*SearchResponse_Result) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SearchResponse_Result) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SearchResponse_Result) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *SearchResponse_Result) GetUrl() string { _ = "STUB: not implemented"; return "" }

func (x *SearchResponse_Result) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (x *SearchResponse_Result) GetSnippets() []string { _ = "STUB: not implemented"; return nil }

var File_reflection_grpc_testing_test_proto protoreflect.FileDescriptor

const file_reflection_grpc_testing_test_proto_rawDesc = "" +
	"\n" +
	"\"reflection/grpc_testing/test.proto\x12\fgrpc.testing\"\x9d\x01\n" +
	"\x0eSearchResponse\x12=\n" +
	"\aresults\x18\x01 \x03(\v2#.grpc.testing.SearchResponse.ResultR\aresults\x1aL\n" +
	"\x06Result\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x1a\n" +
	"\bsnippets\x18\x03 \x03(\tR\bsnippets\"%\n" +
	"\rSearchRequest\x12\x14\n" +
	"\x05query\x18\x01 \x01(\tR\x05query2\xa6\x01\n" +
	"\rSearchService\x12C\n" +
	"\x06Search\x12\x1b.grpc.testing.SearchRequest\x1a\x1c.grpc.testing.SearchResponse\x12P\n" +
	"\x0fStreamingSearch\x12\x1b.grpc.testing.SearchRequest\x1a\x1c.grpc.testing.SearchResponse(\x010\x01B0Z.google.golang.org/grpc/reflection/grpc_testingb\x06proto3"

var (
	file_reflection_grpc_testing_test_proto_rawDescOnce sync.Once
	file_reflection_grpc_testing_test_proto_rawDescData []byte
)

func file_reflection_grpc_testing_test_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_reflection_grpc_testing_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_reflection_grpc_testing_test_proto_goTypes = []any{
	(*SearchResponse)(nil),
	(*SearchRequest)(nil),
	(*SearchResponse_Result)(nil),
}
var file_reflection_grpc_testing_test_proto_depIdxs = []int32{
	2,
	1,
	1,
	0,
	0,
	3,
	1,
	1,
	1,
	0,
}

func init()                                         { file_reflection_grpc_testing_test_proto_init() }
func file_reflection_grpc_testing_test_proto_init() { _ = "STUB: not implemented"; return }
