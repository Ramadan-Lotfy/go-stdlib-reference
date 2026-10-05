package solution

import (
	"strconv"
	"testing"
	"unsafe"
)

func TestSizes(t *testing.T) {
	if strconv.IntSize != 64 {
		t.Skip("sizes below assume a 64-bit platform")
	}
	w, c := Sizes()
	if w != 32 || c != 16 {
		t.Errorf("Sizes() = %d, %d; want 32, 16", w, c)
	}
}

func TestFieldsMatch(t *testing.T) {
	// Compact must hold exactly the same field types as Wasteful.
	var w Wasteful
	var c Compact
	if unsafe.Sizeof(w.A) != unsafe.Sizeof(c.A) || unsafe.Sizeof(w.B) != unsafe.Sizeof(c.B) ||
		unsafe.Sizeof(w.D) != unsafe.Sizeof(c.D) {
		t.Error("Compact must keep the same field types as Wasteful")
	}
}

func TestBytesToString(t *testing.T) {
	if got := BytesToString([]byte("hi")); got != "hi" {
		t.Errorf("BytesToString() = %q", got)
	}
	if got := BytesToString(nil); got != "" {
		t.Errorf("BytesToString(nil) = %q", got)
	}
	if got := BytesToString([]byte{}); got != "" {
		t.Errorf("BytesToString(empty) = %q", got)
	}
}
