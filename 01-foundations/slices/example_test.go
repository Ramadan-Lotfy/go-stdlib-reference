package slicesref_test

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

// Example_sortAndSearch sorts and then binary searches.
func Example_sortAndSearch() {
	s := []int{5, 2, 8, 1}
	slices.Sort(s)
	fmt.Println(s, slices.IsSorted(s))

	idx, found := slices.BinarySearch(s, 5)
	fmt.Println(idx, found)
	idx, found = slices.BinarySearch(s, 3) // not present: insertion point
	fmt.Println(idx, found)
	// Output:
	// [1 2 5 8] true
	// 2 true
	// 2 false
}

type person struct {
	Name string
	Age  int
}

// Example_sortFunc sorts by several keys with cmp.Compare and cmp.Or.
func Example_sortFunc() {
	people := []person{{"Cy", 36}, {"Ada", 36}, {"Bob", 25}}
	slices.SortFunc(people, func(a, b person) int {
		return cmp.Or(
			cmp.Compare(a.Age, b.Age),
			strings.Compare(a.Name, b.Name),
		)
	})
	fmt.Println(people)
	// Output: [{Bob 25} {Ada 36} {Cy 36}]
}

// Example_queries covers the read-only helpers.
func Example_queries() {
	s := []int{3, 1, 4, 1, 5}
	fmt.Println(slices.Contains(s, 4), slices.Index(s, 1), slices.Index(s, 9))
	fmt.Println(slices.IndexFunc(s, func(x int) bool { return x > 3 }))
	fmt.Println(slices.Max(s), slices.Min(s))
	fmt.Println(slices.Equal([]int{1, 2}, []int{1, 2}))
	fmt.Println(slices.Compare([]int{1, 2}, []int{1, 3}))
	// Output:
	// true 1 -1
	// 2
	// 5 1
	// true
	// -1
}

// Example_modify inserts, deletes, reverses, compacts and concatenates.
func Example_modify() {
	s := []int{1, 2, 3, 4, 5}
	s = slices.Insert(s, 2, 99) // insert before index 2
	fmt.Println(s)
	s = slices.Delete(s, 0, 2) // remove s[0:2]
	fmt.Println(s)
	slices.Reverse(s)
	fmt.Println(s)

	fmt.Println(slices.Compact([]int{1, 1, 2, 2, 2, 3, 1})) // consecutive only
	fmt.Println(slices.Concat([]int{1}, []int{2, 3}, nil))
	// Output:
	// [1 2 99 3 4 5]
	// [99 3 4 5]
	// [5 4 3 99]
	// [1 2 3 1]
	// [1 2 3]
}

// Example_aliasing shows why Clone matters: edits can affect the original array.
func Example_aliasing() {
	orig := []int{1, 2, 3}
	c := slices.Clone(orig)
	c[0] = 100
	fmt.Println(orig, c)

	a := []int{1, 2, 3, 4}
	b := slices.Delete(a, 1, 3) // reuses a's array; Go 1.22 zeroes the tail
	fmt.Println(b, a)
	// Output:
	// [1 2 3] [100 2 3]
	// [1 4] [1 4 0 0]
}

// Example_deleteFunc removes elements matching a condition.
func Example_deleteFunc() {
	s := []int{1, 2, 3, 4, 5, 6}
	s = slices.DeleteFunc(s, func(x int) bool { return x%2 == 0 })
	fmt.Println(s)
	// Output: [1 3 5]
}
