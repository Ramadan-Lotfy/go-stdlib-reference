# Stage 1: Core foundations

Text, errors, files, collections and time: the everyday toolkit.

| Package | What it covers |
|---------|----------------|
| [fmt](fmt) | `fmt` implements formatted I/O: printing values to stdout, to strings, or to any `io.Writer`, plus building errors and parsing text back into values. |
| [strings](strings) | `strings` provides functions for searching, splitting, trimming, replacing and building UTF-8 strings. |
| [strconv](strconv) | `strconv` converts between strings and basic types: integers, floats, booleans and quoted strings. |
| [bytes](bytes) | `bytes` is the `[]byte` twin of `strings`: the same searching, splitting and trimming functions, plus `bytes.Buffer`, a growable buffer you can write to and read from, and `bytes.Reader`, a read-only view of a slice. |
| [errors](errors) | `errors` creates, wraps, inspects and combines errors. |
| [os](os) | `os` is Go's interface to the operating system: files and directories, environment variables, process arguments and exit codes. |
| [sort](sort) | `sort` sorts slices and user-defined collections, and does binary search. |
| [slices](slices) | `slices` (Go 1.21+) has generic helpers for any slice: sorting, searching, inserting, deleting, comparing and de-duplicating. |
| [maps](maps) | Go's built-in `map` type is a hash table with unordered iteration. |
| [container/list, container/heap, container/ring](container) | Three small data-structure packages. |
| [time](time) | `time` handles instants, durations, calendars, time zones, timers and tickers. |

Each folder has a `README.md` (concepts, gotchas, Java comparison), runnable `example_test.go`, and `exercises/` with a reference solution.

Run this stage: `go test ./01-foundations/...`
