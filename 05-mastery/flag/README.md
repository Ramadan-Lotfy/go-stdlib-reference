# flag

## Overview
`flag` parses command-line flags such as `-addr :8080 -v`. For small tools it
is all you need; larger CLIs often add a library such as `cobra` for
subcommands and shell completion. Use a `FlagSet` instead of the global
functions in anything that must be tested or has subcommands.

## Core API

| Function / Type | Purpose |
|---|---|
| `flag.String/Int/Bool/Duration(name, default, usage)` | Define a flag; returns a pointer |
| `flag.StringVar(&v, name, default, usage)` | Define a flag that writes into your variable |
| `flag.Var(value, name, usage)`, `flag.Func` | Custom types and repeatable flags |
| `flag.Parse()`, `flag.Args()`, `flag.NArg()` | Parse `os.Args[1:]`; positional arguments after the flags |
| `flag.NewFlagSet(name, errorHandling)` | Independent set: testable, and the basis of subcommands |
| `fs.Parse(args)` | Parse any slice (tests pass explicit slices) |
| `fs.Visit`, `fs.VisitAll`, `fs.PrintDefaults` | Inspect which flags were set; print help |
| `flag.ContinueOnError`, `ExitOnError`, `PanicOnError` | What `Parse` does on bad input |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/flag/
```

## Gotchas
- **Flags must come before positional arguments.** Parsing stops at the first non-flag, so `app file -v` leaves `-v` in `Args()`.
- **Boolean flags take no separate value:** `-v` means true, `-v=false` means false. `-v false` sets `v` to true and leaves `false` as an argument.
- **Call `flag.Parse()` before reading any flag value,** because the pointers are filled in by `Parse`.
- **`-flag` and `--flag` are the same,** and `-flag=value` and `-flag value` are the same (except booleans). There are no short-option groups like `-abc`.
- **Use `NewFlagSet(..., flag.ContinueOnError)` and `SetOutput` in functions you test.** The default `ExitOnError` calls `os.Exit`.
- **`-h` or `-help` return `flag.ErrHelp`** when the flag is not defined. Treat it as success (exit code 0).
- **Defaults are shown by `PrintDefaults`;** write clear usage text because users see it.
- **Priority convention:** flag overrides environment variable overrides config file overrides default. Implement it explicitly.
- **The global `flag.CommandLine` is shared by all imported packages,** which is why libraries should not define global flags.

## Java comparison

| Java | Go |
|---|---|
| picocli / Apache Commons CLI | `flag` (simple) or `cobra` (large CLIs) |
| `@Option(names="-v")` | `fs.Bool("v", false, "verbose")` |
| `@Parameters` | `fs.Args()` |
| Spring `@Value("${...}")` | flags + `os.Getenv` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/flag
- Source: https://cs.opensource.google/go/go/+/master:src/flag/
