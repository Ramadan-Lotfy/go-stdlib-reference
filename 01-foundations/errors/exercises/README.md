# errors exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./01-foundations/errors/exercises/...
```

## Exercise 1: Classify

Write `Classify(err error) string` that returns:

| Error | Result |
|-------|--------|
| `nil` | `ok` |
| wraps `ErrNotFound` | `not found` |
| wraps a `*ValidationError` | `invalid: <field>` |
| anything else | `unknown` |

Hints: `errors.Is`, `errors.As`, and a `switch` with no tag.

## Exercise 2: Retry

Write `Retry(attempts int, fn func() error) error`.

- Call `fn` up to `attempts` times and return `nil` at the first success
- If every attempt fails, return all errors joined, each wrapped as
  `attempt N: <err>` so `errors.Is` still finds the cause
- `attempts <= 0` returns an error

Hints: `fmt.Errorf` with `%w`, `errors.Join`.

## Bonus
Add a `Permanent` error type. If `fn` returns one, stop retrying immediately.
