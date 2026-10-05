# slices

## Overview
`slices` (Go 1.21+) has generic helpers for any slice: sorting, searching,
inserting, deleting, comparing and de-duplicating. It replaces most of
`sort` and a lot of hand-written loops.

## Core API

| Function / Type | Purpose |
|---|---|
| `Sort`, `SortFunc`, `SortStableFunc` | Sort ordered types or with a comparison function |
| `BinarySearch`, `BinarySearchFunc` | Search a sorted slice, returns `(index, found)` |
| `Contains`, `Index`, `IndexFunc`, `ContainsFunc` | Find elements |
| `Equal`, `Compare` | Compare slices element by element |
| `Max`, `Min`, `MaxFunc`, `MinFunc` | Extremes |
| `Insert`, `Delete`, `Replace` | Edit a slice |
| `Reverse`, `Compact`, `CompactFunc` | Reverse, drop consecutive duplicates |
| `Clone`, `Concat`, `Grow`, `Clip` | Copy, join, manage capacity |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/slices/
```

## Gotchas
- **`Compact` only removes *consecutive* duplicates.** Sort first to de-duplicate everything.
- **`Insert`, `Delete` and `Replace` can modify the original array.** Always use the returned slice, and `Clone` first if the original must stay intact.
- **Since Go 1.22, `Delete` zeroes the removed tail** so pointers can be garbage collected. Older versions left stale values behind.
- **`BinarySearch` requires a sorted slice**; on unsorted input the result is meaningless.
- **`Max` and `Min` panic on an empty slice.**
- **Comparison functions return `int`** (negative, zero, positive), like `cmp.Compare`. Combine keys with `cmp.Or`.

## Java comparison

| Java | Go |
|---|---|
| `Collections.sort` | `slices.Sort` |
| `list.contains(x)` | `slices.Contains(s, x)` |
| `list.indexOf(x)` | `slices.Index(s, x)` |
| `Collections.max` | `slices.Max(s)` |
| `Collections.reverse` | `slices.Reverse(s)` |
| `list.removeIf` | `slices.DeleteFunc` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/slices
- Source: https://cs.opensource.google/go/go/+/master:src/slices/
