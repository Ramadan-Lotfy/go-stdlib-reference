package solution

import (
	"path/filepath"
	"testing"
)

func TestSafeJoin(t *testing.T) {
	got, err := SafeJoin("/srv/files", "a/b.txt")
	if err != nil || filepath.ToSlash(got) != "/srv/files/a/b.txt" {
		t.Errorf("SafeJoin() = %q, %v", got, err)
	}
	for _, bad := range []string{"../etc/passwd", "/etc/passwd", "a/../../b", ""} {
		if _, err := SafeJoin("/srv/files", bad); err == nil {
			t.Errorf("SafeJoin(%q) should fail", bad)
		}
	}
}

func TestChangeExt(t *testing.T) {
	tests := []struct{ in, ext, want string }{
		{"report.txt", ".md", "report.md"},
		{"archive.tar.gz", ".md", "archive.tar.md"},
		{"noext", ".md", "noext.md"},
	}
	for _, tt := range tests {
		if got := ChangeExt(tt.in, tt.ext); got != tt.want {
			t.Errorf("ChangeExt(%q, %q) = %q, want %q", tt.in, tt.ext, got, tt.want)
		}
	}
}
