package grpc_testing

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	SearchService_Search_FullMethodName          = "/grpc.testing.SearchService/Search"
	SearchService_StreamingSearch_FullMethodName = "/grpc.testing.SearchService/StreamingSearch"
)

type SearchServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	StreamingSearch(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SearchRequest, SearchResponse], error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	_ = "STUB: not implemented"
	return *new(SearchServiceClient)
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *searchServiceClient) StreamingSearch(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SearchRequest, SearchResponse], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SearchService_StreamingSearchClient = grpc.BidiStreamingClient[SearchRequest, SearchResponse]

type SearchServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	StreamingSearch(grpc.BidiStreamingServer[SearchRequest, SearchResponse]) error
	mustEmbedUnimplementedSearchServiceServer()
}

type UnimplementedSearchServiceServer struct{}

func (UnimplementedSearchServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedSearchServiceServer) StreamingSearch(grpc.BidiStreamingServer[SearchRequest, SearchResponse]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedSearchServiceServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	_ = "STUB: not implemented"
	return
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _SearchService_StreamingSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type SearchService_StreamingSearchServer = grpc.BidiStreamingServer[SearchRequest, SearchResponse]

var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingSearch",
			Handler:       _SearchService_StreamingSearch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "reflection/grpc_testing/test.proto",
}
