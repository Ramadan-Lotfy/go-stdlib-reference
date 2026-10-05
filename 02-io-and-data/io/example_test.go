package ioref_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// Example_copy streams a Reader into a Writer.
func Example_copy() {
	n, err := io.Copy(os.Stdout, strings.NewReader("hello\n"))
	fmt.Println(n, err)
	// Output:
	// hello
	// 6 <nil>
}

// Example_readFull shows ReadFull, LimitReader and the unexpected-EOF case.
func Example_readFull() {
	r := strings.NewReader("abcdefgh")
	buf := make([]byte, 3)
	n, err := io.ReadFull(r, buf)
	fmt.Println(n, err, string(buf))

	rest, _ := io.ReadAll(io.LimitReader(r, 2)) // at most 2 more bytes
	fmt.Println(string(rest))

	n, err = io.ReadFull(strings.NewReader("ab"), make([]byte, 3))
	fmt.Println(n, err)
	// Output:
	// 3 <nil> abc
	// de
	// 2 unexpected EOF
}

// Example_tee copies everything read from a stream into a second place.
func Example_tee() {
	var seen bytes.Buffer
	tr := io.TeeReader(strings.NewReader("payload"), &seen)
	data, _ := io.ReadAll(tr)
	fmt.Println(string(data), seen.String())
	// Output: payload payload
}

// Example_multi fans a write out to several Writers and joins several Readers.
func Example_multi() {
	var a, b bytes.Buffer
	mw := io.MultiWriter(&a, &b)
	io.WriteString(mw, "hi")
	fmt.Println(a.String(), b.String())

	mr := io.MultiReader(strings.NewReader("one "), strings.NewReader("two"))
	data, _ := io.ReadAll(mr)
	fmt.Println(string(data))
	// Output:
	// hi hi
	// one two
}

// Example_pipe connects a producer goroutine to a consumer without a buffer.
func Example_pipe() {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close() // without Close the reader would block forever
		fmt.Fprint(pw, "through the pipe")
	}()
	data, _ := io.ReadAll(pr)
	fmt.Println(string(data))
	// Output: through the pipe
}

// upperReader wraps another Reader and upper-cases what it returns.
type upperReader struct{ r io.Reader }

func (u upperReader) Read(p []byte) (int, error) {
	n, err := u.r.Read(p)
	copy(p[:n], bytes.ToUpper(p[:n]))
	return n, err
}

// Example_decorator shows the standard pattern: a Reader that wraps a Reader.
func Example_decorator() {
	io.Copy(os.Stdout, upperReader{strings.NewReader("hello")})
	fmt.Println()
	// Output: HELLO
}

// Example_sectionReader reads a window of a larger source.
func Example_sectionReader() {
	src := strings.NewReader("0123456789")
	sec := io.NewSectionReader(src, 2, 4) // offset 2, length 4
	data, _ := io.ReadAll(sec)
	fmt.Println(string(data), sec.Size())
	// Output: 2345 4
}

// Example_discard counts bytes without storing them.
func Example_discard() {
	n, _ := io.Copy(io.Discard, strings.NewReader("count me"))
	fmt.Println(n)
	// Output: 8
}
