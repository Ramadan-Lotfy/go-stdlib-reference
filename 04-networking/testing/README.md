# testing

## Overview
The `testing` package and the `go test` command are Go's built-in test
framework: no external library is needed for unit tests, benchmarks, fuzzing,
examples or coverage. Test files end in `_test.go`; test functions are
`TestXxx(t *testing.T)`.

| Command | What it does |
|---------|--------------|
| `go test ./...` | Run all tests in all packages |
| `go test -v -run 'TestAdd/neg'` | Verbose; run tests matching a regexp (`/` selects subtests) |
| `go test -race ./...` | Race detector |
| `go test -count=1 ./...` | Disable result caching |
| `go test -short ./...` | Set `testing.Short()` (skip slow tests) |
| `go test -cover -coverprofile=c.out ./...` | Coverage; view with `go tool cover -html=c.out` |
| `go test -bench=. -benchmem` | Benchmarks with allocation counts |
| `go test -fuzz=FuzzReverse -fuzztime=10s` | Fuzzing |
| `go test -shuffle=on -timeout=30s` | Random order; fail on hangs |

## Core API

| Function / Type | Purpose |
|---|---|
| `func TestXxx(t *testing.T)` | Unit test; `t.Errorf` continues, `t.Fatalf` stops this test |
| `t.Run(name, func(t *testing.T){...})` | Subtests (used for table-driven tests) |
| `t.Parallel()` | Run this test in parallel with other parallel tests |
| `t.Helper()` | Report failures at the caller's line |
| `t.Cleanup`, `t.TempDir`, `t.Setenv` | Automatic cleanup, temp directory, restored env var |
| `t.Skip`, `testing.Short()` | Conditionally skip |
| `func BenchmarkXxx(b *testing.B)` | Benchmark: loop `b.N` times |
| `func FuzzXxx(f *testing.F)` | Fuzz test with seed corpus (`f.Add`) and `f.Fuzz` |
| `func ExampleXxx()` with `// Output:` | Verified examples shown in the docs |
| `func TestMain(m *testing.M)` | Package-level setup and teardown |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./04-networking/testing/
```

## Gotchas
- **`t.Errorf` does not stop the test; `t.Fatalf` does.** Use `Fatal` when continuing would panic (for example after a failed setup).
- **Never call `t.Fatal` from a goroutine other than the test's own.** Send errors back on a channel, or use `t.Error` and return.
- **Before Go 1.22, parallel subtests captured the loop variable** and all ran the last case. Since 1.22 each iteration has its own variable; old code copied it with `tt := tt`.
- **Test names should read like sentences** (`"negative numbers"`); spaces become underscores in `-run` patterns.
- **Mark helpers with `t.Helper()`** so failures point at the test, not at the helper.
- **Prefer `t.TempDir()`, `t.Setenv()` and `t.Cleanup()`** over manual setup and `defer`; they also work in subtests and helpers.
- **Test behavior through the public API.** Put tests in `package foo_test` (external) to keep honest; use `package foo` only for unexported internals.
- **Report both values:** `got X, want Y`, in that order, plus the input. A failure message should be understandable without opening the test.
- **Do not use `time.Sleep` to wait for goroutines.** Use channels, `sync.WaitGroup` or polling with a deadline.
- **Results are cached.** If a test reads files or env vars, `go test -count=1` forces a re-run.
- **Benchmarks must use the result** (assign to a package-level variable) so the compiler cannot delete the work.

## Java comparison

| Java | Go |
|---|---|
| JUnit `@Test` | `func TestXxx(t *testing.T)` |
| `@ParameterizedTest` | table-driven test + `t.Run` |
| AssertJ / Hamcrest | plain `if got != want { t.Errorf(...) }` (or `testify`) |
| `@BeforeEach` / `@AfterEach` | setup in the test + `t.Cleanup` |
| `@BeforeAll` | `TestMain` |
| `@TempDir` | `t.TempDir()` |
| Mockito | small interfaces + hand-written fakes |
| JMH | `testing.B` benchmarks |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/testing
- Source: https://cs.opensource.google/go/go/+/master:src/testing/
