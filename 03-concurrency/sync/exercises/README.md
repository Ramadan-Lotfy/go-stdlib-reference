# sync exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./03-concurrency/sync/exercises/...
```

## Exercise 1: SafeCounter

Write a concurrency-safe counter of string keys:

```go
func NewSafeCounter() *SafeCounter
func (c *SafeCounter) Inc(key string)
func (c *SafeCounter) Value(key string) int
```

Run the tests with the race detector: `go test -race ./03-concurrency/sync/exercises/...`.

Hints: a `map[string]int` and a `sync.Mutex`; `defer Unlock`.

## Exercise 2: ParallelMap

Write a generic `ParallelMap[T, R any](in []T, f func(T) R) []R` that runs `f`
on every element in its own goroutine and returns the results **in input order**.

Hints: preallocate `out := make([]R, len(in))`, have each goroutine write only
its own index, and use a `sync.WaitGroup`.

## Bonus
Limit the number of goroutines running at once to N, using a buffered channel as
a semaphore.
