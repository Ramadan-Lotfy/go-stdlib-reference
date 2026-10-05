# regexp exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/regexp/exercises/...
```

## Exercise 1: ParseLogLine

Write `ParseLogLine(line string) (map[string]string, bool)` for lines like:

```
2026-10-03 ERROR [db] connection lost
```

Return a map with the keys `date`, `level`, `comp` and `msg`, and `false` when
the line does not match.

Hints: a package-level `regexp.MustCompile` with named groups
`(?P<date>...)`, `FindStringSubmatch`, `SubexpNames`.

## Exercise 2: RedactDigits

Write `RedactDigits(s string) string` that replaces every run of four or more
digits with the same number of `*` characters, leaving shorter numbers alone.

`"card 1234567890 id 42"` becomes `"card ********** id 42"`.

Hints: `ReplaceAllStringFunc`, `strings.Repeat`.

## Bonus
Why can't you write a regexp that checks "the same word appears twice in a row"
with RE2? What would you do instead?
