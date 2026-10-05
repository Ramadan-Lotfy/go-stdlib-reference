# io exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/io/exercises/...
```

## Exercise 1: CountingReader

Write a `CountingReader` struct that wraps an `io.Reader` and exposes the total
number of bytes read through it in a field `N int64`.

```go
cr := &CountingReader{R: strings.NewReader("hello world")}
io.Copy(io.Discard, cr)
fmt.Println(cr.N) // 11
```

Hints: implement `Read(p []byte) (int, error)` on a pointer receiver; always add
`n` before returning, even when `err != nil`.

## Exercise 2: Head

Write `Head(r io.Reader, n int64) (string, error)` returning at most the first
`n` bytes of `r`. If the reader is shorter, return what is there.

Hints: `io.LimitReader`, `io.ReadAll`.

## Bonus
Write an `io.Writer` wrapper `UpperWriter` that upper-cases everything written
to the underlying writer.
