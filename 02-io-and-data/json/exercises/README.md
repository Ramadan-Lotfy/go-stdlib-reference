# encoding/json exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/json/exercises/...
```

## Exercise 1: ParseConfig

Write `ParseConfig(data []byte) (Config, error)` for this struct:

```go
type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Debug bool   `json:"debug"`
}
```

Rules:

- Missing fields get defaults: `Host = "localhost"`, `Port = 8080`
- Unknown fields are an error
- Trailing data after the first JSON value is an error
- `Port` must be between 1 and 65535
- Wrap errors as `parse config: ...`

Hints: pre-fill the struct with defaults before decoding, `json.NewDecoder`,
`DisallowUnknownFields`, and a second `Decode` that must return `io.EOF`.

## Exercise 2: PrettyJSON

Write `PrettyJSON(data []byte) (string, error)` that re-indents any valid JSON
with two spaces and returns an error for invalid JSON.

Hints: `json.Indent` into a `bytes.Buffer`.

## Bonus
Add a `Duration` type whose JSON form is a string like `"1m30s"`, using
`time.ParseDuration` and the `Marshaler` / `Unmarshaler` interfaces.
