# encoding/csv exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/csv/exercises/...
```

## Exercise: ParseUsers

Write `ParseUsers(r io.Reader) ([]User, error)` for rows like `name,age`:

```go
type User struct {
	Name string
	Age  int
}
```

Rules:

- The first row is a header. Find the `name` and `age` columns by name,
  ignoring case and spaces, so `Age,NAME` works too
- A missing column or an empty file is an error
- A non-numeric age returns an error mentioning the line, such as
  `line 3: invalid age: ...`, and wraps the `strconv` error with `%w`
- Quoted fields with commas must work

Hints: `csv.NewReader`, a `map[string]int` of header positions, `r.FieldPos`.

## Bonus
Add `WriteUsers(w io.Writer, users []User) error` and check that
`ParseUsers(WriteUsers(x))` returns `x`.
