# Invoice Generator API
Simple JSON API to create an invoice PDF.

## Run with podman (or docker)
```sh
podman build -t invoice-api-web . && podman run -p 8000:8000 invoice-api-web
```

## Documentation
### Environment variables
|Name|Default Value|Description|
|----|-------------|-----------|
|`PORT`|`8000`|Web server port|
|`CHROME_EXECUTABLE`|`google-chrome`|Google Chrome executable|

### Generate a PDF
```sh
curl --request POST \
  --url http://localhost:8000/invoice/generate \
  --header 'Content-Type: application/json' \
  --output 'Invoice #INV-1234.pdf' \
  --data '{
	"number": "INV-1234",
	"date": "2022-10-02",
	"due_date": "11/14/2022",
	"bill_from": {
		"name": "Consultant LLC",
		"address": {
			"line1": "123 Consultant Ln",
			"city": "Austin",
			"state": "TX",
			"postal_code": "12345"
		}
	},
	"bill_to": {
		"name": "Client, Inc.",
		"email_address": "client@example.com",
		"phone_number": "(555) 123-4567",
		"address": {
			"line1": "987 Client Blvd",
			"line2": "Suite 111",
			"city": "New York",
			"state": "NY",
			"postal_code": "98765"
		}
	},
	"line_items": [
		{
			"name": "Consulting Services (Hourly)",
			"quantity": 87.25,
			"unit_price": "$50"
		},
		{
			"name": "Support Services (Hourly)",
			"quantity": 12.35,
			"unit_price": [20, 0]
		},
		{
			"name": "Administrative Services (Hourly)",
			"quantity": 5.75,
			"unit_price": {
				"dollars": 12,
				"cents": 50
			}
		},
		{
			"name": "Incidental Cost",
			"quantity": 1,
			"unit_price": 199.99
		}
	]
}'
```
