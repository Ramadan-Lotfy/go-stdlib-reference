// Package solution holds the reference solutions for the errors exercises.
package solution

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type ValidationError struct{ Field string }

func (e *ValidationError) Error() string { return "invalid " + e.Field }

// Classify returns "ok", "not found", "invalid: <field>" or "unknown".
func Classify(err error) string {
	var ve *ValidationError
	switch {
	case err == nil:
		return "ok"
	case errors.Is(err, ErrNotFound):
		return "not found"
	case errors.As(err, &ve):
		return "invalid: " + ve.Field
	default:
		return "unknown"
	}
}

// Retry calls fn up to attempts times. It returns nil on the first success,
// otherwise all errors joined, each prefixed with its attempt number.
func Retry(attempts int, fn func() error) error {
	if attempts <= 0 {
		return errors.New("attempts must be positive")
	}
	var errs []error
	for i := 1; i <= attempts; i++ {
		err := fn()
		if err == nil {
			return nil
		}
		errs = append(errs, fmt.Errorf("attempt %d: %w", i, err))
	}
	return errors.Join(errs...)
}
