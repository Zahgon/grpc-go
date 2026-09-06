package routeguide

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion9

const (
	RouteGuide_GetFeature_FullMethodName   = "/routeguide.RouteGuide/GetFeature"
	RouteGuide_ListFeatures_FullMethodName = "/routeguide.RouteGuide/ListFeatures"
	RouteGuide_RecordRoute_FullMethodName  = "/routeguide.RouteGuide/RecordRoute"
	RouteGuide_RouteChat_FullMethodName    = "/routeguide.RouteGuide/RouteChat"
)

type RouteGuideClient interface {
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)

	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Feature], error)

	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Point, RouteSummary], error)

	RouteChat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RouteNote, RouteNote], error)
}

type routeGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideClient(cc grpc.ClientConnInterface) RouteGuideClient {
	_ = "STUB: not implemented"
	return *new(RouteGuideClient)
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *routeGuideClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Feature], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RouteGuide_ListFeaturesClient = grpc.ServerStreamingClient[Feature]

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Point, RouteSummary], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RouteGuide_RecordRouteClient = grpc.ClientStreamingClient[Point, RouteSummary]

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RouteNote, RouteNote], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RouteGuide_RouteChatClient = grpc.BidiStreamingClient[RouteNote, RouteNote]

type RouteGuideServer interface {
	GetFeature(context.Context, *Point) (*Feature, error)

	ListFeatures(*Rectangle, grpc.ServerStreamingServer[Feature]) error

	RecordRoute(grpc.ClientStreamingServer[Point, RouteSummary]) error

	RouteChat(grpc.BidiStreamingServer[RouteNote, RouteNote]) error
	mustEmbedUnimplementedRouteGuideServer()
}

type UnimplementedRouteGuideServer struct{}

func (UnimplementedRouteGuideServer) GetFeature(context.Context, *Point) (*Feature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedRouteGuideServer) ListFeatures(*Rectangle, grpc.ServerStreamingServer[Feature]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedRouteGuideServer) RecordRoute(grpc.ClientStreamingServer[Point, RouteSummary]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedRouteGuideServer) RouteChat(grpc.BidiStreamingServer[RouteNote, RouteNote]) error {
	_ = "STUB: not implemented"
	return nil
}

func (UnimplementedRouteGuideServer) mustEmbedUnimplementedRouteGuideServer() {
	_ = "STUB: not implemented"
	return
}
func (UnimplementedRouteGuideServer) testEmbeddedByValue() { _ = "STUB: not implemented"; return }

type UnsafeRouteGuideServer interface {
	mustEmbedUnimplementedRouteGuideServer()
}

func RegisterRouteGuideServer(s grpc.ServiceRegistrar, srv RouteGuideServer) {
	_ = "STUB: not implemented"
	return
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _RouteGuide_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type RouteGuide_ListFeaturesServer = grpc.ServerStreamingServer[Feature]

func _RouteGuide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type RouteGuide_RecordRouteServer = grpc.ClientStreamingServer[Point, RouteSummary]

func _RouteGuide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

type RouteGuide_RouteChatServer = grpc.BidiStreamingServer[RouteNote, RouteNote]

var RouteGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteGuide_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "examples/route_guide/routeguide/route_guide.proto",
}
