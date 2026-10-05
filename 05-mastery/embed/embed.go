// Package embedref shows how //go:embed puts files into the compiled binary.
package embedref

import "embed"

// Hello is a single file embedded as a string.
//
//go:embed testdata/hello.txt
var Hello string

// Data is a single file embedded as bytes.
//
//go:embed testdata/data.bin
var Data []byte

// Files embeds a directory tree and a glob. Names starting with "." or "_"
// are skipped unless the pattern starts with "all:".
//
//go:embed testdata/templates testdata/static/*.css
var Files embed.FS
