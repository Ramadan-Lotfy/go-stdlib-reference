package solution

import (
	"io"
	"strings"
	"testing"
)

func TestCountingReader(t *testing.T) {
	cr := &CountingReader{R: strings.NewReader("hello world")}
	if _, err := io.Copy(io.Discard, cr); err != nil {
		t.Fatal(err)
	}
	if cr.N != 11 {
		t.Errorf("N = %d, want 11", cr.N)
	}
}

func TestHead(t *testing.T) {
	got, err := Head(strings.NewReader("hello"), 3)
	if err != nil || got != "hel" {
		t.Errorf("Head() = %q, %v; want hel", got, err)
	}
	got, _ = Head(strings.NewReader("hi"), 10)
	if got != "hi" {
		t.Errorf("Head(short) = %q, want hi", got)
	}
}
