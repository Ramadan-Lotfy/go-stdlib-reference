// Package solution holds the reference solutions for the generics, cmp and slices exercises.
package solution

import "cmp"

// Map applies f to every element.
func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

// Filter returns the elements for which keep returns true.
func Filter[T any](s []T, keep func(T) bool) []T {
	out := make([]T, 0)
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

// Reduce folds s into a single value, starting from init.
func Reduce[T, A any](s []T, init A, f func(A, T) A) A {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

// Stack is a last-in, first-out collection.
type Stack[T any] struct{ items []T }

// Push adds v on top.
func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

// Pop removes and returns the top element; ok is false when the stack is empty.
func (s *Stack[T]) Pop() (v T, ok bool) {
	if len(s.items) == 0 {
		return v, false
	}
	v = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

// Len returns the number of elements.
func (s *Stack[T]) Len() int { return len(s.items) }

// Set is a collection of unique values.
type Set[T comparable] map[T]struct{}

// Add inserts v.
func (s Set[T]) Add(v T) { s[v] = struct{}{} }

// Has reports whether v is in the set.
func (s Set[T]) Has(v T) bool { _, ok := s[v]; return ok }

// Union returns a new set with the elements of both sets.
func Union[T comparable](a, b Set[T]) Set[T] {
	out := make(Set[T], len(a)+len(b))
	for v := range a {
		out.Add(v)
	}
	for v := range b {
		out.Add(v)
	}
	return out
}

// MaxBy returns the element of s with the largest key and false if s is empty.
func MaxBy[T any, K cmp.Ordered](s []T, key func(T) K) (best T, ok bool) {
	for i, v := range s {
		if i == 0 || key(v) > key(best) {
			best, ok = v, true
		}
	}
	return best, ok
}
