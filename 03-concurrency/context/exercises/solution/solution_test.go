package solution

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	if err := Sleep(context.Background(), time.Millisecond); err != nil {
		t.Errorf("Sleep() = %v, want nil", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Sleep(ctx, time.Hour); !errors.Is(err, context.Canceled) {
		t.Errorf("Sleep(cancelled) = %v, want context.Canceled", err)
	}
}

func slow(ctx context.Context) (string, error) {
	select {
	case <-time.After(5 * time.Second):
		return "slow", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func fast(ctx context.Context) (string, error) { return "fast", nil }

func TestFirstResult(t *testing.T) {
	start := time.Now()
	got, err := FirstResult(context.Background(), slow, fast)
	if err != nil || got != "fast" {
		t.Errorf("FirstResult() = %q, %v; want fast", got, err)
	}
	if time.Since(start) > time.Second {
		t.Error("slow function was not cancelled")
	}
}

func TestFirstResultAllFail(t *testing.T) {
	boom := errors.New("boom")
	failing := func(ctx context.Context) (string, error) { return "", boom }
	_, err := FirstResult(context.Background(), failing, failing)
	if !errors.Is(err, boom) {
		t.Errorf("err = %v, want it to wrap boom", err)
	}

	if _, err := FirstResult(context.Background()); err == nil {
		t.Error("no functions should be an error")
	}
}

func TestFirstResultContextDone(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	block := func(ctx context.Context) (string, error) {
		<-ctx.Done()
		return "", ctx.Err()
	}
	_, err := FirstResult(ctx, block)
	if err == nil {
		t.Error("expected an error when the context ends")
	}
}
