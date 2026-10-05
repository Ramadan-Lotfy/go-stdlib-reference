# runtime and runtime/debug exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/runtime/exercises/...
```

## Exercise 1: LeakCheck

Write `LeakCheck(f func()) (leaked int)` that runs `f` and returns how many
goroutines are still running afterwards.

Rules:

- Count with `runtime.NumGoroutine()` before and after
- Goroutines may need a moment to exit, so wait up to 200 ms for the count to
  fall back to the baseline
- Never return a negative number

Use it in tests to assert that your code does not leak goroutines.

## Exercise 2: CallerInfo

Write `CallerInfo() string` that returns `"file.go:123"` for the line of **the
code that called** `CallerInfo` (not for `CallerInfo` itself).

Hints: `runtime.Caller(1)`, `filepath.Base`.

## Bonus
Write `Where() string` that returns the name of the calling function (use
`runtime.Caller` and `runtime.FuncForPC`), and a `Trace()` that prints the whole
call chain with `runtime.Callers` and `runtime.CallersFrames`.
