# Go Standard Library Reference

A hands-on, organized reference to Go's standard library, from the basics to mastery.
Every example is a real `Example` function that is compiled and verified by `go test`,
so the documentation can never drift out of date.

Written for developers coming from other ecosystems (especially Java/Spring), but useful
for anyone learning Go.

**45 packages, 312 verified examples, 45 exercises with tested solutions.** Requires Go 1.22 or newer.

## How to use this repo

```bash
git clone https://github.com/Ramadan-Lotfy/go-stdlib-reference.git
cd go-stdlib-reference

go test ./...                          # run every example and every exercise solution
go test -race ./...                    # same, with the race detector (recommended)
go test -v ./03-concurrency/context/   # run and watch a single package
go doc -all ./01-foundations/fmt       # read a package's docs in the terminal
```

Each package folder contains:

| File | Purpose |
|------|---------|
| `README.md` | Overview, core API table, gotchas, Java comparison |
| `doc.go` | Package documentation (shows up in `go doc`) |
| `example_test.go` | Runnable examples verified by `go test` |
| `exercises/` | Exercise descriptions plus tested reference solutions in `exercises/solution/` |

## Learning path

Follow the stages in order. Details in [docs/learning-path.md](docs/learning-path.md).

| Stage | Folder | Packages |
|-------|--------|----------|
| 1 | [01-foundations](01-foundations) | fmt, strings, strconv, bytes, errors, os, sort, slices, maps, container/list, time |
| 2 | [02-io-and-data](02-io-and-data) | io, bufio, io/fs, path/filepath, encoding/json, encoding/csv, encoding/xml, encoding/gob, regexp, text/template, html/template, log/slog |
| 3 | [03-concurrency](03-concurrency) | sync, sync/atomic, context, channels |
| 4 | [04-networking](04-networking) | net/http, net/http/httptest, net, net/url, crypto/tls, database/sql, testing |
| 5 | [05-mastery](05-mastery) | reflect, unsafe, generics, runtime, runtime/pprof, runtime/trace, crypto/*, embed, flag, os/exec, os/signal |

For each package: read its `README.md`, run the examples, solve the exercises yourself, then compare
with the solution. Finish a stage by reading some standard library source
([docs/reading-the-source.md](docs/reading-the-source.md)).

## Package index

### Stage 1: Core foundations

| Package | Status |
|---------|--------|
| [fmt](01-foundations/fmt) | ✅ |
| [strings](01-foundations/strings) | ✅ |
| [strconv](01-foundations/strconv) | ✅ |
| [bytes](01-foundations/bytes) | ✅ |
| [errors](01-foundations/errors) | ✅ |
| [os](01-foundations/os) | ✅ |
| [sort](01-foundations/sort) | ✅ |
| [slices](01-foundations/slices) | ✅ |
| [maps](01-foundations/maps) | ✅ |
| [container/list, container/heap, container/ring](01-foundations/container) | ✅ |
| [time](01-foundations/time) | ✅ |

### Stage 2: I/O and data

| Package | Status |
|---------|--------|
| [io](02-io-and-data/io) | ✅ |
| [bufio](02-io-and-data/bufio) | ✅ |
| [io/fs](02-io-and-data/iofs) | ✅ |
| [path/filepath](02-io-and-data/filepath) | ✅ |
| [encoding/json](02-io-and-data/json) | ✅ |
| [encoding/csv](02-io-and-data/csv) | ✅ |
| [encoding/xml](02-io-and-data/xml) | ✅ |
| [encoding/gob](02-io-and-data/gob) | ✅ |
| [regexp](02-io-and-data/regexp) | ✅ |
| [text/template](02-io-and-data/texttemplate) | ✅ |
| [html/template](02-io-and-data/htmltemplate) | ✅ |
| [log/slog](02-io-and-data/slog) | ✅ |

### Stage 3: Concurrency

| Package | Status |
|---------|--------|
| [sync](03-concurrency/sync) | ✅ |
| [sync/atomic](03-concurrency/atomic) | ✅ |
| [context](03-concurrency/context) | ✅ |
| [channels, select and concurrency patterns](03-concurrency/channels) | ✅ |

### Stage 4: Networking and backend

| Package | Status |
|---------|--------|
| [net/http](04-networking/nethttp) | ✅ |
| [net/http/httptest](04-networking/httptest) | ✅ |
| [net](04-networking/net) | ✅ |
| [net/url](04-networking/neturl) | ✅ |
| [crypto/tls](04-networking/tls) | ✅ |
| [database/sql](04-networking/databasesql) | ✅ |
| [testing](04-networking/testing) | ✅ |

### Stage 5: Mastery

| Package | Status |
|---------|--------|
| [reflect](05-mastery/reflect) | ✅ |
| [unsafe](05-mastery/unsafe) | ✅ |
| [generics, cmp and slices](05-mastery/generics) | ✅ |
| [runtime and runtime/debug](05-mastery/runtime) | ✅ |
| [runtime/pprof and net/http/pprof](05-mastery/pprof) | ✅ |
| [runtime/trace](05-mastery/trace) | ✅ |
| [crypto/*](05-mastery/crypto) | ✅ |
| [embed](05-mastery/embed) | ✅ |
| [flag](05-mastery/flag) | ✅ |
| [os/exec](05-mastery/osexec) | ✅ |
| [os/signal](05-mastery/signal) | ✅ |

## More documentation

- [docs/learning-path.md](docs/learning-path.md): the five stages in order, with study tips
- [docs/java-to-go.md](docs/java-to-go.md): cheat sheet for Java developers
- [docs/reading-the-source.md](docs/reading-the-source.md): how to read the standard library source
- [CONTRIBUTING.md](CONTRIBUTING.md): package template and style rules

## Notes on running the tests

- Examples that touch the network use only `127.0.0.1` (loopback) and random ports.
- Timing-based examples use short real timers; run with `-count=1` if you suspect caching.
- Files named `*_unix_test.go` are skipped on Windows.
- `unsafe` examples assume a 64-bit platform.

## License

MIT
