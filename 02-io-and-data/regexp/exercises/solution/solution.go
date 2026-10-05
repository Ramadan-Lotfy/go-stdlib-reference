// Package solution holds the reference solutions for the regexp exercises.
package solution

import (
	"regexp"
	"strings"
)

var logRe = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}) (?P<level>[A-Z]+) \[(?P<comp>\w+)\] (?P<msg>.*)$`)

// ParseLogLine parses "2026-10-03 ERROR [db] connection lost" into its named
// parts (date, level, comp, msg). ok is false if the line does not match.
func ParseLogLine(line string) (fields map[string]string, ok bool) {
	m := logRe.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	fields = make(map[string]string)
	for i, name := range logRe.SubexpNames() {
		if name != "" {
			fields[name] = m[i]
		}
	}
	return fields, true
}

var longDigits = regexp.MustCompile(`\d{4,}`)

// RedactDigits replaces every run of four or more digits with the same number
// of asterisks.
func RedactDigits(s string) string {
	return longDigits.ReplaceAllStringFunc(s, func(m string) string {
		return strings.Repeat("*", len(m))
	})
}
