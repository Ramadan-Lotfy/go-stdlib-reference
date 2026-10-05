package solution

import (
	"reflect"
	"slices"
	"testing"
)

func TestGroupBy(t *testing.T) {
	got := GroupBy([]string{"a", "bb", "cc", "d"}, func(s string) int { return len(s) })
	want := map[int][]string{1: {"a", "d"}, 2: {"bb", "cc"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBy() = %v, want %v", got, want)
	}
}

func TestSortedKeys(t *testing.T) {
	got := SortedKeys(map[string]bool{"b": true, "c": false, "a": true})
	if want := []string{"a", "b", "c"}; !slices.Equal(got, want) {
		t.Errorf("SortedKeys() = %v, want %v", got, want)
	}
	if got := SortedKeys(map[int]string{}); len(got) != 0 {
		t.Errorf("SortedKeys(empty) = %v, want empty", got)
	}
}
