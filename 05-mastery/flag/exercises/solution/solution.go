// Package solution holds the reference solutions for the flag exercises.
package solution

import (
	"errors"
	"flag"
	"io"
	"time"
)

// Config holds the parsed command line.
type Config struct {
	Addr    string
	Verbose bool
	Timeout time.Duration
	Tags    []string
	Files   []string
}

// ParseArgs parses command-line arguments (without the program name):
//
//	-addr string      listen address (default ":8080")
//	-v                verbose output
//	-timeout duration request timeout (default 5s)
//	-tag string       repeatable; an empty value is an error
//
// Everything after the flags is returned in Files. Errors are returned, never
// printed, and the function never exits the program.
func ParseArgs(args []string) (Config, error) {
	var cfg Config
	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.StringVar(&cfg.Addr, "addr", ":8080", "listen address")
	fs.BoolVar(&cfg.Verbose, "v", false, "verbose output")
	fs.DurationVar(&cfg.Timeout, "timeout", 5*time.Second, "request timeout")
	fs.Func("tag", "repeatable tag", func(s string) error {
		if s == "" {
			return errors.New("tag must not be empty")
		}
		cfg.Tags = append(cfg.Tags, s)
		return nil
	})
	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}
	cfg.Files = fs.Args()
	return cfg, nil
}
