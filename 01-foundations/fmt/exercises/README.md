# fmt exercises

## Exercise 1: Receipt printer

Write `Receipt(items []Item) string` that renders an aligned text table.

```go
type Item struct {
	Name  string
	Qty   int
	Price float64 // unit price
}
```

Requirements:

- Header line: `ITEM`, `QTY`, `AMOUNT`
- One line per item; AMOUNT is `Qty * Price` with 2 decimals
- Final `TOTAL` line with the sum, QTY column left blank
- Columns: name left-aligned in 12 characters, qty right-aligned in 3,
  amount right-aligned in 9, separated by one space
- Every line ends with `\n`

Input:

```go
[]Item{{"Coffee", 2, 3.50}, {"Bagel", 1, 2.25}}
```

Expected output (spacing is exact):

```
ITEM         QTY    AMOUNT
Coffee         2      7.00
Bagel          1      2.25
TOTAL                 9.25
```

Hints: `%-12s`, `%3d`, `%9.2f`, and `fmt.Fprintf` into a `strings.Builder`.

Try it yourself first, then compare with
[solution/solution.go](solution/solution.go). Run the checks with:

```bash
go test ./01-foundations/fmt/exercises/...
```
