package orca

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/internal"
	ointernal "google.golang.org/grpc/orca/internal"

	v3orcaservicegrpc "github.com/cncf/xds/go/xds/service/orca/v3"
	v3orcaservicepb "github.com/cncf/xds/go/xds/service/orca/v3"
)

func init() {
	ointernal.AllowAnyMinReportingInterval = func(so *ServiceOptions) {
		so.allowAnyMinReportingInterval = true
	}
	internal.ORCAAllowAnyMinReportingInterval = ointernal.AllowAnyMinReportingInterval
}

const minReportingInterval = 30 * time.Second

type Service struct {
	v3orcaservicegrpc.UnimplementedOpenRcaServiceServer

	minReportingInterval time.Duration

	smProvider ServerMetricsProvider
}

type ServiceOptions struct {
	ServerMetricsProvider ServerMetricsProvider

	MinReportingInterval time.Duration

	allowAnyMinReportingInterval bool
}

type ServerMetricsProvider interface {
	ServerMetrics() *ServerMetrics
}

func NewService(opts ServiceOptions) (*Service, error) { _ = "STUB: not implemented"; return nil, nil }

func Register(s grpc.ServiceRegistrar, opts ServiceOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) determineReportingInterval(req *v3orcaservicepb.OrcaLoadReportRequest) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (s *Service) sendMetricsResponse(stream v3orcaservicegrpc.OpenRcaService_StreamCoreMetricsServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) StreamCoreMetrics(req *v3orcaservicepb.OrcaLoadReportRequest, stream v3orcaservicegrpc.OpenRcaService_StreamCoreMetricsServer) error {
	_ = "STUB: not implemented"
	return nil
}
