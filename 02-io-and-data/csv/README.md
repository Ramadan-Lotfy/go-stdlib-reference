# encoding/csv

## Overview
`encoding/csv` reads and writes comma-separated values as defined by RFC 4180,
including quoted fields that contain commas, quotes and newlines. Rows are
`[]string`; converting to numbers and structs is up to you.

## Core API

| Function / Type | Purpose |
|---|---|
| `csv.NewReader(r)` | Create a reader; `Read` one record or `ReadAll` |
| `Reader.Comma`, `Comment`, `FieldsPerRecord`, `TrimLeadingSpace` | Dialect options |
| `Reader.LazyQuotes`, `ReuseRecord` | Tolerance and performance options |
| `Reader.FieldPos`, `InputOffset` | Line and column of a field |
| `csv.NewWriter(w)` | `Write`, `WriteAll`, then `Flush` |
| `csv.ErrFieldCount`, `*csv.ParseError` | Error values |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/csv/
```

## Gotchas
- **Call `w.Flush()` and then check `w.Error()`.** Output stays in the buffer until you flush.
- **`FieldsPerRecord` is set from the first row.** Later rows with a different count return `ErrFieldCount`; set it to `-1` to allow ragged rows.
- **Use `Read` in a loop for big files**; `ReadAll` loads everything into memory.
- **Everything is a string.** Convert with `strconv` and report the line number on failure.
- **The header row is just another record.** Map column names to indexes yourself so column order can change.
- **A UTF-8 BOM** (common in Excel exports) ends up in the first header name. Strip it.
- **`ReuseRecord` makes `Read` return the same slice** each time; copy fields you keep.

## Java comparison

| Java | Go |
|---|---|
| OpenCSV `CSVReader` | `csv.Reader` |
| Apache Commons CSV `CSVParser` | `csv.Reader` + a loop |
| `CSVWriter.close()` | `w.Flush()` then `w.Error()` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/encoding/csv
- Source: https://cs.opensource.google/go/go/+/master:src/encoding/csv/
