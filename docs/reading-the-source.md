# Reading the Standard Library Source

The Go standard library is mostly plain, readable Go, written by the people who designed
the language. Reading it is the fastest way to learn idiomatic style, concurrency
patterns and performance tricks. It is also the final step of each stage in this repo.

## Where the source lives

```bash
go env GOROOT                 # the Go installation
ls $(go env GOROOT)/src       # every standard library package, as plain directories
go doc -src strings.Cut       # print the source of one function
go doc -all sync.Mutex        # docs plus all methods
```

On pkg.go.dev every function has a link to its source, and
https://cs.opensource.google/go/go has search and cross-references.

## How to read a package

1. **Docs first.** Read `doc.go` or the package comment, then `go doc -all pkg`.
2. **Tests second.** `*_test.go` files and `example_test.go` show intended use and edge cases.
3. **Then the core file** (listed below). Skip `*_windows.go`, `*_plan9.go` and generated files at first.
4. **Follow one function end to end** instead of reading top to bottom.
5. **Ask a question before you start** ("how does it avoid allocating?") and look for the answer.

## Suggested reading list

| Package | Start with | Questions to answer |
|---------|-----------|---------------------|
| `errors` | `errors/wrap.go` | How does `Is` walk a chain, including `Join` trees and an `Is` method? |
| `strings` | `strings/builder.go` | How does `Builder` avoid copying? Why the `copyCheck`? |
| `io` | `io/io.go` | Why are `Reader` and `Writer` so small? How does `Copy` pick `WriterTo` / `ReaderFrom`? |
| `bufio` | `bufio/scan.go` | How does `Scanner` grow its buffer, and where does `ErrTooLong` come from? |
| `sort` / `slices` | `slices/sort.go`, `slices/zsortordered.go` | Which algorithm is used, and why is it generated code? |
| `container/heap` | `container/heap/heap.go` | How small can a heap be? (about 100 lines) |
| `sync` | `sync/mutex.go`, `sync/waitgroup.go`, `sync/once.go` | What are "normal" and "starvation" mode? Why does `Once` use a fast path with an atomic? |
| `sync/atomic` | `sync/atomic/type.go` | How do the typed values wrap the raw functions? |
| `context` | `context/context.go` | How does cancellation propagate to children? What is `propagateCancel`? |
| `time` | `time/sleep.go`, `time/format.go` | How are timers stored? How does layout parsing work? |
| `encoding/json` | `encoding/json/encode.go` | How are encoders built once per type and cached? |
| `net/http` | `net/http/server.go`, `net/http/serve_mux.go` (1.22 routing in `pattern.go`) | What does one connection's goroutine do? How does `Shutdown` wait for idle connections? |
| `net/http/httptest` | `net/http/httptest/recorder.go` | How little is a `ResponseWriter`? |
| `database/sql` | `database/sql/sql.go` | How is the connection pool implemented? When does `Rows.Close` return a connection? |
| `reflect` | `reflect/type.go`, `reflect/value.go` | How are `Type` and `Value` represented in memory? |
| `runtime` | `runtime/proc.go`, `runtime/HACKING.md` | What is the G-M-P scheduler? What does `gopark` do? |
| `testing` | `testing/testing.go` | How do `t.Run` and `t.Parallel` coordinate? |
| `embed` | `embed/embed.go` | How can an `fs.FS` be this small? |

## Reading exercises

1. Find how `strings.Builder` makes `String()` free of copies, then explain why copying a
   `Builder` after first use panics.
2. Trace `sync.Mutex.Lock` from the fast path (one atomic compare-and-swap) to the slow
   path. When does a goroutine spin?
3. Follow an HTTP request from `Server.Serve` to your handler. Where is the goroutine for
   the request started?
4. In `context.WithCancel`, find where a child registers with its parent. What happens when
   the parent is a custom `Context` implementation?
5. In `encoding/json`, find the cache of per-type encoders and explain why the first
   `Marshal` of a type is slower.

## Habits that help

- Keep `$GOROOT/src` open in your editor and use "go to definition" from your own code
  into the library.
- Read the **commit history** of a file (`git log` in a clone of `github.com/golang/go`);
  commit messages explain *why* the code looks the way it does.
- Read design documents for big features: `go.dev/design`, release notes at `go.dev/doc/devel/release`.
- When a behavior surprises you, write a tiny `Example` test that pins it down, then read
  the source to find out why. That is how most of the gotchas in this repo were found.
