// Package solution holds the reference solutions for the text/template exercises.
package solution

import (
	"strings"
	"text/template"
)

// LineItem is one row of an invoice.
type LineItem struct {
	Name  string
	Qty   int
	Price float64
}

// Invoice is the template data.
type Invoice struct {
	Customer string
	Items    []LineItem
}

const invoiceTemplate = `Invoice for {{.Customer}}
{{range .Items}}{{.Name}} x{{.Qty}} @ {{printf "%.2f" .Price}}
{{end}}Total: {{printf "%.2f" (total .Items)}}
`

// RenderInvoice renders an invoice as text.
func RenderInvoice(inv Invoice) (string, error) {
	funcs := template.FuncMap{
		"total": func(items []LineItem) float64 {
			var sum float64
			for _, it := range items {
				sum += float64(it.Qty) * it.Price
			}
			return sum
		},
	}
	t, err := template.New("invoice").Funcs(funcs).Parse(invoiceTemplate)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	if err := t.Execute(&sb, inv); err != nil {
		return "", err
	}
	return sb.String(), nil
}
