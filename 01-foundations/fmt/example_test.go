package fmtref_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

type User struct {
	Name string
	Age  int
}

// Celsius implements fmt.Stringer, so every fmt verb that prints text
// (%v, %s, Println...) uses String() automatically.
type Celsius float64

func (c Celsius) String() string {
	// Convert to float64 first. Calling Sprintf("%v", c) here would call
	// String() again and recurse forever.
	return fmt.Sprintf("%.1f°C", float64(c))
}

// Example_structVerbs shows how %v, %+v, %#v and %T describe the same value.
func Example_structVerbs() {
	u := User{"Ada", 36}
	fmt.Printf("%v\n", u)  // default format
	fmt.Printf("%+v\n", u) // adds field names
	fmt.Printf("%#v\n", u) // Go-syntax representation
	fmt.Printf("%T\n", u)  // type
	// Output:
	// {Ada 36}
	// {Name:Ada Age:36}
	// fmtref_test.User{Name:"Ada", Age:36}
	// fmtref_test.User
}

// Example_collections shows how slices and maps print. Map keys are sorted.
func Example_collections() {
	s := []int{1, 2, 3}
	m := map[string]int{"b": 2, "a": 1}
	fmt.Printf("%v %#v\n", s, s)
	fmt.Printf("%v %#v\n", m, m)
	// Output:
	// [1 2 3] []int{1, 2, 3}
	// map[a:1 b:2] map[string]int{"a":1, "b":2}
}

// Example_numbers covers width, padding, bases and floats.
func Example_numbers() {
	fmt.Printf("%d|%5d|%-5d|%05d\n", 42, 42, 42, 42)
	fmt.Printf("%b %o %x %X\n", 10, 10, 255, 255)
	fmt.Printf("%c %U\n", 'G', 'G')
	fmt.Printf("%.2f|%8.3f|%e\n", 3.14159, 3.14159, 123456.789)
	// Output:
	// 42|   42|42   |00042
	// 1010 12 ff FF
	// G U+0047
	// 3.14|   3.142|1.234568e+05
}

// Example_strings covers %s, %q, width and hex encoding.
func Example_strings() {
	fmt.Printf("%s|%q|%6s|%-6s|\n", "go", "go", "go", "go")
	fmt.Printf("%x\n", "hi")
	// Output:
	// go|"go"|    go|go    |
	// 6869
}

// Example_argumentControl shows explicit argument indexes and width from args.
func Example_argumentControl() {
	fmt.Printf("%[2]d %[1]d\n", 1, 2) // reorder arguments
	fmt.Printf("%*d|\n", 6, 42)       // width supplied as an argument
	// Output:
	// 2 1
	//     42|
}

// Example_sprintfAndFprintf shows the three families: Printf writes to
// stdout, Sprintf returns a string, Fprintf writes to any io.Writer.
func Example_sprintfAndFprintf() {
	s := fmt.Sprintf("user=%s age=%d", "ada", 36)
	fmt.Println(s)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "id=%03d", 7)
	fmt.Println(buf.String())

	fmt.Fprintln(os.Stdout, "written to os.Stdout")
	// Output:
	// user=ada age=36
	// id=007
	// written to os.Stdout
}

// Example_printVersusPrintln shows the spacing rules.
func Example_printVersusPrintln() {
	// Print adds a space between operands only when neither is a string.
	fmt.Print("a", 1, 2, "b", "\n")
	// Println always adds spaces between operands and a trailing newline.
	fmt.Println("a", 1, 2, "b")
	// Output:
	// a1 2b
	// a 1 2 b
}

// Example_stringer shows a custom type controlling its own formatting.
func Example_stringer() {
	t := Celsius(36.6)
	fmt.Println(t)
	fmt.Printf("%v | %s\n", t, t)
	// Output:
	// 36.6°C
	// 36.6°C | 36.6°C
}

// Example_errorf shows error wrapping with %w.
func Example_errorf() {
	err := fmt.Errorf("open config: %w", io.EOF)
	fmt.Println(err)
	fmt.Println(errors.Is(err, io.EOF))
	// Output:
	// open config: EOF
	// true
}

// Example_scanning parses values out of a string.
func Example_scanning() {
	var name string
	var age int
	n, err := fmt.Sscanf("ada 36", "%s %d", &name, &age)
	fmt.Println(n, err, name, age)

	var a, b int
	fmt.Sscan("10 20", &a, &b)
	fmt.Println(a + b)
	// Output:
	// 2 <nil> ada 36
	// 30
}

// Example_misuse shows what happens with a wrong verb or missing argument.
// fmt never panics or fails to compile; it prints a diagnostic instead.
// The format is held in a variable only so that go vet does not reject this
// deliberate mistake.
func Example_misuse() {
	wrongVerb := "%d\n"
	missingArg := "%d %d\n"
	fmt.Printf(wrongVerb, "hi")
	fmt.Printf(missingArg, 1)
	// Output:
	// %!d(string=hi)
	// 1 %!d(MISSING)
}
