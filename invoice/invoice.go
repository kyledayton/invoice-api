package invoice

type Invoice struct {
	Number    string
	Date      Date
	DueDate   Date       `json:"due_date"`
	BillFrom  Contact    `json:"bill_from"`
	BillTo    Contact    `json:"bill_to"`
	LineItems []LineItem `json:"line_items"`
}

func NewInvoice(number string) *Invoice {
	inv := &Invoice{
		Number:    number,
		LineItems: make([]LineItem, 0),
		BillFrom:  Contact{},
		BillTo:    Contact{},
	}

	return inv
}

func (i *Invoice) TotalPrice() Price {
	sum := NewPrice(0, 0)

	for _, line := range i.LineItems {
		sum = sum.Add(line.TotalPrice())
	}

	return sum
}

func (i *Invoice) AddLineItem(item LineItem) {
	i.LineItems = append(i.LineItems, item)
}
