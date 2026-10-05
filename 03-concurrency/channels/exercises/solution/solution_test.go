package solution

import (
	"context"
	"slices"
	"testing"
	"time"
)

func collect[T any](ch <-chan T) []T {
	var out []T
	for v := range ch {
		out = append(out, v)
	}
	return out
}

func TestPipeline(t *testing.T) {
	ctx := context.Background()
	got := collect(Square(ctx, Generate(ctx, 1, 2, 3)))
	if want := []int{1, 4, 9}; !slices.Equal(got, want) {
		t.Errorf("pipeline = %v, want %v", got, want)
	}
}

func TestMerge(t *testing.T) {
	ctx := context.Background()
	got := collect(Merge(Generate(ctx, 1, 2), Generate(ctx, 3), Generate(ctx)))
	slices.Sort(got)
	if want := []int{1, 2, 3}; !slices.Equal(got, want) {
		t.Errorf("Merge() = %v, want %v", got, want)
	}
}

func TestCancelStopsPipeline(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	out := Square(ctx, Generate(ctx, make([]int, 1000)...))
	<-out // take one value, then stop
	cancel()

	done := make(chan struct{})
	go func() {
		for range out { // must end soon after cancel
		}
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("pipeline did not stop after cancel (goroutine leak)")
	}
}
