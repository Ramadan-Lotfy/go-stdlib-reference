# os/signal exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/signal/exercises/...
```

## Exercise: graceful shutdown

Write:

```go
func Serve(ctx context.Context, ln net.Listener, h http.Handler) error
```

It serves `h` on `ln` until `ctx` is cancelled, then shuts down gracefully.

Rules:

- Run `srv.Serve(ln)` in a goroutine and wait for either its error or `ctx.Done()`
- On cancel, call `srv.Shutdown` with a **fresh** context that has a 5 second
  timeout. It must not be derived from the cancelled `ctx` (use
  `context.WithoutCancel`), or shutdown would give up instantly
- In-flight requests must be allowed to finish; the test proves it
- Return `nil` after a clean shutdown (`http.ErrServerClosed` is not a failure);
  return the server's own error if it fails first

Hints: `http.Server`, `Shutdown`, `errors.Is(err, http.ErrServerClosed)`.

In a real `main`, build `ctx` with
`signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)`.

## Bonus
Add a `/healthz` endpoint that starts returning 503 as soon as shutdown begins
(use `srv.RegisterOnShutdown`), so a load balancer can drain traffic first.
