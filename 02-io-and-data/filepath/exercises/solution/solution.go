// Package solution holds the reference solutions for the path/filepath exercises.
package solution

import (
	"fmt"
	"path/filepath"
	"strings"
)

// SafeJoin joins an untrusted relative path onto base, refusing any path that
// would escape base.
func SafeJoin(base, userPath string) (string, error) {
	if !filepath.IsLocal(userPath) {
		return "", fmt.Errorf("unsafe path %q", userPath)
	}
	return filepath.Join(base, userPath), nil
}

// ChangeExt replaces the extension of name (or adds one). newExt includes the dot.
func ChangeExt(name, newExt string) string {
	return strings.TrimSuffix(name, filepath.Ext(name)) + newExt
}
