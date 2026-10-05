# net/http/httptest exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./04-networking/httptest/exercises/...
```

## Exercise: GetJSON

Write:

```go
func GetJSON(ctx context.Context, client *http.Client, url string, v any) error
```

It performs a `GET`, and decodes the JSON response into `v`.

Rules:

- Use the given `client` and `ctx` (`http.NewRequestWithContext`)
- Any status other than 200 is an error that mentions the status code
- Invalid JSON is an error that wraps the decode error with `%w`
- Transport errors (such as a failing `RoundTripper`) are returned and
  `errors.Is` still works on them
- Always close the body; read at most 1 MiB

Test it with `httptest.NewServer` for HTTP behavior and with a fake
`RoundTripper` that returns an error for the transport case.

## Bonus
Add a retry: on a 503 or a timeout, retry up to 3 times with a short backoff,
respecting `ctx`.
