package pprofref_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	_ "net/http/pprof" // registers /debug/pprof/* on http.DefaultServeMux
	"runtime/pprof"
	"slices"
	"strings"
	"time"
)

func busy(d time.Duration) int {
	deadline := time.Now().Add(d)
	n := 0
	for time.Now().Before(deadline) {
		n++
	}
	return n
}

// Example_cpuProfile records a CPU profile into memory. In real code you write
// it to a file and open it with `go tool pprof`.
func Example_cpuProfile() {
	var buf bytes.Buffer
	if err := pprof.StartCPUProfile(&buf); err != nil {
		fmt.Println("cannot profile:", err)
		return
	}
	busy(50 * time.Millisecond)
	pprof.StopCPUProfile() // the profile is complete only after Stop

	fmt.Println(buf.Len() > 0)
	// Output: true
}

// Example_heapProfile writes the heap profile as text (debug=1).
func Example_heapProfile() {
	var buf bytes.Buffer
	if err := pprof.Lookup("heap").WriteTo(&buf, 1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.HasPrefix(buf.String(), "heap profile:"))
	// Output: true
}

// Example_profiles lists the built-in profiles.
func Example_profiles() {
	var names []string
	for _, p := range pprof.Profiles() {
		names = append(names, p.Name())
	}
	slices.Sort(names)
	fmt.Println(names)
	// Output: [allocs block goroutine heap mutex threadcreate]
}

// Example_goroutineProfile is the quickest way to see what every goroutine is doing.
func Example_goroutineProfile() {
	var buf bytes.Buffer
	pprof.Lookup("goroutine").WriteTo(&buf, 1)
	fmt.Println(strings.HasPrefix(buf.String(), "goroutine profile: total"))
	// Output: true
}

// Example_labels tags work so profiles can be filtered per handler or tenant.
func Example_labels() {
	pprof.Do(context.Background(), pprof.Labels("handler", "search"), func(ctx context.Context) {
		v, ok := pprof.Label(ctx, "handler")
		fmt.Println(v, ok)
	})
	// Output: search true
}

// Example_customProfile tracks live resources (here, pretend connections).
func Example_customProfile() {
	const name = "example.connections"
	p := pprof.Lookup(name)
	if p == nil {
		p = pprof.NewProfile(name)
	}
	conn := new(int)
	p.Add(conn, 0) // 0: skip no stack frames
	fmt.Println(p.Count())
	p.Remove(conn)
	fmt.Println(p.Count())
	// Output:
	// 1
	// 0
}

// Example_httpEndpoints shows what the blank import of net/http/pprof exposes.
// Never serve this on a public listener.
func Example_httpEndpoints() {
	srv := httptest.NewServer(http.DefaultServeMux)
	defer srv.Close()

	for _, path := range []string{"/debug/pprof/", "/debug/pprof/goroutine?debug=1"} {
		resp, err := http.Get(srv.URL + path)
		if err != nil {
			fmt.Println(err)
			return
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(path, resp.StatusCode, len(body) > 0)
	}
	// Output:
	// /debug/pprof/ 200 true
	// /debug/pprof/goroutine?debug=1 200 true
}
