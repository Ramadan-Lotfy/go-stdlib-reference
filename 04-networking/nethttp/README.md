# net/http

## Overview
`net/http` is a complete HTTP client and server. A server is a `Handler`
(`ServeHTTP(w, r)`) plus a router (`ServeMux`); since Go 1.22 the router supports
methods and path parameters such as `GET /items/{id}`. Many production Go
services need nothing more than this package. Compare with Spring MVC: handlers
are your controllers, middleware are plain functions that wrap handlers.

## Core API

| Function / Type | Purpose |
|---|---|
| `http.Handler`, `http.HandlerFunc` | The server-side interface: `ServeHTTP(ResponseWriter, *Request)` |
| `http.NewServeMux`, `mux.HandleFunc("GET /items/{id}", h)` | Router with methods, wildcards (`{id}`, `{path...}`, `{$}`) |
| `r.PathValue("id")`, `r.URL.Query()`, `r.FormValue`, `r.Header` | Read the request |
| `w.Header().Set`, `w.WriteHeader`, `w.Write`, `http.Error` | Write the response (in that order) |
| `http.Server{...}`, `ListenAndServe`, `Shutdown` | Configurable server with graceful shutdown |
| `http.Client`, `NewRequestWithContext`, `client.Do` | Client with timeouts and context |
| `http.Redirect`, `http.SetCookie`, `http.FileServer` | Common helpers |
| `http.StatusOK`, `http.MethodGet`, `http.StatusText` | Constants |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./04-networking/nethttp/
```

## Gotchas
- **`http.Get`, `http.DefaultClient` and `http.ListenAndServe` have no timeouts.** In production use your own `http.Client{Timeout: ...}` and `http.Server{ReadHeaderTimeout: ...}`.
- **Always close the response body** (`defer resp.Body.Close()`), and drain it if you want the connection reused. Check `resp.StatusCode`; a 404 is not an error from `Do`.
- **Write the status once, before the body.** Headers set after `WriteHeader` or the first `Write` are ignored. Calling `http.Error` and then continuing to write garbles the response, so `return` after it.
- **Handlers run one goroutine per request.** Shared state needs a mutex or atomics; run tests with `-race`.
- **Limit request bodies** with `http.MaxBytesReader(w, r.Body, n)` before decoding.
- **Method patterns are exact:** `GET /x` also matches `HEAD`; unmatched methods get 405 automatically. A pattern without a method matches all.
- **Patterns in 1.22 are stricter than before:** conflicting patterns panic at registration. `/` matches everything; use `/{$}` for only the root.
- **Middleware is `func(http.Handler) http.Handler`**; order matters (outermost runs first).
- **Do not pass a `*http.Request` to a goroutine that outlives the handler**; the body is closed when the handler returns.
- **Use `Server.Shutdown(ctx)` for graceful stop**; `Serve` then returns `http.ErrServerClosed`, which is not a failure.

## Java comparison

| Java | Go |
|---|---|
| `@RestController` method | `http.Handler` / `HandlerFunc` |
| `@GetMapping("/items/{id}")` | `mux.HandleFunc("GET /items/{id}", h)` |
| `@PathVariable` | `r.PathValue("id")` |
| `@RequestParam` | `r.URL.Query().Get("q")` |
| `@RequestBody` | `json.NewDecoder(r.Body).Decode(&v)` |
| `ResponseEntity.status(201)` | `w.WriteHeader(201)` |
| `Filter` / `HandlerInterceptor` | middleware `func(http.Handler) http.Handler` |
| `RestTemplate` / `WebClient` | `http.Client` |
| embedded Tomcat | `http.Server` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/net/http
- Source: https://cs.opensource.google/go/go/+/master:src/net/http/
