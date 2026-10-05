# bytes

## Overview
`bytes` is the `[]byte` twin of `strings`: the same searching, splitting and
trimming functions, plus `bytes.Buffer`, a growable buffer you can write to and
read from, and `bytes.Reader`, a read-only view of a slice. Most Go I/O APIs
speak `[]byte`, so this package is the glue between them.

## Core API

| Function / Type | Purpose |
|-----------------|---------|
| `Contains`, `Index`, `HasPrefix`, `Count` | Search, same as `strings` |
| `Split`, `Fields`, `Cut`, `Join`, `ReplaceAll`, `TrimSpace` | Transform, same as `strings` |
| `Equal`, `Compare` | Compare slices (`==` does not work) |
| `Buffer` | Growable byte buffer: `io.Reader` and `io.Writer` in one |
| `NewBuffer`, `NewBufferString` | Create a Buffer from existing data |
| `Reader` (`NewReader`) | Read-only `io.Reader`, `io.ReaderAt` and `io.Seeker` over a slice |
| `Runes`, `IndexRune`, `ToUpper` | UTF-8 aware helpers |

## strings vs bytes: which one?
- Text you receive as `string`, or build for display: use `strings` and
  `strings.Builder`.
- Data that arrives as `[]byte` (files, sockets, `json.Marshal`, hashes) or that you
  need to modify in place: use `bytes`.
- Need something you can both write to and read back, or capture the output of an
  `io.Writer`: use `bytes.Buffer`.

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/bytes/
```

## Gotchas
- **`==` does not compare slices.** It only compares a slice to `nil`. Use
  `bytes.Equal`, which also treats `nil` and empty as equal.
- **`Buffer.Bytes()` does not copy.** The slice aliases the buffer's memory and is
  invalidated by the next write. `Buffer.String()` copies.
- **`bytes.Split`, `Fields` and `Cut` return subslices** of the original, so writing
  to a result changes the source. Copy with `append([]byte(nil), b...)` when you
  need independence.
- **Reading consumes a Buffer.** `Len()` and `String()` only cover the *unread*
  part. Use `bytes.NewReader` if you need to read the same data repeatedly.
- **`ReadString` and `ReadBytes` return the data plus `io.EOF`** when the delimiter
  is missing. Handle the data before checking the error.
- **`Next(n)` returns a slice that is only valid until the next read or write.**
- **Conversions copy.** `[]byte(s)` and `string(b)` each allocate. In hot paths,
  keep data in one form.
- **`Buffer` is not safe for concurrent use**, and a used `Buffer` must not be copied.
  Pass `*bytes.Buffer`.
- **`Truncate(n)` panics** if `n` is larger than the unread length.
- **Reuse with `Reset()`** to keep the allocated memory. For high-throughput code,
  pair it with `sync.Pool`.

## Java comparison

| Java | Go |
|------|----|
| `byte[]` | `[]byte` |
| `Arrays.equals(a, b)` | `bytes.Equal(a, b)` |
| `s.getBytes(UTF_8)` | `[]byte(s)` |
| `new String(b, UTF_8)` | `string(b)` |
| `ByteArrayOutputStream` | `bytes.Buffer` (write side) |
| `ByteArrayInputStream` | `bytes.NewReader(b)` |
| `Arrays.copyOf(b, n)` | `append([]byte(nil), b...)` or `copy` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/bytes
- Source: https://cs.opensource.google/go/go/+/master:src/bytes/
