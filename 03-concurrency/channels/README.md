# channels, select and concurrency patterns

## Overview
Channels are typed pipes between goroutines, and `select` waits on several
of them at once. Together with goroutines they are Go's core concurrency tools.
This package covers the language features (`chan`, `select`, `close`, `range`)
and the patterns built from them: worker pools, pipelines, fan-in and
semaphores.

## Core API

| Function / Type | Purpose |
|---|---|
| `make(chan T)`, `make(chan T, n)` | Unbuffered (rendezvous) and buffered channels |
| `ch <- v`, `v := <-ch`, `v, ok := <-ch` | Send, receive, receive with closed check |
| `close(ch)`, `for v := range ch` | Signal "no more values"; loop until closed |
| `select { case ... default: }` | Wait on several channel operations |
| `chan<- T`, `<-chan T` | Send-only and receive-only types for APIs |
| worker pool, pipeline, fan-in/fan-out | The standard patterns |
| `nil` channel | Blocks forever: use it to disable a `select` case |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./03-concurrency/channels/
```

## Gotchas
- **Only the sender closes a channel**, and only once. Receivers cannot know if senders remain. Closing twice or sending on a closed channel panics.
- **A receive from a closed channel returns the zero value immediately.** Use `v, ok := <-ch` to tell the difference.
- **Unbuffered sends block until a receiver is ready.** Forgetting the receiver is a deadlock; the runtime panics with "all goroutines are asleep" only if *every* goroutine is blocked.
- **Goroutine leaks:** a goroutine blocked forever on a channel is never freed. Give every goroutine a way to exit, usually a `context` or a closed channel.
- **`select` picks randomly among ready cases.** Do not depend on ordering.
- **`len(ch)` is a snapshot.** Never use it to make decisions in concurrent code.
- **Buffered channels hide design problems.** Size them for a reason (a known burst, a semaphore), not to "fix" a deadlock.
- **Range loop variables are per-iteration since Go 1.22**, so `go func() { use(v) }()` in a `for` loop is safe. Before 1.22 it was a classic bug.
- **Use a mutex instead** when you are just protecting shared state. Channels are for passing ownership or coordinating.

## Java comparison

| Java | Go |
|---|---|
| `BlockingQueue` | buffered channel |
| `ExecutorService` fixed pool | N goroutines reading one jobs channel |
| `Semaphore` | buffered `chan struct{}` |
| `CompletableFuture.anyOf` | `select` or a result channel |
| Reactive streams | pipeline of channels |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/channels, select and concurrency patterns
- Source: https://cs.opensource.google/go/go/+/master:src/channels, select and concurrency patterns/
