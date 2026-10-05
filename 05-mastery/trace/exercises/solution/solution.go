// Package solution holds the reference solutions for the runtime/trace exercises.
package solution

import (
	"context"
	"io"
	"runtime/trace"
)

// TraceTo runs f while writing an execution trace to w. f receives a context
// that carries a trace task named "work", so regions it creates are grouped.
// It returns the error from trace.Start if tracing is already active.
func TraceTo(w io.Writer, f func(ctx context.Context)) error {
	if err := trace.Start(w); err != nil {
		return err
	}
	defer trace.Stop()

	ctx, task := trace.NewTask(context.Background(), "work")
	defer task.End()
	f(ctx)
	return nil
}
