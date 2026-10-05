# regexp

## Overview
`regexp` implements regular expressions using Go's RE2 syntax. RE2 guarantees
matching time linear in the input size, which makes it safe on untrusted text, at
the cost of not supporting backreferences or lookahead.

## Core API

| Function / Type | Purpose |
|---|---|
| `regexp.MustCompile`, `regexp.Compile` | Build a reusable `*Regexp` (panic or error on bad syntax) |
| `MatchString`, `Match` | Does it match anywhere in the input? |
| `FindString`, `FindStringIndex` | First match, or its position |
| `FindAllString(s, n)` | All matches (`n = -1` for no limit) |
| `FindStringSubmatch` | First match plus capture groups |
| `SubexpNames`, `SubexpIndex` | Work with named groups `(?P<name>...)` |
| `ReplaceAllString`, `ReplaceAllStringFunc` | Replace using `$1` / `${name}` or a function |
| `Split` | Split by a pattern |
| `regexp.QuoteMeta` | Escape a literal for use in a pattern |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/regexp/
```

## Gotchas
- **Compile once.** `MustCompile` is slow; keep the result in a package-level variable, not inside a loop or handler.
- **Patterns are unanchored.** `\d+` matches inside `abc123def`. Use `^...$` to match the whole string.
- **Use raw strings** (backticks) for patterns so you do not have to double every backslash.
- **`$1x` means group named `1x`**, not group 1 followed by `x`. Write `${1}x`.
- **No lookahead, lookbehind or backreferences.** That is the price of linear-time matching.
- **`MustCompile` panics** on a bad pattern; use it only for constants. Use `Compile` for user-supplied patterns.
- **`(?i)` at the start makes a pattern case-insensitive**; `(?s)` lets `.` match newlines; `(?m)` makes `^`/`$` match at lines.
- **Do not parse HTML, JSON or XML with regexp.** Use the real parsers.
- **`Find...` returns `nil` / `-1` / `""` when nothing matches**, so check before indexing a submatch slice.

## Java comparison

| Java | Go |
|---|---|
| `Pattern.compile(re)` | `regexp.MustCompile(re)` |
| `matcher.matches()` | `re.MatchString(s)` with `^...$` |
| `matcher.find()` in a loop | `re.FindAllString(s, -1)` |
| `matcher.group(1)` | `re.FindStringSubmatch(s)[1]` |
| `replaceAll("$1")` | `re.ReplaceAllString(s, "$1")` |
| `(?<name>...)` | `(?P<name>...)` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/regexp
- Source: https://cs.opensource.google/go/go/+/master:src/regexp/
