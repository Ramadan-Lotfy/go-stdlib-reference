# net/http exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./04-networking/nethttp/exercises/...
```

## Exercise: a small API

Write `NewAPI() http.Handler` with these routes, using the Go 1.22 `ServeMux`:

| Request | Result |
|---------|--------|
| `GET /health` | `200`, body `ok` (no key needed) |
| `GET /users/{id}` without header `X-API-Key: secret` | `401 unauthorized` |
| `GET /users/abc` (with key) | `400 invalid id` |
| `GET /users/500` (with key) | `404 user not found` (valid ids are 1 to 100) |
| `GET /users/5` (with key) | `200`, JSON `{"id":5,"name":"user5"}` |
| `POST /users/5` | `405` (automatic from the mux) |

Also write the middleware `RequireAPIKey(key string, next http.Handler) http.Handler`.

Hints: `mux.HandleFunc("GET /health", ...)`, `r.PathValue("id")`, `strconv.Atoi`,
`http.Error`, and test everything with `httptest.NewRecorder`.

## Bonus
Add a logging middleware using `log/slog` that records method, path, status and
duration. You will need a `ResponseWriter` wrapper to capture the status code.
