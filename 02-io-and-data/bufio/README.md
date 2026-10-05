# bufio

## Overview
`bufio` adds buffering to Readers and Writers, which turns many tiny reads
and writes into a few big ones, and provides `Scanner` for reading text line by
line or word by word.

## Core API

| Function / Type | Purpose |
|---|---|
| `bufio.NewScanner` | Read lines, words, runes or custom tokens |
| `Scanner.Split(bufio.ScanWords)` | Choose the token type |
| `Scanner.Buffer(buf, max)` | Raise the maximum token size |
| `bufio.NewReader` | `ReadString`, `ReadLine`, `ReadRune`, `Peek` |
| `bufio.NewWriter` | Buffered output; you must call `Flush` |
| `bufio.NewReadWriter` | Combine a buffered Reader and Writer |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/bufio/
```

## Gotchas
- **`Scanner` rejects lines over 64 KB** and stops with `bufio.ErrTooLong`. Call `sc.Buffer(make([]byte, 0, 64*1024), maxSize)` for bigger lines.
- **Always check `sc.Err()`** after the `Scan` loop; `Scan` returns `false` on both EOF and errors.
- **`sc.Bytes()` is overwritten by the next `Scan`.** Copy it, or use `sc.Text()`, which allocates a string.
- **Forgetting `Flush` loses the last buffered output.** `defer w.Flush()` right after creating the writer.
- **`Reader.ReadString('\n')` returns the data plus `io.EOF`** for a final line without a newline. Use the data, then stop.
- **`Scanner` strips the newline** (and a trailing `\r`); `ReadString` keeps it.
- **Wrap `os.Stdout` in a `bufio.Writer`** when printing many lines. It is much faster.

## Java comparison

| Java | Go |
|---|---|
| `BufferedReader` | `bufio.Reader` / `bufio.Scanner` |
| `reader.readLine()` | `scanner.Scan()` + `scanner.Text()` |
| `java.util.Scanner` | `bufio.Scanner` (tokens only, no number parsing) |
| `BufferedWriter.flush()` | `bufio.Writer.Flush()` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/bufio
- Source: https://cs.opensource.google/go/go/+/master:src/bufio/
