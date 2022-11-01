package render

import (
	"strings"
	"testing"
	"time"

	"invoice-api/invoice"
)

func assert(t *testing.T, html string, content string) {
	if !strings.Contains(html, content) {
		t.Fatalf("Expected document to contain \"%s\", but it did not", content)
	}
}

func TestRenderInvoiceHTML(t *testing.T) {
	inv := invoice.NewInvoice("INV-1234")
	inv.BillFrom = invoice.Contact{
		Name:         "Example Contractor LLC",
		EmailAddress: "contractor@example.com",
		PhoneNumber:  "(555) 123-4567",
		Address: invoice.Address{
			Line1:      "123 Contractor Ln",
			City:       "Austin",
			State:      "TX",
			PostalCode: "12345",
		},
	}

	inv.BillTo = invoice.Contact{
		Name:         "Example Client, Inc.",
		EmailAddress: "client@example.com",
		PhoneNumber:  "(555) 987-6543",
		Address: invoice.Address{
			Line1:      "987 Wall St",
			Line2:      "Suite 9E",
			City:       "New York",
			State:      "NY",
			PostalCode: "98765",
		},
	}

	inv.Date.Time, _ = time.Parse(DATE_FORMAT, "July 23, 2022")
	inv.DueDate.Time, _ = time.Parse(DATE_FORMAT, "August 3, 2022")

	inv.AddLineItem(invoice.NewLineItem("Contractor Services", 123.45, invoice.NewPrice(98, 76)))
	inv.AddLineItem(invoice.NewLineItem("Administrative Services", 3.91, invoice.NewPrice(42, 23)))
	inv.AddLineItem(invoice.NewLineItem("Additional Cost", 1, invoice.NewPrice(999, 99)))

	bytes, err := renderInvoiceHTML(inv)
	if err != nil {
		t.Fatal(err)
	}

	htmlStr := string(bytes)

	assert(t, htmlStr, "INV-1234")
	assert(t, htmlStr, inv.BillFrom.Name)
	assert(t, htmlStr, inv.BillFrom.EmailAddress)
	assert(t, htmlStr, mailto(inv.BillFrom.EmailAddress))
	assert(t, htmlStr, inv.BillFrom.PhoneNumber)
	assert(t, htmlStr, inv.BillFrom.Address.Line1)
	assert(t, htmlStr, addressTopLine(&inv.BillFrom.Address))
	assert(t, htmlStr, addressBottomLine(&inv.BillFrom.Address))

	assert(t, htmlStr, inv.BillTo.Name)
	assert(t, htmlStr, inv.BillTo.EmailAddress)
	assert(t, htmlStr, mailto(inv.BillTo.EmailAddress))
	assert(t, htmlStr, inv.BillTo.PhoneNumber)
	assert(t, htmlStr, addressTopLine(&inv.BillTo.Address))
	assert(t, htmlStr, addressBottomLine(&inv.BillTo.Address))

	assert(t, htmlStr, inv.TotalPrice().String())

	for _, item := range inv.LineItems {
		assert(t, htmlStr, item.Name)
		assert(t, htmlStr, decimalFmt(item.Quantity))
		assert(t, htmlStr, item.UnitPrice.String())
		assert(t, htmlStr, item.TotalPrice().String())
	}

	assert(t, htmlStr, inv.Date.Format(DATE_FORMAT))
	assert(t, htmlStr, inv.DueDate.Format(DATE_FORMAT))
}
