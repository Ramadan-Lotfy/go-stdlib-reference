package solution

import (
	"strings"
	"testing"
)

func TestLongestLine(t *testing.T) {
	got, err := LongestLine(strings.NewReader("a\nlongest line\nmid\nlongest line 2"))
	if err != nil || got != "longest line 2" {
		t.Errorf("LongestLine() = %q, %v", got, err)
	}

	got, _ = LongestLine(strings.NewReader("same\nsize"))
	if got != "same" {
		t.Errorf("tie should keep the first line, got %q", got)
	}

	big := strings.Repeat("x", 100_000)
	got, err = LongestLine(strings.NewReader("short\n" + big + "\n"))
	if err != nil || len(got) != len(big) {
		t.Errorf("big line: len=%d err=%v", len(got), err)
	}

	got, err = LongestLine(strings.NewReader(""))
	if err != nil || got != "" {
		t.Errorf("empty input: %q, %v", got, err)
	}
}
