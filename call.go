package grpc

import (
	"context"
)

func (cc *ClientConn) Invoke(ctx context.Context, method string, args, reply any, opts ...CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

func combine(o1 []CallOption, o2 []CallOption) []CallOption { _ = "STUB: not implemented"; return nil }

func Invoke(ctx context.Context, method string, args, reply any, cc *ClientConn, opts ...CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

var unaryStreamDesc = &StreamDesc{ServerStreams: false, ClientStreams: false}

func invoke(ctx context.Context, method string, req, reply any, cc *ClientConn, opts ...CallOption) error {
	_ = "STUB: not implemented"
	return nil
}
