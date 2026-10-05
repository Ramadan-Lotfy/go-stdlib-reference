# context

## Overview
`context` carries cancellation signals, deadlines and request-scoped values
across API boundaries and between goroutines. By convention the first parameter
of any function that does I/O or may block is `ctx context.Context`. When the
context is cancelled, everything working for it should stop and release resources.

## Core API

| Function / Type | Purpose |
|---|---|
| `context.Background()`, `context.TODO()` | Root contexts |
| `context.WithCancel` | Returns a child and a `cancel` function |
| `context.WithTimeout`, `WithDeadline` | Cancel automatically after a duration or at a time |
| `context.WithValue` | Attach request-scoped data (use sparingly) |
| `context.WithCancelCause`, `context.Cause` | Cancel with a reason |
| `context.AfterFunc` | Run a function after a context is done |
| `context.WithoutCancel` | Keep values but ignore the parent's cancellation |
| `ctx.Done()`, `ctx.Err()` | Channel closed on cancel; `Canceled` or `DeadlineExceeded` |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./03-concurrency/context/
```

## Gotchas
- **Always call `cancel`**, usually `defer cancel()`, even when the context times out. Otherwise timers and resources leak. `go vet` checks this.
- **Pass `ctx` as the first parameter**, named `ctx`. Never store it in a struct field (exceptions are rare).
- **Never pass a nil context.** Use `context.TODO()` when unsure.
- **Cancelling a context does not stop a goroutine by itself.** The goroutine must watch `ctx.Done()` in a `select` and return.
- **Cancelling a parent cancels all of its children**, never the other way around.
- **`WithValue` is for request-scoped data** (trace IDs, auth info), not for optional function arguments. Use an unexported key type to avoid collisions.
- **Check `ctx.Err()` against sentinels** with `errors.Is(err, context.DeadlineExceeded)`.
- **Use `context.WithoutCancel`** for work that must finish after the request ends, such as writing an audit log.

## Java comparison

| Java | Go |
|---|---|
| `Future.cancel`, `Thread.interrupt` | `cancel()` + `<-ctx.Done()` |
| `CompletableFuture.orTimeout` | `context.WithTimeout` |
| `ThreadLocal`, MDC | `context.WithValue` |
| Spring request attributes | values on `r.Context()` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/context
- Source: https://cs.opensource.google/go/go/+/master:src/context/
