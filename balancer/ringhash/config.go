package ringhash

import (
	"encoding/json"

	iringhash "google.golang.org/grpc/internal/ringhash"
)

const (
	defaultMinSize         = 1024
	defaultMaxSize         = 4096
	ringHashSizeUpperBound = 8 * 1024 * 1024
)

func parseConfig(c json.RawMessage) (*iringhash.LBConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
