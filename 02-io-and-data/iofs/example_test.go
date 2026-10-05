package iofsref_test

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing/fstest"
)

func newFS() fstest.MapFS {
	return fstest.MapFS{
		"a.txt":         {Data: []byte("A")},
		"dir/b.txt":     {Data: []byte("BB")},
		"dir/sub/c.txt": {Data: []byte("CCC")},
	}
}

// Example_readFile reads, stats and lists files through the fs.FS interface.
func Example_readFile() {
	fsys := newFS()

	data, _ := fs.ReadFile(fsys, "dir/b.txt")
	fmt.Println(string(data))

	info, _ := fs.Stat(fsys, "dir/sub/c.txt")
	fmt.Println(info.Name(), info.Size())

	entries, _ := fs.ReadDir(fsys, "dir")
	for _, e := range entries {
		fmt.Println(e.Name(), e.IsDir())
	}
	// Output:
	// BB
	// c.txt 3
	// b.txt false
	// sub true
}

// Example_walkDir visits every file and directory in lexical order.
func Example_walkDir() {
	fs.WalkDir(newFS(), ".", func(p string, d fs.DirEntry, err error) error {
		fmt.Println(p, d.IsDir())
		return nil
	})
	// Output:
	// . true
	// a.txt false
	// dir true
	// dir/b.txt false
	// dir/sub true
	// dir/sub/c.txt false
}

// Example_skipDir prunes part of the tree.
func Example_skipDir() {
	fs.WalkDir(newFS(), ".", func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() && d.Name() == "sub" {
			return fs.SkipDir
		}
		if !d.IsDir() {
			fmt.Println(p)
		}
		return nil
	})
	// Output:
	// a.txt
	// dir/b.txt
}

// Example_globAndSub matches patterns and re-roots the file system.
func Example_globAndSub() {
	fsys := newFS()

	matches, _ := fs.Glob(fsys, "dir/*.txt")
	fmt.Println(matches)

	sub, _ := fs.Sub(fsys, "dir")
	data, _ := fs.ReadFile(sub, "b.txt")
	fmt.Println(string(data))
	// Output:
	// [dir/b.txt]
	// BB
}

// Example_validPath shows the path rules every fs.FS enforces.
func Example_validPath() {
	fmt.Println(fs.ValidPath("a/b.txt"))
	fmt.Println(fs.ValidPath("/a/b.txt"))
	fmt.Println(fs.ValidPath("a/../b.txt"))
	fmt.Println(fs.ValidPath("."))
	// Output:
	// true
	// false
	// false
	// true
}

// countTxt takes an fs.FS, so it works on disk, embedded files and in tests.
func countTxt(fsys fs.FS) int {
	n := 0
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() && filepath.Ext(p) == ".txt" {
			n++
		}
		return nil
	})
	return n
}

// Example_dirFS runs the same function against a real directory and a fake one.
func Example_dirFS() {
	dir, _ := os.MkdirTemp("", "dirfs-*")
	defer os.RemoveAll(dir)
	os.WriteFile(filepath.Join(dir, "x.txt"), []byte("x"), 0o644)
	os.WriteFile(filepath.Join(dir, "y.md"), []byte("y"), 0o644)

	fmt.Println(countTxt(os.DirFS(dir)))
	fmt.Println(countTxt(newFS()))
	// Output:
	// 1
	// 3
}
