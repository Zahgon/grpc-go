package gzip

import (
	"compress/gzip"
	"io"
	"sync"

	"google.golang.org/grpc/encoding"
)

const Name = "gzip"

func init() {
	c := &compressor{}
	c.poolCompressor.New = func() any {
		return &writer{Writer: gzip.NewWriter(io.Discard), pool: &c.poolCompressor}
	}
	encoding.RegisterCompressor(c)
}

type writer struct {
	*gzip.Writer
	pool *sync.Pool
}

func SetLevel(level int) error { _ = "STUB: not implemented"; return nil }

func (c *compressor) Compress(w io.Writer) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (z *writer) Close() error { _ = "STUB: not implemented"; return nil }

var _ io.Closer = &reader{}

type reader struct {
	*gzip.Reader
	pool *sync.Pool
}

func (c *compressor) Decompress(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (r *reader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *reader) Close() error { _ = "STUB: not implemented"; return nil }

func (c *compressor) Name() string { _ = "STUB: not implemented"; return "" }

type compressor struct {
	poolCompressor   sync.Pool
	poolDecompressor sync.Pool
}
