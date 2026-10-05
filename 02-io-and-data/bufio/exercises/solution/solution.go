// Package solution holds the reference solutions for the bufio exercises.
package solution

import (
	"bufio"
	"io"
)

// LongestLine returns the first longest line in r. Lines up to 1 MiB are
// supported.
func LongestLine(r io.Reader) (string, error) {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), 1<<20)
	longest := ""
	for sc.Scan() {
		if line := sc.Text(); len(line) > len(longest) {
			longest = line
		}
	}
	return longest, sc.Err()
}
