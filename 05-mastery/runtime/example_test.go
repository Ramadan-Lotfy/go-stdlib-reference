package runtimeref_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// settle waits until the goroutine count drops to want (counts lag slightly).
func settle(want int) bool {
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if runtime.NumGoroutine() <= want {
			return true
		}
		time.Sleep(time.Millisecond)
	}
	return false
}

// Example_goroutines counts goroutines, the basis of a leak check.
func Example_goroutines() {
	base := runtime.NumGoroutine()
	release := make(chan struct{})
	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-release
		}()
	}
	fmt.Println(runtime.NumGoroutine() - base)

	close(release)
	wg.Wait()
	fmt.Println(settle(base))
	// Output:
	// 5
	// true
}

var sink []byte

// Example_memStats measures allocation and a garbage collection.
func Example_memStats() {
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)

	sink = make([]byte, 10<<20) // 10 MiB
	for i := range sink {
		sink[i] = 1
	}

	runtime.ReadMemStats(&after)
	fmt.Println(after.TotalAlloc-before.TotalAlloc >= 10<<20)

	sink = nil
	gcs := after.NumGC
	runtime.GC()
	runtime.ReadMemStats(&after)
	fmt.Println(after.NumGC > gcs)
	// Output:
	// true
	// true
}

// Example_caller finds out where the code is running.
func Example_caller() {
	pc, file, line, ok := runtime.Caller(0) // 0 = this function
	fmt.Println(filepath.Base(file), line > 0, ok)
	fmt.Println(strings.HasSuffix(runtime.FuncForPC(pc).Name(), "Example_caller"))
	// Output:
	// example_test.go true true
	// true
}

// Example_stack captures the current goroutine's stack trace.
func Example_stack() {
	buf := make([]byte, 2048)
	n := runtime.Stack(buf, false) // false: this goroutine only
	trace := string(buf[:n])
	fmt.Println(strings.HasPrefix(trace, "goroutine "), strings.Contains(trace, "Example_stack"))
	// Output: true true
}

// Example_goexit ends a goroutine but still runs its deferred calls.
func Example_goexit() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer fmt.Println("deferred call ran")
		runtime.Goexit()
		fmt.Println("never printed")
	}()
	<-done
	// Output: deferred call ran
}

// Example_platform reads facts about the machine and toolchain.
func Example_platform() {
	fmt.Println(strings.HasPrefix(runtime.Version(), "go"))
	fmt.Println(runtime.GOOS != "", runtime.GOARCH != "")
	fmt.Println(runtime.NumCPU() >= 1, runtime.GOMAXPROCS(0) >= 1)
	// Output:
	// true
	// true true
	// true true
}
