package syncref_test

import (
	"bytes"
	"fmt"
	"slices"
	"sync"
)

type counter struct {
	mu sync.Mutex
	n  int
}

func (c *counter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

// Example_mutex protects shared state from many goroutines.
func Example_mutex() {
	var c counter
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1) // before the goroutine starts
		go func() {
			defer wg.Done()
			c.inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.n)
	// Output: 100
}

// Example_waitGroup collects results by index, so no lock is needed.
func Example_waitGroup() {
	results := make([]int, 5)
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = i * i // each goroutine writes its own element
		}()
	}
	wg.Wait()
	fmt.Println(results)
	// Output: [0 1 4 9 16]
}

type cache struct {
	mu sync.RWMutex
	m  map[string]string
}

func (c *cache) get(k string) (string, bool) {
	c.mu.RLock() // many readers can hold this at once
	defer c.mu.RUnlock()
	v, ok := c.m[k]
	return v, ok
}

func (c *cache) set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[k] = v
}

// Example_rwMutex uses a read-write lock for a read-heavy cache.
func Example_rwMutex() {
	c := &cache{m: map[string]string{}}
	c.set("lang", "go")
	v, ok := c.get("lang")
	_, missing := c.get("nope")
	fmt.Println(v, ok, missing)
	// Output: go true false
}

// Example_once runs initialization exactly once, however many goroutines ask.
func Example_once() {
	var once sync.Once
	calls := 0
	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() { calls++ })
		}()
	}
	wg.Wait()
	fmt.Println("init ran", calls, "time")

	// OnceValue memoizes a computed result.
	n := 0
	get := sync.OnceValue(func() int { n++; return 42 })
	fmt.Println(get(), get(), n)
	// Output:
	// init ran 1 time
	// 42 42 1
}

// Example_cond wakes a waiting goroutine when a condition becomes true.
func Example_cond() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	done := make(chan struct{})

	go func() {
		mu.Lock()
		for !ready { // always recheck in a loop
			cond.Wait()
		}
		mu.Unlock()
		fmt.Println("consumer proceeds")
		close(done)
	}()

	mu.Lock()
	ready = true
	mu.Unlock()
	cond.Broadcast()
	<-done
	// Output: consumer proceeds
}

// Example_map shows the sync.Map API (keys and values are `any`).
func Example_map() {
	var m sync.Map
	m.Store("a", 1)
	v, ok := m.Load("a")
	fmt.Println(v, ok)

	actual, loaded := m.LoadOrStore("b", 2)
	fmt.Println(actual, loaded)
	actual, loaded = m.LoadOrStore("b", 3) // already present
	fmt.Println(actual, loaded)

	m.Delete("a")
	var keys []string
	m.Range(func(k, v any) bool {
		keys = append(keys, k.(string))
		return true
	})
	slices.Sort(keys) // Range order is unspecified
	fmt.Println(keys)
	// Output:
	// 1 true
	// 2 false
	// 2 true
	// [b]
}

var bufPool = sync.Pool{New: func() any { return new(bytes.Buffer) }}

// Example_pool reuses buffers. Always Reset before use and Put when done.
func Example_pool() {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	buf.WriteString("hello")
	fmt.Println(buf.String())
	bufPool.Put(buf)
	// Output: hello
}
