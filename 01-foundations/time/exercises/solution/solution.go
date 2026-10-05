// Package solution holds the reference solutions for the time exercises.
package solution

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// NextWeekday returns midnight of the first date strictly after from that
// falls on weekday wd, in from's location.
func NextWeekday(from time.Time, wd time.Weekday) time.Time {
	days := (int(wd) - int(from.Weekday()) + 7) % 7
	if days == 0 {
		days = 7
	}
	y, m, d := from.Date()
	return time.Date(y, m, d+days, 0, 0, 0, 0, from.Location())
}

// ParseHHMM parses "HH:MM" (for example "01:30") into a Duration.
func ParseHHMM(s string) (time.Duration, error) {
	hh, mm, ok := strings.Cut(s, ":")
	if !ok {
		return 0, errors.New("want HH:MM, got " + strconv.Quote(s))
	}
	h, err := strconv.Atoi(hh)
	if err != nil {
		return 0, fmt.Errorf("hours: %w", err)
	}
	m, err := strconv.Atoi(mm)
	if err != nil {
		return 0, fmt.Errorf("minutes: %w", err)
	}
	if h < 0 || m < 0 || m > 59 {
		return 0, errors.New("out of range: " + s)
	}
	return time.Duration(h)*time.Hour + time.Duration(m)*time.Minute, nil
}
