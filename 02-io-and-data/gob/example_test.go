package gobref_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Point struct{ X, Y int }

// Example_roundTrip encodes a struct to bytes and decodes it back.
func Example_roundTrip() {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(Point{1, 2}); err != nil {
		fmt.Println(err)
		return
	}

	var p Point
	err := gob.NewDecoder(&buf).Decode(&p)
	fmt.Println(p, err)
	// Output: {1 2} <nil>
}

// Example_stream sends several values through one encoder and reads them in order.
func Example_stream() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode("hello")
	enc.Encode(42)
	enc.Encode([]string{"a", "b"})

	dec := gob.NewDecoder(&buf)
	var s string
	var n int
	var list []string
	dec.Decode(&s)
	dec.Decode(&n)
	dec.Decode(&list)
	fmt.Println(s, n, list)
	// Output: hello 42 [a b]
}

type Shape interface{ Area() float64 }
type Rect struct{ W, H float64 }
type Square struct{ S float64 }

func (r Rect) Area() float64   { return r.W * r.H }
func (s Square) Area() float64 { return s.S * s.S }

func init() {
	// Concrete types sent as interface values must be registered.
	gob.Register(Rect{})
	gob.Register(Square{})
}

// Example_interface sends different concrete types through an interface slice.
func Example_interface() {
	var buf bytes.Buffer
	in := []Shape{Rect{2, 3}, Square{4}}
	if err := gob.NewEncoder(&buf).Encode(in); err != nil {
		fmt.Println(err)
		return
	}

	var out []Shape
	if err := gob.NewDecoder(&buf).Decode(&out); err != nil {
		fmt.Println(err)
		return
	}
	for _, s := range out {
		fmt.Printf("%T %.0f\n", s, s.Area())
	}
	// Output:
	// gobref_test.Rect 6
	// gobref_test.Square 16
}

// Example_zeroValues shows the classic gotcha: zero fields are not sent, so
// decoding into a non-empty value keeps its old data.
func Example_zeroValues() {
	type T struct{ A, B int }
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(T{A: 0, B: 5})

	dst := T{A: 9}
	gob.NewDecoder(&buf).Decode(&dst)
	fmt.Println(dst) // A is still 9
	// Output: {9 5}
}

// Example_errors shows a type mismatch between sender and receiver.
func Example_errors() {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode("text")

	var n int
	err := gob.NewDecoder(&buf).Decode(&n)
	fmt.Println(err)
	// Output: gob: decoding into local type *int, received remote type string
}
