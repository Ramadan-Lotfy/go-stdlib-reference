package solution

import (
	"slices"
	"testing"
)

func TestTopN(t *testing.T) {
	words := []string{"b", "a", "c", "b", "a", "b", "d"}
	tests := []struct {
		n    int
		want []string
	}{
		{2, []string{"b", "a"}},
		{3, []string{"b", "a", "c"}},
		{10, []string{"b", "a", "c", "d"}},
		{0, []string{}},
		{-1, []string{}},
	}
	for _, tt := range tests {
		if got := TopN(words, tt.n); !slices.Equal(got, tt.want) {
			t.Errorf("TopN(n=%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
