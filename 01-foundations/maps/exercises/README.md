# maps exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./01-foundations/maps/exercises/...
```

## Exercise 1: GroupBy

Write a generic `GroupBy[K comparable, V any](items []V, key func(V) K) map[K][]V`.
Items in each group keep their original order.

Example: `GroupBy([]string{"a","bb","cc","d"}, len)` gives `map[1:[a d] 2:[bb cc]]`.

## Exercise 2: SortedKeys

Write a generic `SortedKeys[K cmp.Ordered, V any](m map[K]V) []K` that returns
the keys in ascending order.

Hints: `make([]K, 0, len(m))`, `slices.Sort`.

## Bonus
Write `Invert[K, V comparable](m map[K]V) map[V]K`. What should happen when two
keys share a value?
