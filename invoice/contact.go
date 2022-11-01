package invoice

type Contact struct {
	Name         string
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
	Address      Address
}
