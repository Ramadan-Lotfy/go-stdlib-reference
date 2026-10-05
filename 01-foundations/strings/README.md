# strings

## Overview
`strings` provides functions for searching, splitting, trimming, replacing and
building UTF-8 strings. Go strings are immutable byte sequences, so most
functions return a new string. The `bytes` package mirrors this API for `[]byte`.

## Core API

| Function / Type | Purpose |
|-----------------|---------|
| `Contains`, `ContainsAny`, `ContainsRune` | Membership tests |
| `HasPrefix`, `HasSuffix` | Start and end checks |
| `Index`, `LastIndex`, `IndexRune` | Byte offset of a match, or `-1` |
| `Count` | Number of non-overlapping matches |
| `Split`, `SplitN`, `Fields`, `FieldsFunc` | Break a string into pieces |
| `Join`, `Repeat` | Combine pieces |
| `TrimSpace`, `Trim`, `TrimLeft`, `TrimRight` | Remove characters in a cutset |
| `TrimPrefix`, `TrimSuffix`, `CutPrefix`, `CutSuffix` | Remove an exact prefix or suffix |
| `Cut` | Split around the first separator |
| `Replace`, `ReplaceAll`, `NewReplacer`, `Map` | Transform contents |
| `ToUpper`, `ToLower`, `EqualFold` | Case handling |
| `Builder` | Efficient string construction (an `io.Writer`) |
| `NewReader` | Wrap a string as an `io.Reader` |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/strings/
```

## Gotchas
- **`len(s)` counts bytes, not characters.** `len("é")` is 2. Use
  `utf8.RuneCountInString` or `for range` for characters.
- **Slicing is byte-based.** `s[1:2]` can cut a multi-byte character in half.
- **`+=` in a loop is slow.** Each concatenation copies the string. Use
  `strings.Builder`.
- **Do not copy a `Builder` after first use.** Pass a `*strings.Builder`.
- **`Trim`, `TrimLeft`, `TrimRight` take a cutset**, a set of characters, not a
  substring. For an exact prefix or suffix use `TrimPrefix`, `TrimSuffix` or `Cut*`.
- **`Split("", sep)` returns `[""]`**, a slice with one empty string.
- **`Split(s, " ")` keeps empty fields** between repeated spaces. `Fields` collapses
  any run of whitespace.
- **`Index` returns `-1`** when nothing matches. Check before slicing.
- **`Split` is literal.** Java's `split` takes a regex; in Go use `regexp` for
  patterns.
- **`Title` is deprecated.** For correct title casing use
  `golang.org/x/text/cases`.
- **Prefer `==` and `<` over `strings.Compare`**, which exists mainly for symmetry
  with other packages.

## Java comparison

| Java | Go |
|------|----|
| `s.contains(x)` | `strings.Contains(s, x)` |
| `s.startsWith(x)`, `s.endsWith(x)` | `strings.HasPrefix`, `strings.HasSuffix` |
| `s.indexOf(x)` | `strings.Index(s, x)` |
| `s.split(regex)` | `strings.Split(s, sep)` (literal), `regexp` for patterns |
| `String.join(sep, parts)` | `strings.Join(parts, sep)` |
| `s.strip()` / `s.trim()` | `strings.TrimSpace(s)` |
| `s.repeat(n)` | `strings.Repeat(s, n)` |
| `s.replace(a, b)` | `strings.ReplaceAll(s, a, b)` |
| `s.equalsIgnoreCase(t)` | `strings.EqualFold(s, t)` |
| `StringBuilder` | `strings.Builder` |
| `s.length()` (UTF-16 units) | `len(s)` (UTF-8 bytes) |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/strings
- Source: https://cs.opensource.google/go/go/+/master:src/strings/
