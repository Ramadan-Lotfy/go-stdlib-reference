package filepathref_test

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

// Example_parts takes paths apart. ToSlash keeps the output identical on Windows.
func Example_parts() {
	p := "/tmp/x/report.tar.gz"
	fmt.Println(filepath.Base(p), filepath.Ext(p))
	fmt.Println(filepath.ToSlash(filepath.Dir(p)))

	dir, file := filepath.Split("a/b/c.txt")
	fmt.Println(filepath.ToSlash(dir), file)
	// Output:
	// report.tar.gz .gz
	// /tmp/x
	// a/b/ c.txt
}

// Example_joinAndClean builds and normalizes paths.
func Example_joinAndClean() {
	fmt.Println(filepath.ToSlash(filepath.Join("a", "b", "..", "c.txt")))
	fmt.Println(filepath.ToSlash(filepath.Clean("a//b/./c/..")))

	rel, err := filepath.Rel("/a/b", "/a/b/c/d")
	fmt.Println(filepath.ToSlash(rel), err)
	// Output:
	// a/c.txt
	// a/b
	// c/d <nil>
}

// Example_isLocal is the check to run on untrusted relative paths.
func Example_isLocal() {
	for _, p := range []string{"a/b.txt", "../secret", "/etc/passwd", "a/../../b"} {
		fmt.Printf("%-12s %t\n", p, filepath.IsLocal(p))
	}
	// Output:
	// a/b.txt      true
	// ../secret    false
	// /etc/passwd  false
	// a/../../b    false
}

// Example_match uses shell-style patterns.
func Example_match() {
	ok1, _ := filepath.Match("*.go", "main.go")
	ok2, _ := filepath.Match("[a-c]?.txt", "b1.txt")
	ok3, _ := filepath.Match("*.go", "dir/main.go") // * does not cross separators
	fmt.Println(ok1, ok2, ok3)
	// Output: true true false
}

// Example_walkDir lists a tree in lexical order.
func Example_walkDir() {
	dir, _ := os.MkdirTemp("", "walk-*")
	defer os.RemoveAll(dir)
	os.MkdirAll(filepath.Join(dir, "sub"), 0o755)
	os.WriteFile(filepath.Join(dir, "a.txt"), nil, 0o644)
	os.WriteFile(filepath.Join(dir, "sub", "b.go"), nil, 0o644)

	filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		rel, _ := filepath.Rel(dir, p)
		fmt.Println(filepath.ToSlash(rel), d.IsDir())
		return nil
	})
	// Output:
	// . true
	// a.txt false
	// sub true
	// sub/b.go false
}

// Example_glob finds files by pattern in one directory.
func Example_glob() {
	dir, _ := os.MkdirTemp("", "glob-*")
	defer os.RemoveAll(dir)
	for _, name := range []string{"b.txt", "a.txt", "c.log"} {
		os.WriteFile(filepath.Join(dir, name), nil, 0o644)
	}

	matches, _ := filepath.Glob(filepath.Join(dir, "*.txt"))
	for _, m := range matches {
		fmt.Println(filepath.Base(m))
	}
	// Output:
	// a.txt
	// b.txt
}

// Example_pathPackage uses package path, which always uses forward slashes.
func Example_pathPackage() {
	fmt.Println(path.Join("static", "css", "app.css"))
	fmt.Println(path.Base("/x/y.html"), path.Ext("/x/y.html"), path.Dir("/x/y.html"))
	// Output:
	// static/css/app.css
	// y.html .html /x
}
