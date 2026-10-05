package solution

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestServeGracefulShutdown(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	url := "http://" + ln.Addr().String()

	started := make(chan struct{})
	var once sync.Once
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		once.Do(func() { close(started) })
		time.Sleep(200 * time.Millisecond) // a slow in-flight request
		fmt.Fprint(w, "finished")
	})

	done := make(chan error, 1)
	go func() { done <- Serve(ctx, ln, handler) }()

	// Start a slow request, then cancel while it is in flight.
	bodyCh := make(chan string, 1)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			bodyCh <- "error: " + err.Error()
			return
		}
		defer resp.Body.Close()
		b, _ := io.ReadAll(resp.Body)
		bodyCh <- string(b)
	}()
	<-started
	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Errorf("Serve() = %v, want nil after a clean shutdown", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("Serve did not return after the context was cancelled")
	}

	if got := <-bodyCh; got != "finished" {
		t.Errorf("in-flight request got %q, want it to finish normally", got)
	}
	if _, err := http.Get(url); err == nil {
		t.Error("server should refuse new connections after shutdown")
	}
}
