package rls

import (
	"encoding/json"
	"time"

	"google.golang.org/grpc/balancer/rls/internal/keys"
	rlspb "google.golang.org/grpc/internal/proto/grpc_lookup_v1"
	"google.golang.org/grpc/serviceconfig"
	"google.golang.org/protobuf/types/known/durationpb"
)

const (
	maxMaxAge = 5 * time.Minute

	maxCacheSize = 5 * 1024 * 1024 * 8

	defaultLookupServiceTimeout = 10 * time.Second

	dummyChildPolicyTarget = "target_name_to_be_filled_in_later"
)

type lbConfig struct {
	serviceconfig.LoadBalancingConfig

	cacheSizeBytes       int64
	kbMap                keys.BuilderMap
	lookupService        string
	lookupServiceTimeout time.Duration
	maxAge               time.Duration
	staleAge             time.Duration
	defaultTarget        string

	childPolicyName             string
	childPolicyConfig           map[string]json.RawMessage
	childPolicyTargetField      string
	controlChannelServiceConfig string
}

func (lbCfg *lbConfig) Equal(other *lbConfig) bool { _ = "STUB: not implemented"; return false }

func childPolicyConfigEqual(a, b map[string]json.RawMessage) bool {
	_ = "STUB: not implemented"
	return false
}

type lbConfigJSON struct {
	RouteLookupConfig                json.RawMessage
	RouteLookupChannelServiceConfig  json.RawMessage
	ChildPolicy                      []map[string]json.RawMessage
	ChildPolicyConfigTargetFieldName string
}

func (rlsBB) ParseConfig(c json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	_ = "STUB: not implemented"
	return *new(serviceconfig.LoadBalancingConfig), nil
}

func parseRLSProto(rlsProto *rlspb.RouteLookupConfig) (*lbConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseChildPolicyConfigs(childPolicies []map[string]json.RawMessage, targetFieldName string) (string, map[string]json.RawMessage, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func convertDuration(d *durationpb.Duration) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
