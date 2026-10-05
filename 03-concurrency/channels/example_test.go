package channelsref_test

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

// Example_unbuffered passes one value between goroutines.
func Example_unbuffered() {
	ch := make(chan string)
	go func() { ch <- "ping" }()
	fmt.Println(<-ch) // blocks until the goroutine sends
	// Output: ping
}

// Example_buffered shows capacity, closing and ranging.
func Example_buffered() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(len(ch), cap(ch))
	close(ch)

	for v := range ch { // ends once the channel is closed AND drained
		fmt.Println(v)
	}
	v, ok := <-ch
	fmt.Println(v, ok)
	// Output:
	// 2 2
	// 1
	// 2
	// 0 false
}

// Example_closedSend shows the panic from sending on a closed channel.
func Example_closedSend() {
	defer func() { fmt.Println("recovered:", recover()) }()
	ch := make(chan int)
	close(ch)
	ch <- 1
	// Output: recovered: send on closed channel
}

// Example_select tries a receive without blocking and then with a timeout.
func Example_select() {
	ch := make(chan int)

	select {
	case v := <-ch:
		fmt.Println(v)
	default:
		fmt.Println("no value ready")
	}

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(10 * time.Millisecond):
		fmt.Println("timeout")
	}
	// Output:
	// no value ready
	// timeout
}

// Example_workerPool runs jobs on a fixed number of workers.
func Example_workerPool() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup
	for range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				results <- j * j
			}
		}()
	}
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs) // lets the workers' range loops finish
	}()
	go func() {
		wg.Wait()
		close(results) // lets the collector's range loop finish
	}()

	var got []int
	for r := range results {
		got = append(got, r)
	}
	slices.Sort(got) // completion order is not deterministic
	fmt.Println(got)
	// Output: [1 4 9 16 25]
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// Example_pipeline chains stages: generate, square, sum.
func Example_pipeline() {
	sum := 0
	for v := range square(gen(1, 2, 3, 4)) {
		sum += v
	}
	fmt.Println(sum)
	// Output: 30
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, c := range cs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Example_fanIn merges several channels into one.
func Example_fanIn() {
	var got []int
	for v := range merge(gen(1, 2), gen(10, 20), gen(100)) {
		got = append(got, v)
	}
	slices.Sort(got)
	fmt.Println(got)
	// Output: [1 2 10 20 100]
}

func count(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done():
				return // the goroutine always has a way out
			}
		}
	}()
	return out
}

// Example_cancelGenerator stops an infinite producer without leaking it.
func Example_cancelGenerator() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := count(ctx)
	fmt.Println(<-ch, <-ch, <-ch)

	cancel()
	for range ch { // returns once the generator has exited and closed ch
	}
	fmt.Println("generator stopped")
	// Output:
	// 0 1 2
	// generator stopped
}

// Example_semaphore limits how many goroutines run at the same time.
func Example_semaphore() {
	sem := make(chan struct{}, 2) // at most 2 at once
	var current, peak atomic.Int32
	var wg sync.WaitGroup

	for range 6 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}        // acquire
			defer func() { <-sem }() // release

			n := current.Add(1)
			for {
				p := peak.Load()
				if n <= p || peak.CompareAndSwap(p, n) {
					break
				}
			}
			time.Sleep(time.Millisecond)
			current.Add(-1)
		}()
	}
	wg.Wait()
	fmt.Println("peak <= 2:", peak.Load() <= 2)
	// Output: peak <= 2: true
}

// Example_nilChannel disables a select case by setting its channel to nil.
func Example_nilChannel() {
	a, b := gen(1, 2), gen(10)
	var got []int
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil // a nil channel is never ready
				continue
			}
			got = append(got, v)
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			got = append(got, v)
		}
	}
	slices.Sort(got)
	fmt.Println(got)
	// Output: [1 2 10]
}
