// Package solution holds the reference solutions for the io exercises.
package solution

import "io"

// CountingReader wraps a Reader and counts the bytes read through it.
type CountingReader struct {
	R io.Reader
	N int64
}

// Read implements io.Reader.
func (c *CountingReader) Read(p []byte) (int, error) {
	n, err := c.R.Read(p)
	c.N += int64(n)
	return n, err
}

// Head returns at most the first n bytes of r.
func Head(r io.Reader, n int64) (string, error) {
	b, err := io.ReadAll(io.LimitReader(r, n))
	return string(b), err
}
