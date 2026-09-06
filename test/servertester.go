package test

import (
	"bytes"
	"io"
	"testing"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

type serverTester struct {
	cc io.ReadWriteCloser
	t  testing.TB
	fr *http2.Framer

	headerBuf bytes.Buffer
	hpackEnc  *hpack.Encoder

	frc    chan http2.Frame
	frErrc chan error
}

func newServerTesterFromConn(t testing.TB, cc io.ReadWriteCloser) *serverTester {
	_ = "STUB: not implemented"
	return nil
}

func (st *serverTester) readFrame() (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

func (st *serverTester) greet() { _ = "STUB: not implemented"; return }

func (st *serverTester) greetWithSettings(settings ...http2.Setting) {
	_ = "STUB: not implemented"
	return
}

func (st *serverTester) writePreface() { _ = "STUB: not implemented"; return }

func (st *serverTester) writeInitialSettings() { _ = "STUB: not implemented"; return }

func (st *serverTester) writeSettingsAck() { _ = "STUB: not implemented"; return }

func (st *serverTester) wantGoAway(errCode http2.ErrCode) *http2.GoAwayFrame {
	_ = "STUB: not implemented"
	return nil
}

func (st *serverTester) wantPing() *http2.PingFrame { _ = "STUB: not implemented"; return nil }

func (st *serverTester) wantRSTStream(errCode http2.ErrCode) *http2.RSTStreamFrame {
	_ = "STUB: not implemented"
	return nil
}

func (st *serverTester) wantSettings() *http2.SettingsFrame { _ = "STUB: not implemented"; return nil }

func (st *serverTester) wantAnyFrame() http2.Frame {
	_ = "STUB: not implemented"
	return *new(http2.Frame)
}

func (st *serverTester) encodeHeaderField(k, v string) { _ = "STUB: not implemented"; return }

func (st *serverTester) encodeHeader(headers ...string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (st *serverTester) writeHeadersGRPC(streamID uint32, path string, endStream bool) {
	_ = "STUB: not implemented"
	return
}

func (st *serverTester) writeHeaders(p http2.HeadersFrameParam) { _ = "STUB: not implemented"; return }

func (st *serverTester) writeData(streamID uint32, endStream bool, data []byte) {
	_ = "STUB: not implemented"
	return
}

func (st *serverTester) writeRSTStream(streamID uint32, code http2.ErrCode) {
	_ = "STUB: not implemented"
	return
}

func (st *serverTester) writePing(ack bool, data [8]byte) { _ = "STUB: not implemented"; return }
