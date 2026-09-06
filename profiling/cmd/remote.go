package main

import (
	"context"

	ppb "google.golang.org/grpc/profiling/proto"
)

func setEnabled(ctx context.Context, c ppb.ProfilingClient, enabled bool) error {
	_ = "STUB: not implemented"
	return nil
}

func retrieveSnapshot(ctx context.Context, c ppb.ProfilingClient, f string) error {
	_ = "STUB: not implemented"
	return nil
}

func remoteCommand() error { _ = "STUB: not implemented"; return nil }
