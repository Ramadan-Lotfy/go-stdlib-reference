# os/signal

## Overview
`os/signal` lets a program react to operating system signals: `SIGINT`
(Ctrl+C), `SIGTERM` (the polite "please stop" sent by Kubernetes, systemd and
`docker stop`), `SIGHUP` (often "reload config") and others. The standard use is
**graceful shutdown**: on `SIGTERM` stop accepting new work, let in-flight
requests finish, flush buffers, close connections and then exit.

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()

go srv.ListenAndServe()
<-ctx.Done()                         // blocks until a signal arrives
stop()                               // restore default behavior: a second Ctrl+C kills immediately
shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
srv.Shutdown(shutdownCtx)
```

## Core API

| Function / Type | Purpose |
|---|---|
| `signal.NotifyContext(parent, sigs...)` | A context cancelled when a signal arrives (preferred) |
| `signal.Notify(ch, sigs...)` | Deliver signals to a **buffered** channel |
| `signal.Stop(ch)` | Stop delivering to a channel |
| `signal.Ignore`, `signal.Reset` | Ignore or restore default handling |
| `os.Interrupt` (`SIGINT`), `syscall.SIGTERM` | The signals you handle in services |
| `SIGKILL`, `SIGSTOP` | Cannot be caught or ignored |
| `srv.Shutdown(ctx)` | Graceful HTTP server shutdown |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/signal/
```

## Gotchas
- **Use a buffered channel (capacity 1+) with `Notify`.** The package never blocks when delivering, so a signal is dropped if the channel is full and nobody is receiving.
- **`SIGKILL` cannot be handled.** Kubernetes sends `SIGTERM`, waits `terminationGracePeriodSeconds` (30 seconds by default) and then `SIGKILL`. Finish your shutdown within that time.
- **After the first signal call `stop()`,** so a second Ctrl+C terminates the process immediately instead of being swallowed.
- **Windows supports only a small set of signals** (`os.Interrupt` for Ctrl+C and `SIGTERM` in `syscall`). `SIGUSR1`, `SIGHUP` and `syscall.Kill` do not exist there, so the real-signal examples live in a file that is skipped on Windows (`//go:build !windows`).
- **Handle signals only in `main`,** and pass the context down. Library code should never call `os.Exit` or install signal handlers.
- **`os.Exit` skips deferred functions.** Shut down cleanly first, then return from `main` (or exit after cleanup).
- **In containers, make sure your app is PID 1 or uses an init** (`tini`, `docker run --init`), or the signal may never reach it. Use the exec form `CMD ["/app"]`, not a shell wrapper.
- **Shutdown order matters:** stop accepting work (readiness off, listener closed), wait for in-flight work with a timeout, then close databases and flush logs.

## Java comparison

| Java | Go |
|---|---|
| `Runtime.addShutdownHook(Thread)` | `signal.NotifyContext` + cleanup after `<-ctx.Done()` |
| Spring Boot graceful shutdown (`server.shutdown=graceful`) | `srv.Shutdown(ctx)` |
| `@PreDestroy` | `defer` and explicit cleanup in `main` |
| `kill -15` / `SIGTERM` | the same signal; handle it explicitly |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/os/signal
- Source: https://cs.opensource.google/go/go/+/master:src/os/signal/
