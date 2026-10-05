package mapsref_test

import (
	"fmt"
	"maps"
	"slices"
)

// Example_basics covers the comma-ok idiom, delete and clear.
func Example_basics() {
	m := map[string]int{"a": 1, "b": 2}
	v, ok := m["a"]
	fmt.Println(v, ok)
	v, ok = m["zzz"]
	fmt.Println(v, ok) // zero value, false

	delete(m, "a")
	fmt.Println(m, len(m))
	clear(m)
	fmt.Println(m, len(m))
	// Output:
	// 1 true
	// 0 false
	// map[b:2] 1
	// map[] 0
}

// Example_cloneCopyEqual uses the maps package helpers.
func Example_cloneCopyEqual() {
	m := map[string]int{"a": 1}
	c := maps.Clone(m)
	c["b"] = 2
	fmt.Println(m, c, maps.Equal(m, c))

	maps.Copy(m, c) // m now has everything c has
	fmt.Println(maps.Equal(m, c))

	maps.DeleteFunc(m, func(k string, v int) bool { return v%2 == 0 })
	fmt.Println(m)
	// Output:
	// map[a:1] map[a:1 b:2] false
	// true
	// map[a:1]
}

// Example_sortedKeys iterates in a deterministic order.
func Example_sortedKeys() {
	m := map[string]int{"banana": 2, "apple": 1, "cherry": 3}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
	// Output:
	// apple 1
	// banana 2
	// cherry 3
}

// Example_set uses map[T]struct{} as a set.
func Example_set() {
	set := map[string]struct{}{}
	for _, s := range []string{"a", "b", "a", "c", "b"} {
		set[s] = struct{}{}
	}
	_, has := set["a"]
	_, hasZ := set["z"]
	fmt.Println(len(set), has, hasZ)
	// Output: 3 true false
}

// Example_countAndGroup builds frequency and grouping maps.
func Example_countAndGroup() {
	counts := map[string]int{}
	for _, w := range []string{"go", "is", "go", "fun"} {
		counts[w]++ // missing keys start at the zero value
	}
	fmt.Println(counts)

	byLen := map[int][]string{}
	for _, w := range []string{"a", "bb", "cc", "d"} {
		byLen[len(w)] = append(byLen[len(w)], w)
	}
	fmt.Println(byLen)
	// Output:
	// map[fun:1 go:2 is:1]
	// map[1:[a d] 2:[bb cc]]
}

// Example_nilMap shows that a nil map can be read but not written.
func Example_nilMap() {
	var m map[string]int
	fmt.Println(m["x"], len(m), m == nil)

	defer func() { fmt.Println("recovered:", recover()) }()
	m["x"] = 1
	// Output:
	// 0 0 true
	// recovered: assignment to entry in nil map
}

// Example_structValues shows why you cannot assign to a field of a map value.
func Example_structValues() {
	type point struct{ X, Y int }
	m := map[string]point{"p": {1, 2}}
	// m["p"].X = 10 would not compile: map values are not addressable.
	p := m["p"]
	p.X = 10
	m["p"] = p
	fmt.Println(m)
	// Output: map[p:{10 2}]
}
