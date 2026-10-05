# flag exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/flag/exercises/...
```

## Exercise: ParseArgs

Write `ParseArgs(args []string) (Config, error)` for a small server command:

| Flag | Meaning | Default |
|------|---------|---------|
| `-addr string` | listen address | `:8080` |
| `-v` | verbose | `false` |
| `-timeout duration` | request timeout | `5s` |
| `-tag string` | repeatable; empty value is an error | none |

Positional arguments after the flags go to `Config.Files`.

Rules:

- Use a **new `FlagSet`** with `flag.ContinueOnError` and discard its output, so
  the function can be tested and never calls `os.Exit` or prints
- Unknown flags, bad durations, a missing value and an empty `-tag` all return
  an error

Hints: `fs.StringVar`, `fs.BoolVar`, `fs.DurationVar`, `fs.Func` for the
repeatable flag, `fs.Args()`.

## Bonus
Add environment-variable fallbacks (`APP_ADDR`, `APP_TIMEOUT`) so the order of
priority is flag, then environment, then default. Hint: use `fs.Visit` to find
which flags were set explicitly.
