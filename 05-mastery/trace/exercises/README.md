# runtime/trace exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/trace/exercises/...
```

## Exercise: TraceTo

Write:

```go
func TraceTo(w io.Writer, f func(ctx context.Context)) error
```

It must:

- start tracing into `w` (return the error from `trace.Start` if it fails)
- create a task named `"work"` and call `f` with the task's context
- end the task and stop tracing before returning, even if `f` panics

Hints: `trace.Start`, `trace.NewTask`, `defer task.End()`, `defer trace.Stop()`.
Defers run last-in first-out, so register `Stop` first.

Then run the code from a small program, save the output to `trace.out`, and open
it with `go tool trace trace.out`. Find your task under "User-defined tasks".

## Bonus
Write a program with two goroutines contending for one `sync.Mutex`. Use the
trace viewer's "Synchronization blocking profile" to see the waiting.
