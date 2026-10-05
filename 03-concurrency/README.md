# Stage 3: Concurrency

Goroutines' toolbox: locks, atomics, contexts, channels and patterns.

| Package | What it covers |
|---------|----------------|
| [sync](sync) | `sync` provides the classic shared-memory primitives: mutexes, wait groups, one-time initialization, condition variables, a concurrent map and an object pool. |
| [sync/atomic](atomic) | `sync/atomic` provides lock-free operations on single values: counters, flags and pointers. |
| [context](context) | `context` carries cancellation signals, deadlines and request-scoped values across API boundaries and between goroutines. |
| [channels, select and concurrency patterns](channels) | Channels are typed pipes between goroutines, and `select` waits on several of them at once. |

Each folder has a `README.md` (concepts, gotchas, Java comparison), runnable `example_test.go`, and `exercises/` with a reference solution.

Run this stage: `go test ./03-concurrency/...`
