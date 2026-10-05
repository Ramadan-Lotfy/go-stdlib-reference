# log/slog exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/slog/exercises/...
```

## Exercise: NewLogger and Secret

Write:

```go
func NewLogger(w io.Writer, service string) *slog.Logger

type Secret string // never appears in logs
```

`NewLogger` returns a JSON logger that writes to `w`, drops the `time`
attribute (so tests are deterministic) and adds `"service": service` to every
record.

`Secret` implements `slog.LogValuer` so it logs as `[REDACTED]`.

The test expects exactly:

```json
{"level":"INFO","msg":"hi","service":"api","token":"[REDACTED]","n":2}
```

Hints: `slog.NewJSONHandler` with `HandlerOptions.ReplaceAttr`, and
`logger.With`.

## Bonus
Write a handler wrapper that adds a request ID from `context.Context` to every
record logged with `InfoContext`.
