package solution

import "testing"

func TestRenderInvoice(t *testing.T) {
	inv := Invoice{
		Customer: "Ada",
		Items:    []LineItem{{"Pen", 2, 1.5}, {"Book", 1, 12.25}},
	}
	got, err := RenderInvoice(inv)
	want := "Invoice for Ada\nPen x2 @ 1.50\nBook x1 @ 12.25\nTotal: 15.25\n"
	if err != nil || got != want {
		t.Errorf("RenderInvoice() = %q, %v\nwant %q", got, err, want)
	}

	got, _ = RenderInvoice(Invoice{Customer: "Bob"})
	if want := "Invoice for Bob\nTotal: 0.00\n"; got != want {
		t.Errorf("empty invoice = %q, want %q", got, want)
	}
}
