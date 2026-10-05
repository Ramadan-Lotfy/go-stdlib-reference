package embedref_test

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"text/template"

	embedref "github.com/Ramadan-Lotfy/go-stdlib-reference/05-mastery/embed"
)

// Example_string prints a file that was compiled into the binary.
func Example_string() {
	fmt.Print(embedref.Hello)
	// Output: Hello, embedded world!
}

// Example_bytes shows the []byte form.
func Example_bytes() {
	fmt.Println(len(embedref.Data), string(embedref.Data))
	// Output: 3 ABC
}

// Example_fs walks the embedded tree. The hidden file "_hidden.tmpl" is not embedded.
func Example_fs() {
	fs.WalkDir(embedref.Files, ".", func(p string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			fmt.Println(p)
		}
		return nil
	})
	// Output:
	// testdata/static/app.css
	// testdata/templates/greeting.tmpl
}

// Example_template renders a template that ships inside the binary.
func Example_template() {
	t, err := template.ParseFS(embedref.Files, "testdata/templates/*.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.ExecuteTemplate(os.Stdout, "greeting.tmpl", map[string]string{"Name": "Ada"})
	// Output: Hello, Ada!
}

// Example_http serves embedded assets. fs.Sub removes the "testdata/static" prefix.
func Example_http() {
	static, _ := fs.Sub(embedref.Files, "testdata/static")
	h := http.FileServer(http.FS(static))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/app.css", nil))
	fmt.Println(rec.Code, rec.Header().Get("Content-Type"))
	fmt.Print(rec.Body.String())
	// Output:
	// 200 text/css; charset=utf-8
	// body { margin: 0 }
}

// Example_missing shows the error for a file that is not in the FS.
func Example_missing() {
	_, err := fs.ReadFile(embedref.Files, "testdata/nope.txt")
	fmt.Println(err)
	fmt.Println(errors.Is(err, fs.ErrNotExist))
	// Output:
	// open testdata/nope.txt: file does not exist
	// true
}
