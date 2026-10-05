package solution

import (
	"testing"
	"time"
)

func isGzip(b []byte) bool { return len(b) > 2 && b[0] == 0x1f && b[1] == 0x8b }

func TestProfileCPU(t *testing.T) {
	data, err := ProfileCPU(func() {
		deadline := time.Now().Add(30 * time.Millisecond)
		for time.Now().Before(deadline) {
		}
	})
	if err != nil {
		t.Skipf("CPU profiling unavailable (already running?): %v", err)
	}
	if !isGzip(data) {
		t.Errorf("CPU profile should be gzip data, got %d bytes", len(data))
	}
}

func TestHeapProfile(t *testing.T) {
	data, err := HeapProfile()
	if err != nil || !isGzip(data) {
		t.Errorf("HeapProfile() = %d bytes, %v; want gzip data", len(data), err)
	}
}
