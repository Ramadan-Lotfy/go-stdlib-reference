# runtime/pprof and net/http/pprof exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/pprof/exercises/...
```

## Exercise: capture profiles

Write two functions:

```go
func ProfileCPU(f func()) ([]byte, error)
func HeapProfile() ([]byte, error)
```

- `ProfileCPU` starts a CPU profile into a buffer, runs `f`, stops the profile
  and returns the bytes. If a profile is already running it returns the error
  from `StartCPUProfile`
- `HeapProfile` returns the heap profile in the binary format (`debug=0`)
- Profiles are gzip-compressed, so the tests check the magic bytes `0x1f 0x8b`

Hints: `pprof.StartCPUProfile`, `pprof.StopCPUProfile` (call it **before** reading
the buffer), `pprof.Lookup("heap").WriteTo(&buf, 0)`.

Then try it for real: save the bytes to a file and run `go tool pprof -top file`.

## Bonus
Write a small program that serves `net/http/pprof` on `localhost:6060` using its
own `http.ServeMux` (not `DefaultServeMux`), generates load, and lets you grab a
30-second CPU profile with `go tool pprof http://localhost:6060/debug/pprof/profile`.
