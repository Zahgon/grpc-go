package interop

import (
	"context"
	"sync"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/orca"

	v3orcapb "github.com/cncf/xds/go/xds/data/orca/v3"
	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var (
	reqSizes            = []int{27182, 8, 1828, 45904}
	respSizes           = []int{31415, 9, 2653, 58979}
	largeReqSize        = 271828
	largeRespSize       = 314159
	initialMetadataKey  = "x-grpc-test-echo-initial"
	trailingMetadataKey = "x-grpc-test-echo-trailing-bin"

	logger = grpclog.Component("interop")
)

func ClientNewPayload(t testpb.PayloadType, size int) *testpb.Payload {
	_ = "STUB: not implemented"
	return nil
}

func DoEmptyUnaryCall(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoLargeUnaryCall(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoClientStreaming(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoServerStreaming(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoPingPong(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoEmptyStream(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoTimeoutOnSleepingServer(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoComputeEngineCreds(ctx context.Context, tc testgrpc.TestServiceClient, serviceAccount, oauthScope string) {
	_ = "STUB: not implemented"
	return
}

func getServiceAccountJSONKey(keyFile string) []byte { _ = "STUB: not implemented"; return nil }

func DoServiceAccountCreds(ctx context.Context, tc testgrpc.TestServiceClient, serviceAccountKeyFile, oauthScope string) {
	_ = "STUB: not implemented"
	return
}

func DoJWTTokenCreds(ctx context.Context, tc testgrpc.TestServiceClient, serviceAccountKeyFile string) {
	_ = "STUB: not implemented"
	return
}

func GetToken(ctx context.Context, serviceAccountKeyFile string, oauthScope string) *oauth2.Token {
	_ = "STUB: not implemented"
	return nil
}

func DoOauth2TokenCreds(ctx context.Context, tc testgrpc.TestServiceClient, serviceAccountKeyFile, oauthScope string) {
	_ = "STUB: not implemented"
	return
}

func DoPerRPCCreds(ctx context.Context, tc testgrpc.TestServiceClient, serviceAccountKeyFile, oauthScope string) {
	_ = "STUB: not implemented"
	return
}

func DoGoogleDefaultCredentials(ctx context.Context, tc testgrpc.TestServiceClient, defaultServiceAccount string) {
	_ = "STUB: not implemented"
	return
}

func DoComputeEngineChannelCredentials(ctx context.Context, tc testgrpc.TestServiceClient, defaultServiceAccount string) {
	_ = "STUB: not implemented"
	return
}

var testMetadata = metadata.MD{
	"key1": []string{"value1"},
	"key2": []string{"value2"},
}

func DoCancelAfterBegin(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoCancelAfterFirstResponse(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

var (
	initialMetadataValue  = "test_initial_metadata_value"
	trailingMetadataValue = "\x0a\x0b\x0a\x0b\x0a\x0b"
	customMetadata        = metadata.Pairs(
		initialMetadataKey, initialMetadataValue,
		trailingMetadataKey, trailingMetadataValue,
	)
)

func validateMetadata(header, trailer metadata.MD) { _ = "STUB: not implemented"; return }

func DoCustomMetadata(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoStatusCodeAndMessage(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoSpecialStatusMessage(ctx context.Context, tc testgrpc.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

func DoUnimplementedService(ctx context.Context, tc testgrpc.UnimplementedServiceClient) {
	_ = "STUB: not implemented"
	return
}

func DoUnimplementedMethod(ctx context.Context, cc *grpc.ClientConn) {
	_ = "STUB: not implemented"
	return
}

func DoPickFirstUnary(ctx context.Context, tc testgrpc.TestServiceClient) {
	_ = "STUB: not implemented"
	return
}

type testServer struct {
	testgrpc.UnimplementedTestServiceServer

	orcaMu          sync.Mutex
	metricsRecorder orca.ServerMetricsRecorder
}

type NewTestServerOptions struct {
	MetricsRecorder orca.ServerMetricsRecorder
}

func NewTestServer(opts ...NewTestServerOptions) testgrpc.TestServiceServer {
	_ = "STUB: not implemented"
	return *new(testgrpc.TestServiceServer)
}

func (s *testServer) EmptyCall(context.Context, *testpb.Empty) (*testpb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func serverNewPayload(t testpb.PayloadType, size int32) (*testpb.Payload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *testServer) UnaryCall(ctx context.Context, in *testpb.SimpleRequest) (*testpb.SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setORCAMetrics(r orca.ServerMetricsRecorder, orcaData *testpb.TestOrcaReport) {
	_ = "STUB: not implemented"
	return
}

func (s *testServer) StreamingOutputCall(args *testpb.StreamingOutputCallRequest, stream testgrpc.TestService_StreamingOutputCallServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *testServer) StreamingInputCall(stream testgrpc.TestService_StreamingInputCallServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *testServer) FullDuplexCall(stream testgrpc.TestService_FullDuplexCallServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *testServer) HalfDuplexCall(stream testgrpc.TestService_HalfDuplexCallServer) error {
	_ = "STUB: not implemented"
	return nil
}

func DoORCAPerRPCTest(ctx context.Context, tc testgrpc.TestServiceClient) {
	_ = "STUB: not implemented"
	return
}

func DoORCAOOBTest(ctx context.Context, tc testgrpc.TestServiceClient) {
	_ = "STUB: not implemented"
	return
}

func checkORCAMetrics(ctx context.Context, tc testgrpc.TestServiceClient, want *v3orcapb.OrcaLoadReport) {
	_ = "STUB: not implemented"
	return
}
