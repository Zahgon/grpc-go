package ringhash

import (
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const prefix = "[ring-hash-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *ringhashBalancer) *internalgrpclog.PrefixLogger {
	_ = "STUB: not implemented"
	return nil
}
