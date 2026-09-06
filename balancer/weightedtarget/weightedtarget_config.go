package weightedtarget

import (
	"encoding/json"

	internalserviceconfig "google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/serviceconfig"
)

type Target struct {
	Weight uint32 `json:"weight,omitempty"`

	ChildPolicy *internalserviceconfig.BalancerConfig `json:"childPolicy,omitempty"`
}

type LBConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`

	Targets map[string]Target `json:"targets,omitempty"`
}

func parseConfig(c json.RawMessage) (*LBConfig, error) { _ = "STUB: not implemented"; return nil, nil }
