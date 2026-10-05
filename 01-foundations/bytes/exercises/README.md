# bytes exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./01-foundations/bytes/exercises/...
```

## Exercise 1: NormalizeNewlines

Write `NormalizeNewlines(data []byte) []byte` that converts Windows (`\r\n`)
and old Mac (`\r`) line endings to `\n`.

Rules:

- Return a **new** slice. The input must not be modified.
- The result must not share memory with the input.

| Input | Output |
|-------|--------|
| `"a\r\nb\rc\n"` | `"a\nb\nc\n"` |

Hints: `bytes.ReplaceAll`. Which order do the two replacements need to run in?

## Exercise 2: NumberLines

Write `NumberLines(src []byte) []byte` that prefixes every line with its
number, like `1: `.

Rules:

- Keep a trailing newline if the input has one, and add none if it does not
- Empty lines are numbered too
- Empty input gives empty output

| Input | Output |
|-------|--------|
| `"a\nb\n"` | `"1: a\n2: b\n"` |
| `"a\nb"` | `"1: a\n2: b"` |
| `"\n\n"` | `"1: \n2: \n"` |

Hints: `bytes.Cut` in a loop and a `bytes.Buffer` for the output.

## Bonus

Write `Redact(data, secret []byte) []byte` that replaces every occurrence of
`secret` with `***`. Then make it work on a stream by accepting an `io.Reader`
and `io.Writer`. What happens when a secret straddles two reads?
