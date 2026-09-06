package grpc_testing

import (
	sync "sync"

	core "google.golang.org/grpc/interop/grpc_testing/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServerStats struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	TimeElapsed float64 `protobuf:"fixed64,1,opt,name=time_elapsed,json=timeElapsed,proto3" json:"time_elapsed,omitempty"`

	TimeUser float64 `protobuf:"fixed64,2,opt,name=time_user,json=timeUser,proto3" json:"time_user,omitempty"`

	TimeSystem float64 `protobuf:"fixed64,3,opt,name=time_system,json=timeSystem,proto3" json:"time_system,omitempty"`

	TotalCpuTime uint64 `protobuf:"varint,4,opt,name=total_cpu_time,json=totalCpuTime,proto3" json:"total_cpu_time,omitempty"`

	IdleCpuTime uint64 `protobuf:"varint,5,opt,name=idle_cpu_time,json=idleCpuTime,proto3" json:"idle_cpu_time,omitempty"`

	CqPollCount uint64 `protobuf:"varint,6,opt,name=cq_poll_count,json=cqPollCount,proto3" json:"cq_poll_count,omitempty"`

	CoreStats     *core.Stats `protobuf:"bytes,7,opt,name=core_stats,json=coreStats,proto3" json:"core_stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerStats) Reset() { _ = "STUB: not implemented"; return }

func (x *ServerStats) String() string { _ = "STUB: not implemented"; return "" }

func (*ServerStats) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ServerStats) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ServerStats) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ServerStats) GetTimeElapsed() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ServerStats) GetTimeUser() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ServerStats) GetTimeSystem() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ServerStats) GetTotalCpuTime() uint64 { _ = "STUB: not implemented"; return 0 }

func (x *ServerStats) GetIdleCpuTime() uint64 { _ = "STUB: not implemented"; return 0 }

func (x *ServerStats) GetCqPollCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (x *ServerStats) GetCoreStats() *core.Stats { _ = "STUB: not implemented"; return nil }

type HistogramParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Resolution    float64                `protobuf:"fixed64,1,opt,name=resolution,proto3" json:"resolution,omitempty"`
	MaxPossible   float64                `protobuf:"fixed64,2,opt,name=max_possible,json=maxPossible,proto3" json:"max_possible,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HistogramParams) Reset() { _ = "STUB: not implemented"; return }

func (x *HistogramParams) String() string { _ = "STUB: not implemented"; return "" }

func (*HistogramParams) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HistogramParams) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HistogramParams) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HistogramParams) GetResolution() float64 { _ = "STUB: not implemented"; return 0 }

func (x *HistogramParams) GetMaxPossible() float64 { _ = "STUB: not implemented"; return 0 }

type HistogramData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bucket        []uint32               `protobuf:"varint,1,rep,packed,name=bucket,proto3" json:"bucket,omitempty"`
	MinSeen       float64                `protobuf:"fixed64,2,opt,name=min_seen,json=minSeen,proto3" json:"min_seen,omitempty"`
	MaxSeen       float64                `protobuf:"fixed64,3,opt,name=max_seen,json=maxSeen,proto3" json:"max_seen,omitempty"`
	Sum           float64                `protobuf:"fixed64,4,opt,name=sum,proto3" json:"sum,omitempty"`
	SumOfSquares  float64                `protobuf:"fixed64,5,opt,name=sum_of_squares,json=sumOfSquares,proto3" json:"sum_of_squares,omitempty"`
	Count         float64                `protobuf:"fixed64,6,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HistogramData) Reset() { _ = "STUB: not implemented"; return }

func (x *HistogramData) String() string { _ = "STUB: not implemented"; return "" }

func (*HistogramData) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HistogramData) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HistogramData) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HistogramData) GetBucket() []uint32 { _ = "STUB: not implemented"; return nil }

func (x *HistogramData) GetMinSeen() float64 { _ = "STUB: not implemented"; return 0 }

func (x *HistogramData) GetMaxSeen() float64 { _ = "STUB: not implemented"; return 0 }

func (x *HistogramData) GetSum() float64 { _ = "STUB: not implemented"; return 0 }

func (x *HistogramData) GetSumOfSquares() float64 { _ = "STUB: not implemented"; return 0 }

func (x *HistogramData) GetCount() float64 { _ = "STUB: not implemented"; return 0 }

type RequestResultCount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Count         int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestResultCount) Reset() { _ = "STUB: not implemented"; return }

func (x *RequestResultCount) String() string { _ = "STUB: not implemented"; return "" }

func (*RequestResultCount) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *RequestResultCount) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*RequestResultCount) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *RequestResultCount) GetStatusCode() int32 { _ = "STUB: not implemented"; return 0 }

func (x *RequestResultCount) GetCount() int64 { _ = "STUB: not implemented"; return 0 }

type ClientStats struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	Latencies *HistogramData `protobuf:"bytes,1,opt,name=latencies,proto3" json:"latencies,omitempty"`

	TimeElapsed float64 `protobuf:"fixed64,2,opt,name=time_elapsed,json=timeElapsed,proto3" json:"time_elapsed,omitempty"`
	TimeUser    float64 `protobuf:"fixed64,3,opt,name=time_user,json=timeUser,proto3" json:"time_user,omitempty"`
	TimeSystem  float64 `protobuf:"fixed64,4,opt,name=time_system,json=timeSystem,proto3" json:"time_system,omitempty"`

	RequestResults []*RequestResultCount `protobuf:"bytes,5,rep,name=request_results,json=requestResults,proto3" json:"request_results,omitempty"`

	CqPollCount uint64 `protobuf:"varint,6,opt,name=cq_poll_count,json=cqPollCount,proto3" json:"cq_poll_count,omitempty"`

	CoreStats     *core.Stats `protobuf:"bytes,7,opt,name=core_stats,json=coreStats,proto3" json:"core_stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientStats) Reset() { _ = "STUB: not implemented"; return }

func (x *ClientStats) String() string { _ = "STUB: not implemented"; return "" }

func (*ClientStats) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *ClientStats) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*ClientStats) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *ClientStats) GetLatencies() *HistogramData { _ = "STUB: not implemented"; return nil }

func (x *ClientStats) GetTimeElapsed() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ClientStats) GetTimeUser() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ClientStats) GetTimeSystem() float64 { _ = "STUB: not implemented"; return 0 }

func (x *ClientStats) GetRequestResults() []*RequestResultCount {
	_ = "STUB: not implemented"
	return nil
}

func (x *ClientStats) GetCqPollCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (x *ClientStats) GetCoreStats() *core.Stats { _ = "STUB: not implemented"; return nil }

var File_grpc_testing_stats_proto protoreflect.FileDescriptor

const file_grpc_testing_stats_proto_rawDesc = "" +
	"\n" +
	"\x18grpc/testing/stats.proto\x12\fgrpc.testing\x1a\x15grpc/core/stats.proto\"\x8d\x02\n" +
	"\vServerStats\x12!\n" +
	"\ftime_elapsed\x18\x01 \x01(\x01R\vtimeElapsed\x12\x1b\n" +
	"\ttime_user\x18\x02 \x01(\x01R\btimeUser\x12\x1f\n" +
	"\vtime_system\x18\x03 \x01(\x01R\n" +
	"timeSystem\x12$\n" +
	"\x0etotal_cpu_time\x18\x04 \x01(\x04R\ftotalCpuTime\x12\"\n" +
	"\ridle_cpu_time\x18\x05 \x01(\x04R\vidleCpuTime\x12\"\n" +
	"\rcq_poll_count\x18\x06 \x01(\x04R\vcqPollCount\x12/\n" +
	"\n" +
	"core_stats\x18\a \x01(\v2\x10.grpc.core.StatsR\tcoreStats\"T\n" +
	"\x0fHistogramParams\x12\x1e\n" +
	"\n" +
	"resolution\x18\x01 \x01(\x01R\n" +
	"resolution\x12!\n" +
	"\fmax_possible\x18\x02 \x01(\x01R\vmaxPossible\"\xab\x01\n" +
	"\rHistogramData\x12\x16\n" +
	"\x06bucket\x18\x01 \x03(\rR\x06bucket\x12\x19\n" +
	"\bmin_seen\x18\x02 \x01(\x01R\aminSeen\x12\x19\n" +
	"\bmax_seen\x18\x03 \x01(\x01R\amaxSeen\x12\x10\n" +
	"\x03sum\x18\x04 \x01(\x01R\x03sum\x12$\n" +
	"\x0esum_of_squares\x18\x05 \x01(\x01R\fsumOfSquares\x12\x14\n" +
	"\x05count\x18\x06 \x01(\x01R\x05count\"K\n" +
	"\x12RequestResultCount\x12\x1f\n" +
	"\vstatus_code\x18\x01 \x01(\x05R\n" +
	"statusCode\x12\x14\n" +
	"\x05count\x18\x02 \x01(\x03R\x05count\"\xc9\x02\n" +
	"\vClientStats\x129\n" +
	"\tlatencies\x18\x01 \x01(\v2\x1b.grpc.testing.HistogramDataR\tlatencies\x12!\n" +
	"\ftime_elapsed\x18\x02 \x01(\x01R\vtimeElapsed\x12\x1b\n" +
	"\ttime_user\x18\x03 \x01(\x01R\btimeUser\x12\x1f\n" +
	"\vtime_system\x18\x04 \x01(\x01R\n" +
	"timeSystem\x12I\n" +
	"\x0frequest_results\x18\x05 \x03(\v2 .grpc.testing.RequestResultCountR\x0erequestResults\x12\"\n" +
	"\rcq_poll_count\x18\x06 \x01(\x04R\vcqPollCount\x12/\n" +
	"\n" +
	"core_stats\x18\a \x01(\v2\x10.grpc.core.StatsR\tcoreStatsB\x1f\n" +
	"\x0fio.grpc.testingB\n" +
	"StatsProtoP\x01b\x06proto3"

var (
	file_grpc_testing_stats_proto_rawDescOnce sync.Once
	file_grpc_testing_stats_proto_rawDescData []byte
)

func file_grpc_testing_stats_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_grpc_testing_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_testing_stats_proto_goTypes = []any{
	(*ServerStats)(nil),
	(*HistogramParams)(nil),
	(*HistogramData)(nil),
	(*RequestResultCount)(nil),
	(*ClientStats)(nil),
	(*core.Stats)(nil),
}
var file_grpc_testing_stats_proto_depIdxs = []int32{
	5,
	2,
	3,
	5,
	4,
	4,
	4,
	4,
	0,
}

func init()                               { file_grpc_testing_stats_proto_init() }
func file_grpc_testing_stats_proto_init() { _ = "STUB: not implemented"; return }
