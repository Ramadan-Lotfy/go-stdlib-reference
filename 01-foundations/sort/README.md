# sort

## Overview
`sort` sorts slices and user-defined collections, and does binary search.
Since Go 1.21 the generic `slices` package is usually the better choice for
slices, but `sort` is everywhere in existing code and its `Interface` still
matters for custom collections.

## Core API

| Function / Type | Purpose |
|---|---|
| `sort.Ints`, `sort.Strings`, `sort.Float64s` | Sort the common slice types in place |
| `sort.Slice`, `sort.SliceStable` | Sort any slice with a `less(i, j)` function |
| `sort.Sort`, `sort.Stable` | Sort anything implementing `sort.Interface` |
| `sort.Reverse` | Wrap an `Interface` to reverse the order |
| `sort.Search`, `sort.SearchInts` | Binary search for an insertion point |
| `sort.IsSorted`, `sort.SliceIsSorted` | Check order |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/sort/
```

## Gotchas
- **`sort.Slice` is not stable.** Equal elements may be reordered. Use `SliceStable` when ties matter.
- **The `less` function takes indexes, not values.** Use `s[i]` and `s[j]`.
- **`Search` needs a sorted slice** and returns the insertion point, which may be `len(s)`. Check that the element is really there.
- **Prefer `slices.Sort` and `slices.SortFunc`** in new code: they are generic, faster and do not need index closures.
- **Sorting is in place** and mutates the slice. Clone first if you need the original order.

## Java comparison

| Java | Go |
|---|---|
| `Collections.sort(list)` | `slices.Sort(s)` or `sort.Ints(s)` |
| `list.sort(Comparator)` | `sort.Slice(s, less)` or `slices.SortFunc` |
| `Comparator.reversed()` | `sort.Sort(sort.Reverse(...))` |
| `Collections.binarySearch` | `sort.SearchInts` / `slices.BinarySearch` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/sort
- Source: https://cs.opensource.google/go/go/+/master:src/sort/
