package solution

import (
	"maps"
	"testing"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"Hello, World!", "hello-world"},
		{"  Go   1.22 Released  ", "go-1-22-released"},
		{"Spring Boot & Go", "spring-boot-go"},
		{"---", ""},
		{"", ""},
	}
	for _, tt := range tests {
		if got := Slugify(tt.in); got != tt.want {
			t.Errorf("Slugify(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCountWords(t *testing.T) {
	got := CountWords(`The cat and the hat. The "end"!`)
	want := map[string]int{"the": 3, "cat": 1, "and": 1, "hat": 1, "end": 1}
	if !maps.Equal(got, want) {
		t.Errorf("CountWords() = %v, want %v", got, want)
	}
	if got := CountWords("  ...  "); len(got) != 0 {
		t.Errorf("CountWords(punctuation only) = %v, want empty map", got)
	}
}
