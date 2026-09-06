package latency

import (
	"bytes"
	"context"
	"net"
	"time"
)

type Dialer func(network, address string) (net.Conn, error)

type TimeoutDialer func(network, address string, timeout time.Duration) (net.Conn, error)

type ContextDialer func(ctx context.Context, network, address string) (net.Conn, error)

type Network struct {
	Kbps    int
	Latency time.Duration
	MTU     int
}

var (
	Local = Network{0, 0, 0}

	LAN = Network{100 * 1024, 2 * time.Millisecond, 1500}

	WAN = Network{20 * 1024, 30 * time.Millisecond, 1500}

	Longhaul = Network{1000 * 1024, 200 * time.Millisecond, 9000}
)

func (n *Network) isLocal() bool { _ = "STUB: not implemented"; return false }

func (n *Network) Conn(c net.Conn) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

type conn struct {
	net.Conn
	network *Network

	readBuf     *bytes.Buffer
	lastSendEnd time.Time
	delay       time.Duration
}

type header struct {
	ReadTime int64
	Sz       int32
}

func (c *conn) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) sync() error { _ = "STUB: not implemented"; return nil }

func (n *Network) Listener(l net.Listener) net.Listener {
	_ = "STUB: not implemented"
	return *new(net.Listener)
}

type listener struct {
	net.Listener
	network *Network
}

func (l *listener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (n *Network) Dialer(d Dialer) Dialer { _ = "STUB: not implemented"; return *new(Dialer) }

func (n *Network) TimeoutDialer(d TimeoutDialer) TimeoutDialer {
	_ = "STUB: not implemented"
	return *new(TimeoutDialer)
}

func (n *Network) ContextDialer(d ContextDialer) ContextDialer {
	_ = "STUB: not implemented"
	return *new(ContextDialer)
}

func (n *Network) pktTime(b int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

var now = time.Now
var sleep = time.Sleep
