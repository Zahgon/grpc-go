package stats

import (
	"io"
)

type Histogram struct {
	Count int64

	Sum int64

	SumOfSquares int64

	Min int64

	Max int64

	Buckets []HistogramBucket

	opts                          HistogramOptions
	logBaseBucketSize             float64
	oneOverLogOnePlusGrowthFactor float64
}

type HistogramOptions struct {
	NumBuckets int

	GrowthFactor float64

	BaseBucketSize float64

	MinValue int64
}

type HistogramBucket struct {
	LowBound float64

	Count int64
}

func NewHistogram(opts HistogramOptions) *Histogram { _ = "STUB: not implemented"; return nil }

func (h *Histogram) Print(w io.Writer) { _ = "STUB: not implemented"; return }

func (h *Histogram) PrintWithUnit(w io.Writer, unit float64) { _ = "STUB: not implemented"; return }

func (h *Histogram) String() string { _ = "STUB: not implemented"; return "" }

func (h *Histogram) Clear() { _ = "STUB: not implemented"; return }

func (h *Histogram) Opts() HistogramOptions {
	_ = "STUB: not implemented"
	return *new(HistogramOptions)
}

func (h *Histogram) Add(value int64) error { _ = "STUB: not implemented"; return nil }

func (h *Histogram) findBucket(value int64) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (h *Histogram) Merge(h2 *Histogram) { _ = "STUB: not implemented"; return }
