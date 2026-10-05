# path/filepath

## Overview
`path/filepath` manipulates file names in the way the current operating system
expects: backslashes on Windows, slashes elsewhere. Its sibling `path` always
uses forward slashes and is for URLs and `fs.FS` names. Most functions here are
pure string operations that never touch the disk.

## Core API

| Function / Type | Purpose |
|---|---|
| `filepath.Join`, `Split`, `Dir`, `Base`, `Ext` | Build and take apart paths |
| `filepath.Clean` | Normalize `a//b/./c/..` to `a/b` |
| `filepath.Abs`, `Rel` | Make a path absolute, or relative to a base |
| `filepath.IsLocal` | Is the path safely inside the current directory? (Go 1.20+) |
| `filepath.Match`, `Glob` | Shell-style patterns |
| `filepath.WalkDir` | Walk a directory tree (faster than `Walk`) |
| `filepath.ToSlash`, `FromSlash` | Convert separators |
| `path.Join`, `path.Base` | Same ideas, always `/` |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/filepath/
```

## Gotchas
- **Never build paths with `+` or `"/"`.** Use `filepath.Join`.
- **`Join` cleans its result**, so `Join("a", "../b")` is `b`. That is how path traversal attacks sneak through; validate user input with `IsLocal` **before** joining.
- **`Ext` includes the dot** (`.gz`) and only returns the last extension of `archive.tar.gz`.
- **`Base("")` is `.`, and `Dir("file")` is `.`** Edge cases return a dot, not an empty string.
- **`Glob` and `Match` have no `**`.** Use `WalkDir` for recursion.
- **Prefer `WalkDir` over `Walk`**; it avoids a `Stat` per file. Entries are visited in lexical order.
- **Use `path` (not `filepath`) for URL paths and `fs.FS` names**, so they stay slash-separated on Windows.
- **These examples print `filepath.ToSlash(...)`** so output is identical on every OS.

## Java comparison

| Java | Go |
|---|---|
| `Paths.get(a, b)`, `Path.resolve` | `filepath.Join(a, b)` |
| `path.getFileName()` | `filepath.Base(p)` |
| `path.getParent()` | `filepath.Dir(p)` |
| `path.relativize` | `filepath.Rel(base, p)` |
| `path.normalize()` | `filepath.Clean(p)` |
| `Files.walkFileTree` | `filepath.WalkDir` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/path/filepath
- Source: https://cs.opensource.google/go/go/+/master:src/path/filepath/
