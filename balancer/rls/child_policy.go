package rls

import (
	"unsafe"

	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

type childPolicyWrapper struct {
	logger *internalgrpclog.PrefixLogger
	target string
	refCnt int

	state unsafe.Pointer
}

func newChildPolicyWrapper(target string) *childPolicyWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (c *childPolicyWrapper) acquireRef() { _ = "STUB: not implemented"; return }

func (c *childPolicyWrapper) releaseRef() bool { _ = "STUB: not implemented"; return false }

func (c *childPolicyWrapper) lamify(err error) { _ = "STUB: not implemented"; return }
