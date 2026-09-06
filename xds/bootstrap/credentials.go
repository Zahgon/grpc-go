package bootstrap

import (
	"encoding/json"

	"google.golang.org/grpc/credentials"
)

func init() {
	RegisterChannelCredentials(&insecureCredsBuilder{})
	RegisterChannelCredentials(&googleDefaultCredsBuilder{})
	RegisterChannelCredentials(&tlsCredsBuilder{})

	RegisterCallCredentials(&jwtCallCredsBuilder{})
}

type insecureCredsBuilder struct{}

func (i *insecureCredsBuilder) Build(json.RawMessage) (credentials.Bundle, func(), error) {
	_ = "STUB: not implemented"
	return *new(credentials.Bundle), nil, nil
}

func (i *insecureCredsBuilder) Name() string { _ = "STUB: not implemented"; return "" }

type tlsCredsBuilder struct{}

func (t *tlsCredsBuilder) Build(config json.RawMessage) (credentials.Bundle, func(), error) {
	_ = "STUB: not implemented"
	return *new(credentials.Bundle), nil, nil
}

func (t *tlsCredsBuilder) Name() string { _ = "STUB: not implemented"; return "" }

type googleDefaultCredsBuilder struct{}

func (d *googleDefaultCredsBuilder) Build(json.RawMessage) (credentials.Bundle, func(), error) {
	_ = "STUB: not implemented"
	return *new(credentials.Bundle), nil, nil
}

func (d *googleDefaultCredsBuilder) Name() string { _ = "STUB: not implemented"; return "" }

type jwtCallCredsBuilder struct{}

func (j *jwtCallCredsBuilder) Build(configJSON json.RawMessage) (credentials.PerRPCCredentials, func(), error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil, nil
}

func (j *jwtCallCredsBuilder) Name() string { _ = "STUB: not implemented"; return "" }
