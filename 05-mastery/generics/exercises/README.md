# generics, cmp and slices exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/generics/exercises/...
```

## Exercise: a small generic toolkit

Implement in `solution.go`:

```go
func Map[T, U any](s []T, f func(T) U) []U
func Filter[T any](s []T, keep func(T) bool) []T      // empty non-nil slice when nothing matches
func Reduce[T, A any](s []T, init A, f func(A, T) A) A

type Stack[T any] struct{ ... }                         // Push, Pop() (T, bool), Len
type Set[T comparable] map[T]struct{}                   // Add, Has
func Union[T comparable](a, b Set[T]) Set[T]            // does not modify a or b
func MaxBy[T any, K cmp.Ordered](s []T, key func(T) K) (T, bool)
```

Hints:

- `Pop` on an empty stack returns the zero value of `T` (`var zero T`) and `false`
- `Set` is a named map type, so it can have methods
- `MaxBy` needs `cmp.Ordered` on the key type to use `>`

## Bonus
Write `GroupBy[T any, K comparable](s []T, key func(T) K) map[K][]T` and
`Keys[M ~map[K]V, K comparable, V any](m M) []K`. Why is the `~` needed in `Keys`?
