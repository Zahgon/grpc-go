package opencensus

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	keyClientMethod = tag.MustNewKey("grpc_client_method")
	keyClientStatus = tag.MustNewKey("grpc_client_status")
)

var (
	clientSentMessagesPerRPC            = stats.Int64("grpc.io/client/sent_messages_per_rpc", "Number of messages sent in the RPC (always 1 for non-streaming RPCs).", stats.UnitDimensionless)
	clientSentBytesPerRPC               = stats.Int64("grpc.io/client/sent_bytes_per_rpc", "Total bytes sent across all request messages per RPC.", stats.UnitBytes)
	clientSentCompressedBytesPerRPC     = stats.Int64("grpc.io/client/sent_compressed_message_bytes_per_rpc", "Total compressed bytes sent across all request messages per RPC.", stats.UnitBytes)
	clientReceivedMessagesPerRPC        = stats.Int64("grpc.io/client/received_messages_per_rpc", "Number of response messages received per RPC (always 1 for non-streaming RPCs).", stats.UnitDimensionless)
	clientReceivedBytesPerRPC           = stats.Int64("grpc.io/client/received_bytes_per_rpc", "Total bytes received across all response messages per RPC.", stats.UnitBytes)
	clientReceivedCompressedBytesPerRPC = stats.Int64("grpc.io/client/received_compressed_message_bytes_per_rpc", "Total compressed bytes received across all response messages per RPC.", stats.UnitBytes)
	clientRoundtripLatency              = stats.Float64("grpc.io/client/roundtrip_latency", "Time between first byte of request sent to last byte of response received, or terminal error.", stats.UnitMilliseconds)
	clientStartedRPCs                   = stats.Int64("grpc.io/client/started_rpcs", "The total number of client RPCs ever opened, including those that have not completed.", stats.UnitDimensionless)
	clientServerLatency                 = stats.Float64("grpc.io/client/server_latency", `Propagated from the server and should have the same value as "grpc.io/server/latency".`, stats.UnitMilliseconds)

	clientAPILatency = stats.Float64("grpc.io/client/api_latency", "The end-to-end time the gRPC library takes to complete an RPC from the application’s perspective", stats.UnitMilliseconds)
)

var (
	ClientSentMessagesPerRPCView = &view.View{
		Measure:     clientSentMessagesPerRPC,
		Name:        "grpc.io/client/sent_messages_per_rpc",
		Description: "Distribution of sent messages per RPC, by method.",
		TagKeys:     []tag.Key{keyClientMethod},
		Aggregation: countDistribution,
	}

	ClientReceivedMessagesPerRPCView = &view.View{
		Measure:     clientReceivedMessagesPerRPC,
		Name:        "grpc.io/client/received_messages_per_rpc",
		Description: "Distribution of received messages per RPC, by method.",
		TagKeys:     []tag.Key{keyClientMethod},
		Aggregation: countDistribution,
	}

	ClientSentBytesPerRPCView = &view.View{
		Measure:     clientSentBytesPerRPC,
		Name:        "grpc.io/client/sent_bytes_per_rpc",
		Description: "Distribution of sent bytes per RPC, by method.",
		TagKeys:     []tag.Key{keyClientMethod},
		Aggregation: bytesDistribution,
	}

	ClientSentCompressedMessageBytesPerRPCView = &view.View{
		Measure:     clientSentCompressedBytesPerRPC,
		Name:        "grpc.io/client/sent_compressed_message_bytes_per_rpc",
		Description: "Distribution of sent compressed message bytes per RPC, by method.",
		TagKeys:     []tag.Key{keyClientMethod},
		Aggregation: bytesDistribution,
	}

	ClientReceivedBytesPerRPCView = &view.View{
		Measure:     clientReceivedBytesPerRPC,
		Name:        "grpc.io/client/received_bytes_per_rpc",
		Description: "Distribution of received bytes per RPC, by method.",
		TagKeys:     []tag.Key{keyClientMethod},
		Aggregation: bytesDistribution,
	}

	ClientReceivedCompressedMessageBytesPerRPCView = &view.View{
		Measure:     clientReceivedCompressedBytesPerRPC,
		Name:        "grpc.io/client/received_compressed_message_bytes_per_rpc",
		Description: "Distribution of received compressed message bytes per RPC, by method.",
		TagKeys:     []tag.Key{keyClientMethod},
		Aggregation: bytesDistribution,
	}

	ClientStartedRPCsView = &view.View{
		Measure:     clientStartedRPCs,
		Name:        "grpc.io/client/started_rpcs",
		Description: "Number of opened client RPCs, by method.",
		TagKeys:     []tag.Key{keyClientMethod},
		Aggregation: view.Count(),
	}

	ClientCompletedRPCsView = &view.View{
		Measure:     clientRoundtripLatency,
		Name:        "grpc.io/client/completed_rpcs",
		Description: "Number of completed RPCs by method and status.",
		TagKeys:     []tag.Key{keyClientMethod, keyClientStatus},
		Aggregation: view.Count(),
	}

	ClientRoundtripLatencyView = &view.View{
		Measure:     clientRoundtripLatency,
		Name:        "grpc.io/client/roundtrip_latency",
		Description: "Distribution of round-trip latency, by method.",
		TagKeys:     []tag.Key{keyClientMethod},
		Aggregation: millisecondsDistribution,
	}

	ClientAPILatencyView = &view.View{
		Measure:     clientAPILatency,
		Name:        "grpc.io/client/api_latency",
		Description: "Distribution of client api latency, by method and status",
		TagKeys:     []tag.Key{keyClientMethod, keyClientStatus},
		Aggregation: millisecondsDistribution,
	}
)

var DefaultClientViews = []*view.View{
	ClientSentBytesPerRPCView,
	ClientReceivedBytesPerRPCView,
	ClientRoundtripLatencyView,
	ClientCompletedRPCsView,
	ClientStartedRPCsView,
}
