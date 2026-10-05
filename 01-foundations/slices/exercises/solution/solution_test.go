package solution

import (
	"reflect"
	"slices"
	"testing"
)

func TestUnique(t *testing.T) {
	in := []int{3, 1, 3, 2, 1}
	got := Unique(in)
	if want := []int{1, 2, 3}; !slices.Equal(got, want) {
		t.Errorf("Unique() = %v, want %v", got, want)
	}
	if want := []int{3, 1, 3, 2, 1}; !slices.Equal(in, want) {
		t.Errorf("input was modified: %v", in)
	}
}

func TestChunk(t *testing.T) {
	got := Chunk([]int{1, 2, 3, 4, 5}, 2)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Chunk() = %v, want %v", got, want)
	}
	if Chunk([]int{1}, 0) != nil {
		t.Error("Chunk with n=0 should return nil")
	}
	if got := Chunk([]string{}, 3); len(got) != 0 {
		t.Errorf("Chunk(empty) = %v, want empty", got)
	}
}
