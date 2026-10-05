# sync/atomic

## Overview
`sync/atomic` provides lock-free operations on single values: counters,
flags and pointers. Use it for simple shared state where a mutex would be
overkill, such as a request counter or a "stop" flag. For anything involving
more than one variable, use a mutex.

## Core API

| Function / Type | Purpose |
|---|---|
| `atomic.Int32`, `Int64`, `Uint32`, `Uint64` | Typed counters: `Add`, `Load`, `Store`, `Swap`, `CompareAndSwap` |
| `atomic.Bool` | Atomic flag |
| `atomic.Pointer[T]` | Atomically swap a pointer to an immutable value |
| `atomic.Value` | Store any value of one consistent type |
| `atomic.AddInt64(&x, 1)` and friends | Older function style on plain variables |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./03-concurrency/atomic/
```

## Gotchas
- **Prefer the typed values** (`atomic.Int64`) over the function style; they cannot be accidentally read non-atomically.
- **Never mix atomic and plain access** to the same variable. Every read and write must be atomic.
- **Atomic operations protect one value only.** Two atomics updated back to back are not a transaction.
- **Store *immutable* data behind `atomic.Pointer`.** Build a new struct and swap the pointer; never mutate the old one.
- **Use a `CompareAndSwap` loop for read-modify-write** of anything more complex than an add.
- **`atomic.Value` panics** if you store values of different types, or `nil`.
- **64-bit alignment:** on 32-bit platforms, 64-bit atomic fields must be 8-byte aligned. The typed values handle this for you.

## Java comparison

| Java | Go |
|---|---|
| `AtomicInteger` / `AtomicLong` | `atomic.Int32` / `atomic.Int64` |
| `AtomicBoolean` | `atomic.Bool` |
| `AtomicReference<T>` | `atomic.Pointer[T]` |
| `volatile` | No equivalent: use `atomic` or a mutex |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/sync/atomic
- Source: https://cs.opensource.google/go/go/+/master:src/sync/atomic/
