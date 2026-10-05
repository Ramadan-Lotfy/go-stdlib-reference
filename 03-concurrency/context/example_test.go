package contextref_test

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Example_withCancel stops a goroutine on demand.
func Example_withCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	stopped := make(chan struct{})

	go func() {
		<-ctx.Done() // blocks until cancel is called
		close(stopped)
	}()

	cancel()
	<-stopped
	fmt.Println(ctx.Err())
	// Output: context canceled
}

// Example_withTimeout gives up on slow work after a deadline.
func Example_withTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(time.Second): // slow work
		fmt.Println("finished")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println(errors.Is(ctx.Err(), context.DeadlineExceeded))
	}
	// Output:
	// context deadline exceeded
	// true
}

// work is the standard shape of a cancellable function.
func work(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Example_cancellableFunc shows a function that returns as soon as the context ends.
func Example_cancellableFunc() {
	fmt.Println(work(context.Background(), time.Millisecond))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	fmt.Println(work(ctx, time.Hour))
	// Output:
	// <nil>
	// context canceled
}

type ctxKey string // unexported key type avoids collisions

// Example_withValue passes request-scoped data down the call chain.
func Example_withValue() {
	ctx := context.WithValue(context.Background(), ctxKey("requestID"), "req-42")

	id, ok := ctx.Value(ctxKey("requestID")).(string)
	fmt.Println(id, ok)

	_, ok = ctx.Value(ctxKey("missing")).(string)
	fmt.Println(ok)
	// Output:
	// req-42 true
	// false
}

// Example_propagation shows that cancelling a parent cancels its children.
func Example_propagation() {
	parent, cancelParent := context.WithCancel(context.Background())
	child, cancelChild := context.WithTimeout(parent, time.Hour)
	defer cancelChild()

	fmt.Println(child.Err())
	cancelParent()
	<-child.Done()
	fmt.Println(child.Err())
	// Output:
	// <nil>
	// context canceled
}

// Example_cause records why a context was cancelled.
func Example_cause() {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(errors.New("upstream failed"))

	fmt.Println(ctx.Err())
	fmt.Println(context.Cause(ctx))
	// Output:
	// context canceled
	// upstream failed
}

// Example_afterFunc runs cleanup when a context ends.
func Example_afterFunc() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	context.AfterFunc(ctx, func() {
		fmt.Println("cleanup")
		close(done)
	})
	cancel()
	<-done
	// Output: cleanup
}

// Example_withoutCancel detaches work from the request's cancellation.
func Example_withoutCancel() {
	parent, cancel := context.WithCancel(context.Background())
	detached := context.WithoutCancel(parent)
	cancel()
	fmt.Println(parent.Err(), detached.Err())
	// Output: context canceled <nil>
}

// Example_deadline reads the deadline so callees can adapt.
func Example_deadline() {
	_, has := context.Background().Deadline()
	fmt.Println(has)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	d, has := ctx.Deadline()
	fmt.Println(has, time.Until(d) > 59*time.Second)
	// Output:
	// false
	// true true
}
