package health

import (
	"context"
	"sync"

	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	maxAllowedServices = 100
)

type Server struct {
	healthgrpc.UnimplementedHealthServer
	mu sync.RWMutex

	shutdown bool

	statusMap map[string]healthpb.HealthCheckResponse_ServingStatus
	updates   map[string]map[healthgrpc.Health_WatchServer]chan healthpb.HealthCheckResponse_ServingStatus
}

func NewServer() *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Check(_ context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) List(_ context.Context, _ *healthpb.HealthListRequest) (*healthpb.HealthListResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) Watch(in *healthpb.HealthCheckRequest, stream healthgrpc.Health_WatchServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) SetServingStatus(service string, servingStatus healthpb.HealthCheckResponse_ServingStatus) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) setServingStatusLocked(service string, servingStatus healthpb.HealthCheckResponse_ServingStatus) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) Shutdown() { _ = "STUB: not implemented"; return }

func (s *Server) Resume() { _ = "STUB: not implemented"; return }
