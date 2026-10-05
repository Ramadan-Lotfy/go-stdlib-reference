package solution

import (
	"strings"
	"sync"
	"testing"
)

func TestLeakCheck(t *testing.T) {
	release := make(chan struct{})
	got := LeakCheck(func() {
		go func() { <-release }() // never finishes until we say so
	})
	close(release)
	if got != 1 {
		t.Errorf("LeakCheck(leaky) = %d, want 1", got)
	}

	got = LeakCheck(func() {
		var wg sync.WaitGroup
		for range 3 {
			wg.Add(1)
			go func() { defer wg.Done() }()
		}
		wg.Wait()
	})
	if got != 0 {
		t.Errorf("LeakCheck(clean) = %d, want 0", got)
	}
}

func TestCallerInfo(t *testing.T) {
	got := CallerInfo()
	if !strings.HasPrefix(got, "solution_test.go:") {
		t.Errorf("CallerInfo() = %q, want it to start with solution_test.go:", got)
	}
}
