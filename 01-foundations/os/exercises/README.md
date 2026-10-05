# os exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./01-foundations/os/exercises/...
```

## Exercise 1: WordCountFile

Write `WordCountFile(path string) (int, error)` that returns the number of
words in a file. Wrap any error with `%w` so callers can still use
`errors.Is(err, fs.ErrNotExist)`.

Hints: `os.ReadFile`, `strings.Fields`.

## Exercise 2: AppendLine

Write `AppendLine(path, line string) error` that adds `line` and a newline to
the end of a file, creating the file if it does not exist.

Rules:

- Use `os.OpenFile` with the right flags (`O_APPEND`, `O_CREATE`, `O_WRONLY`)
- Return the error from `Close`

Hints: tests use `t.TempDir()`, which Go deletes automatically.

## Bonus
Write `CopyFile(dst, src string) error` with `os.Open`, `os.Create` and `io.Copy`.
