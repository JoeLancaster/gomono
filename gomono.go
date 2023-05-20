package gomono

import "time"

func init() {
	if !checkVdsoAvailable() {
		panic("No VDSO clock_gettime available")
	}
}

// Monotonic is a type that represents a measurement from the OS's monotonic clock
// on its own, it is not useful for telling the time, only measuring durations
type Monotonic int64

// Now gets the elapsed time according to clock_gettime(CLOCK_MONOTONIC)
func Now() Monotonic {
	return Monotonic(getMono())

}

// Before returns true when t happened before u, false otherwise
func (t Monotonic) Before(u Monotonic) bool {
	return t < u
}

// After returns true when t happened after u, false otherwise
func (t Monotonic) After(u Monotonic) bool {
	return t > u
}

// Diff returns the absolute duration between two clock measurements
func (t Monotonic) Diff(u Monotonic) time.Duration {
	if t > u {
		return time.Duration(t - u)
	}
	return time.Duration(u - t)

}
