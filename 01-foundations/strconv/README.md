# strconv

## Overview
`strconv` converts between strings and basic types: integers, floats, booleans
and quoted strings. Every parse function returns an error instead of throwing,
so bad input is handled with an ordinary `if err != nil`.

## Core API

| Function | Purpose |
|----------|---------|
| `Atoi`, `Itoa` | String to `int` and back (base 10) |
| `ParseInt`, `ParseUint` | Parse with a base and a bit size |
| `FormatInt`, `FormatUint` | Integer to text in any base 2 to 36 |
| `ParseFloat`, `FormatFloat` | Float conversions with format and precision |
| `ParseBool`, `FormatBool` | Boolean conversions |
| `Quote`, `Unquote`, `QuoteToASCII` | Go-syntax string escaping |
| `AppendInt`, `AppendBool`, `AppendQuote`... | Write into a `[]byte` without extra allocations |
| `NumError`, `ErrSyntax`, `ErrRange` | Error type and sentinel errors |

`Atoi(s)` is `ParseInt(s, 10, 0)` converted to `int`, and `Itoa(i)` is
`FormatInt(int64(i), 10)`.

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/strconv/
```

## Errors in detail
Parse failures return `*strconv.NumError` with three fields: `Func` (which
function), `Num` (the input) and `Err`, which is either `strconv.ErrSyntax` or
`strconv.ErrRange`. Test them with `errors.Is`, or extract the struct with
`errors.As`.

## Gotchas
- **Never ignore the error.** `n, _ := strconv.Atoi(s)` silently turns bad input
  into `0`.
- **Do not use the value when `err != nil`.** After a syntax error it is `0`;
  after a range error it is the nearest limit.
- **`Atoi` is strict.** It rejects whitespace, empty strings and decimals such as
  `"3.5"`. Trim first with `strings.TrimSpace` if input is untrusted.
- **Base 0 infers the base from a prefix.** `0x` is hex, `0b` is binary, `0o` or a
  leading `0` is octal, so `ParseInt("010", 0, 64)` is `8`. Underscores such as
  `1_000` are accepted only with base 0.
- **`bitSize` is a range check, not a return type.** `ParseInt` always returns
  `int64`; `bitSize 8` just makes values above 127 fail with `ErrRange`.
- **`ParseFloat` accepts `inf`, `infinity` and `nan`** (any case). Reject them if
  your data should never contain them.
- **`ParseBool` accepts only** `1 t T TRUE true True 0 f F FALSE false False`.
  `"yes"` and `"on"` are errors.
- **`string(65)` is `"A"`, not `"65"`.** Use `strconv.Itoa`. `go vet` flags this.
- **Precision `-1` in `FormatFloat`** gives the shortest string that parses back
  to the same value. Use a fixed precision for display.

## Java comparison

| Java | Go |
|------|----|
| `Integer.parseInt(s)` (throws `NumberFormatException`) | `strconv.Atoi(s)` (returns an error) |
| `Integer.toString(n)`, `String.valueOf(n)` | `strconv.Itoa(n)` |
| `Integer.toString(n, radix)` | `strconv.FormatInt(n, base)` |
| `Integer.toBinaryString(n)` | `strconv.FormatInt(n, 2)` |
| `Double.parseDouble(s)` | `strconv.ParseFloat(s, 64)` |
| `Boolean.parseBoolean(s)` (anything unknown is `false`) | `strconv.ParseBool(s)` (unknown is an error) |
| `NumberFormatException` | `*strconv.NumError` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/strconv
- Source: https://cs.opensource.google/go/go/+/master:src/strconv/
