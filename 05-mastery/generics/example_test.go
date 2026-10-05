package genericsref_test

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
)

func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

func Filter[T any](s []T, keep func(T) bool) []T {
	var out []T
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func Reduce[T, A any](s []T, init A, f func(A, T) A) A {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

// Example_functions chains generic helpers. Type arguments are inferred.
func Example_functions() {
	nums := []int{1, 2, 3, 4, 5}
	strs := Map(nums, strconv.Itoa)
	fmt.Printf("%q\n", strs)

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens)

	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println(sum)

	// Explicit instantiation, when inference is not possible or for clarity.
	fmt.Println(Map[int, float64](nums[:2], func(n int) float64 { return float64(n) / 2 }))
	// Output:
	// ["1" "2" "3" "4" "5"]
	// [2 4]
	// 15
	// [0.5 1]
}

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](xs []T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total
}

type Celsius float64

// Example_constraints uses a type-set constraint; ~ admits named types.
func Example_constraints() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]float64{0.5, 0.25}))
	fmt.Println(Sum([]Celsius{20.5, 1.5})) // Celsius has underlying type float64
	// Output:
	// 6
	// 0.75
	// 22
}

type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func (s *Stack[T]) Pop() (T, bool) {
	var zero T // the zero value of any T
	if len(s.items) == 0 {
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

// Example_genericType is a type-safe stack for any element type.
func Example_genericType() {
	var s Stack[string]
	s.Push("a")
	s.Push("b")
	v, ok := s.Pop()
	fmt.Println(v, ok)
	s.Pop()
	v, ok = s.Pop()
	fmt.Printf("%q %v\n", v, ok)
	// Output:
	// b true
	// "" false
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T)      { s[v] = struct{}{} }
func (s Set[T]) Has(v T) bool { _, ok := s[v]; return ok }

// Keys returns the keys of any map type, including named map types.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Example_comparable needs `comparable` for map keys and == checks.
func Example_comparable() {
	s := Set[string]{}
	s.Add("go")
	fmt.Println(s.Has("go"), s.Has("rust"), len(s))

	keys := Keys(Set[int]{3: {}, 1: {}, 2: {}})
	slices.Sort(keys)
	fmt.Println(keys)
	// Output:
	// true false 1
	// [1 2 3]
}

func MaxOf[T cmp.Ordered](first T, rest ...T) T {
	m := first
	for _, v := range rest {
		m = max(m, v)
	}
	return m
}

type Pair[K comparable, V any] struct {
	Key K
	Val V
}

func (p Pair[K, V]) String() string { return fmt.Sprintf("%v=%v", p.Key, p.Val) }

// Example_cmp orders values with cmp.Ordered, cmp.Compare and cmp.Or.
func Example_cmp() {
	fmt.Println(MaxOf(3, 9, 4), MaxOf("b", "a"), MaxOf(2.5))
	fmt.Println(cmp.Compare(1, 2), cmp.Compare("b", "a"), cmp.Less(1.5, 2.5))

	pairs := []Pair[string, int]{{"b", 2}, {"a", 2}, {"c", 1}}
	slices.SortFunc(pairs, func(x, y Pair[string, int]) int {
		return cmp.Or(
			cmp.Compare(x.Val, y.Val), // by value first
			cmp.Compare(x.Key, y.Key), // then by key
		)
	})
	fmt.Println(pairs)
	fmt.Println(cmp.Or("", "default"), cmp.Or(0, 0, 7))
	// Output:
	// 9 b 2.5
	// -1 1 true
	// [c=1 a=2 b=2]
	// default 7
}

func First[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[0], true
}

func Zero[T any]() T {
	var z T
	return z
}

// Example_inference shows what inference can and cannot do.
func Example_inference() {
	v, ok := First([]string{"x", "y"}) // T inferred from the argument
	fmt.Println(v, ok)

	n, ok := First([]int(nil))
	fmt.Println(n, ok)

	// Nothing to infer from: the type argument must be explicit.
	fmt.Printf("%d %q %v\n", Zero[int](), Zero[string](), Zero[*int]())
	// Output:
	// x true
	// 0 false
	// 0 "" <nil>
}

// Example_stdlib uses the generic slices helpers built on the same features.
func Example_stdlib() {
	words := []string{"banana", "kiwi", "apple"}
	slices.SortFunc(words, func(a, b string) int {
		return cmp.Or(cmp.Compare(len(a), len(b)), cmp.Compare(a, b))
	})
	fmt.Println(words)
	fmt.Println(slices.MaxFunc(words, func(a, b string) int { return cmp.Compare(len(a), len(b)) }))
	i, found := slices.BinarySearchFunc(words, "banana", func(e, t string) int {
		return cmp.Or(cmp.Compare(len(e), len(t)), cmp.Compare(e, t))
	})
	fmt.Println(i, found)
	// Output:
	// [kiwi apple banana]
	// banana
	// 2 true
}
