// Package solution holds the reference solutions for the os/exec exercises.
package solution

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Run executes name with args and returns its combined output with
// surrounding whitespace trimmed.
//
//   - The command is killed after timeout (or when ctx ends) and the error
//     wraps ctx.Err(), so errors.Is(err, context.DeadlineExceeded) works
//   - A non-zero exit returns an error that wraps the *exec.ExitError and
//     includes the output
//   - A missing program returns an error wrapping exec.ErrNotFound
func Run(ctx context.Context, timeout time.Duration, name string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	cmd.WaitDelay = time.Second
	out, err := cmd.CombinedOutput()
	text := strings.TrimSpace(string(out))

	if ctxErr := ctx.Err(); ctxErr != nil {
		return text, fmt.Errorf("run %s: %w", name, ctxErr)
	}
	if err != nil {
		return text, fmt.Errorf("run %s: %w: %s", name, err, text)
	}
	return text, nil
}
