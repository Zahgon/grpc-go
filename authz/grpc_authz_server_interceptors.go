package authz

import (
	"context"
	"time"
	"unsafe"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/xds/rbac"
)

var logger = grpclog.Component("authz")

type StaticInterceptor struct {
	engines rbac.ChainEngine
}

func NewStatic(authzPolicy string) (*StaticInterceptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *StaticInterceptor) UnaryInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (i *StaticInterceptor) StreamInterceptor(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}

type FileWatcherInterceptor struct {
	options             FileWatcherOptions
	internalInterceptor unsafe.Pointer
	policyContents      []byte
	cancel              context.CancelFunc
}

type FileWatcherOptions struct {
	PolicyFile string

	RefreshDuration time.Duration

	OnPolicyUpdate func(string)
}

func NewFileWatcher(file string, duration time.Duration) (*FileWatcherInterceptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFileWatcherWithOptions(options FileWatcherOptions) (*FileWatcherInterceptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *FileWatcherInterceptor) run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (i *FileWatcherInterceptor) updateInternalInterceptor() error {
	_ = "STUB: not implemented"
	return nil
}

func (i *FileWatcherInterceptor) Close() { _ = "STUB: not implemented"; return }

func (i *FileWatcherInterceptor) UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (i *FileWatcherInterceptor) StreamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}
