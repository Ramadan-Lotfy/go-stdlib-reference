package httptestref_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Version", "1")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "queued")
}

// Example_recorder calls a handler directly and inspects what it wrote.
func Example_recorder() {
	rec := httptest.NewRecorder()
	hello(rec, httptest.NewRequest("POST", "/jobs", nil))

	fmt.Println(rec.Code, rec.Header().Get("X-Version"), rec.Body.String())

	res := rec.Result() // a real *http.Response
	body, _ := io.ReadAll(res.Body)
	fmt.Println(res.StatusCode, res.Status, string(body))
	// Output:
	// 202 1 queued
	// 202 202 Accepted queued
}

// Example_newRequest shows the request fields httptest fills in for you.
func Example_newRequest() {
	req := httptest.NewRequest("PUT", "/a?b=1", strings.NewReader("data"))
	fmt.Println(req.Method, req.URL.Path, req.URL.RawQuery, req.ContentLength)
	fmt.Println(req.Host, req.RemoteAddr)
	// Output:
	// PUT /a b=1 4
	// example.com 192.0.2.1:1234
}

// Example_defaultCode shows that an untouched recorder reports 200.
func Example_defaultCode() {
	rec := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}).ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	fmt.Println(rec.Code)
	// Output: 200
}

// Example_tableTest runs many cases against one handler.
func Example_tableTest() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "pong") })

	tests := []struct {
		method, path string
		want         int
	}{
		{"GET", "/ping", 200},
		{"POST", "/ping", 405},
		{"GET", "/missing", 404},
	}
	for _, tt := range tests {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest(tt.method, tt.path, nil))
		if rec.Code != tt.want {
			fmt.Printf("%s %s: got %d, want %d\n", tt.method, tt.path, rec.Code, tt.want)
		}
	}
	fmt.Println("all cases passed")
	// Output: all cases passed
}

// Example_server runs a real server on a free local port and counts requests.
func Example_server() {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		fmt.Fprint(w, "hi")
	}))
	defer srv.Close()

	for range 3 {
		resp, err := http.Get(srv.URL)
		if err != nil {
			fmt.Println(err)
			return
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
	fmt.Println(strings.HasPrefix(srv.URL, "http://127.0.0.1:"), hits.Load())
	// Output: true 3
}

// roundTripFunc turns a function into an http.RoundTripper.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// Example_roundTripper fakes the remote side of a client with no server at all.
func Example_roundTripper() {
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(`{"ok":true}`)),
		}, nil
	})}

	resp, err := client.Get("http://anything.invalid/x")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode, string(body))
	// Output: 200 {"ok":true}
}

// Example_unstartedServer lets you tweak the server before it starts.
func Example_unstartedServer() {
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "configured")
	}))
	srv.Config.ReadHeaderTimeout = 0 // customize srv.Config here
	srv.Start()
	defer srv.Close()

	resp, _ := srv.Client().Get(srv.URL)
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body))
	// Output: configured
}
