package weightedroundrobin

import (
	iserviceconfig "google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/serviceconfig"
)

type lbConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`

	EnableOOBLoadReport bool `json:"enableOobLoadReport,omitempty"`

	OOBReportingPeriod iserviceconfig.Duration `json:"oobReportingPeriod,omitempty"`

	BlackoutPeriod iserviceconfig.Duration `json:"blackoutPeriod,omitempty"`

	WeightExpirationPeriod iserviceconfig.Duration `json:"weightExpirationPeriod,omitempty"`

	WeightUpdatePeriod iserviceconfig.Duration `json:"weightUpdatePeriod,omitempty"`

	ErrorUtilizationPenalty float64 `json:"errorUtilizationPenalty,omitempty"`
}
