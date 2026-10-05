package slogref_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

// newLogger builds a logger without timestamps so output is deterministic.
func newLogger(w io.Writer, json bool, level slog.Leveler) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{} // drop the time attribute
			}
			return a
		},
	}
	if json {
		return slog.New(slog.NewJSONHandler(w, opts))
	}
	return slog.New(slog.NewTextHandler(w, opts))
}

// Example_text logs with the key=value text handler.
func Example_text() {
	l := newLogger(os.Stdout, false, slog.LevelInfo)
	l.Info("user logged in", "user", "ada", "attempts", 2)
	// Output: level=INFO msg="user logged in" user=ada attempts=2
}

// Example_json logs the same record as one JSON object per line.
func Example_json() {
	l := newLogger(os.Stdout, true, slog.LevelInfo)
	l.Info("hello", "user", "ada")
	l.WithGroup("req").Info("done", "id", 7)
	// Output:
	// {"level":"INFO","msg":"hello","user":"ada"}
	// {"level":"INFO","msg":"done","req":{"id":7}}
}

// Example_levels filters by level and changes it at runtime.
func Example_levels() {
	lv := new(slog.LevelVar) // starts at Info
	l := newLogger(os.Stdout, false, lv)

	l.Debug("hidden")
	l.Info("shown")
	lv.Set(slog.LevelDebug)
	l.Debug("now visible")
	// Output:
	// level=INFO msg=shown
	// level=DEBUG msg="now visible"
}

// Example_withAndGroups adds shared context and namespaces attributes.
func Example_withAndGroups() {
	l := newLogger(os.Stdout, false, slog.LevelInfo).With("service", "api")
	l.Info("started")
	l.WithGroup("req").Info("done", "id", 7)
	l.Info("http", slog.Group("http", "method", "GET", "status", 200))
	// Output:
	// level=INFO msg=started service=api
	// level=INFO msg=done service=api req.id=7
	// level=INFO msg=http service=api http.method=GET http.status=200
}

// Example_attrs uses typed attributes and logs an error value.
func Example_attrs() {
	l := newLogger(os.Stdout, false, slog.LevelInfo)
	l.LogAttrs(context.Background(), slog.LevelInfo, "typed",
		slog.String("s", "x"),
		slog.Int("n", 3),
		slog.Duration("d", 1500*time.Millisecond),
		slog.Bool("ok", true),
	)
	l.Error("failed", "err", errors.New("disk full"))
	// Output:
	// level=INFO msg=typed s=x n=3 d=1.5s ok=true
	// level=ERROR msg=failed err="disk full"
}

// Secret hides its value from logs.
type Secret string

func (Secret) LogValue() slog.Value { return slog.StringValue("[REDACTED]") }

// User logs as a group of attributes.
type User struct {
	ID   int
	Name string
}

func (u User) LogValue() slog.Value {
	return slog.GroupValue(slog.Int("id", u.ID), slog.String("name", u.Name))
}

// Example_logValuer redacts secrets and expands structs with LogValuer.
func Example_logValuer() {
	l := newLogger(os.Stdout, false, slog.LevelInfo)
	l.Info("login", "password", Secret("hunter2"))
	l.Info("created", "user", User{1, "ada"})
	// Output:
	// level=INFO msg=login password=[REDACTED]
	// level=INFO msg=created user.id=1 user.name=ada
}

// Example_setDefault installs a logger for the package-level functions.
func Example_setDefault() {
	old := slog.Default()
	defer slog.SetDefault(old)

	slog.SetDefault(newLogger(os.Stdout, false, slog.LevelInfo))
	slog.Info("via default", "k", "v")
	fmt.Println("done")
	// Output:
	// level=INFO msg="via default" k=v
	// done
}
