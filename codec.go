package grpc

import (
	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/proto"
	"google.golang.org/grpc/mem"
)

type baseCodec interface {
	Marshal(v any) (mem.BufferSlice, error)
	Unmarshal(data mem.BufferSlice, v any) error
}

func getCodec(name string) encoding.CodecV2 {
	_ = "STUB: not implemented"
	return *new(encoding.CodecV2)
}

func newCodecV0Bridge(c Codec) baseCodec { _ = "STUB: not implemented"; return *new(baseCodec) }

func newCodecV1Bridge(c encoding.Codec) encoding.CodecV2 {
	_ = "STUB: not implemented"
	return *new(encoding.CodecV2)
}

var _ baseCodec = codecV0Bridge{}

type codecV0Bridge struct {
	codec interface {
		Marshal(v any) ([]byte, error)
		Unmarshal(data []byte, v any) error
	}
}

func (c codecV0Bridge) Marshal(v any) (mem.BufferSlice, error) {
	_ = "STUB: not implemented"
	return *new(mem.BufferSlice), nil
}

func (c codecV0Bridge) Unmarshal(data mem.BufferSlice, v any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

var _ encoding.CodecV2 = codecV1Bridge{}

type codecV1Bridge struct {
	codecV0Bridge
	name string
}

func (c codecV1Bridge) Name() string { _ = "STUB: not implemented"; return "" }

type Codec interface {
	Marshal(v any) ([]byte, error)

	Unmarshal(data []byte, v any) error

	String() string
}
