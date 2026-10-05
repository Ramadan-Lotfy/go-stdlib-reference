package solution

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestClassify(t *testing.T) {
	tests := []struct {
		err  error
		want string
	}{
		{nil, "ok"},
		{fmt.Errorf("ctx: %w", ErrNotFound), "not found"},
		{fmt.Errorf("ctx: %w", &ValidationError{"age"}), "invalid: age"},
		{errors.New("boom"), "unknown"},
	}
	for _, tt := range tests {
		if got := Classify(tt.err); got != tt.want {
			t.Errorf("Classify(%v) = %q, want %q", tt.err, got, tt.want)
		}
	}
}

func TestRetry(t *testing.T) {
	boom := errors.New("boom")

	calls := 0
	err := Retry(3, func() error {
		calls++
		if calls < 3 {
			return boom
		}
		return nil
	})
	if err != nil || calls != 3 {
		t.Errorf("Retry succeeded case: err=%v calls=%d", err, calls)
	}

	err = Retry(2, func() error { return boom })
	if !errors.Is(err, boom) {
		t.Errorf("err = %v, want it to wrap boom", err)
	}
	if !strings.Contains(err.Error(), "attempt 1") || !strings.Contains(err.Error(), "attempt 2") {
		t.Errorf("err = %q, want both attempts listed", err)
	}

	if Retry(0, func() error { return nil }) == nil {
		t.Error("Retry(0) should return an error")
	}
}
