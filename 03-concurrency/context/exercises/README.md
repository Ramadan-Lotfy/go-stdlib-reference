# context exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./03-concurrency/context/exercises/...
```

## Exercise 1: Sleep

Write `Sleep(ctx context.Context, d time.Duration) error` that waits for `d`,
but returns `ctx.Err()` immediately if the context ends first. Unlike
`time.Sleep`, it can be cancelled.

Hints: `time.NewTimer`, `defer t.Stop()`, `select` on `t.C` and `ctx.Done()`.

## Exercise 2: FirstResult

Write:

```go
func FirstResult(ctx context.Context, fns ...func(context.Context) (string, error)) (string, error)
```

It runs all functions concurrently and returns the first **successful** result.

Rules:

- Cancel the remaining functions once a result is found
- If every function fails, return all errors joined (`errors.Is` must still work)
- If `ctx` ends first, return `ctx.Err()`
- With no functions, return an error
- Never leak a goroutine blocked on a send

Hints: derive a cancellable child context, use a **buffered** result channel of
size `len(fns)`, and `errors.Join`.

## Bonus
Add a per-call timeout: each function gets at most `perCall` before it is
considered failed.
