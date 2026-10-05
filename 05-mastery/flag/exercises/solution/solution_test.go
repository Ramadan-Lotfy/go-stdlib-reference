package solution

import (
	"reflect"
	"testing"
	"time"
)

func TestParseArgsDefaults(t *testing.T) {
	got, err := ParseArgs(nil)
	want := Config{Addr: ":8080", Timeout: 5 * time.Second}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("ParseArgs(nil) = %+v, %v; want %+v", got, err, want)
	}
}

func TestParseArgsFull(t *testing.T) {
	got, err := ParseArgs([]string{"-addr=:9000", "-v", "-timeout", "2m", "-tag", "a", "-tag", "b", "x.txt", "y.txt"})
	want := Config{
		Addr: ":9000", Verbose: true, Timeout: 2 * time.Minute,
		Tags: []string{"a", "b"}, Files: []string{"x.txt", "y.txt"},
	}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("ParseArgs() = %+v, %v\nwant %+v", got, err, want)
	}
}

func TestParseArgsErrors(t *testing.T) {
	for _, args := range [][]string{
		{"-nope"},
		{"-timeout", "soon"},
		{"-tag", ""},
		{"-addr"}, // missing value
	} {
		if _, err := ParseArgs(args); err == nil {
			t.Errorf("ParseArgs(%v) should fail", args)
		}
	}
}
