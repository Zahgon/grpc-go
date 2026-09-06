package grpc

import (
	"encoding/json"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal"
	internalserviceconfig "google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/serviceconfig"
)

const maxInt = int(^uint(0) >> 1)

type MethodConfig = internalserviceconfig.MethodConfig

type ServiceConfig struct {
	serviceconfig.Config

	lbConfig serviceconfig.LoadBalancingConfig

	Methods map[string]MethodConfig

	retryThrottling *retryThrottlingPolicy

	healthCheckConfig *healthCheckConfig

	rawJSONString string
}

type healthCheckConfig struct {
	ServiceName string
}

type jsonRetryPolicy struct {
	MaxAttempts          int
	InitialBackoff       internalserviceconfig.Duration
	MaxBackoff           internalserviceconfig.Duration
	BackoffMultiplier    float64
	RetryableStatusCodes []codes.Code
}

type retryThrottlingPolicy struct {
	MaxTokens float64

	TokenRatio float64
}

type jsonName struct {
	Service string
	Method  string
}

var (
	errDuplicatedName             = errors.New("duplicated name")
	errEmptyServiceNonEmptyMethod = errors.New("cannot combine empty 'service' and non-empty 'method'")
)

func (j jsonName) generatePath() (string, error) { _ = "STUB: not implemented"; return "", nil }

type jsonMC struct {
	Name                    *[]jsonName
	WaitForReady            *bool
	Timeout                 *internalserviceconfig.Duration
	MaxRequestMessageBytes  *int64
	MaxResponseMessageBytes *int64
	RetryPolicy             *jsonRetryPolicy
}

type jsonSC struct {
	LoadBalancingPolicy *string
	LoadBalancingConfig *json.RawMessage
	MethodConfig        *[]jsonMC
	RetryThrottling     *retryThrottlingPolicy
	HealthCheckConfig   *healthCheckConfig
}

func init() {
	internal.ParseServiceConfig = func(js string) *serviceconfig.ParseResult {
		return parseServiceConfig(js, defaultMaxCallAttempts)
	}
}

func parseServiceConfig(js string, maxAttempts int) *serviceconfig.ParseResult {
	_ = "STUB: not implemented"
	return nil
}

func isValidRetryPolicy(jrp *jsonRetryPolicy) bool { _ = "STUB: not implemented"; return false }

func convertRetryPolicy(jrp *jsonRetryPolicy, maxAttempts int) (p *internalserviceconfig.RetryPolicy, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func minPointers(a, b *int) *int { _ = "STUB: not implemented"; return nil }

func getMaxSize(mcMax, doptMax *int, defaultVal int) *int { _ = "STUB: not implemented"; return nil }

func newInt(b int) *int { _ = "STUB: not implemented"; return nil }

func init() {
	internal.EqualServiceConfigForTesting = equalServiceConfig
}

func equalServiceConfig(a, b serviceconfig.Config) bool { _ = "STUB: not implemented"; return false }
