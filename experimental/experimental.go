package experimental

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/mem"
)

func SetDefaultBufferPool(bufferPool mem.BufferPool) { _ = "STUB: not implemented"; return }

func WithBufferPool(bufferPool mem.BufferPool) grpc.DialOption {
	_ = "STUB: not implemented"
	return *new(grpc.DialOption)
}

func BufferPool(bufferPool mem.BufferPool) grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}

func AcceptCompressors(names ...string) grpc.CallOption {
	_ = "STUB: not implemented"
	return *new(grpc.CallOption)
}
