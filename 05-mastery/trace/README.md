# runtime/trace

## Overview
`runtime/trace` records an **execution trace**: a timeline of everything the
runtime does (goroutine creation, blocking, scheduling, garbage collection,
system calls) with nanosecond timestamps. Where a profile tells you *where
time goes in aggregate*, a trace tells you *what happened, in order*: why a
request was slow, whether goroutines wait on each other, how well all CPUs are
used. You can add your own **tasks**, **regions** and **log** messages.

```bash
go test -trace trace.out ./...      # or write one with trace.Start in a program
go tool trace trace.out             # opens a browser UI: timeline, goroutine analysis
```

## Core API

| Function / Type | Purpose |
|---|---|
| `trace.Start(w)` / `trace.Stop()` | Begin and end writing a trace to `w` |
| `trace.NewTask(ctx, name)` / `task.End()` | A logical operation that may span goroutines (a request) |
| `trace.WithRegion(ctx, name, f)` | Mark a span of code inside a task |
| `trace.StartRegion(ctx, name)` / `region.End()` | Same, without a closure |
| `trace.Log(ctx, category, message)` | Add a message to the timeline |
| `trace.IsEnabled()` | Is tracing active right now? |
| `go tool trace` | The viewer |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/trace/
```

## Gotchas
- **Traces are large and slow the program down.** Capture a few seconds of a representative workload, not minutes.
- **Only one trace can run at a time,** and `trace.Start` returns an error if one is already active (for example under `go test -trace`).
- **Regions, tasks and logs cost almost nothing when tracing is off,** so it is fine to leave them in production code.
- **Use a trace when a profile does not explain latency:** waiting, lock contention, scheduler delays and idle CPUs do not show up in CPU profiles.
- **Pass the task's `ctx` down the call chain,** so regions in other goroutines attach to the same task.
- **Always call `trace.Stop()` before reading the buffer,** or the trace is truncated.
- **`go tool trace` needs a browser** and the same Go version that created the trace (or newer).

## Java comparison

| Java | Go |
|---|---|
| JFR events | `trace.Log` / regions + `go tool trace` |
| OpenTelemetry spans | `trace.NewTask` is local to one process; use OpenTelemetry for distributed tracing |
| Thread timeline views | the goroutine and processor timelines in `go tool trace` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/runtime/trace
- Source: https://cs.opensource.google/go/go/+/master:src/runtime/trace/
