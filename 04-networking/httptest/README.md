# net/http/httptest

## Overview
`net/http/httptest` makes HTTP code testable without real networks or ports.
A `ResponseRecorder` stands in for the response writer so you can call a handler
directly; `NewServer` starts a real server on a local port for testing clients and
full stacks; and a custom `RoundTripper` lets you fake the remote side of an
`http.Client`.

## Core API

| Function / Type | Purpose |
|---|---|
| `httptest.NewRecorder()` | Fake `ResponseWriter`: `Code`, `Header()`, `Body`, `Result()` |
| `httptest.NewRequest(method, target, body)` | Build an incoming request (panics on bad input) |
| `httptest.NewServer(handler)` | Real local HTTP server; `srv.URL`, `srv.Client()`, `srv.Close()` |
| `httptest.NewTLSServer` | Same with HTTPS and a throwaway certificate |
| `httptest.NewUnstartedServer` | Configure `srv.Config` before `Start()` |
| custom `http.RoundTripper` | Fake transport that never touches the network |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./04-networking/httptest/
```

## Gotchas
- **`NewRequest` is for servers; `http.NewRequest` is for clients.** `httptest.NewRequest` sets `Host` and `RemoteAddr` for you and panics on invalid input.
- **`rec.Code` defaults to 200** even if the handler never wrote anything.
- **`rec.Body` is a `*bytes.Buffer`.** Read it with `.String()`; use `rec.Result()` to get a full `*http.Response`.
- **Always `defer srv.Close()`** or the test leaks a listener and goroutines.
- **Use `srv.Client()` with `NewTLSServer`;** it already trusts the test certificate.
- **Prefer testing handlers directly with a recorder** (fast, no ports) and use `NewServer` only when you need real HTTP behavior such as redirects, cookies or timeouts.
- **Do not call `t.Fatal` from a handler goroutine;** record the error and check it in the test.
- **A fake `RoundTripper` is the best way to test client code** that calls third-party APIs.

## Java comparison

| Java | Go |
|---|---|
| Spring `MockMvc` | `httptest.NewRecorder` + `ServeHTTP` |
| `@SpringBootTest(webEnvironment = RANDOM_PORT)` | `httptest.NewServer` |
| WireMock / MockWebServer | `httptest.NewServer` or a fake `RoundTripper` |
| `TestRestTemplate` | `srv.Client()` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/net/http/httptest
- Source: https://cs.opensource.google/go/go/+/master:src/net/http/httptest/
