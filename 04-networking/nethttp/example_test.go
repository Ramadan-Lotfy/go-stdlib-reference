package nethttpref_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing/fstest"
	"time"
)

// Example_handler is the smallest handler, exercised without a network.
func Example_handler() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %s", r.URL.Query().Get("name"))
	})
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/?name=ada", nil))
	fmt.Println(rec.Code, rec.Body.String())
	// Output: 200 hello ada
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "item %s", r.PathValue("id"))
	})
	mux.HandleFunc("GET /files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "file %s", r.PathValue("path"))
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "home")
	})
	return mux
}

// Example_mux shows Go 1.22 routing: methods, path parameters, wildcards.
func Example_mux() {
	mux := newMux()
	do := func(method, target string) {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest(method, target, nil))
		fmt.Println(method, target, rec.Code, strings.TrimSpace(rec.Body.String()))
	}
	do("GET", "/items/42")
	do("GET", "/files/a/b.txt")
	do("GET", "/")
	do("POST", "/items/42") // right path, wrong method
	do("GET", "/nope")
	// Output:
	// GET /items/42 200 item 42
	// GET /files/a/b.txt 200 file a/b.txt
	// GET / 200 home
	// POST /items/42 405 Method Not Allowed
	// GET /nope 404 404 page not found
}

func stamp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Seen", "true")
		next.ServeHTTP(w, r)
	})
}

func requireToken(token string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer "+token {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return // never continue after writing an error
		}
		next.ServeHTTP(w, r)
	})
}

// Example_middleware wraps a handler; the outermost wrapper runs first.
func Example_middleware() {
	ok := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "secret data") })
	h := stamp(requireToken("s3", ok))

	for _, token := range []string{"", "s3"} {
		req := httptest.NewRequest("GET", "/", nil)
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		fmt.Println(rec.Code, rec.Header().Get("X-Seen"), strings.TrimSpace(rec.Body.String()))
	}
	// Output:
	// 401 true unauthorized
	// 200 true secret data
}

// Example_jsonAPI decodes a JSON request and writes a JSON response.
func Example_jsonAPI() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var in struct {
			Name string `json:"name"`
		}
		body := http.MaxBytesReader(w, r.Body, 1<<20)
		if err := json.NewDecoder(body).Decode(&in); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")                  // headers first...
		w.WriteHeader(http.StatusCreated)                                   // ...then status...
		json.NewEncoder(w).Encode(map[string]any{"id": 1, "name": in.Name}) // ...then body
	})

	for _, body := range []string{`{"name":"ada"}`, `not json`} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("POST", "/users", strings.NewReader(body)))
		fmt.Println(rec.Code, rec.Header().Get("Content-Type"), strings.TrimSpace(rec.Body.String()))
	}
	// Output:
	// 201 application/json {"id":1,"name":"ada"}
	// 400 text/plain; charset=utf-8 bad json
}

// Example_client talks to a real (local, test) server with a configured client.
func Example_client() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %s", r.Header.Get("X-Name"))
	}))
	defer srv.Close()

	client := &http.Client{Timeout: 2 * time.Second} // never rely on DefaultClient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL, nil)
	req.Header.Set("X-Name", "ada")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close() // always close the body

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode, string(body))
	// Output: 200 hello ada
}

// Example_shutdown starts a server on a free port and stops it gracefully.
func Example_shutdown() {
	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	srv := &http.Server{
		Handler:           newMux(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	done := make(chan error, 1)
	go func() { done <- srv.Serve(ln) }()

	resp, err := http.Get("http://" + ln.Addr().String() + "/items/9")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println(srv.Shutdown(ctx))
	fmt.Println(<-done == http.ErrServerClosed)
	// Output:
	// item 9
	// <nil>
	// true
}

// Example_redirectAndCookies covers two common response helpers.
func Example_redirectAndCookies() {
	rec := httptest.NewRecorder()
	http.Redirect(rec, httptest.NewRequest("GET", "/old", nil), "/new", http.StatusFound)
	fmt.Println(rec.Code, rec.Header().Get("Location"))

	rec = httptest.NewRecorder()
	http.SetCookie(rec, &http.Cookie{Name: "sid", Value: "abc", HttpOnly: true, Path: "/"})
	fmt.Println(rec.Header().Get("Set-Cookie"))
	c := rec.Result().Cookies()[0]
	fmt.Println(c.Name, c.Value, c.HttpOnly)
	// Output:
	// 302 /new
	// sid=abc; Path=/; HttpOnly
	// sid abc true
}

// Example_formAndQuery reads query string and form values.
func Example_formAndQuery() {
	req := httptest.NewRequest("POST", "/?q=go", strings.NewReader("name=ada&name=bob"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ParseForm()

	fmt.Println(req.URL.Query().Get("q"))
	fmt.Println(req.PostForm["name"], req.PostForm.Get("name"))
	fmt.Println(req.FormValue("q")) // FormValue also searches the query string
	// Output:
	// go
	// [ada bob] ada
	// go
}

// Example_fileServer serves any fs.FS, here an in-memory one.
func Example_fileServer() {
	h := http.FileServer(http.FS(fstest.MapFS{"a.txt": {Data: []byte("A")}}))
	for _, p := range []string{"/a.txt", "/missing"} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("GET", p, nil))
		fmt.Println(p, rec.Code)
	}
	// Output:
	// /a.txt 200
	// /missing 404
}

// Example_constants shows the helpers that replace magic numbers and strings.
func Example_constants() {
	fmt.Println(http.StatusText(http.StatusNotFound), http.MethodPost)
	fmt.Println(http.CanonicalHeaderKey("content-type"))
	// Output:
	// Not Found POST
	// Content-Type
}
