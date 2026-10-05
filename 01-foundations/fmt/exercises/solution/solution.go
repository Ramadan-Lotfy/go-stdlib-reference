// Package solution holds the reference solution for the fmt exercise.
package solution

import (
	"fmt"
	"strings"
)

// Item is one line on a receipt.
type Item struct {
	Name  string
	Qty   int
	Price float64 // unit price
}

// Receipt renders items as an aligned text table with a total line.
func Receipt(items []Item) string {
	var sb strings.Builder
	var total float64

	fmt.Fprintf(&sb, "%-12s %3s %9s\n", "ITEM", "QTY", "AMOUNT")
	for _, it := range items {
		amount := float64(it.Qty) * it.Price
		fmt.Fprintf(&sb, "%-12s %3d %9.2f\n", it.Name, it.Qty, amount)
		total += amount
	}
	fmt.Fprintf(&sb, "%-12s %3s %9.2f\n", "TOTAL", "", total)
	return sb.String()
}
