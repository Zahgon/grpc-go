package grpc

import (
	"context"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/mem"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

func init() {
	internal.AcceptCompressors = acceptCompressors
}

type Compressor interface {
	Do(w io.Writer, p []byte) error

	Type() string
}

type gzipCompressor struct {
	pool sync.Pool
}

func NewGZIPCompressor() Compressor { _ = "STUB: not implemented"; return *new(Compressor) }

func NewGZIPCompressorWithLevel(level int) (Compressor, error) {
	_ = "STUB: not implemented"
	return *new(Compressor), nil
}

func (c *gzipCompressor) Do(w io.Writer, p []byte) error { _ = "STUB: not implemented"; return nil }

func (c *gzipCompressor) Type() string { _ = "STUB: not implemented"; return "" }

type Decompressor interface {
	Do(r io.Reader) ([]byte, error)

	Type() string
}

type gzipDecompressor struct {
	pool sync.Pool
}

func NewGZIPDecompressor() Decompressor { _ = "STUB: not implemented"; return *new(Decompressor) }

func (d *gzipDecompressor) Do(r io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *gzipDecompressor) doWithMaxSize(r io.Reader, maxMessageSize int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *gzipDecompressor) Type() string { _ = "STUB: not implemented"; return "" }

type callInfo struct {
	compressorName              string
	failFast                    bool
	maxReceiveMessageSize       *int
	maxSendMessageSize          *int
	creds                       credentials.PerRPCCredentials
	contentSubtype              string
	codec                       baseCodec
	maxRetryRPCBufferSize       int
	onFinish                    []func(err error)
	authority                   string
	acceptedResponseCompressors []string
}

func acceptedCompressorAllows(allowed []string, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func defaultCallInfo() *callInfo { _ = "STUB: not implemented"; return nil }

func newAcceptedCompressionConfig(names []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CallOption interface {
	before(*callInfo) error

	after(*callInfo, *csAttempt)
}

type EmptyCallOption struct{}

func (EmptyCallOption) before(*callInfo) error      { _ = "STUB: not implemented"; return nil }
func (EmptyCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func StaticMethod() CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type StaticMethodCallOption struct {
	EmptyCallOption
}

func Header(md *metadata.MD) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type HeaderCallOption struct {
	HeaderAddr *metadata.MD
}

func (o HeaderCallOption) before(*callInfo) error                { _ = "STUB: not implemented"; return nil }
func (o HeaderCallOption) after(_ *callInfo, attempt *csAttempt) { _ = "STUB: not implemented"; return }

func Trailer(md *metadata.MD) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type TrailerCallOption struct {
	TrailerAddr *metadata.MD
}

func (o TrailerCallOption) before(*callInfo) error { _ = "STUB: not implemented"; return nil }
func (o TrailerCallOption) after(_ *callInfo, attempt *csAttempt) {
	_ = "STUB: not implemented"
	return
}

func Peer(p *peer.Peer) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type PeerCallOption struct {
	PeerAddr *peer.Peer
}

func (o PeerCallOption) before(*callInfo) error                { _ = "STUB: not implemented"; return nil }
func (o PeerCallOption) after(_ *callInfo, attempt *csAttempt) { _ = "STUB: not implemented"; return }

func WaitForReady(waitForReady bool) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

func FailFast(failFast bool) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type FailFastCallOption struct {
	FailFast bool
}

func (o FailFastCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o FailFastCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func OnFinish(onFinish func(err error)) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

type OnFinishCallOption struct {
	OnFinish func(error)
}

func (o OnFinishCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o OnFinishCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func MaxCallRecvMsgSize(bytes int) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type MaxRecvMsgSizeCallOption struct {
	MaxRecvMsgSize int
}

func (o MaxRecvMsgSizeCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o MaxRecvMsgSizeCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func CallAuthority(authority string) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type AuthorityOverrideCallOption struct {
	Authority string
}

func (o AuthorityOverrideCallOption) before(c *callInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (o AuthorityOverrideCallOption) after(*callInfo, *csAttempt) {
	_ = "STUB: not implemented"
	return
}

func MaxCallSendMsgSize(bytes int) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type MaxSendMsgSizeCallOption struct {
	MaxSendMsgSize int
}

func (o MaxSendMsgSizeCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o MaxSendMsgSizeCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func PerRPCCredentials(creds credentials.PerRPCCredentials) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

type PerRPCCredsCallOption struct {
	Creds credentials.PerRPCCredentials
}

func (o PerRPCCredsCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o PerRPCCredsCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func UseCompressor(name string) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type CompressorCallOption struct {
	CompressorType string
}

func (o CompressorCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o CompressorCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func acceptCompressors(names ...string) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

type acceptCompressorsCallOption struct {
	names []string
}

func (o acceptCompressorsCallOption) before(c *callInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (acceptCompressorsCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func CallContentSubtype(contentSubtype string) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

type ContentSubtypeCallOption struct {
	ContentSubtype string
}

func (o ContentSubtypeCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o ContentSubtypeCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func ForceCodec(codec encoding.Codec) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

type ForceCodecCallOption struct {
	Codec encoding.Codec
}

func (o ForceCodecCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o ForceCodecCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func ForceCodecV2(codec encoding.CodecV2) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

type ForceCodecV2CallOption struct {
	CodecV2 encoding.CodecV2
}

func (o ForceCodecV2CallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o ForceCodecV2CallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func CallCustomCodec(codec Codec) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type CustomCodecCallOption struct {
	Codec Codec
}

func (o CustomCodecCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o CustomCodecCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

func MaxRetryRPCBufferSize(bytes int) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

type MaxRetryRPCBufferSizeCallOption struct {
	MaxRetryRPCBufferSize int
}

func (o MaxRetryRPCBufferSizeCallOption) before(c *callInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (o MaxRetryRPCBufferSizeCallOption) after(*callInfo, *csAttempt) {
	_ = "STUB: not implemented"
	return
}

type payloadFormat uint8

const (
	compressionNone payloadFormat = 0
	compressionMade payloadFormat = 1
)

func (pf payloadFormat) isCompressed() bool { _ = "STUB: not implemented"; return false }

type streamReader interface {
	ReadMessageHeader(header []byte) error
	Read(n int) (mem.BufferSlice, error)
}

type noCopy struct {
}

func (*noCopy) Lock()   { _ = "STUB: not implemented"; return }
func (*noCopy) Unlock() { _ = "STUB: not implemented"; return }

type parser struct {
	_ noCopy

	r streamReader

	header [5]byte

	bufferPool mem.BufferPool
}

func (p *parser) recvMsg(maxReceiveMessageSize int) (payloadFormat, mem.BufferSlice, error) {
	_ = "STUB: not implemented"
	return *new(payloadFormat), *new(mem.BufferSlice), nil
}

func encode(c baseCodec, msg any) (mem.BufferSlice, error) {
	_ = "STUB: not implemented"
	return *new(mem.BufferSlice), nil
}

func compress(in mem.BufferSlice, cp Compressor, compressor encoding.Compressor, pool mem.BufferPool) (mem.BufferSlice, payloadFormat, error) {
	_ = "STUB: not implemented"
	return *new(mem.BufferSlice), *new(payloadFormat), nil
}

const (
	payloadLen = 1
	sizeLen    = 4
	headerLen  = payloadLen + sizeLen
)

func msgHeader(data, compData mem.BufferSlice, pf payloadFormat) (hdr []byte, payload mem.BufferSlice) {
	_ = "STUB: not implemented"
	return nil, *new(mem.BufferSlice)
}

func outPayload(client bool, msg any, dataLength, payloadLength int, t time.Time) *stats.OutPayload {
	_ = "STUB: not implemented"
	return nil
}

func checkRecvPayload(pf payloadFormat, recvCompress string, haveCompressor bool, isServer bool) *status.Status {
	_ = "STUB: not implemented"
	return nil
}

type payloadInfo struct {
	compressedLength  int
	uncompressedBytes mem.BufferSlice
}

func (p *payloadInfo) free() { _ = "STUB: not implemented"; return }

func recvAndDecompress(p *parser, s recvCompressor, dc Decompressor, maxReceiveMessageSize int, payInfo *payloadInfo, compressor encoding.Compressor, isServer bool) (out mem.BufferSlice, err error) {
	_ = "STUB: not implemented"
	return *new(mem.BufferSlice), nil
}

func decompress(compressor encoding.Compressor, d mem.BufferSlice, dc Decompressor, maxReceiveMessageSize int, pool mem.BufferPool) (mem.BufferSlice, error) {
	_ = "STUB: not implemented"
	return *new(mem.BufferSlice), nil
}

type recvCompressor interface {
	RecvCompress() string
}

func recv(p *parser, c baseCodec, s recvCompressor, dc Decompressor, m any, maxReceiveMessageSize int, payInfo *payloadInfo, compressor encoding.Compressor, isServer bool) error {
	_ = "STUB: not implemented"
	return nil
}

type rpcInfo struct {
	failfast      bool
	preloaderInfo compressorInfo
}

type compressorInfo struct {
	codec baseCodec
	cp    Compressor
	comp  encoding.Compressor
}

type rpcInfoContextKey struct{}

func newContextWithRPCInfo(ctx context.Context, failfast bool, codec baseCodec, cp Compressor, comp encoding.Compressor) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func rpcInfoFromContext(ctx context.Context) (s *rpcInfo, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func Code(err error) codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }

func ErrorDesc(err error) string { _ = "STUB: not implemented"; return "" }

func Errorf(c codes.Code, format string, a ...any) error { _ = "STUB: not implemented"; return nil }

var errContextCanceled = status.Error(codes.Canceled, context.Canceled.Error())
var errContextDeadline = status.Error(codes.DeadlineExceeded, context.DeadlineExceeded.Error())

func toRPCErr(err error) error { _ = "STUB: not implemented"; return nil }

func setCallInfoCodec(c *callInfo) error { _ = "STUB: not implemented"; return nil }

const (
	SupportPackageIsVersion3 = true
	SupportPackageIsVersion4 = true
	SupportPackageIsVersion5 = true
	SupportPackageIsVersion6 = true
	SupportPackageIsVersion7 = true
	SupportPackageIsVersion8 = true
	SupportPackageIsVersion9 = true
)

const grpcUA = "grpc-go/" + Version
