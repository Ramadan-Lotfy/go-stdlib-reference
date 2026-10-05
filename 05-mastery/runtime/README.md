# runtime and runtime/debug

## Overview
`runtime` is the interface to Go's run-time system: the scheduler, garbage
collector and memory allocator. It answers practical questions (how many
goroutines are running? how much memory am I using? where was this called
from?) and exposes knobs such as `GOMAXPROCS`. Go's scheduler multiplexes many
goroutines (G) onto a few OS threads (M) using logical processors (P): that G-M-P
model is why goroutines are cheap and why blocking calls do not block the world.

## Core API

| Function / Type | Purpose |
|---|---|
| `runtime.NumGoroutine()` | Goroutines currently alive (great for leak checks) |
| `runtime.GOMAXPROCS(n)`, `runtime.NumCPU()` | Logical processors used; `GOMAXPROCS(0)` reads the value |
| `runtime.ReadMemStats(&m)` | Heap size, total allocations, GC count |
| `runtime.GC()` | Force a collection (for tests and diagnostics only) |
| `runtime.Caller(skip)`, `runtime.Callers`, `FuncForPC` | Who called me? File, line, function name |
| `runtime.Stack(buf, all)`, `debug.Stack()` | Goroutine stack traces |
| `runtime.Gosched()`, `runtime.Goexit()` | Yield the processor; end the goroutine after running defers |
| `runtime.GOOS`, `GOARCH`, `Version()` | Platform information |
| `debug.SetGCPercent`, `debug.SetMemoryLimit` | Tune the garbage collector |
| `debug.ReadBuildInfo()` | Module and VCS information embedded in the binary |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/runtime/
```

## Gotchas
- **`runtime.GC()` and `ReadMemStats` are for diagnostics,** not control flow. `ReadMemStats` briefly stops the world.
- **Do not use `Gosched` for synchronization.** Use channels, mutexes or `WaitGroup`; scheduling order is not guaranteed.
- **Goroutine counts settle asynchronously.** A goroutine that has returned may still be counted for a moment, so poll with a deadline in tests.
- **`GOMAXPROCS` follows the host CPU count, not the container limit,** in Go versions before 1.25. In Kubernetes set it from the CPU limit (or use `automaxprocs`) to avoid CPU throttling.
- **`GOGC` and `GOMEMLIMIT`** (or `debug.SetGCPercent` / `SetMemoryLimit`) are the GC knobs. Setting a memory limit below the live heap makes the GC run constantly.
- **`runtime.Caller(skip)` counts stack frames:** `0` is the function calling `Caller`, `1` is its caller. Wrapping the call in another helper changes the right value of `skip`. Inlining does not change the result.
- **Goroutine leaks are the most common production bug:** a goroutine blocked forever on a channel or lock is never freed. Watch `runtime.NumGoroutine()` as a metric.
- **Finalizers (`SetFinalizer`) are not destructors:** they may run late or never. Use `defer` and `Close`.
- **`runtime.KeepAlive(x)`** keeps `x` alive until that point; needed with `unsafe` or cgo.

## Java comparison

| Java | Go |
|---|---|
| `Thread.activeCount()` | `runtime.NumGoroutine()` |
| `Runtime.availableProcessors()` | `runtime.NumCPU()` |
| `Runtime.totalMemory()/freeMemory()` | `runtime.ReadMemStats` |
| `System.gc()` | `runtime.GC()` |
| `Thread.dumpStack()` | `debug.PrintStack()` / `runtime.Stack` |
| `StackWalker` | `runtime.Caller` / `runtime.Callers` |
| JVM flags `-Xmx` | `GOMEMLIMIT`, `GOGC` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/runtime and runtime/debug
- Source: https://cs.opensource.google/go/go/+/master:src/runtime and runtime/debug/
