package dns

import (
	"time"

	"google.golang.org/grpc/resolver"
)

func SetResolvingTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

func NewBuilder() resolver.Builder { _ = "STUB: not implemented"; return *new(resolver.Builder) }

func SetMinResolutionInterval(d time.Duration) { _ = "STUB: not implemented"; return }
