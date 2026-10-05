package solution

import (
	"slices"
	"testing"
	"testing/fstest"
)

func TestListGoFiles(t *testing.T) {
	fsys := fstest.MapFS{
		"main.go":             {},
		"README.md":           {},
		"pkg/a.go":            {},
		"pkg/a_test.go":       {},
		"vendor/lib/lib.go":   {},
		".git/hooks/hook.go":  {},
		"internal/deep/b.go":  {},
		"internal/deep/c.txt": {},
	}
	got, err := ListGoFiles(fsys)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"internal/deep/b.go", "main.go", "pkg/a.go", "pkg/a_test.go"}
	if !slices.Equal(got, want) {
		t.Errorf("ListGoFiles() = %v, want %v", got, want)
	}
}
