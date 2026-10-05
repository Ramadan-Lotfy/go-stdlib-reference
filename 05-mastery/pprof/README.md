# runtime/pprof and net/http/pprof

## Overview
`runtime/pprof` records profiles of a running program: where CPU time goes,
which code allocates memory, which goroutines exist, and where they block. The
`go tool pprof` command turns a profile into a call graph, flame graph or
ranked list. In services, importing `net/http/pprof` exposes the same data over
HTTP at `/debug/pprof/`.

```bash
go test -bench . -cpuprofile cpu.out -memprofile mem.out
go tool pprof -http=:8080 cpu.out          # interactive web UI with flame graph
go tool pprof http://localhost:6060/debug/pprof/heap
go tool pprof -top -nodecount=15 cpu.out   # quick text summary
```

## Core API

| Function / Type | Purpose |
|---|---|
| `pprof.StartCPUProfile(w)` / `StopCPUProfile()` | Sample CPU usage (about 100 times a second) into `w` |
| `pprof.Lookup("heap").WriteTo(w, 0)` | Write a named profile: `heap`, `allocs`, `goroutine`, `block`, `mutex`, `threadcreate` |
| `pprof.Profiles()` | List all profiles |
| `pprof.Do(ctx, pprof.Labels(k, v), f)` | Attach labels to CPU samples (per handler, per tenant) |
| `pprof.NewProfile(name)` | Custom profile, for example open resources |
| `runtime.SetBlockProfileRate`, `SetMutexProfileFraction` | Enable block and mutex profiles (off by default) |
| `import _ "net/http/pprof"` | Registers `/debug/pprof/*` on `http.DefaultServeMux` |
| `go test -cpuprofile/-memprofile` | Profile tests and benchmarks |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/pprof/
```

## Gotchas
- **Never expose `/debug/pprof` on a public port.** It leaks internals and allows expensive requests. Serve it on a separate localhost or admin listener with its own `ServeMux`.
- **Profile a realistic workload for at least 10 to 30 seconds.** A short CPU profile has too few samples to be meaningful.
- **Only one CPU profile can run at a time.** `StartCPUProfile` returns an error if one is active (for example under `go test -cpuprofile`).
- **Block and mutex profiles are off by default.** Turn them on with `SetBlockProfileRate(1)` / `SetMutexProfileFraction(1)` while investigating, then off again; they cost performance.
- **The heap profile shows allocations as of the last GC** and samples roughly every 512 KB. Use `alloc_space` to see total allocation volume and `inuse_space` for live memory.
- **Measure first, optimize second.** Optimize the top entries in the profile, then re-measure. Intuition about hot spots is often wrong.
- **The `goroutine` profile with `debug=2` is the quickest way to debug a hang or a leak:** it prints every goroutine's stack.
- **Profiles are gzip-compressed protobuf** (`debug=0`); `debug=1` gives readable text.

## Java comparison

| Java | Go |
|---|---|
| JFR / VisualVM / async-profiler | `go tool pprof` + `runtime/pprof` |
| Thread dump (`jstack`) | `/debug/pprof/goroutine?debug=2` |
| Heap dump + MAT | `/debug/pprof/heap` (sampled, not a full dump) |
| Spring Boot Actuator | `net/http/pprof` on an admin port |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/runtime/pprof and net/http/pprof
- Source: https://cs.opensource.google/go/go/+/master:src/runtime/pprof and net/http/pprof/
