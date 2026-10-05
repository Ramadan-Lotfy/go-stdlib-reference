# Stage 5: Mastery

Reflection, generics, the runtime, profiling, crypto and production concerns.

| Package | What it covers |
|---------|----------------|
| [reflect](reflect) | `reflect` lets a program inspect and modify values whose types are not known at compile time: list struct fields and tags, call methods by name, build slices and maps dynamically. |
| [unsafe](unsafe) | `unsafe` steps outside Go's type and memory safety rules: sizes and offsets, pointer arithmetic, and zero-copy conversions. |
| [generics, cmp and slices](generics) | Generics (Go 1.18+) let functions and types work over many types while keeping compile-time type safety. |
| [runtime and runtime/debug](runtime) | `runtime` is the interface to Go's run-time system: the scheduler, garbage collector and memory allocator. |
| [runtime/pprof and net/http/pprof](pprof) | `runtime/pprof` records profiles of a running program: where CPU time goes, which code allocates memory, which goroutines exist, and where they block. |
| [runtime/trace](trace) | `runtime/trace` records an **execution trace**: a timeline of everything the runtime does (goroutine creation, blocking, scheduling, garbage collection, system calls) with nanosecond timestamps. |
| [crypto/*](crypto) | The `crypto` packages provide hashing, message authentication, random numbers, encryption and digital signatures. |
| [embed](embed) | `embed` compiles files into the program binary, so a single executable can carry its own templates, SQL migrations, static web assets and default config. |
| [flag](flag) | `flag` parses command-line flags such as `-addr :8080 -v`. |
| [os/exec](osexec) | `os/exec` runs external programs. |
| [os/signal](signal) | `os/signal` lets a program react to operating system signals: `SIGINT` (Ctrl+C), `SIGTERM` (the polite "please stop" sent by Kubernetes, systemd and `docker stop`), `SIGHUP` (often "reload config") and others. |

Each folder has a `README.md` (concepts, gotchas, Java comparison), runnable `example_test.go`, and `exercises/` with a reference solution.

Run this stage: `go test ./05-mastery/...`
