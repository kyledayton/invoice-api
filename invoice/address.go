package invoice

type Address struct {
	Line1      string
	Line2      string
	City       string
	State      string
	PostalCode string `json:"postal_code"`
}
