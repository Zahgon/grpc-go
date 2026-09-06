package codes

import (
	"google.golang.org/grpc/internal"
)

func init() {
	internal.CanonicalString = canonicalString
}

func (c Code) String() string { _ = "STUB: not implemented"; return "" }

func canonicalString(c Code) string { _ = "STUB: not implemented"; return "" }
