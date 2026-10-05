# encoding/json

## Overview
`encoding/json` converts Go values to JSON and back using reflection and
struct tags. `Marshal` and `Unmarshal` work on byte slices; `Encoder` and
`Decoder` work on streams such as files and HTTP bodies.

## Core API

| Function / Type | Purpose |
|---|---|
| `json.Marshal`, `json.MarshalIndent` | Value to JSON bytes |
| `json.Unmarshal` | JSON bytes into a pointer to a value |
| `json.NewEncoder`, `json.NewDecoder` | Stream to a Writer / from a Reader |
| `` `json:"name,omitempty"` `` | Struct tags: rename, omit empty, skip with `-` |
| `json.RawMessage` | Delay decoding part of a document |
| `json.Marshaler`, `json.Unmarshaler` | Custom JSON for your own types |
| `Decoder.DisallowUnknownFields`, `UseNumber` | Strict decoding, exact numbers |
| `json.Valid`, `json.Indent`, `json.Compact` | Validate and reformat |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/json/
```

## Gotchas
- **Only exported fields are encoded.** A lowercase field name is silently ignored.
- **Unmarshal needs a pointer**: `json.Unmarshal(data, &v)`. Without `&` you get `InvalidUnmarshalError`.
- **Numbers decode to `float64` in `any`.** Large integers lose precision; use `Decoder.UseNumber()` or a typed struct.
- **`nil` slices encode as `null`, empty slices as `[]`.** Initialize with `[]T{}` if clients expect an array.
- **Field matching is case-insensitive on decode**, which can surprise you.
- **Unknown fields are ignored by default.** Use `DisallowUnknownFields` for strict APIs.
- **`omitempty` does not omit zero structs or zero `time.Time`.** Use a pointer, or `omitzero` in Go 1.24+.
- **Marshal escapes `<`, `>` and `&`** as `\u003c` etc. Use an Encoder with `SetEscapeHTML(false)` to turn that off.
- **Check the decode error and then trailing data.** `Decode` reads one value and leaves the rest of the stream.
- **Maps encode with sorted keys**, so output is deterministic.

## Java comparison

| Java | Go |
|---|---|
| Jackson `ObjectMapper.writeValueAsString` | `json.Marshal` |
| `readValue(json, User.class)` | `json.Unmarshal(data, &user)` |
| `@JsonProperty("name")` | `` `json:"name"` `` |
| `@JsonIgnore` | `` `json:"-"` `` |
| `@JsonInclude(NON_EMPTY)` | `omitempty` |
| `JsonNode` | `any` / `map[string]any` / `json.RawMessage` |
| `FAIL_ON_UNKNOWN_PROPERTIES` | `Decoder.DisallowUnknownFields()` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/encoding/json
- Source: https://cs.opensource.google/go/go/+/master:src/encoding/json/
