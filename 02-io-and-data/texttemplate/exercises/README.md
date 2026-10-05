# text/template exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/texttemplate/exercises/...
```

## Exercise: RenderInvoice

Write `RenderInvoice(inv Invoice) (string, error)` using `text/template`:

```go
type LineItem struct {
	Name  string
	Qty   int
	Price float64
}

type Invoice struct {
	Customer string
	Items    []LineItem
}
```

Expected output for customer `Ada`, items `Pen x2 @ 1.50` and `Book x1 @ 12.25`:

```
Invoice for Ada
Pen x2 @ 1.50
Book x1 @ 12.25
Total: 15.25
```

Rules:

- Prices use two decimals (`printf "%.2f"`)
- The total is computed by a template function registered with `FuncMap`
- No items still prints the header and `Total: 0.00`

Hints: `range`, `Funcs` before `Parse`, a `strings.Builder` as the writer.

## Bonus
Move the template into a `templates/invoice.tmpl` file and load it with `embed`
and `template.ParseFS`.
