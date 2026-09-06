package googledirectpath

import (
	"bytes"
	"strings"
	"sync"
	"time"
)

func getFromMetadata(timeout time.Duration, urlStr string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	zone     string
	zoneOnce sync.Once
)

var getZone = func(timeout time.Duration) string {
	zoneOnce.Do(func() {
		qualifiedZone, err := getFromMetadata(timeout, zoneURL)
		if err != nil {
			logger.Warningf("could not discover instance zone: %v", err)
			return
		}
		i := bytes.LastIndexByte(qualifiedZone, '/')
		if i == -1 {
			logger.Warningf("could not parse zone from metadata server: %s", qualifiedZone)
			return
		}
		zone = string(qualifiedZone[i+1:])
	})
	return zone
}

var (
	ipv6Capable     bool
	ipv6CapableOnce sync.Once
)

var getIPv6Capable = func(timeout time.Duration) bool {
	ipv6CapableOnce.Do(func() {
		addr, err := getFromMetadata(timeout, ipv6URL)
		if err != nil {
			logger.Warningf("could not discover ipv6 capability: %v", err)
			return
		}
		if trimmedAddr := strings.TrimSpace(string(addr)); trimmedAddr == "" {
			logger.Warningf("metadata server returned empty ipv6 address")
			return
		}
		ipv6Capable = true
	})
	return ipv6Capable
}
