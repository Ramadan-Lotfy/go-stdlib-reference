// Package solution holds the reference solutions for the slices exercises.
package solution

import "slices"

// Unique returns the sorted distinct values of s. It does not modify s.
func Unique(s []int) []int {
	out := slices.Clone(s)
	slices.Sort(out)
	return slices.Compact(out)
}

// Chunk splits s into consecutive pieces of at most n elements.
// n <= 0 returns nil.
func Chunk[T any](s []T, n int) [][]T {
	if n <= 0 {
		return nil
	}
	var out [][]T
	for len(s) > 0 {
		size := min(n, len(s))
		out = append(out, s[:size:size])
		s = s[size:]
	}
	return out
}
