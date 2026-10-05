// Package testingref contains the small functions that example_test.go tests.
package testingref

import (
	"errors"
	"unicode"
)

// ErrDivideByZero is returned by Divide when the divisor is zero.
var ErrDivideByZero = errors.New("divide by zero")

// Add returns a + b.
func Add(a, b int) int { return a + b }

// Divide returns a / b, or ErrDivideByZero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// Reverse returns s with its characters (runes) in reverse order.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// IsPalindrome reports whether s reads the same backwards, ignoring case,
// spaces and punctuation.
func IsPalindrome(s string) bool {
	var clean []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			clean = append(clean, unicode.ToLower(r))
		}
	}
	return string(clean) == Reverse(string(clean))
}
