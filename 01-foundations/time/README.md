# time

## Overview
`time` handles instants, durations, calendars, time zones, timers and
tickers. A `time.Time` is an instant plus a location; a `time.Duration` is a
number of nanoseconds. Formatting uses a reference date instead of letters like
`yyyy`: you write the layout using `Mon Jan 2 15:04:05 MST 2006`.

## Core API

| Function / Type | Purpose |
|---|---|
| `time.Now`, `time.Since`, `time.Until` | Current time and elapsed time |
| `time.Date`, `time.Unix` | Construct a `Time` |
| `t.Format`, `time.Parse` | Text conversion using a layout |
| `t.Add`, `t.AddDate`, `t.Sub` | Arithmetic |
| `t.Before`, `t.After`, `t.Equal` | Comparison |
| `t.In`, `time.FixedZone`, `time.LoadLocation` | Time zones |
| `time.Duration`, `time.ParseDuration` | Spans of time |
| `time.NewTimer`, `time.NewTicker`, `time.After` | Delays and periodic events |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/time/
```

## Gotchas
- **The layout is the reference time**: `2006-01-02 15:04:05`. `yyyy-MM-dd` does not work. Constants such as `time.RFC3339`, `time.DateOnly` and `time.DateTime` cover common cases.
- **Compare with `t.Equal`, not `==`.** `==` also compares the location and monotonic clock reading.
- **`AddDate` normalizes overflow.** January 31 plus one month is March 3 (non-leap year), not February 28.
- **Convert ints before multiplying durations**: `time.Duration(n) * time.Second`. `n * time.Second` fails to compile when `n` is an `int` variable.
- **`LoadLocation` needs tz data.** On Windows or minimal containers import `_ "time/tzdata"` to embed it.
- **Always `Stop` tickers** (`defer ticker.Stop()`), or they leak. Since Go 1.23 unreferenced timers are collected, but explicit cleanup is still good practice.
- **Do not use `time.Sleep` for synchronization in tests.** Use channels, `sync.WaitGroup` or a fake clock.
- **Store and exchange instants in UTC** (RFC 3339 strings, or Unix timestamps) and convert to local time only for display.

## Java comparison

| Java | Go |
|---|---|
| `Instant.now()`, `LocalDateTime` | `time.Now()` |
| `Duration.ofMinutes(90)` | `90 * time.Minute` |
| `DateTimeFormatter.ofPattern("yyyy-MM-dd")` | `t.Format("2006-01-02")` |
| `ZoneId`, `ZonedDateTime` | `*time.Location`, `t.In(loc)` |
| `ScheduledExecutorService` | `time.Ticker` / `time.AfterFunc` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/time
- Source: https://cs.opensource.google/go/go/+/master:src/time/
