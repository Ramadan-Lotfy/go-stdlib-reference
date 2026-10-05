# os/exec exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/osexec/exercises/...
```

## Exercise: Run with a timeout

Write:

```go
func Run(ctx context.Context, timeout time.Duration, name string, args ...string) (string, error)
```

It runs a command and returns its combined stdout and stderr, trimmed.

Rules:

- Kill the command after `timeout` (or when `ctx` ends). The returned error must
  wrap the context error: `errors.Is(err, context.DeadlineExceeded)`
- On a non-zero exit, return an error that wraps the `*exec.ExitError`
  (`errors.As` works) and includes the command's output in the message; also
  return the output
- A program that does not exist returns an error wrapping `exec.ErrNotFound`
- Set `cmd.WaitDelay` so a stuck process cannot block forever

The tests use the helper-process technique: `TestMain` makes the test binary act
as a child command when `OSEXEC_HELPER=1`, so no external tools are needed.

Hints: `context.WithTimeout`, `exec.CommandContext`, `cmd.CombinedOutput()`, check
`ctx.Err()` **before** the command error (a killed process also returns an error).

## Bonus
Add a `RunWithStdin` variant, and a way to stream output line by line to a
callback while the command is running.
