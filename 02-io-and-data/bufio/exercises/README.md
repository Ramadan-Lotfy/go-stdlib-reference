# bufio exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/bufio/exercises/...
```

## Exercise: LongestLine

Write `LongestLine(r io.Reader) (string, error)` that returns the first longest
line in the input.

Rules:

- Ties keep the earliest line
- Lines up to 1 MiB must work (the default `Scanner` limit is 64 KB)
- Return the scanner's error, if any
- Empty input returns `""` and no error

Hints: `bufio.NewScanner`, `sc.Buffer`, `sc.Err()`.

## Bonus
Rewrite it with `bufio.Reader.ReadString` so there is no line-length limit at all.
What do you have to do differently for the final line?
