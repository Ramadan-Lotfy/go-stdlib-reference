// Package solution holds the reference solutions for the unsafe exercises.
package solution

import "unsafe"

// Wasteful has padding between its fields.
type Wasteful struct {
	A bool
	B int64
	C bool
	D int32
	E bool
}

// Compact has the same fields as Wasteful, reordered so it is as small as
// possible on a 64-bit platform (16 bytes instead of 32).
type Compact struct {
	B int64
	D int32
	A bool
	C bool
	E bool
}

// BytesToString converts b to a string without copying. The caller must not
// modify b afterwards.
func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(&b[0], len(b))
}

// Sizes returns the sizes of Wasteful and Compact.
func Sizes() (wasteful, compact uintptr) {
	return unsafe.Sizeof(Wasteful{}), unsafe.Sizeof(Compact{})
}
