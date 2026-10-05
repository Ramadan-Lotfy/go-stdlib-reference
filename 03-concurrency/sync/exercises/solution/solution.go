// Package solution holds the reference solutions for the sync exercises.
package solution

import "sync"

// SafeCounter counts occurrences of string keys and is safe for concurrent use.
type SafeCounter struct {
	mu sync.Mutex
	m  map[string]int
}

// NewSafeCounter returns an empty counter.
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{m: make(map[string]int)}
}

// Inc adds one to key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key]++
}

// Value returns the current count for key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.m[key]
}

// ParallelMap applies f to every element concurrently and returns the results
// in the same order as in.
func ParallelMap[T, R any](in []T, f func(T) R) []R {
	out := make([]R, len(in))
	var wg sync.WaitGroup
	for i, v := range in {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out[i] = f(v)
		}()
	}
	wg.Wait()
	return out
}
