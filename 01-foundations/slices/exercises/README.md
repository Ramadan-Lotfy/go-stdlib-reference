# slices exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./01-foundations/slices/exercises/...
```

## Exercise 1: Unique

Write `Unique(s []int) []int` that returns the sorted distinct values of `s`
**without modifying `s`**.

Example: `[3 1 3 2 1]` gives `[1 2 3]`.

Hints: `slices.Clone`, `slices.Sort`, `slices.Compact`.

## Exercise 2: Chunk

Write a generic `Chunk[T any](s []T, n int) [][]T` that splits `s` into
consecutive pieces of at most `n` elements. `n <= 0` returns `nil`.

Example: `Chunk([]int{1,2,3,4,5}, 2)` gives `[[1 2] [3 4] [5]]`.

Hints: reslice in a loop, and use the full slice expression `s[:size:size]` so
appending to one chunk cannot overwrite the next.

## Bonus
Write `GroupConsecutive` that groups equal neighbors, such as `[1 1 2 2 2 1]` into `[[1 1] [2 2 2] [1]]`.
