package opencensus

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	keyServerMethod = tag.MustNewKey("grpc_server_method")
	keyServerStatus = tag.MustNewKey("grpc_server_status")
)

var (
	serverReceivedMessagesPerRPC        = stats.Int64("grpc.io/server/received_messages_per_rpc", "Number of messages received in each RPC. Has value 1 for non-streaming RPCs.", stats.UnitDimensionless)
	serverReceivedBytesPerRPC           = stats.Int64("grpc.io/server/received_bytes_per_rpc", "Total bytes received across all messages per RPC.", stats.UnitBytes)
	serverReceivedCompressedBytesPerRPC = stats.Int64("grpc.io/server/received_compressed_bytes_per_rpc", "Total compressed bytes received across all messages per RPC.", stats.UnitBytes)
	serverSentMessagesPerRPC            = stats.Int64("grpc.io/server/sent_messages_per_rpc", "Number of messages sent in each RPC. Has value 1 for non-streaming RPCs.", stats.UnitDimensionless)
	serverSentBytesPerRPC               = stats.Int64("grpc.io/server/sent_bytes_per_rpc", "Total bytes sent in across all response messages per RPC.", stats.UnitBytes)
	serverSentCompressedBytesPerRPC     = stats.Int64("grpc.io/server/sent_compressed_bytes_per_rpc", "Total compressed bytes sent in across all response messages per RPC.", stats.UnitBytes)
	serverStartedRPCs                   = stats.Int64("grpc.io/server/started_rpcs", "The total number of server RPCs ever opened, including those that have not completed.", stats.UnitDimensionless)
	serverLatency                       = stats.Float64("grpc.io/server/server_latency", "Time between first byte of request received to last byte of response sent, or terminal error.", stats.UnitMilliseconds)
)

var (
	ServerSentMessagesPerRPCView = &view.View{
		Name:        "grpc.io/server/sent_messages_per_rpc",
		Description: "Distribution of sent messages per RPC, by method.",
		TagKeys:     []tag.Key{keyServerMethod},
		Measure:     serverSentMessagesPerRPC,
		Aggregation: countDistribution,
	}

	ServerReceivedMessagesPerRPCView = &view.View{
		Name:        "grpc.io/server/received_messages_per_rpc",
		Description: "Distribution of received messages per RPC, by method.",
		TagKeys:     []tag.Key{keyServerMethod},
		Measure:     serverReceivedMessagesPerRPC,
		Aggregation: countDistribution,
	}

	ServerSentBytesPerRPCView = &view.View{
		Name:        "grpc.io/server/sent_bytes_per_rpc",
		Description: "Distribution of sent bytes per RPC, by method.",
		Measure:     serverSentBytesPerRPC,
		TagKeys:     []tag.Key{keyServerMethod},
		Aggregation: bytesDistribution,
	}

	ServerSentCompressedMessageBytesPerRPCView = &view.View{
		Name:        "grpc.io/server/sent_compressed_message_bytes_per_rpc",
		Description: "Distribution of sent compressed message bytes per RPC, by method.",
		Measure:     serverSentCompressedBytesPerRPC,
		TagKeys:     []tag.Key{keyServerMethod},
		Aggregation: bytesDistribution,
	}

	ServerReceivedBytesPerRPCView = &view.View{
		Name:        "grpc.io/server/received_bytes_per_rpc",
		Description: "Distribution of received bytes per RPC, by method.",
		Measure:     serverReceivedBytesPerRPC,
		TagKeys:     []tag.Key{keyServerMethod},
		Aggregation: bytesDistribution,
	}

	ServerReceivedCompressedMessageBytesPerRPCView = &view.View{
		Name:        "grpc.io/server/received_compressed_message_bytes_per_rpc",
		Description: "Distribution of received compressed message bytes per RPC, by method.",
		Measure:     serverReceivedCompressedBytesPerRPC,
		TagKeys:     []tag.Key{keyServerMethod},
		Aggregation: bytesDistribution,
	}

	ServerStartedRPCsView = &view.View{
		Measure:     serverStartedRPCs,
		Name:        "grpc.io/server/started_rpcs",
		Description: "Number of opened server RPCs, by method.",
		TagKeys:     []tag.Key{keyServerMethod},
		Aggregation: view.Count(),
	}

	ServerCompletedRPCsView = &view.View{
		Name:        "grpc.io/server/completed_rpcs",
		Description: "Number of completed RPCs by method and status.",
		TagKeys:     []tag.Key{keyServerMethod, keyServerStatus},
		Measure:     serverLatency,
		Aggregation: view.Count(),
	}

	ServerLatencyView = &view.View{
		Name:        "grpc.io/server/server_latency",
		Description: "Distribution of server latency in milliseconds, by method.",
		TagKeys:     []tag.Key{keyServerMethod},
		Measure:     serverLatency,
		Aggregation: millisecondsDistribution,
	}
)

var DefaultServerViews = []*view.View{
	ServerReceivedBytesPerRPCView,
	ServerSentBytesPerRPCView,
	ServerLatencyView,
	ServerCompletedRPCsView,
	ServerStartedRPCsView,
}
