# encoding/gob

## Overview
`encoding/gob` is Go's own binary serialization format. It is compact, handles
Go types (including maps, slices and interfaces) with no struct tags, and sends
type information once per stream. Because it is Go-specific, use it between Go
programs; for anything that other languages read, use JSON or Protocol Buffers.

## Core API

| Function / Type | Purpose |
|---|---|
| `gob.NewEncoder(w).Encode(v)` | Write a value (the first call also sends type info) |
| `gob.NewDecoder(r).Decode(&v)` | Read the next value into a pointer |
| `gob.Register(v)` | Required for concrete types sent as interface values |
| `gob.GobEncoder`, `gob.GobDecoder` | Custom encoding for a type |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/gob/
```

## Gotchas
- **Zero values are not transmitted.** Decoding into a pre-filled struct leaves old values where the sender had zero values. Always decode into a fresh value.
- **Only exported fields are encoded.** A type with no exported fields is an error.
- **Register concrete types before sending them as interface values**, on both the sender and the receiver.
- **Encoder and decoder keep state.** Use one encoder per stream; create a new pair for each independent message.
- **Empty slices and maps decode as `nil`.**
- **Do not use gob as a long-term storage or wire format.** Changing struct layouts across versions is possible but fragile, and no other language reads it.
- **Never decode untrusted input.** Treat it like any binary format parser.

## Java comparison

| Java | Go |
|---|---|
| `ObjectOutputStream`/`ObjectInputStream` (`Serializable`) | `gob.Encoder` / `gob.Decoder` |
| `serialVersionUID` | None: matching is by field name |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/encoding/gob
- Source: https://cs.opensource.google/go/go/+/master:src/encoding/gob/
