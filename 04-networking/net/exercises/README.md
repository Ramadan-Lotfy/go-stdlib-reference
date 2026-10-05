# net exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./04-networking/net/exercises/...
```

## Exercise 1: Serve

Write `Serve(ln net.Listener) error`, an upper-casing line server:

- Accept connections in a loop, handling each in its own goroutine
- For every line a client sends, reply with that line upper-cased
- When `ln` is closed, `Serve` returns `nil` (not an error)
- Close each connection when its client disconnects

Hints: `ln.Accept`, `errors.Is(err, net.ErrClosed)`, `bufio.Scanner` on the
connection, `fmt.Fprintln(conn, ...)`. Test with `net.Listen("tcp", "127.0.0.1:0")`.

## Exercise 2: IsPrivateIP

Write `IsPrivateIP(s string) (bool, error)` that returns true for private
(`10.x`, `172.16-31.x`, `192.168.x`), loopback and link-local addresses, false
for public ones, and an error for text that is not an IP.

Hints: `net.ParseIP` returns `nil` on bad input; `ip.IsPrivate`, `IsLoopback`,
`IsLinkLocalUnicast`.

## Bonus
Add an idle timeout: close a connection if the client sends nothing for 30
seconds, by resetting `conn.SetReadDeadline` before every read.
