package errorsref_test

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type ValidationError struct{ Field, Reason string }

func (e *ValidationError) Error() string { return e.Field + ": " + e.Reason }

func find(id int) (string, error) {
	if id != 1 {
		return "", ErrNotFound
	}
	return "ada", nil
}

func validate(age int) error {
	if age < 0 {
		return &ValidationError{"age", "must be >= 0"}
	}
	return nil
}

// Example_sentinel compares against a package-level sentinel error.
func Example_sentinel() {
	_, err := find(2)
	fmt.Println(err == ErrNotFound, err)
	// Output: true not found
}

// Example_wrapping adds context with %w while keeping the cause reachable.
func Example_wrapping() {
	_, err := find(2)
	err = fmt.Errorf("load user 2: %w", err)
	fmt.Println(err)
	fmt.Println(err == ErrNotFound)
	fmt.Println(errors.Is(err, ErrNotFound))
	fmt.Println(errors.Unwrap(err) == ErrNotFound)
	// Output:
	// load user 2: not found
	// false
	// true
	// true
}

// Example_customType extracts a typed error from a chain with errors.As.
func Example_customType() {
	err := fmt.Errorf("signup: %w", validate(-1))
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("field:", ve.Field)
		fmt.Println("reason:", ve.Reason)
	}
	// Output:
	// field: age
	// reason: must be >= 0
}

type tempErr struct{}

func (tempErr) Error() string   { return "network down" }
func (tempErr) Temporary() bool { return true }

// Example_asInterface matches on behavior (an interface) instead of a concrete type.
func Example_asInterface() {
	err := fmt.Errorf("fetch: %w", tempErr{})
	var t interface{ Temporary() bool }
	fmt.Println(errors.As(err, &t) && t.Temporary())
	// Output: true
}

// Example_join combines independent errors. Is and As search all of them.
func Example_join() {
	e1 := errors.New("disk full")
	e2 := errors.New("network down")
	err := errors.Join(e1, e2)
	fmt.Println(err)
	fmt.Println(errors.Is(err, e1), errors.Is(err, e2))
	// Output:
	// disk full
	// network down
	// true true
}

// Example_multipleWrap wraps two errors in one fmt.Errorf call (Go 1.20+).
func Example_multipleWrap() {
	e1 := errors.New("disk full")
	e2 := errors.New("network down")
	err := fmt.Errorf("both failed: %w, %w", e1, e2)
	fmt.Println(err)
	fmt.Println(errors.Is(err, e1), errors.Is(err, e2))
	// Output:
	// both failed: disk full, network down
	// true true
}

func mayFail(fail bool) error {
	var ve *ValidationError // nil pointer
	if fail {
		ve = &ValidationError{"x", "bad"}
	}
	return ve // BUG: a non-nil interface holding a nil pointer
}

// Example_typedNilTrap shows why you should return a literal nil.
func Example_typedNilTrap() {
	fmt.Println(mayFail(false) == nil)
	// Output: false
}
