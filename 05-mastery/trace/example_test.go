package traceref_test

import (
	"bytes"
	"context"
	"fmt"
	"runtime/trace"
	"time"
)

// Example_trace records a trace with a task, a region and a log message.
func Example_trace() {
	var buf bytes.Buffer
	if err := trace.Start(&buf); err != nil {
		fmt.Println("cannot trace:", err)
		return
	}

	ctx, task := trace.NewTask(context.Background(), "request")
	trace.WithRegion(ctx, "parse", func() { time.Sleep(time.Millisecond) })
	region := trace.StartRegion(ctx, "query")
	time.Sleep(time.Millisecond)
	region.End()
	trace.Log(ctx, "user", "ada")
	task.End()

	trace.Stop() // stop before reading the buffer
	fmt.Println(buf.Len() > 0, trace.IsEnabled())
	// Output: true false
}

// Example_disabled shows that instrumentation is harmless when tracing is off.
func Example_disabled() {
	fmt.Println(trace.IsEnabled())

	ctx, task := trace.NewTask(context.Background(), "noop")
	defer task.End()
	trace.WithRegion(ctx, "work", func() { fmt.Println("ran") })
	trace.Log(ctx, "category", "message")
	// Output:
	// false
	// ran
}
