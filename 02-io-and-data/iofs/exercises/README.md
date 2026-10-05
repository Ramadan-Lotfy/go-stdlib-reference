# io/fs exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/iofs/exercises/...
```

## Exercise: ListGoFiles

Write `ListGoFiles(fsys fs.FS) ([]string, error)` that returns the path of every
`.go` file in the file system, in walk order.

Rules:

- Skip directories named `vendor`
- Skip directories whose name starts with `.` (like `.git`)
- Do not skip the root `"."`
- Return any walk error

Hints: `fs.WalkDir`, `fs.SkipDir`, package `path` (not `filepath`). Test with
`fstest.MapFS` so no disk is needed.

## Bonus
Add a `maxDepth int` parameter and stop descending below it.
