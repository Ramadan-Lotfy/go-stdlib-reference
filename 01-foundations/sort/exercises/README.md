# sort exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./01-foundations/sort/exercises/...
```

## Exercise: TopN

Write `TopN(words []string, n int) []string` that returns the `n` most
frequent words.

Rules:

- Most frequent first
- Ties are broken alphabetically
- If `n` is larger than the number of distinct words, return all of them
- `n <= 0` returns an empty (non-nil or nil) slice

Example: words `b a c b a b d`, `n = 3` gives `[b a c]` (`b` x3, `a` x2, then `c` before `d`).

Hints: count with a map, copy the keys into a slice, then `sort.Slice`.

## Bonus
Rewrite it with `slices.SortFunc` and `cmp.Or`. Which version reads better?
