# io/fs

## Overview
`io/fs` defines `fs.FS`, a read-only file system interface with a single
method, `Open(name)`. Real directories (`os.DirFS`), embedded files (`embed.FS`),
zip archives and in-memory test file systems all implement it, so code that takes
an `fs.FS` is easy to test and works with all of them.

## Core API

| Function / Type | Purpose |
|---|---|
| `fs.FS`, `fs.File`, `fs.DirEntry`, `fs.FileInfo` | The core interfaces |
| `fs.ReadFile`, `fs.ReadDir`, `fs.Stat` | Helper functions that work on any `FS` |
| `fs.WalkDir` | Walk a tree in lexical order; return `fs.SkipDir` to skip |
| `fs.Glob`, `fs.Sub` | Pattern matching and sub-trees |
| `fs.ValidPath` | Check a path name (slash-separated, no `..`) |
| `os.DirFS(dir)` | Real directory as an `fs.FS` |
| `testing/fstest.MapFS` | In-memory `FS` for tests |
| `embed.FS` | Files compiled into the binary |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/iofs/
```

## Gotchas
- **Paths are always slash-separated and unrooted**: `dir/file.txt`, never `\dir`, `/dir` or `./x`. Use package `path`, not `filepath`, with an `fs.FS`.
- **The root is `"."`.** `fs.ReadDir(fsys, ".")` lists the top level.
- **`os.DirFS` is read-only and does not stop symlink escapes.** Do not treat it as a security sandbox.
- **`fs.WalkDir` calls the callback with `err != nil`** if a directory cannot be read. Handle it; `d` may be nil when the root itself fails.
- **Use `fstest.MapFS` in tests** instead of touching the disk, and `fstest.TestFS` to verify your own `FS` implementation.
- **`fs.Sub` does not copy.** It is a view onto the same files.

## Java comparison

| Java | Go |
|---|---|
| `java.nio.file.FileSystem` | `fs.FS` |
| `Files.walk` | `fs.WalkDir` |
| `Class.getResourceAsStream` | `embed.FS` + `fs.ReadFile` |
| `Jimfs` (in-memory FS for tests) | `fstest.MapFS` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/io/fs
- Source: https://cs.opensource.google/go/go/+/master:src/io/fs/
