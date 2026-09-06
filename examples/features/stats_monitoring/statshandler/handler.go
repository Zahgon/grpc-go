package statshandler

import (
	"context"

	"google.golang.org/grpc/stats"
)

type Handler struct{}

type connStatCtxKey struct{}

func (st *Handler) TagConn(ctx context.Context, stat *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (st *Handler) HandleConn(ctx context.Context, stat stats.ConnStats) {
	_ = "STUB: not implemented"
	return
}

type rpcStatCtxKey struct{}

func (st *Handler) TagRPC(ctx context.Context, stat *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (st *Handler) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}

func New() *Handler { _ = "STUB: not implemented"; return nil }
