# Learning Path

Work through the stages in order. For each package:

1. Read the package `README.md` (overview, core API, gotchas).
2. Run the examples: `go test -v ./<stage>/<package>/` and read `example_test.go` next to the output.
3. Change an example and predict the output before running it.
4. Solve the exercises in `exercises/README.md` **before** looking at `exercises/solution/`.
5. Read the package docs on pkg.go.dev, then the source of one key function ([how](reading-the-source.md)).

## Stage 1: Core foundations

Text, errors, files, collections and time: the everyday toolkit.

1. [fmt](../01-foundations/fmt): `fmt` implements formatted I/O: printing values to stdout, to strings, or to any `io.Writer`, plus building errors and parsing text back into values.
2. [strings](../01-foundations/strings): `strings` provides functions for searching, splitting, trimming, replacing and building UTF-8 strings.
3. [strconv](../01-foundations/strconv): `strconv` converts between strings and basic types: integers, floats, booleans and quoted strings.
4. [bytes](../01-foundations/bytes): `bytes` is the `[]byte` twin of `strings`: the same searching, splitting and trimming functions, plus `bytes.Buffer`, a growable buffer you can write to and read from, and `bytes.Reader`, a read-only view of a slice.
5. [errors](../01-foundations/errors): `errors` creates, wraps, inspects and combines errors.
6. [os](../01-foundations/os): `os` is Go's interface to the operating system: files and directories, environment variables, process arguments and exit codes.
7. [sort](../01-foundations/sort): `sort` sorts slices and user-defined collections, and does binary search.
8. [slices](../01-foundations/slices): `slices` (Go 1.21+) has generic helpers for any slice: sorting, searching, inserting, deleting, comparing and de-duplicating.
9. [maps](../01-foundations/maps): Go's built-in `map` type is a hash table with unordered iteration.
10. [container/list, container/heap, container/ring](../01-foundations/container): Three small data-structure packages.
11. [time](../01-foundations/time): `time` handles instants, durations, calendars, time zones, timers and tickers.

## Stage 2: I/O and data

Streams, file systems, encodings, regular expressions, templates and logging.

1. [io](../02-io-and-data/io): `io` defines the two interfaces that hold Go's I/O together, `io.Reader` and `io.Writer`, plus helpers to compose them.
2. [bufio](../02-io-and-data/bufio): `bufio` adds buffering to Readers and Writers, which turns many tiny reads and writes into a few big ones, and provides `Scanner` for reading text line by line or word by word.
3. [io/fs](../02-io-and-data/iofs): `io/fs` defines `fs.FS`, a read-only file system interface with a single method, `Open(name)`.
4. [path/filepath](../02-io-and-data/filepath): `path/filepath` manipulates file names in the way the current operating system expects: backslashes on Windows, slashes elsewhere.
5. [encoding/json](../02-io-and-data/json): `encoding/json` converts Go values to JSON and back using reflection and struct tags.
6. [encoding/csv](../02-io-and-data/csv): `encoding/csv` reads and writes comma-separated values as defined by RFC 4180, including quoted fields that contain commas, quotes and newlines.
7. [encoding/xml](../02-io-and-data/xml): `encoding/xml` marshals structs to XML and parses XML into structs, and offers a token-level `Decoder` for large documents.
8. [encoding/gob](../02-io-and-data/gob): `encoding/gob` is Go's own binary serialization format.
9. [regexp](../02-io-and-data/regexp): `regexp` implements regular expressions using Go's RE2 syntax.
10. [text/template](../02-io-and-data/texttemplate): `text/template` fills a text template with data.
11. [html/template](../02-io-and-data/htmltemplate): `html/template` has the same API as `text/template` but understands HTML, CSS, JavaScript and URL contexts and escapes values automatically for the place they appear.
12. [log/slog](../02-io-and-data/slog): `log/slog` (Go 1.21+) is structured logging: each record has a level, a message and key-value attributes, and a pluggable `Handler` decides the format (human-readable text or JSON).

## Stage 3: Concurrency

Goroutines' toolbox: locks, atomics, contexts, channels and patterns.

1. [sync](../03-concurrency/sync): `sync` provides the classic shared-memory primitives: mutexes, wait groups, one-time initialization, condition variables, a concurrent map and an object pool.
2. [sync/atomic](../03-concurrency/atomic): `sync/atomic` provides lock-free operations on single values: counters, flags and pointers.
3. [context](../03-concurrency/context): `context` carries cancellation signals, deadlines and request-scoped values across API boundaries and between goroutines.
4. [channels, select and concurrency patterns](../03-concurrency/channels): Channels are typed pipes between goroutines, and `select` waits on several of them at once.

## Stage 4: Networking and backend

HTTP, sockets, URLs, TLS, SQL and testing.

1. [net/http](../04-networking/nethttp): `net/http` is a complete HTTP client and server.
2. [net/http/httptest](../04-networking/httptest): `net/http/httptest` makes HTTP code testable without real networks or ports.
3. [net](../04-networking/net): `net` is the portable interface to network I/O: TCP and UDP sockets, Unix sockets, DNS lookups, IP address and CIDR handling.
4. [net/url](../04-networking/neturl): `net/url` parses, builds and escapes URLs and query strings.
5. [crypto/tls](../04-networking/tls): `crypto/tls` implements TLS, the protocol behind HTTPS.
6. [database/sql](../04-networking/databasesql): `database/sql` is Go's database access layer.
7. [testing](../04-networking/testing): The `testing` package and the `go test` command are Go's built-in test framework: no external library is needed for unit tests, benchmarks, fuzzing, examples or coverage.

## Stage 5: Mastery

Reflection, generics, the runtime, profiling, crypto and production concerns.

1. [reflect](../05-mastery/reflect): `reflect` lets a program inspect and modify values whose types are not known at compile time: list struct fields and tags, call methods by name, build slices and maps dynamically.
2. [unsafe](../05-mastery/unsafe): `unsafe` steps outside Go's type and memory safety rules: sizes and offsets, pointer arithmetic, and zero-copy conversions.
3. [generics, cmp and slices](../05-mastery/generics): Generics (Go 1.18+) let functions and types work over many types while keeping compile-time type safety.
4. [runtime and runtime/debug](../05-mastery/runtime): `runtime` is the interface to Go's run-time system: the scheduler, garbage collector and memory allocator.
5. [runtime/pprof and net/http/pprof](../05-mastery/pprof): `runtime/pprof` records profiles of a running program: where CPU time goes, which code allocates memory, which goroutines exist, and where they block.
6. [runtime/trace](../05-mastery/trace): `runtime/trace` records an **execution trace**: a timeline of everything the runtime does (goroutine creation, blocking, scheduling, garbage collection, system calls) with nanosecond timestamps.
7. [crypto/*](../05-mastery/crypto): The `crypto` packages provide hashing, message authentication, random numbers, encryption and digital signatures.
8. [embed](../05-mastery/embed): `embed` compiles files into the program binary, so a single executable can carry its own templates, SQL migrations, static web assets and default config.
9. [flag](../05-mastery/flag): `flag` parses command-line flags such as `-addr :8080 -v`.
10. [os/exec](../05-mastery/osexec): `os/exec` runs external programs.
11. [os/signal](../05-mastery/signal): `os/signal` lets a program react to operating system signals: `SIGINT` (Ctrl+C), `SIGTERM` (the polite "please stop" sent by Kubernetes, systemd and `docker stop`), `SIGHUP` (often "reload config") and others.

## After the five stages

- Read `net/http/server.go`, `sync/mutex.go` and `io/io.go` end to end.
- Rebuild a small piece yourself: a rate limiter (`sync`, `time`), a worker pool (`channels`, `context`), a tiny router (`net/http`).
- Profile one of your own services with `pprof` and `trace`.
