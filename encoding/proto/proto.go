package proto

import (
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/mem"
	"google.golang.org/protobuf/proto"
)

const Name = "proto"

func init() {
	encoding.RegisterCodecV2(&codecV2{})
}

type codecV2 struct{}

func (c *codecV2) Marshal(v any) (data mem.BufferSlice, err error) {
	_ = "STUB: not implemented"
	return *new(mem.BufferSlice), nil
}

func (c *codecV2) Unmarshal(data mem.BufferSlice, v any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func messageV2Of(v any) proto.Message { _ = "STUB: not implemented"; return *new(proto.Message) }

func (c *codecV2) Name() string { _ = "STUB: not implemented"; return "" }
