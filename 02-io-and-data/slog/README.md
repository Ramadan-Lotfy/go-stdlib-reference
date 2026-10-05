# log/slog

## Overview
`log/slog` (Go 1.21+) is structured logging: each record has a level, a
message and key-value attributes, and a pluggable `Handler` decides the format
(human-readable text or JSON). It replaces the old `log` package for services,
and works with log aggregation tools.

## Core API

| Function / Type | Purpose |
|---|---|
| `slog.Info`, `Warn`, `Error`, `Debug` | Log with the default logger |
| `slog.New(handler)` | Create a logger |
| `slog.NewTextHandler`, `slog.NewJSONHandler` | Built-in handlers; `HandlerOptions` sets level and attribute rewriting |
| `logger.With(...)`, `WithGroup(...)` | Add attributes to every later record |
| `slog.String`, `Int`, `Duration`, `Any`, `Group` | Typed attributes |
| `slog.LevelVar` | Change the level at runtime |
| `slog.LogValuer` | Control how your type appears in logs (redaction) |
| `slog.SetDefault` | Install the default logger (also redirects package `log`) |
| `logger.InfoContext(ctx, ...)` | Pass a context to handlers |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/slog/
```

## Gotchas
- **Pass attributes as alternating key, value pairs.** An odd number of arguments produces a `!BADKEY` entry; `go vet` checks this.
- **Prefer `slog.Attr` helpers (`slog.Int`, `slog.String`) in hot paths**; they avoid allocations and mistakes.
- **The default level is Info**, so `Debug` calls are dropped unless you lower the level.
- **Do not log secrets.** Implement `LogValuer` on sensitive types so they print as `[REDACTED]`.
- **Include the error as `"err", err`** and add context with `With`, instead of formatting it into the message.
- **Remove the `time` attribute in tests** with `ReplaceAttr`, so output is deterministic (these examples do).
- **Handlers are shared across goroutines;** the built-in ones lock the writer for you.
- **Keep messages constant** (`"user created"`) and put variable data in attributes, so logs can be searched and counted.

## Java comparison

| Java | Go |
|---|---|
| SLF4J `log.info("x {}", v)` | `slog.Info("x", "v", v)` |
| MDC | `logger.With(...)` / `slog.Group` |
| Logback JSON encoder | `slog.NewJSONHandler` |
| Log levels (`TRACE..ERROR`) | `slog.LevelDebug/Info/Warn/Error` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/log/slog
- Source: https://cs.opensource.google/go/go/+/master:src/log/slog/
