package grpc

import (
	"google.golang.org/grpc/mem"
)

type PreparedMsg struct {
	encodedData mem.BufferSlice
	hdr         []byte
	payload     mem.BufferSlice
	pf          payloadFormat
}

func (p *PreparedMsg) Encode(s Stream, msg any) error { _ = "STUB: not implemented"; return nil }
