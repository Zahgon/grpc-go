package test

import (
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	defaultTestTimeout = 10 * time.Second
)

type ExpectedStatusCodes struct {
	ChannelzCode codes.Code
	CSDSCode     codes.Code
}

func RunRegisterTests(t *testing.T, ec ExpectedStatusCodes) { _ = "STUB: not implemented"; return }

func RunChannelz(conn *grpc.ClientConn) error { _ = "STUB: not implemented"; return nil }

func RunCSDS(conn *grpc.ClientConn) error { _ = "STUB: not implemented"; return nil }
