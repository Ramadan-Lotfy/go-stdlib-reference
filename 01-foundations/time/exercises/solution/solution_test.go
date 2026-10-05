package solution

import (
	"testing"
	"time"
)

func TestNextWeekday(t *testing.T) {
	sat := time.Date(2026, time.October, 3, 14, 30, 0, 0, time.UTC)
	tests := []struct {
		wd   time.Weekday
		want string
	}{
		{time.Monday, "2026-10-05"},
		{time.Saturday, "2026-10-10"}, // same weekday: strictly after
		{time.Sunday, "2026-10-04"},
		{time.Friday, "2026-10-09"},
	}
	for _, tt := range tests {
		got := NextWeekday(sat, tt.wd)
		if got.Format(time.DateOnly) != tt.want || got.Hour() != 0 {
			t.Errorf("NextWeekday(%v) = %v, want %s at midnight", tt.wd, got, tt.want)
		}
	}
}

func TestParseHHMM(t *testing.T) {
	d, err := ParseHHMM("01:30")
	if err != nil || d != 90*time.Minute {
		t.Errorf("ParseHHMM(01:30) = %v, %v", d, err)
	}
	for _, bad := range []string{"", "1", "ab:cd", "01:75", "-1:00"} {
		if _, err := ParseHHMM(bad); err == nil {
			t.Errorf("ParseHHMM(%q) should fail", bad)
		}
	}
}
