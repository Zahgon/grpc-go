package state

import (
	"google.golang.org/grpc/resolver"
)

type keyType string

const key = keyType("grpc.grpclb.state")

type State struct {
	BalancerAddresses []resolver.Address
}

func Set(state resolver.State, s *State) resolver.State {
	_ = "STUB: not implemented"
	return *new(resolver.State)
}

func Get(state resolver.State) *State { _ = "STUB: not implemented"; return nil }
