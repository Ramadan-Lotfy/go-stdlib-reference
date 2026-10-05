// Package solution holds the reference solutions for the bytes exercises.
package solution

import (
	"bytes"
	"fmt"
)

// NormalizeNewlines converts "\r\n" and lone "\r" to "\n". It returns a new
// slice and never modifies data.
func NormalizeNewlines(data []byte) []byte {
	out := bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))
	return bytes.ReplaceAll(out, []byte("\r"), []byte("\n"))
}

// NumberLines prefixes each line with "N: ". A trailing newline in the input
// is preserved; an input without one produces output without one.
func NumberLines(src []byte) []byte {
	var out bytes.Buffer
	n := 0
	for len(src) > 0 {
		line, rest, found := bytes.Cut(src, []byte("\n"))
		n++
		fmt.Fprintf(&out, "%d: %s", n, line)
		if found {
			out.WriteByte('\n')
		}
		src = rest
	}
	return out.Bytes()
}
