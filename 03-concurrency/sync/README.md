# sync

## Overview
`sync` provides the classic shared-memory primitives: mutexes, wait groups,
one-time initialization, condition variables, a concurrent map and an object
pool. Go's motto is "share memory by communicating" (channels), but a mutex is
often the simplest correct tool for protecting a small piece of state.

## Core API

| Function / Type | Purpose |
|---|---|
| `sync.Mutex`, `sync.RWMutex` | Mutual exclusion; `RWMutex` allows many readers or one writer |
| `sync.WaitGroup` | Wait for a set of goroutines to finish (`Add`, `Done`, `Wait`) |
| `sync.Once`, `OnceFunc`, `OnceValue` | Run initialization exactly once |
| `sync.Cond` | Wait for and signal a condition |
| `sync.Map` | Concurrent map for append-mostly or disjoint-key workloads |
| `sync.Pool` | Reuse temporary objects to reduce allocation |
| `go test -race` | The race detector: always run it |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./03-concurrency/sync/
```

## Gotchas
- **Always `defer mu.Unlock()`** right after `Lock()`, so every return path unlocks.
- **Never copy a Mutex or WaitGroup** (or a struct containing one). Pass pointers. `go vet` catches this.
- **Call `wg.Add` *before* starting the goroutine**, not inside it, or `Wait` may return too early.
- **A `Mutex` is not reentrant.** Locking twice in the same goroutine deadlocks.
- **Do not hold a lock while doing I/O or calling unknown code.** Keep critical sections tiny.
- **Prefer a plain `map` + `Mutex` over `sync.Map`** unless you measured a need; `sync.Map` has no type safety.
- **`sync.Pool` may drop items at any time** (every GC). It is a cache, not a free list. Reset objects before reuse.
- **`Cond.Wait` must be called in a loop** that rechecks the condition, with the lock held.
- **Run tests with `-race`.** Data races are undefined behavior and often invisible without it.

## Java comparison

| Java | Go |
|---|---|
| `synchronized`, `ReentrantLock` | `sync.Mutex` (not reentrant) |
| `ReadWriteLock` | `sync.RWMutex` |
| `CountDownLatch` | `sync.WaitGroup` |
| `ConcurrentHashMap` | `map` + `sync.Mutex`, or `sync.Map` |
| `Condition` | `sync.Cond` |
| Holder/lazy singleton | `sync.Once` / `sync.OnceValue` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/sync
- Source: https://cs.opensource.google/go/go/+/master:src/sync/
