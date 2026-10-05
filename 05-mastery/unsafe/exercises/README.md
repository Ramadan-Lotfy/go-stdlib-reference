# unsafe exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/unsafe/exercises/...
```

## Exercise 1: reorder a struct

`Wasteful` has five fields (`bool`, `int64`, `bool`, `int32`, `bool`) and takes
32 bytes on a 64-bit platform because of padding. Define `Compact` with the
**same fields** in a better order so it takes 16 bytes.

Check your answer with `unsafe.Sizeof` and `unsafe.Offsetof`. Why does the
order matter? (Each field must start at a multiple of its alignment, and the
struct size is rounded up to its largest alignment.)

## Exercise 2: BytesToString

Write `BytesToString(b []byte) string` that converts without copying, using
`unsafe.String`. Handle empty and nil slices (taking `&b[0]` would panic).

Then write down, in a comment, one situation where this is unsafe.

## Bonus
Benchmark `string(b)` against your `BytesToString` for 1 KB and 1 MB slices
(`go test -bench . -benchmem`). When does the zero-copy version win?
