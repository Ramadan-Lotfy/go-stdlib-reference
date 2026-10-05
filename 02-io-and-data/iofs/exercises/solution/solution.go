// Package solution holds the reference solutions for the io/fs exercises.
package solution

import (
	"io/fs"
	"path"
	"strings"
)

// ListGoFiles returns every .go file in fsys in walk order. Directories named
// "vendor" and directories starting with "." are skipped.
func ListGoFiles(fsys fs.FS) ([]string, error) {
	var files []string
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if p != "." && (name == "vendor" || strings.HasPrefix(name, ".")) {
				return fs.SkipDir
			}
			return nil
		}
		if path.Ext(p) == ".go" {
			files = append(files, p)
		}
		return nil
	})
	return files, err
}
