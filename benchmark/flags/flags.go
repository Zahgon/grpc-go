package flags

import (
	"time"
)

type stringFlagWithAllowedValues struct {
	val     string
	allowed []string
}

func StringWithAllowedValues(name, defaultVal, usage string, allowed []string) *string {
	_ = "STUB: not implemented"
	return nil
}

func (as *stringFlagWithAllowedValues) String() string { _ = "STUB: not implemented"; return "" }

func (as *stringFlagWithAllowedValues) Set(val string) error { _ = "STUB: not implemented"; return nil }

type durationSliceValue []time.Duration

func DurationSlice(name string, defaultVal []time.Duration, usage string) *[]time.Duration {
	_ = "STUB: not implemented"
	return nil
}

func (dsv *durationSliceValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (dsv *durationSliceValue) String() string { _ = "STUB: not implemented"; return "" }

type intSliceValue []int

func IntSlice(name string, defaultVal []int, usage string) *[]int {
	_ = "STUB: not implemented"
	return nil
}

func (isv *intSliceValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (isv *intSliceValue) String() string { _ = "STUB: not implemented"; return "" }

type stringSliceValue []string

func StringSlice(name string, defaultVal []string, usage string) *[]string {
	_ = "STUB: not implemented"
	return nil
}

func escapedCommaSplit(str string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (ss *stringSliceValue) Set(str string) error { _ = "STUB: not implemented"; return nil }

func (ss *stringSliceValue) String() string { _ = "STUB: not implemented"; return "" }
