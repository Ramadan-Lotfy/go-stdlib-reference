// Package solution holds the reference solutions for the strings exercises.
package solution

import (
	"strings"
	"unicode"
)

// Slugify turns a title into a URL-friendly slug: letters and digits are
// kept (lowercased), and every run of other characters becomes one hyphen.
func Slugify(s string) string {
	mapped := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return ' '
	}, s)
	return strings.Join(strings.Fields(mapped), "-")
}

// CountWords counts how often each word appears in text. Words are
// case-insensitive and surrounding punctuation is ignored.
func CountWords(text string) map[string]int {
	counts := make(map[string]int)
	for _, w := range strings.Fields(strings.ToLower(text)) {
		w = strings.Trim(w, `.,;:!?"'()`)
		if w != "" {
			counts[w]++
		}
	}
	return counts
}
