package solution

import (
	"slices"
	"strconv"
	"testing"
)

func TestMapFilterReduce(t *testing.T) {
	got := Map([]int{1, 2, 3}, strconv.Itoa)
	if !slices.Equal(got, []string{"1", "2", "3"}) {
		t.Errorf("Map() = %v", got)
	}
	evens := Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
	if !slices.Equal(evens, []int{2, 4}) {
		t.Errorf("Filter() = %v", evens)
	}
	if none := Filter([]int{1, 3}, func(n int) bool { return n%2 == 0 }); len(none) != 0 || none == nil {
		t.Errorf("Filter(no match) = %#v, want empty non-nil", none)
	}
	if sum := Reduce([]int{1, 2, 3}, 10, func(a, n int) int { return a + n }); sum != 16 {
		t.Errorf("Reduce() = %d, want 16", sum)
	}
}

func TestStack(t *testing.T) {
	var s Stack[string]
	if _, ok := s.Pop(); ok {
		t.Error("Pop on empty stack should report false")
	}
	s.Push("a")
	s.Push("b")
	if s.Len() != 2 {
		t.Errorf("Len() = %d", s.Len())
	}
	if v, ok := s.Pop(); !ok || v != "b" {
		t.Errorf("Pop() = %q, %v", v, ok)
	}
}

func TestSetAndUnion(t *testing.T) {
	a, b := Set[int]{}, Set[int]{}
	a.Add(1)
	a.Add(2)
	b.Add(2)
	b.Add(3)
	u := Union(a, b)
	if len(u) != 3 || !u.Has(1) || !u.Has(3) || u.Has(9) {
		t.Errorf("Union() = %v", u)
	}
	if len(a) != 2 || len(b) != 2 {
		t.Error("Union must not modify its inputs")
	}
}

func TestMaxBy(t *testing.T) {
	words := []string{"go", "gopher", "gc"}
	best, ok := MaxBy(words, func(s string) int { return len(s) })
	if !ok || best != "gopher" {
		t.Errorf("MaxBy() = %q, %v", best, ok)
	}
	if _, ok := MaxBy([]string{}, func(s string) int { return len(s) }); ok {
		t.Error("MaxBy(empty) should report false")
	}
}
