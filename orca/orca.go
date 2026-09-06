package orca

import (
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/balancerload"
	"google.golang.org/grpc/metadata"
)

var logger = grpclog.Component("orca-backend-metrics")

type loadParser struct{}

func (loadParser) Parse(md metadata.MD) any { _ = "STUB: not implemented"; return *new(any) }

func init() {
	balancerload.SetParser(loadParser{})
}
