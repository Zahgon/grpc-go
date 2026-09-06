package main

import (
	ppb "google.golang.org/grpc/profiling/proto"
)

type jsonNode struct {
	Name      string  `json:"name"`
	Cat       string  `json:"cat"`
	ID        string  `json:"id"`
	Cname     string  `json:"cname"`
	Phase     string  `json:"ph"`
	Timestamp float64 `json:"ts"`
	PID       string  `json:"pid"`
	TID       string  `json:"tid"`
}

func hashCname(tag string) string { _ = "STUB: not implemented"; return "" }

func filterCounter(stat *ppb.Stat, filter string, counter int) int {
	_ = "STUB: not implemented"
	return 0
}

type counter struct {
	c map[string]int
}

func newCounter() *counter { _ = "STUB: not implemented"; return nil }

func (c *counter) GetAndInc(s string) int { _ = "STUB: not implemented"; return 0 }

func catapultNs(sec int64, nsec int32) float64 { _ = "STUB: not implemented"; return 0 }

func streamStatsCatapultJSONSingle(stat *ppb.Stat, baseSec int64, baseNsec int32) []jsonNode {
	_ = "STUB: not implemented"
	return nil
}

func timerBeginIsBefore(ti *ppb.Timer, tj *ppb.Timer) bool { _ = "STUB: not implemented"; return false }

func streamStatsCatapultJSON(s *snapshot, streamStatsCatapultJSONFileName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
