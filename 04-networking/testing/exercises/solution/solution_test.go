package solution

import (
	"errors"
	"slices"
	"testing"
)

func TestClamp(t *testing.T) {
	tests := []struct {
		name            string
		x, lo, hi, want int
	}{
		{"inside", 5, 0, 10, 5},
		{"below", -3, 0, 10, 0},
		{"above", 30, 0, 10, 10},
		{"at lower edge", 0, 0, 10, 0},
		{"single point range", 7, 4, 4, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.x, tt.lo, tt.hi); got != tt.want {
				t.Errorf("Clamp(%d, %d, %d) = %d, want %d", tt.x, tt.lo, tt.hi, got, tt.want)
			}
		})
	}
}

func TestClampPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Clamp(1, 5, 0) should panic")
		}
	}()
	Clamp(1, 5, 0)
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name string
		in   []float64
		want float64
	}{
		{"odd count", []float64{3, 1, 2}, 2},
		{"even count", []float64{4, 1, 3, 2}, 2.5},
		{"single", []float64{9}, 9},
		{"negatives", []float64{-5, -1, -3}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Median(tt.in)
			if err != nil || got != tt.want {
				t.Errorf("Median(%v) = %v, %v; want %v", tt.in, got, err, tt.want)
			}
		})
	}
}

func TestMedianEmptyAndInputUnchanged(t *testing.T) {
	if _, err := Median(nil); !errors.Is(err, ErrEmpty) {
		t.Errorf("Median(nil) error = %v, want ErrEmpty", err)
	}
	in := []float64{3, 1, 2}
	Median(in)
	if !slices.Equal(in, []float64{3, 1, 2}) {
		t.Errorf("input was modified: %v", in)
	}
}
