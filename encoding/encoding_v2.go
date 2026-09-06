package encoding

import (
	"google.golang.org/grpc/mem"
)

type CodecV2 interface {
	Marshal(v any) (out mem.BufferSlice, err error)

	Unmarshal(data mem.BufferSlice, v any) error

	Name() string
}

func RegisterCodecV2(codec CodecV2) { _ = "STUB: not implemented"; return }

func GetCodecV2(contentSubtype string) CodecV2 { _ = "STUB: not implemented"; return *new(CodecV2) }
