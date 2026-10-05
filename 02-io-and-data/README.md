# Stage 2: I/O and data

Streams, file systems, encodings, regular expressions, templates and logging.

| Package | What it covers |
|---------|----------------|
| [io](io) | `io` defines the two interfaces that hold Go's I/O together, `io.Reader` and `io.Writer`, plus helpers to compose them. |
| [bufio](bufio) | `bufio` adds buffering to Readers and Writers, which turns many tiny reads and writes into a few big ones, and provides `Scanner` for reading text line by line or word by word. |
| [io/fs](iofs) | `io/fs` defines `fs.FS`, a read-only file system interface with a single method, `Open(name)`. |
| [path/filepath](filepath) | `path/filepath` manipulates file names in the way the current operating system expects: backslashes on Windows, slashes elsewhere. |
| [encoding/json](json) | `encoding/json` converts Go values to JSON and back using reflection and struct tags. |
| [encoding/csv](csv) | `encoding/csv` reads and writes comma-separated values as defined by RFC 4180, including quoted fields that contain commas, quotes and newlines. |
| [encoding/xml](xml) | `encoding/xml` marshals structs to XML and parses XML into structs, and offers a token-level `Decoder` for large documents. |
| [encoding/gob](gob) | `encoding/gob` is Go's own binary serialization format. |
| [regexp](regexp) | `regexp` implements regular expressions using Go's RE2 syntax. |
| [text/template](texttemplate) | `text/template` fills a text template with data. |
| [html/template](htmltemplate) | `html/template` has the same API as `text/template` but understands HTML, CSS, JavaScript and URL contexts and escapes values automatically for the place they appear. |
| [log/slog](slog) | `log/slog` (Go 1.21+) is structured logging: each record has a level, a message and key-value attributes, and a pluggable `Handler` decides the format (human-readable text or JSON). |

Each folder has a `README.md` (concepts, gotchas, Java comparison), runnable `example_test.go`, and `exercises/` with a reference solution.

Run this stage: `go test ./02-io-and-data/...`
