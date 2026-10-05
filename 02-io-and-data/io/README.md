# io

## Overview
`io` defines the two interfaces that hold Go's I/O together, `io.Reader` and
`io.Writer`, plus helpers to compose them. Files, network connections, HTTP
bodies, buffers, compressors and hashes all implement these interfaces, so code
written against them works with all of them.

## Core API

| Function / Type | Purpose |
|---|---|
| `io.Reader`, `io.Writer`, `io.Closer` | The core interfaces (`Read(p) (n, err)`, `Write(p) (n, err)`) |
| `io.Copy`, `io.CopyN` | Stream from a Reader to a Writer |
| `io.ReadAll`, `io.ReadFull` | Read everything, or exactly `len(buf)` bytes |
| `io.LimitReader`, `io.SectionReader` | Restrict how much or which part is read |
| `io.TeeReader`, `io.MultiWriter`, `io.MultiReader` | Duplicate or concatenate streams |
| `io.Pipe` | Connect a Writer to a Reader in memory |
| `io.Discard`, `io.NopCloser`, `io.WriteString` | Small utilities |
| `io.EOF`, `io.ErrUnexpectedEOF` | End-of-stream signals |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/io/
```

## Gotchas
- **`Read` may return fewer bytes than asked, and data together with an error.** Process `n` bytes first, then look at `err`.
- **`io.EOF` is normal, not a failure.** Compare with `err == io.EOF` (or `errors.Is`).
- **`ReadAll` loads everything into memory.** Wrap untrusted input in `io.LimitReader` first.
- **`io.Copy` returns the byte count and the first error.** It allocates a 32 KB buffer unless a side implements `WriterTo` or `ReaderFrom`.
- **Always close what you open**, and close the write end of an `io.Pipe` or the reader never sees EOF.
- **A Reader can usually be read once.** To read twice, buffer it, use `io.TeeReader`, or use a `Seeker`.
- **`ReadFull` returns `ErrUnexpectedEOF`** when the stream ends mid-buffer, and plain `EOF` only when no bytes were read.

## Java comparison

| Java | Go |
|---|---|
| `InputStream` / `OutputStream` | `io.Reader` / `io.Writer` |
| `InputStream.transferTo(out)` | `io.Copy(dst, src)` |
| `readAllBytes()` | `io.ReadAll(r)` |
| `BoundedInputStream` | `io.LimitReader` |
| `PipedInputStream`/`PipedOutputStream` | `io.Pipe` |
| `OutputStream.nullOutputStream()` | `io.Discard` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/io
- Source: https://cs.opensource.google/go/go/+/master:src/io/
