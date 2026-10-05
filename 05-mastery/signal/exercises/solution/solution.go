// Package solution holds the reference solutions for the os/signal exercises.
package solution

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

// Serve runs h on ln until ctx is cancelled, then shuts the server down
// gracefully, waiting up to five seconds for in-flight requests.
//
// It returns nil after a clean shutdown, the server's error if it fails
// earlier, and the shutdown error if shutdown times out.
func Serve(ctx context.Context, ln net.Listener, h http.Handler) error {
	srv := &http.Server{Handler: h, ReadHeaderTimeout: 5 * time.Second}
	errCh := make(chan error, 1)
	go func() { errCh <- srv.Serve(ln) }()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
	}

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
