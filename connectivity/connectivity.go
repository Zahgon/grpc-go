package connectivity

import (
	"google.golang.org/grpc/grpclog"
)

var logger = grpclog.Component("core")

type State int

func (s State) String() string { _ = "STUB: not implemented"; return "" }

const (
	Idle State = iota

	Connecting

	Ready

	TransientFailure

	Shutdown
)

type ServingMode int

const (
	ServingModeStarting ServingMode = iota

	ServingModeServing

	ServingModeNotServing
)

func (s ServingMode) String() string { _ = "STUB: not implemented"; return "" }
