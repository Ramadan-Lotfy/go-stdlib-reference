package solution

import (
	"bytes"
	"testing"
)

func TestNewLogger(t *testing.T) {
	var buf bytes.Buffer
	l := NewLogger(&buf, "api")
	l.Info("hi", "token", Secret("abc123"), "n", 2)

	want := `{"level":"INFO","msg":"hi","service":"api","token":"[REDACTED]","n":2}` + "\n"
	if got := buf.String(); got != want {
		t.Errorf("log line mismatch\n got: %s\nwant: %s", got, want)
	}
}
