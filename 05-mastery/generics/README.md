# generics, cmp and slices

## Overview
Generics (Go 1.18+) let functions and types work over many types while
keeping compile-time type safety. A type parameter is declared in square
brackets with a **constraint**: `any`, `comparable`, an interface of allowed
types (`~int | ~float64`), or `cmp.Ordered`. The standard library uses them in
`slices`, `maps`, `cmp`, `sync/atomic` and `iter` (1.23+). This package covers
writing your own generic code and combining it with `cmp` and `slices`.

## Core API

| Function / Type | Purpose |
|---|---|
| `func Map[T, U any](s []T, f func(T) U) []U` | Generic function; type arguments are usually inferred |
| `type Stack[T any] struct { ... }` | Generic type with methods `func (s *Stack[T]) Push(v T)` |
| `any`, `comparable` | Built-in constraints (no operations / `==` allowed) |
| `cmp.Ordered`, `cmp.Compare`, `cmp.Less`, `cmp.Or` | Ordering for `<` and comparison functions |
| `interface{ ~int | ~float64 }` | Type-set constraint; `~` includes named types with that underlying type |
| `[M ~map[K]V, K comparable, V any]` | Constraint that refers to other type parameters |
| `var zero T` | The zero value of a type parameter |
| `slices`, `maps` | Generic standard-library helpers |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/generics/
```

## Gotchas
- **Do not reach for generics too early.** If an interface with a method does the job, use it. Use generics when the same algorithm works over different types (containers, `Map`/`Filter`, `Max`).
- **Methods cannot have their own type parameters.** Only the type itself (`Stack[T]`) can; for `Map` over a `List[T]`, write a top-level function.
- **`any` allows no operations** except assignment. To use `==` you need `comparable`; to use `<` you need `cmp.Ordered` or a type-set constraint.
- **`~int` includes named types** (`type Celsius float64` satisfies `~float64`); plain `int` does not.
- **Type inference is partial.** It works from function arguments, not from return types, so `Zero[int]()` needs the explicit `[int]`.
- **Generic types are instantiated at use:** `Stack[int]` and `Stack[string]` are different types. You cannot have a `Stack[any]` hold both and still get type safety.
- **Comparing a `T` to `nil` does not compile.** Use `var zero T` and `reflect` or a pointer type parameter if you really need it.
- **Generic code can be slower than hand-written code** for pointer-shaped types (Go shares code between them). Benchmark if it is a hot path.

## Java comparison

| Java | Go |
|---|---|
| `<T extends Comparable<T>>` | `[T cmp.Ordered]` |
| `List<T>` | `[]T` or `type List[T any] ...` |
| wildcards `? extends T` | no wildcards: use constraints or interfaces |
| type erasure | no erasure: separate instantiations (`Stack[int]` is a real type) |
| `Collections.max` | `slices.Max` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/generics, cmp and slices
- Source: https://cs.opensource.google/go/go/+/master:src/generics, cmp and slices/
