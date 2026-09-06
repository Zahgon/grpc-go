package test

import (
	"bytes"
	"io"
	"net"
	"sync"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

type listenerWrapper struct {
	net.Listener
	mu  sync.Mutex
	rcw *rawConnWrapper
}

func listenWithConnControl(network, address string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (l *listenerWrapper) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (l *listenerWrapper) getLastConn() *rawConnWrapper { _ = "STUB: not implemented"; return nil }

type dialerWrapper struct {
	c   net.Conn
	rcw *rawConnWrapper
}

func (d *dialerWrapper) dialer(target string, t time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d *dialerWrapper) getRawConnWrapper() *rawConnWrapper { _ = "STUB: not implemented"; return nil }

type rawConnWrapper struct {
	cc io.ReadWriteCloser
	fr *http2.Framer

	headerBuf bytes.Buffer
	hpackEnc  *hpack.Encoder

	frc    chan http2.Frame
	frErrc chan error
}

func newRawConnWrapperFromConn(cc io.ReadWriteCloser) *rawConnWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (rcw *rawConnWrapper) Close() error { _ = "STUB: not implemented"; return nil }

func (rcw *rawConnWrapper) encodeHeaderField(k, v string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rcw *rawConnWrapper) encodeRawHeader(headers ...string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (rcw *rawConnWrapper) encodeHeader(headers ...string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (rcw *rawConnWrapper) writeHeaders(p http2.HeadersFrameParam) error {
	_ = "STUB: not implemented"
	return nil
}

func (rcw *rawConnWrapper) writeRSTStream(streamID uint32, code http2.ErrCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (rcw *rawConnWrapper) writeGoAway(maxStreamID uint32, code http2.ErrCode, debugData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (rcw *rawConnWrapper) writeDataFrame(streamID uint32, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}
