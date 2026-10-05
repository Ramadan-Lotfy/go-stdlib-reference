package timeref_test

import (
	"fmt"
	"time"
)

// Example_durations shows Duration arithmetic, parsing and rounding.
func Example_durations() {
	d := 90 * time.Minute
	fmt.Println(d, d.Hours(), d.Minutes())

	n := 3
	fmt.Println(time.Duration(n) * time.Second) // convert the int first

	p, err := time.ParseDuration("1h15m30.5s")
	fmt.Println(p, err)
	fmt.Println(p.Round(time.Minute), p.Truncate(time.Hour))
	// Output:
	// 1h30m0s 1.5 90
	// 3s
	// 1h15m30.5s <nil>
	// 1h16m0s 1h0m0s
}

// Example_dateAndFormat builds a time and formats it with layouts.
func Example_dateAndFormat() {
	t := time.Date(2026, time.October, 3, 14, 30, 0, 0, time.UTC)
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Weekday(), t.YearDay())
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format("Mon Jan 2, 3:04PM"))
	fmt.Println(time.UnixMilli(t.UnixMilli()).UTC().Equal(t))
	// Output:
	// 2026 October 3 Saturday 276
	// 2026-10-03 14:30:00
	// 2026-10-03T14:30:00Z
	// Sat Oct 3, 2:30PM
	// true
}

// Example_parse converts text to a Time and shows the error for bad input.
func Example_parse() {
	t, err := time.Parse("2006-01-02", "2026-12-25")
	fmt.Println(t.Weekday(), err)

	_, err = time.Parse("2006-01-02", "2026-13-01")
	fmt.Println(err)

	t2, _ := time.Parse(time.RFC3339, "2026-10-03T14:30:00+02:00")
	fmt.Println(t2.UTC())
	// Output:
	// Friday <nil>
	// parsing time "2026-13-01": month out of range
	// 2026-10-03 12:30:00 +0000 UTC
}

// Example_arithmetic adds durations and calendar units, and compares times.
func Example_arithmetic() {
	t := time.Date(2026, time.January, 31, 9, 0, 0, 0, time.UTC)
	fmt.Println(t.Add(48 * time.Hour).Format(time.DateOnly))
	fmt.Println(t.AddDate(0, 1, 0).Format(time.DateOnly)) // Feb 31 normalizes to Mar 3

	end := time.Date(2026, time.March, 1, 9, 0, 0, 0, time.UTC)
	fmt.Println(end.Sub(t))
	fmt.Println(t.Before(end), t.After(end), t.Equal(t))
	fmt.Println(t.Truncate(24 * time.Hour).Format(time.DateTime))
	// Output:
	// 2026-02-02
	// 2026-03-03
	// 696h0m0s
	// true false true
	// 2026-01-31 00:00:00
}

// Example_zones converts between zones and shows why == is the wrong comparison.
func Example_zones() {
	utc := time.Date(2026, time.October, 3, 12, 0, 0, 0, time.UTC)
	plus2 := time.FixedZone("UTC+2", 2*60*60)
	local := utc.In(plus2)

	fmt.Println(local.Format(time.RFC3339))
	fmt.Println(utc == local, utc.Equal(local))
	// Output:
	// 2026-10-03T14:00:00+02:00
	// false true
}

// Example_timers covers Timer, Ticker, After and Stop.
func Example_timers() {
	timer := time.NewTimer(10 * time.Millisecond)
	<-timer.C
	fmt.Println("timer fired")

	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	for range 3 {
		<-ticker.C
	}
	fmt.Println("3 ticks")

	select {
	case <-time.After(10 * time.Millisecond):
		fmt.Println("timeout")
	case <-make(chan struct{}): // never ready
	}

	pending := time.NewTimer(time.Hour)
	fmt.Println(pending.Stop()) // true: stopped before firing
	// Output:
	// timer fired
	// 3 ticks
	// timeout
	// true
}

// Example_elapsed measures how long something took.
func Example_elapsed() {
	start := time.Now()
	time.Sleep(2 * time.Millisecond)
	fmt.Println(time.Since(start) >= 2*time.Millisecond)
	// Output: true
}
