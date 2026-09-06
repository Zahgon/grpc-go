package test

import (
	"net"
	"testing"

	"golang.org/x/net/http2"
)

var (
	clientPreface = []byte(http2.ClientPreface)
)

func newClientTester(t *testing.T, conn net.Conn) *clientTester {
	_ = "STUB: not implemented"
	return nil
}

type clientTester struct {
	t    *testing.T
	conn net.Conn
	fr   *http2.Framer
}

func (ct *clientTester) greet() { _ = "STUB: not implemented"; return }

func (ct *clientTester) wantClientPreface() { _ = "STUB: not implemented"; return }

func (ct *clientTester) wantSettingsFrame() { _ = "STUB: not implemented"; return }

func (ct *clientTester) writeSettingsFrame() { _ = "STUB: not implemented"; return }

func (ct *clientTester) writeSettingsAck() { _ = "STUB: not implemented"; return }

func (ct *clientTester) writeGoAway(maxStreamID uint32, code http2.ErrCode, debugData []byte) {
	_ = "STUB: not implemented"
	return
}
