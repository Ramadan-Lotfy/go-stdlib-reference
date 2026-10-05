package testingref_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"unicode/utf8"

	testingref "github.com/Ramadan-Lotfy/go-stdlib-reference/04-networking/testing"
)

// This file is a tour of testing features. Run it with:
//
//	go test -v ./04-networking/testing/
//	go test -bench=. -benchmem ./04-networking/testing/
//	go test -fuzz=FuzzReverse -fuzztime=5s ./04-networking/testing/

// TestAdd is the standard table-driven test with named subtests.
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positives", 1, 2, 3},
		{"negatives", -1, -2, -3},
		{"with zero", 0, 5, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testingref.Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// assertNoError is a helper. t.Helper makes failures point at the caller.
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// TestDivide checks both the success path and a sentinel error.
func TestDivide(t *testing.T) {
	got, err := testingref.Divide(6, 3)
	assertNoError(t, err)
	if got != 2 {
		t.Errorf("Divide(6, 3) = %v, want 2", got)
	}

	_, err = testingref.Divide(1, 0)
	if !errors.Is(err, testingref.ErrDivideByZero) {
		t.Errorf("Divide(1, 0) error = %v, want ErrDivideByZero", err)
	}
}

// TestParallel runs subtests concurrently. The loop variable is per-iteration
// in Go 1.22+, so no `s := s` copy is needed.
func TestParallel(t *testing.T) {
	for _, s := range []string{"a", "ab", "abc", "héllo"} {
		t.Run(s, func(t *testing.T) {
			t.Parallel()
			if got := testingref.Reverse(testingref.Reverse(s)); got != s {
				t.Errorf("double reverse of %q = %q", s, got)
			}
		})
	}
}

// TestFilesAndEnv uses the automatic helpers that replace manual cleanup.
func TestFilesAndEnv(t *testing.T) {
	dir := t.TempDir() // deleted when the test ends
	path := filepath.Join(dir, "data.txt")
	if err := os.WriteFile(path, []byte("hi"), 0o644); err != nil {
		t.Fatal(err)
	}

	t.Setenv("APP_ENV", "test") // restored automatically
	if os.Getenv("APP_ENV") != "test" {
		t.Error("env var not set")
	}

	t.Cleanup(func() { t.Log("cleanup runs after the test, in LIFO order") })
}

// TestSlow is skipped with `go test -short`.
func TestSlow(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping slow test in -short mode")
	}
	// ... long-running work would go here
}

// BenchmarkReverse measures speed. The loop runs b.N times; the framework
// chooses N so the benchmark lasts long enough to be reliable.
var sink string

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = testingref.Reverse("hello, world")
	}
}

// FuzzReverse checks a property for generated inputs. Under plain `go test`
// it only runs the seed corpus, so it is fast and deterministic.
func FuzzReverse(f *testing.F) {
	f.Add("hello")
	f.Add("héllo, 世界")
	f.Add("")
	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip() // Reverse works on valid UTF-8 text
		}
		if got := testingref.Reverse(testingref.Reverse(s)); got != s {
			t.Errorf("double reverse of %q = %q", s, got)
		}
	})
}

// ExampleReverse is documentation that go test verifies.
func ExampleReverse() {
	fmt.Println(testingref.Reverse("héllo"))
	// Output: olléh
}

// ExampleIsPalindrome shows the behavior for sentences.
func ExampleIsPalindrome() {
	fmt.Println(testingref.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(testingref.IsPalindrome("Gopher"))
	// Output:
	// true
	// false
}
