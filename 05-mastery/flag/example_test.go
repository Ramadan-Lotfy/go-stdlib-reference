package flagref_test

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// newSet returns a FlagSet that never exits the program and stays quiet.
func newSet() *flag.FlagSet {
	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	return fs
}

// Example_basic defines flags, parses an explicit slice and reads the results.
func Example_basic() {
	fs := newSet()
	name := fs.String("name", "world", "who to greet")
	count := fs.Int("count", 1, "how many times")
	verbose := fs.Bool("v", false, "verbose output")

	err := fs.Parse([]string{"-name", "ada", "-count=3", "-v", "rest1", "rest2"})
	fmt.Println(*name, *count, *verbose, fs.Args(), err)
	// Output: ada 3 true [rest1 rest2] <nil>
}

// Example_defaults shows that unset flags keep their defaults.
func Example_defaults() {
	fs := newSet()
	name := fs.String("name", "world", "who to greet")
	timeout := fs.Duration("timeout", 5*time.Second, "request timeout")

	fs.Parse(nil)
	fmt.Println(*name, *timeout)

	fs.Parse([]string{"-timeout", "1m30s"})
	fmt.Println(*timeout)
	// Output:
	// world 5s
	// 1m30s
}

// listFlag collects a flag that can be repeated.
type listFlag []string

func (l *listFlag) String() string     { return strings.Join(*l, ",") }
func (l *listFlag) Set(v string) error { *l = append(*l, v); return nil }

// Example_customValue implements flag.Value for a repeatable flag.
func Example_customValue() {
	fs := newSet()
	var tags listFlag
	fs.Var(&tags, "tag", "repeatable tag")

	fs.Parse([]string{"-tag", "a", "-tag", "b"})
	fmt.Println(tags, len(tags))
	// Output: [a b] 2
}

// Example_errors shows the three error cases of Parse.
func Example_errors() {
	fs := newSet()
	fs.Int("count", 1, "how many")

	fmt.Println(fs.Parse([]string{"-nope"}))
	fmt.Println(fs.Parse([]string{"-count=x"}))
	fmt.Println(fs.Parse([]string{"-h"}) == flag.ErrHelp)
	// Output:
	// flag provided but not defined: -nope
	// invalid value "x" for flag -count: parse error
	// true
}

// Example_orderMatters shows the two classic surprises.
func Example_orderMatters() {
	fs := newSet()
	v := fs.Bool("v", false, "verbose")
	fs.Parse([]string{"file.txt", "-v"}) // flag after a positional argument
	fmt.Println(*v, fs.Args())

	fs = newSet()
	v = fs.Bool("v", false, "verbose")
	fs.Parse([]string{"-v", "false"}) // boolean flags do not take a separate value
	fmt.Println(*v, fs.Args())
	// Output:
	// false [file.txt -v]
	// true [false]
}

// Example_visit lists which flags the user actually set.
func Example_visit() {
	fs := newSet()
	fs.String("name", "world", "who to greet")
	fs.Int("count", 1, "how many")
	fs.Parse([]string{"-count", "2"})

	fs.Visit(func(f *flag.Flag) { fmt.Println("set:", f.Name, f.Value) })
	fs.VisitAll(func(f *flag.Flag) { fmt.Println("all:", f.Name, f.Value) })
	// Output:
	// set: count 2
	// all: count 2
	// all: name world
}

// Example_usage prints the generated help text.
func Example_usage() {
	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.String("name", "world", "who to greet")
	fs.Int("count", 1, "how many times")

	var sb strings.Builder
	fs.SetOutput(&sb)
	fs.PrintDefaults()
	help := sb.String()

	fmt.Println(strings.Contains(help, "-count int"), strings.Contains(help, `(default "world")`))
	fmt.Fprint(os.Stdout, "")
	// Output: true true
}

func runCommand(args []string) string {
	add := flag.NewFlagSet("add", flag.ContinueOnError)
	n := add.Int("n", 1, "amount")
	switch args[0] {
	case "add":
		add.Parse(args[1:])
		return fmt.Sprintf("add n=%d args=%v", *n, add.Args())
	case "version":
		return "v1.0.0"
	default:
		return "unknown command " + args[0]
	}
}

// Example_subcommands dispatches on the first argument, each with its own FlagSet.
func Example_subcommands() {
	fmt.Println(runCommand([]string{"add", "-n", "5", "item"}))
	fmt.Println(runCommand([]string{"version"}))
	fmt.Println(runCommand([]string{"push"}))
	// Output:
	// add n=5 args=[item]
	// v1.0.0
	// unknown command push
}
