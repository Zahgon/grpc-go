package weightedtarget

import (
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[weighted-target-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *weightedTargetBalancer) *internalgrpclog.PrefixLogger {
	_ = "STUB: not implemented"
	return nil
}
