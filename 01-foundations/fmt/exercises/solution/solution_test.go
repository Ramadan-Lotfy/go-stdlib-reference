package solution

import "testing"

func TestReceipt(t *testing.T) {
	items := []Item{
		{"Coffee", 2, 3.50},
		{"Bagel", 1, 2.25},
	}
	want := "ITEM         QTY    AMOUNT\nCoffee         2      7.00\nBagel          1      2.25\nTOTAL                 9.25\n"
	if got := Receipt(items); got != want {
		t.Errorf("Receipt() mismatch\n got: %q\nwant: %q", got, want)
	}
}

func TestReceiptEmpty(t *testing.T) {
	want := "ITEM         QTY    AMOUNT\nTOTAL                 0.00\n"
	if got := Receipt(nil); got != want {
		t.Errorf("Receipt(nil) mismatch\n got: %q\nwant: %q", got, want)
	}
}
