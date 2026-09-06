package status

import (
	spb "google.golang.org/genproto/googleapis/rpc/status"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
)

type Status = status.Status

func New(c codes.Code, msg string) *Status { _ = "STUB: not implemented"; return nil }

func Newf(c codes.Code, format string, a ...any) *Status { _ = "STUB: not implemented"; return nil }

func Error(c codes.Code, msg string) error { _ = "STUB: not implemented"; return nil }

func Errorf(c codes.Code, format string, a ...any) error { _ = "STUB: not implemented"; return nil }

func ErrorProto(s *spb.Status) error { _ = "STUB: not implemented"; return nil }

func FromProto(s *spb.Status) *Status { _ = "STUB: not implemented"; return nil }

func FromError(err error) (s *Status, ok bool) { _ = "STUB: not implemented"; return nil, false }

func Convert(err error) *Status { _ = "STUB: not implemented"; return nil }

func Code(err error) codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }

func FromContextError(err error) *Status { _ = "STUB: not implemented"; return nil }
