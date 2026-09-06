package bootstrap

import (
	"encoding/json"

	"google.golang.org/grpc/credentials"
)

var channelCredsRegistry = make(map[string]ChannelCredentials)

var callCredsRegistry = make(map[string]CallCredentials)

type ChannelCredentials interface {
	Build(config json.RawMessage) (credentials.Bundle, func(), error)

	Name() string
}

func RegisterChannelCredentials(c ChannelCredentials) { _ = "STUB: not implemented"; return }

func GetChannelCredentials(name string) ChannelCredentials {
	_ = "STUB: not implemented"
	return *new(ChannelCredentials)
}

type CallCredentials interface {
	Build(config json.RawMessage) (credentials.PerRPCCredentials, func(), error)

	Name() string
}

func RegisterCallCredentials(c CallCredentials) { _ = "STUB: not implemented"; return }

func GetCallCredentials(name string) CallCredentials {
	_ = "STUB: not implemented"
	return *new(CallCredentials)
}
