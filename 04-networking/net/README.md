# net

## Overview
`net` is the portable interface to network I/O: TCP and UDP sockets, Unix
sockets, DNS lookups, IP address and CIDR handling. `net/http` is built on top of
it. A `net.Conn` is just an `io.Reader` + `io.Writer` + deadlines, so everything
from `io` and `bufio` works on connections.

## Core API

| Function / Type | Purpose |
|---|---|
| `net.Listen`, `ln.Accept` | Start a TCP server and accept connections |
| `net.Dial`, `net.DialTimeout`, `net.Dialer` | Connect as a client (`Dialer` supports context) |
| `net.Conn` | `Read`, `Write`, `Close`, `SetDeadline`, `LocalAddr`, `RemoteAddr` |
| `net.Pipe` | In-memory synchronous `Conn` pair for tests |
| `net.ListenPacket`, `net.DialUDP` | UDP |
| `net.ParseIP`, `net.ParseCIDR`, `IPNet.Contains` | IP and subnet handling |
| `net.SplitHostPort`, `net.JoinHostPort` | Correct `host:port` handling (IPv6-safe) |
| `net.LookupHost`, `net.Resolver` | DNS |
| `net.Error` (`Timeout()`), `*net.OpError` | Error inspection |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./04-networking/net/
```

## Gotchas
- **A TCP connection is a byte stream, not messages.** One `Write` may arrive as several `Read`s and vice versa. Frame your protocol (newlines, length prefixes) and use `bufio` or `io.ReadFull`.
- **Always set deadlines** (`conn.SetDeadline`) on connections to untrusted peers, or a slow client holds a goroutine forever.
- **Each accepted connection needs its own goroutine and its own `Close`.**
- **Use `net.JoinHostPort`, not string concatenation**; IPv6 addresses need brackets (`[::1]:80`).
- **Listen on `127.0.0.1:0` in tests** to get a free port, then read it from `ln.Addr()`.
- **`ln.Accept` returns `net.ErrClosed`** after `ln.Close()`. That is how a server loop shuts down cleanly.
- **Check timeouts with `errors.Is(err, os.ErrDeadlineExceeded)`** or `net.Error.Timeout()`.
- **DNS results depend on the machine,** so do not assert on them in tests.
- **`ParseIP` returns nil on invalid input**, never an error. Check for `nil`.

## Java comparison

| Java | Go |
|---|---|
| `ServerSocket` / `Socket` | `net.Listener` / `net.Conn` |
| `InetAddress.getByName` | `net.ParseIP` / `net.LookupHost` |
| `socket.setSoTimeout` | `conn.SetReadDeadline` |
| `DatagramSocket` | `net.ListenPacket("udp", ...)` |
| Netty / NIO | goroutine per connection on `net.Conn` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/net
- Source: https://cs.opensource.google/go/go/+/master:src/net/
