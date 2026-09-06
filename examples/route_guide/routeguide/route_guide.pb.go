package routeguide

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Point struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Latitude      int32                  `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude     int32                  `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Point) Reset() { _ = "STUB: not implemented"; return }

func (x *Point) String() string { _ = "STUB: not implemented"; return "" }

func (*Point) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Point) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Point) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Point) GetLatitude() int32 { _ = "STUB: not implemented"; return 0 }

func (x *Point) GetLongitude() int32 { _ = "STUB: not implemented"; return 0 }

type Rectangle struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Lo *Point `protobuf:"bytes,1,opt,name=lo" json:"lo,omitempty"`

	Hi            *Point `protobuf:"bytes,2,opt,name=hi" json:"hi,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Rectangle) Reset() { _ = "STUB: not implemented"; return }

func (x *Rectangle) String() string { _ = "STUB: not implemented"; return "" }

func (*Rectangle) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Rectangle) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Rectangle) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Rectangle) GetLo() *Point { _ = "STUB: not implemented"; return nil }

func (x *Rectangle) GetHi() *Point { _ = "STUB: not implemented"; return nil }

type Feature struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`

	Location      *Point `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Feature) Reset() { _ = "STUB: not implemented"; return }

func (x *Feature) String() string { _ = "STUB: not implemented"; return "" }

func (*Feature) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Feature) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Feature) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Feature) GetName() string { _ = "STUB: not implemented"; return "" }

func (x *Feature) GetLocation() *Point { _ = "STUB: not implemented"; return nil }

type RouteNote struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Location *Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`

	Message       string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RouteNote) Reset() { _ = "STUB: not implemented"; return }

func (x *RouteNote) String() string { _ = "STUB: not implemented"; return "" }

func (*RouteNote) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *RouteNote) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*RouteNote) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *RouteNote) GetLocation() *Point { _ = "STUB: not implemented"; return nil }

func (x *RouteNote) GetMessage() string { _ = "STUB: not implemented"; return "" }

type RouteSummary struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	PointCount int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount" json:"point_count,omitempty"`

	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount" json:"feature_count,omitempty"`

	Distance int32 `protobuf:"varint,3,opt,name=distance" json:"distance,omitempty"`

	ElapsedTime   int32 `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime" json:"elapsed_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RouteSummary) Reset() { _ = "STUB: not implemented"; return }

func (x *RouteSummary) String() string { _ = "STUB: not implemented"; return "" }

func (*RouteSummary) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *RouteSummary) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*RouteSummary) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *RouteSummary) GetPointCount() int32 { _ = "STUB: not implemented"; return 0 }

func (x *RouteSummary) GetFeatureCount() int32 { _ = "STUB: not implemented"; return 0 }

func (x *RouteSummary) GetDistance() int32 { _ = "STUB: not implemented"; return 0 }

func (x *RouteSummary) GetElapsedTime() int32 { _ = "STUB: not implemented"; return 0 }

var File_examples_route_guide_routeguide_route_guide_proto protoreflect.FileDescriptor

const file_examples_route_guide_routeguide_route_guide_proto_rawDesc = "" +
	"\n" +
	"1examples/route_guide/routeguide/route_guide.proto\x12\n" +
	"routeguide\"A\n" +
	"\x05Point\x12\x1a\n" +
	"\blatitude\x18\x01 \x01(\x05R\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x02 \x01(\x05R\tlongitude\"Q\n" +
	"\tRectangle\x12!\n" +
	"\x02lo\x18\x01 \x01(\v2\x11.routeguide.PointR\x02lo\x12!\n" +
	"\x02hi\x18\x02 \x01(\v2\x11.routeguide.PointR\x02hi\"L\n" +
	"\aFeature\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12-\n" +
	"\blocation\x18\x02 \x01(\v2\x11.routeguide.PointR\blocation\"T\n" +
	"\tRouteNote\x12-\n" +
	"\blocation\x18\x01 \x01(\v2\x11.routeguide.PointR\blocation\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"\x93\x01\n" +
	"\fRouteSummary\x12\x1f\n" +
	"\vpoint_count\x18\x01 \x01(\x05R\n" +
	"pointCount\x12#\n" +
	"\rfeature_count\x18\x02 \x01(\x05R\ffeatureCount\x12\x1a\n" +
	"\bdistance\x18\x03 \x01(\x05R\bdistance\x12!\n" +
	"\felapsed_time\x18\x04 \x01(\x05R\velapsedTime2\x85\x02\n" +
	"\n" +
	"RouteGuide\x126\n" +
	"\n" +
	"GetFeature\x12\x11.routeguide.Point\x1a\x13.routeguide.Feature\"\x00\x12>\n" +
	"\fListFeatures\x12\x15.routeguide.Rectangle\x1a\x13.routeguide.Feature\"\x000\x01\x12>\n" +
	"\vRecordRoute\x12\x11.routeguide.Point\x1a\x18.routeguide.RouteSummary\"\x00(\x01\x12?\n" +
	"\tRouteChat\x12\x15.routeguide.RouteNote\x1a\x15.routeguide.RouteNote\"\x00(\x010\x01Bm\n" +
	"\x1bio.grpc.examples.routeguideB\x0fRouteGuideProtoP\x01Z6google.golang.org/grpc/examples/route_guide/routeguide\x92\x03\x02\b\x02b\beditionsp\xe8\a"

var (
	file_examples_route_guide_routeguide_route_guide_proto_rawDescOnce sync.Once
	file_examples_route_guide_routeguide_route_guide_proto_rawDescData []byte
)

func file_examples_route_guide_routeguide_route_guide_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_examples_route_guide_routeguide_route_guide_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_examples_route_guide_routeguide_route_guide_proto_goTypes = []any{
	(*Point)(nil),
	(*Rectangle)(nil),
	(*Feature)(nil),
	(*RouteNote)(nil),
	(*RouteSummary)(nil),
}
var file_examples_route_guide_routeguide_route_guide_proto_depIdxs = []int32{
	0,
	0,
	0,
	0,
	0,
	1,
	0,
	3,
	2,
	2,
	4,
	3,
	8,
	4,
	4,
	4,
	0,
}

func init() { file_examples_route_guide_routeguide_route_guide_proto_init() }
func file_examples_route_guide_routeguide_route_guide_proto_init() {
	_ = "STUB: not implemented"
	return
}
