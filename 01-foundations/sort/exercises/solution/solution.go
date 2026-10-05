// Package solution holds the reference solutions for the sort exercises.
package solution

import "sort"

// TopN returns the n most frequent words, most frequent first. Ties are
// broken alphabetically. If n is larger than the number of distinct words,
// all of them are returned.
func TopN(words []string, n int) []string {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	uniq := make([]string, 0, len(counts))
	for w := range counts {
		uniq = append(uniq, w)
	}
	sort.Slice(uniq, func(i, j int) bool {
		if counts[uniq[i]] != counts[uniq[j]] {
			return counts[uniq[i]] > counts[uniq[j]]
		}
		return uniq[i] < uniq[j]
	})
	if n < 0 {
		n = 0
	}
	if n > len(uniq) {
		n = len(uniq)
	}
	return uniq[:n]
}
