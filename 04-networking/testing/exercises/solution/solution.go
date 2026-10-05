// Package solution holds the reference solutions for the testing exercises.
package solution

import (
	"errors"
	"slices"
)

// ErrEmpty is returned by Median for an empty input.
var ErrEmpty = errors.New("empty input")

// Clamp limits x to the range [lo, hi]. It panics if lo > hi.
func Clamp(x, lo, hi int) int {
	if lo > hi {
		panic("Clamp: lo > hi")
	}
	return min(max(x, lo), hi)
}

// Median returns the median of xs without modifying xs. For an even number of
// values it is the mean of the two middle ones.
func Median(xs []float64) (float64, error) {
	if len(xs) == 0 {
		return 0, ErrEmpty
	}
	s := slices.Clone(xs)
	slices.Sort(s)
	mid := len(s) / 2
	if len(s)%2 == 1 {
		return s[mid], nil
	}
	return (s[mid-1] + s[mid]) / 2, nil
}
