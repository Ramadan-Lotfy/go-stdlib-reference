package bufioref_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// Example_scannerLines reads text line by line. Empty lines are returned too.
func Example_scannerLines() {
	sc := bufio.NewScanner(strings.NewReader("one\ntwo\n\nthree"))
	for sc.Scan() {
		fmt.Printf("%q\n", sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Println(err)
	}
	// Output:
	// "one"
	// "two"
	// ""
	// "three"
}

// Example_scannerWords counts words instead of lines.
func Example_scannerWords() {
	sc := bufio.NewScanner(strings.NewReader("the quick  brown\nfox"))
	sc.Split(bufio.ScanWords)
	n := 0
	for sc.Scan() {
		n++
	}
	fmt.Println(n)
	// Output: 4
}

// splitComma is a custom SplitFunc that tokenizes on commas.
func splitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[:i], nil
	}
	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}
	return 0, nil, nil // ask for more data
}

// Example_customSplit plugs in a custom tokenizer.
func Example_customSplit() {
	sc := bufio.NewScanner(strings.NewReader("a,b,c"))
	sc.Split(splitComma)
	var parts []string
	for sc.Scan() {
		parts = append(parts, sc.Text())
	}
	fmt.Printf("%q\n", parts)
	// Output: ["a" "b" "c"]
}

// Example_tooLong shows the 64 KB default limit and how to raise it.
func Example_tooLong() {
	long := strings.Repeat("x", 70_000)

	sc := bufio.NewScanner(strings.NewReader(long))
	fmt.Println(sc.Scan(), sc.Err())

	sc = bufio.NewScanner(strings.NewReader(long))
	sc.Buffer(make([]byte, 0, 4096), 1<<20) // allow tokens up to 1 MiB
	fmt.Println(sc.Scan(), len(sc.Text()), sc.Err())
	// Output:
	// false bufio.Scanner: token too long
	// true 70000 <nil>
}

// Example_reader shows Peek, ReadString, ReadRune and the final-line EOF.
func Example_reader() {
	br := bufio.NewReader(strings.NewReader("hello world\nsecond"))

	peek, _ := br.Peek(5) // look ahead without consuming
	fmt.Printf("%s\n", peek)

	line, err := br.ReadString('\n')
	fmt.Printf("%q %v\n", line, err)

	r, size, _ := br.ReadRune()
	fmt.Println(string(r), size)

	rest, err := br.ReadString('\n') // no newline left: data AND io.EOF
	fmt.Printf("%q %v\n", rest, err)
	// Output:
	// hello
	// "hello world\n" <nil>
	// s 1
	// "econd" EOF
}

// Example_writer shows that buffered output is invisible until Flush.
func Example_writer() {
	var out bytes.Buffer
	w := bufio.NewWriter(&out)
	w.WriteString("hi")
	fmt.Println(out.Len()) // still buffered
	w.Flush()
	fmt.Println(out.Len())
	// Output:
	// 0
	// 2
}
