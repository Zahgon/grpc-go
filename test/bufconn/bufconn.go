package bufconn

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Listener struct {
	mu   sync.Mutex
	sz   int
	ch   chan net.Conn
	done chan struct{}
}

type netErrorTimeout struct {
	error
}

func (e netErrorTimeout) Timeout() bool   { _ = "STUB: not implemented"; return false }
func (e netErrorTimeout) Temporary() bool { _ = "STUB: not implemented"; return false }

var errClosed = fmt.Errorf("closed")
var errTimeout net.Error = netErrorTimeout{error: fmt.Errorf("i/o timeout")}

func Listen(sz int) *Listener { _ = "STUB: not implemented"; return nil }

func (l *Listener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (l *Listener) Close() error { _ = "STUB: not implemented"; return nil }

func (l *Listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (l *Listener) Dial() (net.Conn, error) { _ = "STUB: not implemented"; return *new(net.Conn), nil }

func (l *Listener) DialContext(ctx context.Context) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

type pipe struct {
	mu sync.Mutex

	buf  []byte
	w, r int

	wwait sync.Cond
	rwait sync.Cond

	wtimedout bool
	rtimedout bool

	wtimer *time.Timer
	rtimer *time.Timer

	closed      bool
	writeClosed bool
}

func newPipe(sz int) *pipe { _ = "STUB: not implemented"; return nil }

func (p *pipe) empty() bool { _ = "STUB: not implemented"; return false }

func (p *pipe) full() bool { _ = "STUB: not implemented"; return false }

func (p *pipe) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pipe) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pipe) Close() error { _ = "STUB: not implemented"; return nil }

func (p *pipe) closeWrite() error { _ = "STUB: not implemented"; return nil }

type conn struct {
	io.Reader
	io.Writer
}

func (c *conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (*conn) LocalAddr() net.Addr  { _ = "STUB: not implemented"; return *new(net.Addr) }
func (*conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

type addr struct{}

func (addr) Network() string { _ = "STUB: not implemented"; return "" }
func (addr) String() string  { _ = "STUB: not implemented"; return "" }
