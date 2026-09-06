package mem

import (
	"io"
)

const (
	readAllBufSize = 32 * 1024
)

type BufferSlice []Buffer

func (s BufferSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (s BufferSlice) Ref() { _ = "STUB: not implemented"; return }

func (s BufferSlice) Free() { _ = "STUB: not implemented"; return }

func (s BufferSlice) CopyTo(dst []byte) int { _ = "STUB: not implemented"; return 0 }

func (s BufferSlice) Materialize() []byte { _ = "STUB: not implemented"; return nil }

func (s BufferSlice) MaterializeToBuffer(pool BufferPool) Buffer {
	_ = "STUB: not implemented"
	return *new(Buffer)
}

func (s BufferSlice) Reader() *Reader { _ = "STUB: not implemented"; return nil }

type Reader struct {
	data BufferSlice
	len  int

	bufferIdx int
}

func (r *Reader) Remaining() int { _ = "STUB: not implemented"; return 0 }

func (r *Reader) Reset(s BufferSlice) { _ = "STUB: not implemented"; return }

func (r *Reader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *Reader) freeFirstBufferIfEmpty() bool { _ = "STUB: not implemented"; return false }

func (r *Reader) Read(buf []byte) (n int, _ error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Reader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

var _ io.Writer = (*writer)(nil)

type writer struct {
	buffers *BufferSlice
	pool    BufferPool
}

func (w *writer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func NewWriter(buffers *BufferSlice, pool BufferPool) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func ReadAll(r io.Reader, pool BufferPool) (BufferSlice, error) {
	_ = "STUB: not implemented"
	return *new(BufferSlice), nil
}

func (r *Reader) Discard(n int) (discarded int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *Reader) Peek(n int, res [][]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
