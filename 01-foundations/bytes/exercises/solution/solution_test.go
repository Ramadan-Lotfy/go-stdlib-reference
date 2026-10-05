package solution

import (
	"bytes"
	"testing"
)

func TestNormalizeNewlines(t *testing.T) {
	in := []byte("a\r\nb\rc\n")
	orig := append([]byte(nil), in...)

	got := NormalizeNewlines(in)
	if want := []byte("a\nb\nc\n"); !bytes.Equal(got, want) {
		t.Errorf("NormalizeNewlines() = %q, want %q", got, want)
	}
	if !bytes.Equal(in, orig) {
		t.Errorf("input was modified: %q", in)
	}

	// The result must not share memory with the input.
	got[0] = 'X'
	if !bytes.Equal(in, orig) {
		t.Errorf("result shares memory with input: %q", in)
	}
}

func TestNumberLines(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"a\nb\n", "1: a\n2: b\n"},
		{"a\nb", "1: a\n2: b"},
		{"\n\n", "1: \n2: \n"},
		{"", ""},
	}
	for _, tt := range tests {
		got := NumberLines([]byte(tt.in))
		if !bytes.Equal(got, []byte(tt.want)) {
			t.Errorf("NumberLines(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
