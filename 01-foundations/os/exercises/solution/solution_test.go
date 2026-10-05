package solution

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestWordCountFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "words.txt")
	if err := os.WriteFile(path, []byte("one two\nthree\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	n, err := WordCountFile(path)
	if err != nil || n != 3 {
		t.Errorf("WordCountFile() = %d, %v; want 3, nil", n, err)
	}

	_, err = WordCountFile(filepath.Join(t.TempDir(), "missing.txt"))
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("err = %v, want it to wrap fs.ErrNotExist", err)
	}
}

func TestAppendLine(t *testing.T) {
	path := filepath.Join(t.TempDir(), "log.txt")
	for _, line := range []string{"a", "b"} {
		if err := AppendLine(path, line); err != nil {
			t.Fatal(err)
		}
	}
	data, _ := os.ReadFile(path)
	if string(data) != "a\nb\n" {
		t.Errorf("file = %q, want %q", data, "a\nb\n")
	}
}
