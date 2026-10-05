# sync/atomic exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./03-concurrency/atomic/exercises/...
```

## Exercise 1: Stats

Write a `Stats` type with lock-free counters:

```go
func (s *Stats) Record(failed bool)
func (s *Stats) Snapshot() (requests, failures int64)
```

`Record` counts a request, and a failure when `failed` is true.

Hints: two `atomic.Int64` fields. Run with `go test -race`.

## Exercise 2: ConfigStore

Write a `ConfigStore` holding an immutable `Config` behind an `atomic.Pointer`:

```go
func NewConfigStore(c Config) *ConfigStore
func (s *ConfigStore) Get() Config
func (s *ConfigStore) Update(f func(Config) Config)
```

`Update` must be safe when 100 goroutines each increment `Limit` at once, so
the final value is exactly 100.

Hints: copy the config, apply `f`, and retry with `CompareAndSwap` until it succeeds.

## Bonus
Rewrite `ConfigStore` with a `sync.RWMutex`. Which version is simpler? Which is
faster under heavy reads?
