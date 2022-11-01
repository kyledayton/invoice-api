package invoice

import "testing"

func TestNewLineItem(t *testing.T) {
	name := "Widget"
	quantity := float64(10)
	unitPrice := NewPrice(9, 99)

	subject := NewLineItem(name, quantity, unitPrice)

	if subject.Name != name {
		t.Fatalf("Expected name: %v, got: %v", name, subject.Name)
	}

	if subject.Quantity != quantity {
		t.Fatalf("Expected quantity: %v, got: %v", quantity, subject.Quantity)
	}

	if subject.UnitPrice != unitPrice {
		t.Fatalf("Expected unit price: %v, got: %v", unitPrice, subject.UnitPrice)
	}
}

func TestLineItem_TotalPrice(t *testing.T) {
	item := NewLineItem("Widget", 98.25, NewPrice(9, 99))
	expectedPrice := NewPrice(981, 52)

	if item.TotalPrice() != expectedPrice {
		t.Fatalf("Expected total price: %v, got: %v", expectedPrice, item.TotalPrice())
	}
}
