# maps

## Overview
Go's built-in `map` type is a hash table with unordered iteration. The `maps`
package adds generic helpers: cloning, copying, comparing and filtering. In
Go 1.22 it has no key/value iterators (those arrived in 1.23), so the idiomatic
way to iterate in order is: collect keys, sort them, loop.

## Core API

| Function / Type | Purpose |
|---|---|
| `map[K]V`, `make`, `delete`, `clear`, `len` | Built-in operations |
| `v, ok := m[k]` | Distinguish a missing key from a zero value |
| `maps.Clone` | Shallow copy |
| `maps.Copy(dst, src)` | Add all entries of `src` to `dst` |
| `maps.Equal`, `maps.EqualFunc` | Compare two maps |
| `maps.DeleteFunc` | Remove entries matching a condition |
| `map[T]struct{}` | The idiomatic set |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/maps/
```

## Gotchas
- **Iteration order is random** and changes between runs. Sort the keys if you need order. `fmt` prints maps sorted by key.
- **Writing to a nil map panics.** Reading from one is fine and returns the zero value. Always `make` before writing.
- **Maps are not safe for concurrent use.** Guard with a `sync.Mutex` or use `sync.Map`.
- **Keys must be comparable**: no slices, maps or functions as keys.
- **Values are not addressable.** `m[k].field = x` does not compile for struct values. Store pointers or reassign the whole value.
- **`maps.Clone` is shallow.** Nested slices and maps are shared with the copy.
- **Deleting during `range` is safe**; adding during `range` may or may not be visited.

## Java comparison

| Java | Go |
|---|---|
| `HashMap<K,V>` | `map[K]V` |
| `map.getOrDefault(k, d)` | `v, ok := m[k]` then default if `!ok` |
| `map.containsKey(k)` | `_, ok := m[k]` |
| `new HashMap<>(other)` | `maps.Clone(other)` |
| `HashSet<T>` | `map[T]struct{}` |
| `TreeMap` (sorted) | map + sorted key slice, or a third-party tree |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/maps
- Source: https://cs.opensource.google/go/go/+/master:src/maps/
