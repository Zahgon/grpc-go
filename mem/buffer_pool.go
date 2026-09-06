package mem

import (
	"fmt"

	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/internal/mem"
)

type BufferPool interface {
	Get(length int) *[]byte

	Put(*[]byte)
}

var (
	defaultBufferPoolSizeExponents = []uint8{
		8,
		12,
		14,
		15,
		20,
	}
	defaultBufferPool BufferPool
)

func init() {
	var err error
	defaultBufferPool, err = NewBinaryTieredBufferPool(defaultBufferPoolSizeExponents...)
	if err != nil {
		panic(fmt.Sprintf("Failed to create default buffer pool: %v", err))
	}

	internal.SetDefaultBufferPool = func(pool BufferPool) {
		defaultBufferPool = pool
	}
}

func DefaultBufferPool() BufferPool { _ = "STUB: not implemented"; return *new(BufferPool) }

func NewTieredBufferPool(poolSizes ...int) BufferPool {
	_ = "STUB: not implemented"
	return *new(BufferPool)
}

func NewBinaryTieredBufferPool(powerOfTwoExponents ...uint8) (BufferPool, error) {
	_ = "STUB: not implemented"
	return *new(BufferPool), nil
}

type NopBufferPool struct {
	mem.NopBufferPool
}

var _ BufferPool = NopBufferPool{}
