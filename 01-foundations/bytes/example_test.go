package bytesref_test

import (
	"bytes"
	"fmt"
	"io"
)

// Example_mirrorsStrings shows that bytes offers the same functions as
// strings, but for []byte values.
func Example_mirrorsStrings() {
	data := []byte("  Go is fun  ")
	fmt.Printf("%q\n", bytes.TrimSpace(data))
	fmt.Println(bytes.Contains(data, []byte("is")))
	fmt.Println(bytes.HasPrefix(data, []byte("  Go")))
	fmt.Println(bytes.Index(data, []byte("fun")), bytes.IndexByte(data, 'G'))
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))

	// %q on a [][]byte shows each element quoted.
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Fields([]byte(" a  b ")))

	// %s prints a []byte as text.
	fmt.Printf("%s\n", bytes.Join([][]byte{[]byte("a"), []byte("b")}, []byte("-")))
	fmt.Printf("%s\n", bytes.ReplaceAll([]byte("oink oink"), []byte("oink"), []byte("moo")))

	before, after, found := bytes.Cut([]byte("key=value"), []byte("="))
	fmt.Printf("%s %s %t\n", before, after, found)
	// Output:
	// "Go is fun"
	// true
	// true
	// 8 2
	// 3
	// ["a" "b" "c"]
	// ["a" "b"]
	// a-b
	// moo moo
	// key value true
}

// Example_equal shows how to compare byte slices. The == operator only
// compares a slice with nil, so use bytes.Equal or bytes.Compare.
func Example_equal() {
	a := []byte("go")
	b := []byte("go")
	fmt.Println(bytes.Equal(a, b))
	fmt.Println(bytes.Equal(nil, []byte{})) // nil and empty are equal
	fmt.Println(bytes.Compare([]byte("a"), []byte("b")))
	fmt.Println(string(a) == string(b)) // also fine, and the compiler avoids a copy
	// Output:
	// true
	// true
	// -1
	// true
}

// Example_buffer shows the zero-value Buffer: write bytes, strings, runes
// and formatted text, then read the result.
func Example_buffer() {
	var buf bytes.Buffer // ready to use, no constructor needed
	buf.WriteString("hello ")
	buf.Write([]byte("world"))
	buf.WriteByte('!')
	fmt.Fprintf(&buf, " (%d)", 42) // Buffer is an io.Writer
	fmt.Println(buf.String(), buf.Len())
	// Output:
	// hello world! (42) 17
}

// Example_bufferAsReader shows that reading from a Buffer consumes the data.
func Example_bufferAsReader() {
	buf := bytes.NewBufferString("line1\nline2\nrest")

	line, _ := buf.ReadString('\n')
	fmt.Printf("%q\n", line)

	b, _ := buf.ReadByte()
	fmt.Println(string(b))

	fmt.Printf("%q\n", buf.Next(3)) // next 3 bytes, valid until the next read or write
	fmt.Println(buf.Len())          // Len is the number of UNREAD bytes

	l2, err := buf.ReadString('\n')
	fmt.Printf("%q %v\n", l2, err)

	// No delimiter left: you get the remaining data and io.EOF.
	tail, err := buf.ReadString('\n')
	fmt.Printf("%q %v\n", tail, err)
	// Output:
	// "line1\n"
	// l
	// "ine"
	// 6
	// "2\n" <nil>
	// "rest" EOF
}

// Example_bufferRead uses the lower-level Read method and the io.EOF signal.
func Example_bufferRead() {
	buf := bytes.NewBufferString("abc")
	p := make([]byte, 2)

	n, err := buf.Read(p)
	fmt.Println(n, err, string(p[:n]))

	n, err = buf.Read(p)
	fmt.Println(n, err, string(p[:n]))

	n, err = buf.Read(p)
	fmt.Println(n, err)
	// Output:
	// 2 <nil> ab
	// 1 <nil> c
	// 0 EOF
}

// greet writes to any io.Writer, which makes it easy to test.
func greet(w io.Writer, name string) {
	fmt.Fprintf(w, "hello, %s\n", name)
}

// Example_bufferAsWriter captures the output of code that writes to an
// io.Writer. This is the standard way to unit test such code.
func Example_bufferAsWriter() {
	var buf bytes.Buffer
	greet(&buf, "ada")
	fmt.Printf("%q\n", buf.String())
	// Output:
	// "hello, ada\n"
}

// Example_bufferManagement shows Truncate and Reset. Reset keeps the
// allocated memory, so a Buffer can be reused cheaply.
func Example_bufferManagement() {
	var buf bytes.Buffer
	buf.WriteString("hello world")
	buf.Truncate(5) // keep only the first 5 unread bytes
	fmt.Println(buf.String())

	buf.Reset()
	fmt.Println(buf.Len())
	// Output:
	// hello
	// 0
}

// Example_copies shows that converting between string and []byte copies the data.
func Example_copies() {
	b := []byte("hi")
	s := string(b) // copy
	b[0] = 'H'
	fmt.Println(s, string(b))
	// Output:
	// hi Hi
}

// Example_sharedMemory shows two places where a result shares memory with its
// source. Buffer.Bytes and bytes.Split do NOT copy, so writes show through.
func Example_sharedMemory() {
	var buf bytes.Buffer
	buf.WriteString("abc")
	view := buf.Bytes()
	view[0] = 'X'
	fmt.Println(buf.String())

	data := []byte("a,b")
	parts := bytes.Split(data, []byte(","))
	parts[0][0] = 'Z'
	fmt.Println(string(data))
	// Output:
	// Xbc
	// Z,b
}

// Example_reader shows bytes.Reader, a read-only view with Seek and ReadAt.
func Example_reader() {
	r := bytes.NewReader([]byte("0123456789"))
	r.Seek(5, io.SeekStart)
	b := make([]byte, 3)
	r.Read(b)
	fmt.Println(string(b), r.Len()) // 3 bytes read, 2 unread

	r.ReadAt(b, 0) // ReadAt ignores and does not move the read position
	fmt.Println(string(b))
	// Output:
	// 567 2
	// 012
}

// Example_unicode shows that bytes functions are UTF-8 aware where it matters.
func Example_unicode() {
	b := []byte("héy")
	fmt.Println(len(b), len(bytes.Runes(b)))
	fmt.Println(string(bytes.ToUpper([]byte("héllo"))))
	fmt.Println(bytes.IndexRune(b, 'y')) // byte offset
	// Output:
	// 4 3
	// HÉLLO
	// 3
}
