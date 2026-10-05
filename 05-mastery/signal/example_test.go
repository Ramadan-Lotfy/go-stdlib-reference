package signalref_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Example_notifyContext shows the shape of the API without sending a signal.
// The returned context ends when a listed signal arrives or stop is called.
func Example_notifyContext() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	fmt.Println(ctx.Err())
	stop() // also restores default signal behavior
	<-ctx.Done()
	fmt.Println(ctx.Err())
	// Output:
	// <nil>
	// context canceled
}

// serve runs h until ctx ends, then shuts down gracefully.
func serve(ctx context.Context, ln net.Listener, h http.Handler) error {
	srv := &http.Server{Handler: h, ReadHeaderTimeout: 5 * time.Second}
	errCh := make(chan error, 1)
	go func() { errCh <- srv.Serve(ln) }()

	select {
	case err := <-errCh: // the server failed on its own
		return err
	case <-ctx.Done():
	}

	// Give in-flight requests time to finish. The shutdown context must NOT be
	// derived from the cancelled ctx, so detach it.
	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}
	if err := <-errCh; !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Example_gracefulShutdown serves requests until the context is cancelled. In
// a real program the context comes from signal.NotifyContext, so Ctrl+C or
// SIGTERM triggers the same path; here cancel() stands in for the signal.
func Example_gracefulShutdown() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // safe to call twice; also covers the early-return path
	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	url := "http://" + ln.Addr().String()

	done := make(chan error, 1)
	go func() {
		done <- serve(ctx, ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "alive")
		}))
	}()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body))

	cancel() // "SIGTERM arrives"
	fmt.Println(<-done)

	_, err = http.Get(url)
	fmt.Println(err != nil) // the listener is closed now
	// Output:
	// alive
	// <nil>
	// true
}

// Example_names prints how signals describe themselves.
func Example_names() {
	fmt.Println(os.Interrupt)
	fmt.Println(os.Interrupt == syscall.SIGINT)
	// Output:
	// interrupt
	// true
}
