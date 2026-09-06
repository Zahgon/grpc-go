package health

import (
	"context"
	"time"

	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/internal/backoff"
)

var (
	backoffStrategy = backoff.DefaultExponential
	backoffFunc     = func(ctx context.Context, retries int) bool {
		d := backoffStrategy.Backoff(retries)
		timer := time.NewTimer(d)
		select {
		case <-timer.C:
			return true
		case <-ctx.Done():
			timer.Stop()
			return false
		}
	}
)

func init() {
	internal.HealthCheckFunc = clientHealthCheck
}

const healthCheckMethod = "/grpc.health.v1.Health/Watch"

func clientHealthCheck(ctx context.Context, newStream func(string) (any, error), setConnectivityState func(connectivity.State, error), service string) error {
	_ = "STUB: not implemented"
	return nil
}
