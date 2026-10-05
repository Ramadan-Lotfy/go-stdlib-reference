package osref_test

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Example_readWrite writes a file and reads it back.
func Example_readWrite() {
	dir, _ := os.MkdirTemp("", "demo-*")
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "a.txt")
	if err := os.WriteFile(path, []byte("hello\n"), 0o644); err != nil {
		fmt.Println(err)
		return
	}
	data, _ := os.ReadFile(path)
	fmt.Printf("%q\n", data)

	info, _ := os.Stat(path)
	fmt.Println(info.Name(), info.Size(), info.IsDir())
	// Output:
	// "hello\n"
	// a.txt 6 false
}

// Example_openFile appends to an existing file with explicit flags.
func Example_openFile() {
	dir, _ := os.MkdirTemp("", "demo-*")
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "log.txt")
	os.WriteFile(path, []byte("one\n"), 0o644)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(f, "two")
	if err := f.Close(); err != nil { // check Close on files you wrote to
		fmt.Println(err)
	}

	data, _ := os.ReadFile(path)
	fmt.Printf("%q\n", data)
	// Output: "one\ntwo\n"
}

// Example_readDir lists a directory. Entries come back sorted by name.
func Example_readDir() {
	dir, _ := os.MkdirTemp("", "demo-*")
	defer os.RemoveAll(dir)
	os.WriteFile(filepath.Join(dir, "b.txt"), []byte("b"), 0o644)
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("a"), 0o644)
	os.Mkdir(filepath.Join(dir, "sub"), 0o755)

	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		fmt.Println(e.Name(), e.IsDir())
	}
	// Output:
	// a.txt false
	// b.txt false
	// sub true
}

// Example_errors shows how to detect and inspect file errors.
func Example_errors() {
	_, err := os.ReadFile("definitely-missing.txt")
	fmt.Println(errors.Is(err, fs.ErrNotExist))

	var pe *fs.PathError
	if errors.As(err, &pe) {
		fmt.Println(pe.Op, pe.Path)
	}
	// Output:
	// true
	// open definitely-missing.txt
}

// Example_tempFile creates a uniquely named temporary file.
func Example_tempFile() {
	f, err := os.CreateTemp("", "demo-*.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(f.Name())
	defer f.Close()

	base := filepath.Base(f.Name())
	fmt.Println(strings.HasPrefix(base, "demo-"), strings.HasSuffix(base, ".txt"))
	// Output: true true
}

// Example_env reads and writes environment variables.
func Example_env() {
	os.Setenv("APP_MODE", "dev")
	defer os.Unsetenv("APP_MODE")

	fmt.Println(os.Getenv("APP_MODE"))

	v, ok := os.LookupEnv("APP_SURELY_UNSET")
	fmt.Printf("%q %t\n", v, ok)

	os.Setenv("APP_NAME", "ada")
	defer os.Unsetenv("APP_NAME")
	fmt.Println(os.ExpandEnv("hi ${APP_NAME}"))

	fmt.Println(os.Expand("$a-$b", func(k string) string { return strings.ToUpper(k) }))
	// Output:
	// dev
	// "" false
	// hi ada
	// A-B
}
