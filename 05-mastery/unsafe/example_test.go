package unsaferef_test

import (
	"fmt"
	"math"
	"unsafe"
)

type wasteful struct {
	a bool
	b int64
	c bool
}

type compact struct {
	b    int64
	a, c bool
}

// Example_sizes prints the memory size of common types (64-bit platform).
func Example_sizes() {
	fmt.Println(unsafe.Sizeof(int32(0)), unsafe.Sizeof(int64(0)))
	fmt.Println(unsafe.Sizeof(""), unsafe.Sizeof([]int{}), unsafe.Sizeof(any(nil)))
	// Output:
	// 4 8
	// 16 24 16
}

// Example_padding shows how field order changes size through alignment padding.
func Example_padding() {
	fmt.Println(unsafe.Sizeof(wasteful{}), unsafe.Sizeof(compact{}))

	var w wasteful
	fmt.Println(unsafe.Offsetof(w.a), unsafe.Offsetof(w.b), unsafe.Offsetof(w.c))
	fmt.Println(unsafe.Alignof(w.b))
	// Output:
	// 24 16
	// 0 8 16
	// 8
}

// Example_pointerMath reads the second element of an array through a raw pointer.
func Example_pointerMath() {
	arr := [3]int32{10, 20, 30}
	p := unsafe.Pointer(&arr[0])
	second := (*int32)(unsafe.Add(p, unsafe.Sizeof(arr[0])))
	fmt.Println(*second)

	s := unsafe.Slice(&arr[0], 2) // a slice view over the first two elements
	fmt.Println(s)
	// Output:
	// 20
	// [10 20]
}

// Example_zeroCopy converts bytes to a string without copying, and shows the
// danger: the string changes when the bytes change.
func Example_zeroCopy() {
	b := []byte("hello")
	s := unsafe.String(&b[0], len(b))
	fmt.Println(s)

	b[0] = 'J'
	fmt.Println(s) // the "immutable" string changed
	// Output:
	// hello
	// Jello
}

// Example_floatBits reinterprets memory. math.Float32bits does the same, safely.
func Example_floatBits() {
	f := float32(1.0)
	bits := *(*uint32)(unsafe.Pointer(&f))
	fmt.Printf("%#x %v\n", bits, bits == math.Float32bits(f))
	// Output: 0x3f800000 true
}
