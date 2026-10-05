package sortref_test

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

// Example_basic sorts the common slice types and reverses an order.
func Example_basic() {
	ints := []int{3, 1, 2}
	words := []string{"pear", "apple", "fig"}
	sort.Ints(ints)
	sort.Strings(words)
	fmt.Println(ints, words)

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)
	fmt.Println(sort.IntsAreSorted(ints))
	// Output:
	// [1 2 3] [apple fig pear]
	// [3 2 1]
	// false
}

// Example_stable shows that SliceStable keeps the original order of equal elements.
func Example_stable() {
	people := []person{{"Ada", 36}, {"Bob", 25}, {"Cy", 36}, {"Di", 25}}
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
	// Output: [{Bob 25} {Di 25} {Ada 36} {Cy 36}]
}

type byLen []string

func (s byLen) Len() int           { return len(s) }
func (s byLen) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
func (s byLen) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Example_interface sorts a custom type by implementing sort.Interface.
func Example_interface() {
	words := byLen{"ccc", "a", "bb"}
	sort.Sort(words)
	fmt.Println(words)
	// Output: [a bb ccc]
}

// Example_search finds insertion points with binary search.
func Example_search() {
	a := []int{1, 3, 5, 7}
	fmt.Println(sort.SearchInts(a, 5)) // found at index 2
	fmt.Println(sort.SearchInts(a, 4)) // not present: insertion point is 2
	fmt.Println(sort.SearchInts(a, 9)) // beyond the end: len(a)

	i := sort.Search(len(a), func(i int) bool { return a[i] >= 6 })
	fmt.Println(i, i < len(a) && a[i] == 6)
	// Output:
	// 2
	// 2
	// 4
	// 3 false
}
