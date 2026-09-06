package grpc

type ServerStreamingClient[Res any] interface {
	Recv() (*Res, error)

	ClientStream
}

type ServerStreamingServer[Res any] interface {
	Send(*Res) error

	ServerStream
}

type ClientStreamingClient[Req any, Res any] interface {
	Send(*Req) error

	CloseAndRecv() (*Res, error)

	ClientStream
}

type ClientStreamingServer[Req any, Res any] interface {
	Recv() (*Req, error)

	SendAndClose(*Res) error

	ServerStream
}

type BidiStreamingClient[Req any, Res any] interface {
	Send(*Req) error

	Recv() (*Res, error)

	ClientStream
}

type BidiStreamingServer[Req any, Res any] interface {
	Recv() (*Req, error)

	Send(*Res) error

	ServerStream
}

type GenericClientStream[Req any, Res any] struct {
	ClientStream
}

var _ ServerStreamingClient[string] = (*GenericClientStream[int, string])(nil)
var _ ClientStreamingClient[int, string] = (*GenericClientStream[int, string])(nil)
var _ BidiStreamingClient[int, string] = (*GenericClientStream[int, string])(nil)

func (x *GenericClientStream[Req, Res]) Send(m *Req) error { _ = "STUB: not implemented"; return nil }

func (x *GenericClientStream[Req, Res]) Recv() (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *GenericClientStream[Req, Res]) CloseAndRecv() (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServerStream[Req any, Res any] struct {
	ServerStream
}

var _ ServerStreamingServer[string] = (*GenericServerStream[int, string])(nil)
var _ ClientStreamingServer[int, string] = (*GenericServerStream[int, string])(nil)
var _ BidiStreamingServer[int, string] = (*GenericServerStream[int, string])(nil)

func (x *GenericServerStream[Req, Res]) Send(m *Res) error { _ = "STUB: not implemented"; return nil }

func (x *GenericServerStream[Req, Res]) SendAndClose(m *Res) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *GenericServerStream[Req, Res]) Recv() (*Req, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
