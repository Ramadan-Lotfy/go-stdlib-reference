package atomicref_test

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Example_counter increments one shared counter from many goroutines.
func Example_counter() {
	var n atomic.Int64
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(n.Load())
	// Output: 1000
}

// Example_compareAndSwap updates a value only if it still has the expected one.
func Example_compareAndSwap() {
	var v atomic.Int32
	v.Store(5)
	fmt.Println(v.CompareAndSwap(5, 10)) // 5 -> 10
	fmt.Println(v.CompareAndSwap(5, 20)) // value is 10 now, so no change
	fmt.Println(v.Load())
	fmt.Println(v.Swap(99)) // returns the old value
	// Output:
	// true
	// false
	// 10
	// 10
}

// Example_flag uses a boolean as a stop signal.
func Example_flag() {
	var stop atomic.Bool
	done := make(chan struct{})
	go func() {
		for !stop.Load() {
			// do work
		}
		close(done)
	}()
	stop.Store(true)
	<-done
	fmt.Println("stopped")
	// Output: stopped
}

type Config struct {
	Debug bool
	Limit int
}

// Example_pointer swaps an immutable configuration without locking readers.
func Example_pointer() {
	var cfg atomic.Pointer[Config]
	cfg.Store(&Config{Debug: false, Limit: 10})

	old := cfg.Swap(&Config{Debug: true, Limit: 20})
	fmt.Println(old.Debug, old.Limit)
	fmt.Println(cfg.Load().Debug, cfg.Load().Limit)
	// Output:
	// false 10
	// true 20
}

// Example_updateLoop is the compare-and-swap loop for a read-modify-write.
func Example_updateLoop() {
	var cfg atomic.Pointer[Config]
	cfg.Store(&Config{Limit: 1})

	var wg sync.WaitGroup
	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				old := cfg.Load()
				updated := *old // copy, never mutate the old value
				updated.Limit++
				if cfg.CompareAndSwap(old, &updated) {
					return
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println(cfg.Load().Limit)
	// Output: 51
}

// Example_value stores any single type with atomic.Value.
func Example_value() {
	var v atomic.Value
	v.Store("first")
	fmt.Println(v.Load().(string))
	v.Store("second")
	fmt.Println(v.Load().(string))
	// Output:
	// first
	// second
}
