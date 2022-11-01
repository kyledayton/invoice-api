package invoice

import "testing"

func TestNewInvoice(t *testing.T) {
	number := "INV-001"
	inv := NewInvoice(number)

	if inv.Number != number {
		t.Fatalf("Got: %v, expected: %v", inv.Number, number)
	}
}

func TestInvoice_AddLineItem(t *testing.T) {
	item1 := NewLineItem("Foo", 10, NewPrice(9, 99))
	item2 := NewLineItem("Bar", 5, NewPrice(123, 45))

	inv := NewInvoice("INV-001")
	inv.AddLineItem(item1)
	inv.AddLineItem(item2)

	if len(inv.LineItems) != 2 {
		t.Fatalf("Got len(LineItems): %v, expected: %v", len(inv.LineItems), 2)
	}

	if inv.LineItems[0] != item1 {
		t.Fatalf("Got LineItems[0]: %v, expected: %v", inv.LineItems[0], item1)
	}

	if inv.LineItems[1] != item2 {
		t.Fatalf("Got LineItems[1]: %v, expected: %v", inv.LineItems[1], item2)
	}
}

func TestInvoice_TotalPrice(t *testing.T) {
	item1 := NewLineItem("Foo", 10, NewPrice(9, 99))
	item2 := NewLineItem("Bar", 5, NewPrice(123, 45))

	inv := NewInvoice("INV-001")
	inv.AddLineItem(item1)
	inv.AddLineItem(item2)

	total := inv.TotalPrice()
	expected := NewPrice(717, 15)

	if total != expected {
		t.Fatalf("Got TotalPrice: %v, expected: %v", total, expected)
	}
}

func TestInvoice_TotalPrice_Zero(t *testing.T) {
	inv := NewInvoice("INV-001")

	total := inv.TotalPrice()
	expected := NewPrice(0, 0)

	if total != expected {
		t.Fatalf("Got TotalPrice: %v, expected: %v", total, expected)
	}
}
