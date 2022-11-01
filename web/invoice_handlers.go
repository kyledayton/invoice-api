package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"invoice-api/invoice"
	"invoice-api/invoice/render"
)

func InvoiceGenerateHandler(res http.ResponseWriter, req *http.Request) {
	// Parse input JSON
	var inv invoice.Invoice
	err := json.NewDecoder(req.Body).Decode(&inv)
	if err != nil {
		log.Printf("JSON Decoding failed: %v", err)
		renderError(res, errors.New("failed to parse request"), http.StatusBadRequest)
		return
	}

	// Render PDF
	pdf, err := render.PDF(&inv)
	if err != nil {
		log.Printf("PDF generation failed: %v", err)
		renderError(res, errors.New("failed to generate PDF"), http.StatusInternalServerError)
		return
	}

	// Send PDF file as download
	filename := fmt.Sprintf("Invoice #%s.pdf", inv.Number)

	res.Header().Set("Content-Type", "application/pdf")
	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	res.WriteHeader(200)
	res.Write(pdf)
}
