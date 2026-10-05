# os

## Overview
`os` is Go's interface to the operating system: files and directories,
environment variables, process arguments and exit codes. File operations
return `*fs.PathError`, which works with `errors.Is(err, fs.ErrNotExist)`.

## Core API

| Function / Type | Purpose |
|---|---|
| `os.ReadFile`, `os.WriteFile` | Read or write a whole file in one call |
| `os.Open`, `os.Create`, `os.OpenFile` | Get an `*os.File` for streaming I/O |
| `os.Stat` | File size, mode, modification time |
| `os.ReadDir` | List a directory, sorted by name |
| `os.MkdirAll`, `os.Remove`, `os.RemoveAll`, `os.Rename` | Manage files and directories |
| `os.MkdirTemp`, `os.CreateTemp` | Temporary files and directories |
| `os.Getenv`, `os.LookupEnv`, `os.Setenv` | Environment variables |
| `os.Args`, `os.Exit` | Command-line arguments and exit status |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/os/
```

## Gotchas
- **Always `Close` files**, normally with `defer f.Close()` right after the error check. For writes, check the error from `Close` too.
- **`os.Exit` skips deferred functions.** Do cleanup first, or keep `main` small and return an exit code from a `run()` function.
- **Test for missing files with `errors.Is(err, fs.ErrNotExist)`**, not by comparing error strings. `os.IsNotExist` is the older form.
- **`Getenv` cannot tell empty from unset.** Use `LookupEnv` when the difference matters.
- **Permission bits are `0o644`-style octal** and are filtered by the process umask on Unix. They mostly do nothing on Windows.
- **`ReadFile` loads everything into memory.** Stream large files with `Open` and `bufio`.
- **Use `filepath.Join`** to build paths. Never concatenate with `/`.

## Java comparison

| Java | Go |
|---|---|
| `Files.readAllBytes`, `Files.writeString` | `os.ReadFile`, `os.WriteFile` |
| `Files.createTempFile` | `os.CreateTemp` |
| `System.getenv("X")` | `os.Getenv("X")` |
| `File.exists()` | `os.Stat` and `errors.Is(err, fs.ErrNotExist)` |
| `System.exit(1)` | `os.Exit(1)` |
| `args` in `main(String[] args)` | `os.Args[1:]` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/os
- Source: https://cs.opensource.google/go/go/+/master:src/os/
