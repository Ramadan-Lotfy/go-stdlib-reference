# net/url

## Overview
`net/url` parses, builds and escapes URLs and query strings. Use it instead
of string concatenation whenever you build a URL from user input.

## Core API

| Function / Type | Purpose |
|---|---|
| `url.Parse`, `url.ParseRequestURI` | Parse a URL into a `*url.URL` |
| `u.Scheme`, `Host`, `Hostname()`, `Port()`, `Path`, `RawQuery`, `Fragment`, `User` | The parts |
| `u.Query()`, `url.Values` | Parse and build query strings (`Get`, `Set`, `Add`, `Encode`) |
| `url.QueryEscape`, `url.PathEscape` and the `Unescape` versions | Escape a single component |
| `u.String()`, `u.ResolveReference`, `url.JoinPath` | Build and combine URLs |
| `u.Redacted()` | String with the password hidden (for logs) |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./04-networking/neturl/
```

## Gotchas
- **`url.Parse` accepts almost anything.** `"example.com/path"` parses without error but has no scheme or host. Check `u.Scheme` and `u.Host`, or use `ParseRequestURI`.
- **Escape per component:** `QueryEscape` for query values (space becomes `+`), `PathEscape` for path segments (space becomes `%20`). They are not interchangeable.
- **Build query strings with `url.Values`,** never by concatenation. `Encode` sorts keys and escapes everything.
- **`u.Query()` re-parses on every call** and returns a copy; edit it, then assign `u.RawQuery = q.Encode()`.
- **`u.Path` is decoded; `u.RawPath` / `u.EscapedPath()` keep the original.** Use `EscapedPath` when you need the encoded form.
- **Never log a URL with credentials;** use `u.Redacted()`.
- **Validate redirect targets** (open-redirect attacks): parse them and check the host against an allow-list.
- **`Hostname()` strips the port and the IPv6 brackets;** `Host` keeps both.

## Java comparison

| Java | Go |
|---|---|
| `java.net.URI` / `URL` | `*url.URL` |
| `UriComponentsBuilder` | `url.URL` + `url.Values` |
| `URLEncoder.encode(s, UTF_8)` | `url.QueryEscape(s)` |
| `URI.resolve` | `base.ResolveReference(ref)` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/net/url
- Source: https://cs.opensource.google/go/go/+/master:src/net/url/
