package invoice

type LineItem struct {
	Name      string
	Quantity  float64
	UnitPrice Price `json:"unit_price"`
}

func NewLineItem(name string, quantity float64, unitPrice Price) LineItem {
	return LineItem{
		name,
		quantity,
		unitPrice,
	}
}

func (i *LineItem) TotalPrice() Price {
	return i.UnitPrice.Mul(i.Quantity)
}
