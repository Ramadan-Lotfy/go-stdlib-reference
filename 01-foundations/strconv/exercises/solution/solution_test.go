package solution

import (
	"errors"
	"maps"
	"strconv"
	"testing"
)

func TestParseKV(t *testing.T) {
	got, err := ParseKV("a=1,b=2, c = 3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := map[string]int{"a": 1, "b": 2, "c": 3}
	if !maps.Equal(got, want) {
		t.Errorf("ParseKV() = %v, want %v", got, want)
	}

	got, err = ParseKV("")
	if err != nil || len(got) != 0 {
		t.Errorf("ParseKV(\"\") = %v, %v; want empty map, nil", got, err)
	}
}

func TestParseKVErrors(t *testing.T) {
	if _, err := ParseKV("a=1,b"); err == nil {
		t.Error("missing '=' should return an error")
	}

	_, err := ParseKV("a=x")
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("err = %v, want it to wrap strconv.ErrSyntax", err)
	}

	_, err = ParseKV("a=99999999999999999999")
	if !errors.Is(err, strconv.ErrRange) {
		t.Errorf("err = %v, want it to wrap strconv.ErrRange", err)
	}
}

func TestHumanBytes(t *testing.T) {
	tests := []struct {
		in   int64
		want string
	}{
		{0, "0 B"},
		{999, "999 B"},
		{1023, "1023 B"},
		{1024, "1 KB"},
		{1536, "1.5 KB"},
		{1048575, "1 MB"}, // rounds up to 1024.0 KB, so promote
		{1048576, "1 MB"},
		{1500000, "1.4 MB"},
		{5 * 1024 * 1024 * 1024, "5 GB"},
	}
	for _, tt := range tests {
		if got := HumanBytes(tt.in); got != tt.want {
			t.Errorf("HumanBytes(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
