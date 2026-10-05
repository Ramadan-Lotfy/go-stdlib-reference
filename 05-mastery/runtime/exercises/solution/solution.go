// Package solution holds the reference solutions for the runtime and runtime/debug exercises.
package solution

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

// LeakCheck runs f and returns how many goroutines it left running. It waits
// up to 200 ms for goroutines that are shutting down to finish.
func LeakCheck(f func()) (leaked int) {
	before := runtime.NumGoroutine()
	f()
	deadline := time.Now().Add(200 * time.Millisecond)
	for runtime.NumGoroutine() > before && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	return max(0, runtime.NumGoroutine()-before)
}

// CallerInfo returns "file.go:line" of the code that called CallerInfo.
func CallerInfo() string {
	_, file, line, ok := runtime.Caller(1) // 1 = the caller of CallerInfo
	if !ok {
		return "unknown"
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}
