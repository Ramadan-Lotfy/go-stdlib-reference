# errors

## Overview
`errors` creates, wraps, inspects and combines errors. Go has no exceptions: a
function that can fail returns an `error` value, and callers decide what to do
with it. Wrapping with `%w` keeps the cause reachable for `errors.Is` and `errors.As`.

## Core API

| Function / Type | Purpose |
|---|---|
| `errors.New` | Create a simple error (use for sentinel values) |
| `errors.Is(err, target)` | Does any error in the chain equal `target`? |
| `errors.As(err, &target)` | Find the first error in the chain of a given type |
| `errors.Unwrap` | Return the next error in the chain |
| `errors.Join` | Combine several errors into one |
| `fmt.Errorf("...%w", err)` | Wrap an error with context |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./01-foundations/errors/
```

## Gotchas
- **Wrap with `%w`, not `%v`.** `%v` keeps the text but loses the chain, so `errors.Is` fails.
- **Compare with `errors.Is`, not `==`**, once errors may be wrapped.
- **`errors.As` needs a pointer** to the target type (`&ve`, where `ve` is `*ValidationError`). `go vet` checks this.
- **Typed nil trap.** Returning a nil `*MyErr` as `error` gives a non-nil interface. Return a literal `nil` instead.
- **Error strings are lowercase with no trailing punctuation**, because they get chained: `load user 2: not found`.
- **Add context at each layer**, but do not log and return the same error, or it is reported twice.

## Java comparison

| Java | Go |
|---|---|
| `try { } catch (IOException e)` | `if err != nil { }` |
| `throw new X()` | `return X{}` (an error value) |
| `catch (NotFound e)` by type | `errors.As(err, &target)` |
| `e.getCause()` | `errors.Unwrap(err)` |
| `addSuppressed` | `errors.Join` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/errors
- Source: https://cs.opensource.google/go/go/+/master:src/errors/
