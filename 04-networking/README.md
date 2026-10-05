# Stage 4: Networking and backend

HTTP, sockets, URLs, TLS, SQL and testing.

| Package | What it covers |
|---------|----------------|
| [net/http](nethttp) | `net/http` is a complete HTTP client and server. |
| [net/http/httptest](httptest) | `net/http/httptest` makes HTTP code testable without real networks or ports. |
| [net](net) | `net` is the portable interface to network I/O: TCP and UDP sockets, Unix sockets, DNS lookups, IP address and CIDR handling. |
| [net/url](neturl) | `net/url` parses, builds and escapes URLs and query strings. |
| [crypto/tls](tls) | `crypto/tls` implements TLS, the protocol behind HTTPS. |
| [database/sql](databasesql) | `database/sql` is Go's database access layer. |
| [testing](testing) | The `testing` package and the `go test` command are Go's built-in test framework: no external library is needed for unit tests, benchmarks, fuzzing, examples or coverage. |

Each folder has a `README.md` (concepts, gotchas, Java comparison), runnable `example_test.go`, and `exercises/` with a reference solution.

Run this stage: `go test ./04-networking/...`
