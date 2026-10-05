package solution

import (
	"bytes"
	"context"
	"runtime/trace"
	"testing"
)

func TestTraceTo(t *testing.T) {
	var buf bytes.Buffer
	ran := false
	err := TraceTo(&buf, func(ctx context.Context) {
		trace.WithRegion(ctx, "step", func() { ran = true })
	})
	if err != nil {
		t.Skipf("tracing unavailable (already running?): %v", err)
	}
	if !ran {
		t.Error("f was not called")
	}
	if trace.IsEnabled() {
		t.Error("tracing should be stopped after TraceTo returns")
	}
	if !bytes.HasPrefix(buf.Bytes(), []byte("go 1.")) {
		t.Errorf("trace should start with a \"go 1.x trace\" header, got %q", buf.Bytes()[:min(16, buf.Len())])
	}
}
