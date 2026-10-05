# unsafe

## Overview
`unsafe` steps outside Go's type and memory safety rules: sizes and offsets,
pointer arithmetic, and zero-copy conversions. It is used inside the standard
library and in performance-critical or system-level code, and almost never in
application code. Learn it to read other people's code and to understand memory
layout, not to use daily.

**The examples assume a 64-bit platform** (amd64 or arm64), where `int` and
pointers are 8 bytes.

## Core API

| Function / Type | Purpose |
|---|---|
| `unsafe.Sizeof(x)`, `Alignof(x)`, `Offsetof(s.f)` | Memory layout (compile-time constants) |
| `unsafe.Pointer` | A pointer of any type; converts to and from `uintptr` |
| `unsafe.Add(p, n)` | Pointer plus `n` bytes |
| `unsafe.Slice(ptr, n)`, `unsafe.SliceData(s)` | Slice from a pointer / pointer to the slice's first element |
| `unsafe.String(ptr, n)`, `unsafe.StringData(s)` | Zero-copy string conversions (Go 1.20+) |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/unsafe/
```

## Gotchas
- **Go gives no compatibility guarantee for `unsafe` code** other than the six documented pointer patterns. It can break with a new Go version.
- **A `uintptr` is just a number; the garbage collector does not follow it.** Never store a pointer in a `uintptr` across statements.
- **Zero-copy `string` from `[]byte` shares memory.** Changing the bytes later changes the "immutable" string, which breaks map keys and data races. Only do it when the bytes are never modified.
- **Struct field order changes size** because of alignment padding. Order fields from largest to smallest to minimize size (or use `fieldalignment` from `golang.org/x/tools`).
- **Do not use `unsafe` to "speed up" code without a benchmark** that proves it matters. The safe alternatives (`math.Float32bits`, `strings.Builder`, `encoding/binary`) are usually as fast.
- **`go vet` checks some misuse** (`unsafeptr`). Run the race detector with `-d=checkptr` (enabled by `-race`) to catch invalid pointer conversions.
- **cgo and `syscall` need it;** that is where you will meet it first.

## Java comparison

| Java | Go |
|---|---|
| `sun.misc.Unsafe` / `VarHandle` | `unsafe.Pointer` |
| `ByteBuffer` views over memory | `unsafe.Slice(ptr, n)` |
| `Unsafe.objectFieldOffset` | `unsafe.Offsetof` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/unsafe
- Source: https://cs.opensource.google/go/go/+/master:src/unsafe/
