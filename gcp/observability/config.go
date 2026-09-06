package observability

import (
	"context"
	"encoding/json"
)

const envProjectID = "GOOGLE_CLOUD_PROJECT"

func fetchDefaultProjectID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func validateMethodString(method string) error { _ = "STUB: not implemented"; return nil }

func validateLogEventMethod(methods []string, exclude bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateLoggingEvents(config *config) error { _ = "STUB: not implemented"; return nil }

func unmarshalAndVerifyConfig(rawJSON json.RawMessage) (*config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseObservabilityConfig() (*config, error) { _ = "STUB: not implemented"; return nil, nil }

func ensureProjectIDInObservabilityConfig(ctx context.Context, config *config) error {
	_ = "STUB: not implemented"
	return nil
}

type clientRPCEvents struct {
	Methods []string `json:"methods,omitempty"`

	Exclude bool `json:"exclude,omitempty"`

	MaxMetadataBytes int `json:"max_metadata_bytes"`

	MaxMessageBytes int `json:"max_message_bytes"`
}

type serverRPCEvents struct {
	Methods []string `json:"methods,omitempty"`

	Exclude bool `json:"exclude,omitempty"`

	MaxMetadataBytes int `json:"max_metadata_bytes"`

	MaxMessageBytes int `json:"max_message_bytes"`
}

type cloudLogging struct {
	ClientRPCEvents []clientRPCEvents `json:"client_rpc_events,omitempty"`

	ServerRPCEvents []serverRPCEvents `json:"server_rpc_events,omitempty"`
}

type cloudMonitoring struct{}

type cloudTrace struct {
	SamplingRate float64 `json:"sampling_rate,omitempty"`
}

type config struct {
	ProjectID string `json:"project_id,omitempty"`

	CloudLogging *cloudLogging `json:"cloud_logging,omitempty"`

	CloudMonitoring *cloudMonitoring `json:"cloud_monitoring,omitempty"`

	CloudTrace *cloudTrace `json:"cloud_trace,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`
}
