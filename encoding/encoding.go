package encoding

import (
	"io"
	"slices"

	"google.golang.org/grpc/encoding/internal"
	"google.golang.org/grpc/internal/grpcutil"
)

const Identity = "identity"

func init() {
	internal.RegisterCompressorForTesting = func(c Compressor) func() {
		name := c.Name()
		curCompressor, found := registeredCompressor[name]
		RegisterCompressor(c)
		return func() {
			if found {
				registeredCompressor[name] = curCompressor
				return
			}
			delete(registeredCompressor, name)
			grpcutil.RegisteredCompressorNames = slices.DeleteFunc(grpcutil.RegisteredCompressorNames, func(s string) bool {
				return s == name
			})
		}
	}
}

type Compressor interface {
	Compress(w io.Writer) (io.WriteCloser, error)

	Decompress(r io.Reader) (io.Reader, error)

	Name() string
}

var registeredCompressor = make(map[string]Compressor)

func RegisterCompressor(c Compressor) { _ = "STUB: not implemented"; return }

func GetCompressor(name string) Compressor { _ = "STUB: not implemented"; return *new(Compressor) }

type Codec interface {
	Marshal(v any) ([]byte, error)

	Unmarshal(data []byte, v any) error

	Name() string
}

var registeredCodecs = make(map[string]any)

func RegisterCodec(codec Codec) { _ = "STUB: not implemented"; return }

func GetCodec(contentSubtype string) Codec { _ = "STUB: not implemented"; return *new(Codec) }
