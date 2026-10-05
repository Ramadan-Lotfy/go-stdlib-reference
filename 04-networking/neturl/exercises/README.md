# net/url exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./04-networking/neturl/exercises/...
```

## Exercise 1: AddQuery

Write `AddQuery(rawURL string, params map[string]string) (string, error)` that
adds (or replaces) query parameters while keeping existing ones and the
fragment.

| Input | Params | Result |
|-------|--------|--------|
| `https://x.com/p?a=1` | `b: "2 3"` | `https://x.com/p?a=1&b=2+3` |
| `https://x.com/p?a=1` | `a: "9"` | `https://x.com/p?a=9` |
| `https://x.com/p#top` | `q: "go"` | `https://x.com/p?q=go#top` |
| `x.com/p` | any | error (no scheme or host) |

Hints: `url.Parse`, `u.Query()`, `Set`, `Encode`, `u.String()`.

## Exercise 2: Domain

Write `Domain(rawURL string) (string, error)` that returns the host name without
port or credentials, and an error when there is none.

Hints: `u.Hostname()`.

## Bonus
Write `SafeRedirect(target, allowedHost string) (string, bool)` that accepts only
relative paths or URLs on `allowedHost`. Why is `//evil.com` dangerous?
