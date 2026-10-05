# channels, select and concurrency patterns exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./03-concurrency/channels/exercises/...
```

## Exercise: a cancellable pipeline

Implement three functions in `solution.go`:

```go
func Generate(ctx context.Context, nums ...int) <-chan int
func Square(ctx context.Context, in <-chan int) <-chan int
func Merge[T any](cs ...<-chan T) <-chan T
```

Rules:

- `Generate` sends each number and then closes its channel
- `Square` reads from `in`, sends the square of each value, and closes its output
  when `in` is closed
- `Merge` forwards values from all inputs and closes its output only after
  **every** input is closed
- `Generate` and `Square` must stop and close their outputs when `ctx` is
  cancelled, even if nobody is reading (no goroutine leaks): every send is a
  `select` with `ctx.Done()`

Run with the race detector: `go test -race ./03-concurrency/channels/exercises/...`.

Hints: `defer close(out)` as the first line of each goroutine, and a
`sync.WaitGroup` inside `Merge`.

## Bonus
Write `FanOut(ctx, in, workers, f)` that runs `f` on N goroutines reading the
same input channel and merges the results.
