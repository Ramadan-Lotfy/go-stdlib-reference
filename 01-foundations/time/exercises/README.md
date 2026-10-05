# time exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./01-foundations/time/exercises/...
```

## Exercise 1: NextWeekday

Write `NextWeekday(from time.Time, wd time.Weekday) time.Time` that returns
midnight of the first date **strictly after** `from` that falls on `wd`, in
`from`'s location.

| from | wd | Result |
|------|----|--------|
| Sat 2026-10-03 | Monday | 2026-10-05 |
| Sat 2026-10-03 | Saturday | 2026-10-10 |

Hints: `from.Weekday()`, `time.Date` (it normalizes day overflow), modular arithmetic.

## Exercise 2: ParseHHMM

Write `ParseHHMM(s string) (time.Duration, error)` that parses `"01:30"` into
`1h30m`. Return an error for a missing colon, non-numbers, negative values or
minutes above 59.

Hints: `strings.Cut`, `strconv.Atoi`, `time.Duration(h) * time.Hour`.

## Bonus
Write a `Countdown` that prints the time left until a deadline every second using `time.Ticker`, stopping with `context`.
