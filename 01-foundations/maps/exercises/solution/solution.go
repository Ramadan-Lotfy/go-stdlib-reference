// Package solution holds the reference solutions for the maps exercises.
package solution

import (
	"cmp"
	"slices"
)

// GroupBy groups items by the key returned by key. Order within each group
// follows the order of items.
func GroupBy[K comparable, V any](items []V, key func(V) K) map[K][]V {
	groups := make(map[K][]V)
	for _, it := range items {
		k := key(it)
		groups[k] = append(groups[k], it)
	}
	return groups
}

// SortedKeys returns the keys of m in ascending order.
func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
