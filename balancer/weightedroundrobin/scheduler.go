package weightedroundrobin

import (
	"math"
)

type scheduler interface {
	nextIndex() int
}

func (p *picker) newScheduler(recordMetrics bool) scheduler {
	_ = "STUB: not implemented"
	return *new(scheduler)
}

const maxWeight = math.MaxUint16

type edfScheduler struct {
	inc     func() uint32
	weights []uint16
}

func (s *edfScheduler) nextIndex() int { _ = "STUB: not implemented"; return 0 }

type rrScheduler struct {
	inc    func() uint32
	numSCs uint32
}

func (s *rrScheduler) nextIndex() int { _ = "STUB: not implemented"; return 0 }
