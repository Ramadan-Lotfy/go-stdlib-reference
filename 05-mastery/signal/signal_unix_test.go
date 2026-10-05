//go:build !windows

package signalref_test

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Example_unixSignal sends a real signal to this process and observes it.
// (This file is skipped on Windows, where there is no SIGUSR1.)
func Example_unixSignal() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGUSR1)
	defer stop()

	syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)
	<-ctx.Done()
	fmt.Println("received:", ctx.Err())
	// Output: received: context canceled
}

// Example_channel is the classic form: a buffered channel receives signals.
func Example_channel() {
	sigs := make(chan os.Signal, 1) // must be buffered: Notify never blocks
	signal.Notify(sigs, syscall.SIGUSR2)
	defer signal.Stop(sigs)

	syscall.Kill(syscall.Getpid(), syscall.SIGUSR2)
	s := <-sigs
	fmt.Println(s == syscall.SIGUSR2)
	// Output: true
}
