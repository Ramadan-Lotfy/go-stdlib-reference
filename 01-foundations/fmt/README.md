# fmt

## Overview
`fmt` implements formatted I/O: printing values to stdout, to strings, or to any
`io.Writer`, plus building errors and parsing text back into values. It is the
first package every Go program uses.

## Core API

| Function | Purpose |
|----------|---------|
| `Println`, `Printf`, `Print` | Write to stdout |
| `Sprintf`, `Sprint`, `Sprintln` | Return a formatted `string` |
| `Fprintf`, `Fprint`, `Fprintln` | Write to any `io.Writer` (file, buffer, HTTP response) |
| `Errorf` | Build an error; `%w` wraps another error |
| `Sscanf`, `Sscan` | Parse values from a string |
| `Stringer` | Interface with `String() string` to customize printing |

The three families (`Print`, `Sprint`, `Fprint`) each come in plain, `f`
(format string) and `ln` (spaces and newline) variants.

## Verb cheat sheet

| Verb | Meaning | Example |
|------|---------|---------|
| `%v` | Default format | `{Ada 36}` |
| `%+v` | Struct with field names | `{Name:Ada Age:36}` |
| `%#v` | Go-syntax representation | `main.User{Name:"Ada", Age:36}` |
| `%T` | Type | `main.User` |
| `%d` `%b` `%o` `%x` | Integer in base 10, 2, 8, 16 | `42` `101010` `52` `2a` |
| `%f` `%.2f` `%e` | Float | `3.141590` `3.14` `3.141590e+00` |
| `%s` `%q` | String, quoted string | `go` `"go"` |
| `%c` `%U` | Character, Unicode code point | `G` `U+0047` |
| `%p` | Pointer address | `0xc000012345` |
| `%t` | Bool | `true` |
| `%%` | Literal percent sign | `%` |

Modifiers go between `%` and the verb: `%5d` (width), `%-5d` (left-align),
`%05d` (zero-pad), `%.2f` (precision), `%*d` (width from an argument),
`%[2]d` (explicit argument index).

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/fmt/
```

## Gotchas
- **Wrong verbs compile fine.** `Printf("%d", "hi")` prints `%!d(string=hi)`.
  `go vet` catches this, and `go test` runs vet automatically.
- **Missing or extra arguments** print `%!d(MISSING)` or `%!(EXTRA ...)`.
- **`Print` vs `Println` spacing.** `Print` adds spaces only between operands
  when neither is a string; `Println` always adds spaces and a newline.
- **`String()` recursion.** Calling `Sprintf("%v", x)` inside `x.String()`
  recurses forever. Convert to the underlying type first.
- **Never pass user input as the format string.** Use `Println(s)` or
  `Printf("%s", s)`.
- **Pointers to structs** print as `&{...}` at the top level but as raw addresses
  when nested.
- **Use `\n`, not `%n`.** Go has no `%n` verb.
- **Map output is sorted by key**, so printed maps are deterministic.
- **`%w`** in `Errorf` wraps an error so `errors.Is` and `errors.As` can see it.
- For reading input line by line, prefer `bufio.Scanner` over `Scan` functions.

## Java comparison

| Java | Go |
|------|----|
| `String.format(...)` | `fmt.Sprintf(...)` |
| `System.out.printf(...)` | `fmt.Printf(...)` |
| `System.out.println(x)` | `fmt.Println(x)` |
| `toString()` | `String() string` (`fmt.Stringer`) |
| `%n` | `\n` |
| `new RuntimeException("x", cause)` | `fmt.Errorf("x: %w", err)` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/fmt
- Source: https://cs.opensource.google/go/go/+/master:src/fmt/
