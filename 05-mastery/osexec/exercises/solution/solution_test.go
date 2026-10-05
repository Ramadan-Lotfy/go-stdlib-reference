package solution

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// TestMain lets this test binary act as the child program (see the osexec README).
func TestMain(m *testing.M) {
	if os.Getenv("OSEXEC_HELPER") == "1" {
		os.Exit(helper(os.Args[1:]))
	}
	os.Exit(m.Run())
}

func helper(args []string) int {
	switch args[0] {
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "fail":
		fmt.Println("partial output")
		fmt.Fprintln(os.Stderr, "bad")
		return 3
	case "sleep":
		time.Sleep(time.Minute)
	}
	return 0
}

func TestRun(t *testing.T) {
	t.Setenv("OSEXEC_HELPER", "1") // the child inherits this
	self := os.Args[0]

	out, err := Run(context.Background(), 5*time.Second, self, "echo", "hello")
	if err != nil || out != "hello" {
		t.Errorf("echo: %q, %v", out, err)
	}

	out, err = Run(context.Background(), 5*time.Second, self, "fail")
	var ee *exec.ExitError
	if !errors.As(err, &ee) || ee.ExitCode() != 3 {
		t.Errorf("fail: err = %v, want ExitError with code 3", err)
	}
	if !strings.Contains(out, "partial output") || !strings.Contains(err.Error(), "bad") {
		t.Errorf("fail: output %q and error %q should include the command output", out, err)
	}

	start := time.Now()
	_, err = Run(context.Background(), 100*time.Millisecond, self, "sleep")
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("sleep: err = %v, want DeadlineExceeded", err)
	}
	if time.Since(start) > 5*time.Second {
		t.Error("the command was not killed on timeout")
	}

	_, err = Run(context.Background(), time.Second, "definitely-not-a-real-program")
	if !errors.Is(err, exec.ErrNotFound) {
		t.Errorf("missing program: err = %v, want exec.ErrNotFound", err)
	}
}
