// Package solution holds the reference solutions for the context exercises.
package solution

import (
	"context"
	"errors"
	"time"
)

// Sleep waits for d or until ctx is done, whichever comes first. It returns
// ctx.Err() when interrupted.
func Sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FirstResult runs every function concurrently and returns the first
// successful result, cancelling the rest. If all fail it returns their errors
// joined. If ctx ends first it returns ctx.Err().
func FirstResult(ctx context.Context, fns ...func(context.Context) (string, error)) (string, error) {
	if len(fns) == 0 {
		return "", errors.New("no functions to run")
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	type result struct {
		val string
		err error
	}
	ch := make(chan result, len(fns)) // buffered: late senders never block
	for _, fn := range fns {
		go func() {
			v, err := fn(ctx)
			ch <- result{v, err}
		}()
	}

	var errs []error
	for range fns {
		select {
		case r := <-ch:
			if r.err == nil {
				return r.val, nil
			}
			errs = append(errs, r.err)
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
	return "", errors.Join(errs...)
}
