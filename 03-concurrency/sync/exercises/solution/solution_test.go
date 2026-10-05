package solution

import (
	"slices"
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	c := NewSafeCounter()
	var wg sync.WaitGroup
	for range 200 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("a")
			c.Inc("b")
		}()
	}
	wg.Wait()
	if c.Value("a") != 200 || c.Value("b") != 200 || c.Value("zzz") != 0 {
		t.Errorf("a=%d b=%d zzz=%d", c.Value("a"), c.Value("b"), c.Value("zzz"))
	}
}

func TestParallelMap(t *testing.T) {
	got := ParallelMap([]int{1, 2, 3, 4}, func(x int) string { return string(rune('a' + x)) })
	if want := []string{"b", "c", "d", "e"}; !slices.Equal(got, want) {
		t.Errorf("ParallelMap() = %v, want %v", got, want)
	}
	if got := ParallelMap([]int{}, func(x int) int { return x }); len(got) != 0 {
		t.Errorf("empty input gave %v", got)
	}
}
