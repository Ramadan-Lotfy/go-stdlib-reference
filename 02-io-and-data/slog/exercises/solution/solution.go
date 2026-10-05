// Package solution holds the reference solutions for the log/slog exercises.
package solution

import (
	"io"
	"log/slog"
)

// NewLogger returns a JSON logger that writes to w, omits the timestamp and
// adds a "service" attribute to every record.
func NewLogger(w io.Writer, service string) *slog.Logger {
	h := slog.NewJSONHandler(w, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{}
			}
			return a
		},
	})
	return slog.New(h).With("service", service)
}

// Secret is a string that never appears in logs.
type Secret string

// LogValue implements slog.LogValuer.
func (Secret) LogValue() slog.Value { return slog.StringValue("[REDACTED]") }
