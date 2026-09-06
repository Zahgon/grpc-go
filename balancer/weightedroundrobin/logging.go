package weightedroundrobin

import (
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[%p] "

var logger = grpclog.Component("weighted-round-robin")

func prefixLogger(p *wrrBalancer) *internalgrpclog.PrefixLogger {
	_ = "STUB: not implemented"
	return nil
}
