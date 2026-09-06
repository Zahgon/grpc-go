package mem

import (
	"sync"
	"sync/atomic"
)

type Buffer interface {
	ReadOnlyData() []byte

	Ref()

	Free()

	Len() int

	Slice(start, end int) Buffer

	split(n int) (left, right Buffer)
	read(buf []byte) (int, Buffer)
}

var (
	bufferObjectPool = sync.Pool{New: func() any { return new(buffer) }}
)

func IsBelowBufferPoolingThreshold(size int) bool { _ = "STUB: not implemented"; return false }

type buffer struct {
	refs atomic.Int32
	data []byte

	rootBuf *buffer

	origData *[]byte
	pool     BufferPool
}

func newBuffer() *buffer { _ = "STUB: not implemented"; return nil }

func NewBuffer(data *[]byte, pool BufferPool) Buffer {
	_ = "STUB: not implemented"
	return *new(Buffer)
}

func Copy(data []byte, pool BufferPool) Buffer { _ = "STUB: not implemented"; return *new(Buffer) }

func (b *buffer) ReadOnlyData() []byte { _ = "STUB: not implemented"; return nil }

func (b *buffer) Ref() { _ = "STUB: not implemented"; return }

func (b *buffer) Free() { _ = "STUB: not implemented"; return }

func (b *buffer) Len() int { _ = "STUB: not implemented"; return 0 }

func (b *buffer) Slice(start, end int) Buffer { _ = "STUB: not implemented"; return *new(Buffer) }

func (b *buffer) split(n int) (Buffer, Buffer) {
	_ = "STUB: not implemented"
	return *new(Buffer), *new(Buffer)
}

func (b *buffer) read(buf []byte) (int, Buffer) { _ = "STUB: not implemented"; return 0, *new(Buffer) }

func (b *buffer) String() string { _ = "STUB: not implemented"; return "" }

func ReadUnsafe(dst []byte, buf Buffer) (int, Buffer) {
	_ = "STUB: not implemented"
	return 0, *new(Buffer)
}

func SplitUnsafe(buf Buffer, n int) (left, right Buffer) {
	_ = "STUB: not implemented"
	return *new(Buffer), *new(Buffer)
}

type emptyBuffer struct{}

func (e emptyBuffer) ReadOnlyData() []byte { _ = "STUB: not implemented"; return nil }

func (e emptyBuffer) Ref()  { _ = "STUB: not implemented"; return }
func (e emptyBuffer) Free() { _ = "STUB: not implemented"; return }

func (e emptyBuffer) Len() int { _ = "STUB: not implemented"; return 0 }

func (e emptyBuffer) Slice(start, end int) Buffer { _ = "STUB: not implemented"; return *new(Buffer) }

func (e emptyBuffer) split(int) (left, right Buffer) {
	_ = "STUB: not implemented"
	return *new(Buffer), *new(Buffer)
}

func (e emptyBuffer) read([]byte) (int, Buffer) { _ = "STUB: not implemented"; return 0, *new(Buffer) }

type SliceBuffer []byte

func (s SliceBuffer) ReadOnlyData() []byte { _ = "STUB: not implemented"; return nil }

func (s SliceBuffer) Ref() { _ = "STUB: not implemented"; return }

func (s SliceBuffer) Free() { _ = "STUB: not implemented"; return }

func (s SliceBuffer) Len() int { _ = "STUB: not implemented"; return 0 }

func (s SliceBuffer) Slice(start, end int) Buffer { _ = "STUB: not implemented"; return *new(Buffer) }

func (s SliceBuffer) split(n int) (left, right Buffer) {
	_ = "STUB: not implemented"
	return *new(Buffer), *new(Buffer)
}

func (s SliceBuffer) read(buf []byte) (int, Buffer) {
	_ = "STUB: not implemented"
	return 0, *new(Buffer)
}
