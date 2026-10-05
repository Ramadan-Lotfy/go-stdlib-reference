# os/exec

## Overview
`os/exec` runs external programs. A `*exec.Cmd` describes the command, its
arguments, environment, working directory and input/output; `Run`, `Output`,
`CombinedOutput` or `Start` + `Wait` execute it. Arguments are passed directly to
the program, **without a shell**, so there is no quoting or globbing and no
shell injection unless you deliberately run `sh -c`.

The examples need a program that behaves the same on every operating system, so
the test binary turns itself into a tiny helper: `TestMain` checks an environment
variable and, if set, acts as the child command. This is the technique the Go
standard library itself uses to test `os/exec`.

## Core API

| Function / Type | Purpose |
|---|---|
| `exec.Command(name, args...)` | Describe a command (not started yet) |
| `cmd.Run()` | Start and wait; error if it cannot start or exits non-zero |
| `cmd.Output()`, `cmd.CombinedOutput()` | Run and capture stdout (and stderr) |
| `cmd.Start()` / `cmd.Wait()` | Run in the background; use pipes in between |
| `cmd.StdinPipe/StdoutPipe/StderrPipe`, `cmd.Stdin/Stdout/Stderr` | Connect input and output |
| `cmd.Env`, `cmd.Dir`, `cmd.Environ()` | Environment and working directory |
| `exec.CommandContext(ctx, ...)`, `cmd.WaitDelay` | Kill the process when the context ends |
| `*exec.ExitError`, `ExitCode()` | Non-zero exit status |
| `exec.LookPath`, `exec.ErrNotFound` | Find an executable in `PATH` |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/osexec/
```

## Gotchas
- **There is no shell.** `exec.Command("ls", "*.go")` passes the literal `*.go`. That is also what protects you from injection. Use `sh -c` only with trusted, constant strings.
- **Never build a command line by concatenating user input;** pass each value as a separate argument.
- **A non-zero exit status is an error (`*exec.ExitError`),** not a normal return. Use `errors.As` and `ExitCode()`.
- **Use `CommandContext` for anything that can hang,** and set `cmd.WaitDelay` so a stuck process that ignores the kill cannot block `Wait` forever.
- **`Output()` captures only stdout;** stderr is saved in `ExitError.Stderr` when `cmd.Stderr` is nil. Set `cmd.Stderr` yourself for logging.
- **A `Cmd` can only be run once.** Create a new one for each run.
- **`cmd.Env = nil` inherits the parent environment;** setting it replaces it entirely. Start from `os.Environ()` (or `cmd.Environ()`) and append.
- **With `Start`, always call `Wait`,** or you leak the process and its pipes. Read all output before `Wait` when using `StdoutPipe`.
- **Since Go 1.19 `LookPath` will not run executables found in the current directory** (`exec.ErrDot`), which prevents a class of attacks on Windows.
- **Killing the child does not kill its children.** Process groups need OS-specific code (`SysProcAttr`).

## Java comparison

| Java | Go |
|---|---|
| `ProcessBuilder` / `Runtime.exec` | `exec.Command` |
| `Process.waitFor()` | `cmd.Wait()` / `cmd.Run()` |
| `process.getInputStream()` | `cmd.StdoutPipe()` or `cmd.Output()` |
| `Process.destroy()` | `cmd.Process.Kill()` or `CommandContext` |
| `pb.environment()` | `cmd.Env` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/os/exec
- Source: https://cs.opensource.google/go/go/+/master:src/os/exec/
