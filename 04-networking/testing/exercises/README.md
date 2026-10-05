# testing exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./04-networking/testing/exercises/...
```

## Exercise: implement and test

The reference solution (`solution/solution.go`) and its table-driven tests
(`solution/solution_test.go`) show the style. Try the exercise in this order:

1. Read the tests in `solution_test.go` **first**; they are the specification.
2. In a scratch copy of the package, implement:
   - `Clamp(x, lo, hi int) int`, which limits `x` to `[lo, hi]` and panics if `lo > hi`
   - `Median(xs []float64) (float64, error)`, which returns the median without
     modifying `xs`; an empty input returns the exported error `ErrEmpty`
3. Run `go test -v -run TestMedian ./04-networking/testing/exercises/...`.
4. Run `go test -race -cover ./04-networking/testing/exercises/...` and look at the coverage.

Hints: `min` and `max` are built in since Go 1.21; `slices.Clone` and
`slices.Sort` avoid mutating the input.

## Bonus
1. Add a `FuzzMedian` that checks the result always lies between the minimum and maximum of the input.
2. Add a benchmark for `Median` with 1,000 elements and compare with and without the `Clone`.
3. Write a golden-file test that compares output to a file in a `testdata/` directory.
