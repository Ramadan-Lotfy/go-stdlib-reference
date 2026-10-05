// Package solution holds the reference solutions for the os exercises.
package solution

import (
	"fmt"
	"os"
	"strings"
)

// WordCountFile returns the number of whitespace-separated words in a file.
func WordCountFile(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("word count: %w", err)
	}
	return len(strings.Fields(string(data))), nil
}

// AppendLine appends line plus a newline to path, creating the file if needed.
func AppendLine(path, line string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintln(f, line); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
