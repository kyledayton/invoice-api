package render

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"math"
	"strings"

	"invoice-api/invoice"
)

//go:embed invoice.tmpl.html
var invoiceTemplateString string
var invoiceTemplate *template.Template

func init() {
	invoiceTemplate = template.New("invoice").Funcs(template.FuncMap{
		"addressTopLine":    addressTopLine,
		"addressBottomLine": addressBottomLine,
		"mailto":            mailto,
		"date":              dateFmt,
		"decimal":           decimalFmt,
	})

	invoiceTemplate = template.Must(invoiceTemplate.Parse(invoiceTemplateString))
}

func renderInvoiceHTML(invoice *invoice.Invoice) ([]byte, error) {
	output := bytes.NewBuffer(nil)

	err := invoiceTemplate.Execute(output, invoice)
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

func addressTopLine(a *invoice.Address) string {
	if strings.TrimSpace(a.Line2) != "" {
		return fmt.Sprintf("%s, %s", a.Line1, a.Line2)
	} else {
		return a.Line1
	}
}

func addressBottomLine(a *invoice.Address) string {
	return fmt.Sprintf("%s, %s %s", a.City, a.State, a.PostalCode)
}

func mailto(email string) string {
	return fmt.Sprintf("mailto:%s", email)
}

const DATE_FORMAT = "January 2, 2006"

func dateFmt(date invoice.Date) string {
	return date.Time.Format(DATE_FORMAT)
}

func decimalFmt(f float64) string {
	const EPSILLON float64 = 1e-100

	rounded := math.Round(f)

	if math.Abs(f-rounded) < EPSILLON {
		return fmt.Sprintf("%.0f", rounded)
	} else {
		return fmt.Sprintf("%.2f", f)
	}
}
