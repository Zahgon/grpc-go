package metadata

import (
	"context"

	"google.golang.org/grpc/internal"
)

func init() {
	internal.FromOutgoingContextRaw = fromOutgoingContextRaw
}

func DecodeKeyValue(k, v string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type MD map[string][]string

func New(m map[string]string) MD { _ = "STUB: not implemented"; return *new(MD) }

func Pairs(kv ...string) MD { _ = "STUB: not implemented"; return *new(MD) }

var loggableMetadataKeys = map[string]bool{
	"content-type":               true,
	"te":                         true,
	"user-agent":                 true,
	"grpc-encoding":              true,
	"grpc-accept-encoding":       true,
	"grpc-timeout":               true,
	"grpc-status":                true,
	"grpc-message-type":          true,
	"grpc-previous-rpc-attempts": true,
	"grpc-retry-pushback-ms":     true,
}

func (md MD) String() string { _ = "STUB: not implemented"; return "" }

func (md MD) Len() int { _ = "STUB: not implemented"; return 0 }

func (md MD) Copy() MD { _ = "STUB: not implemented"; return *new(MD) }

func (md MD) Get(k string) []string { _ = "STUB: not implemented"; return nil }

func (md MD) Set(k string, vals ...string) { _ = "STUB: not implemented"; return }

func (md MD) Append(k string, vals ...string) { _ = "STUB: not implemented"; return }

func (md MD) Delete(k string) { _ = "STUB: not implemented"; return }

func Join(mds ...MD) MD { _ = "STUB: not implemented"; return *new(MD) }

type mdIncomingKey struct{}
type mdOutgoingKey struct{}

func NewIncomingContext(ctx context.Context, md MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func NewOutgoingContext(ctx context.Context, md MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func AppendToOutgoingContext(ctx context.Context, kv ...string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromIncomingContext(ctx context.Context) (MD, bool) {
	_ = "STUB: not implemented"
	return *new(MD), false
}

func ValueFromIncomingContext(ctx context.Context, key string) []string {
	_ = "STUB: not implemented"
	return nil
}

func copyOf(v []string) []string { _ = "STUB: not implemented"; return nil }

func fromOutgoingContextRaw(ctx context.Context) (MD, [][]string, bool) {
	_ = "STUB: not implemented"
	return *new(MD), nil, false
}

func FromOutgoingContext(ctx context.Context) (MD, bool) {
	_ = "STUB: not implemented"
	return *new(MD), false
}

func ValueFromOutgoingContext(ctx context.Context, key string) []string {
	_ = "STUB: not implemented"
	return nil
}

type rawMD struct {
	md    MD
	added [][]string
}
