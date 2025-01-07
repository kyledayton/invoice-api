package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"invoice-api/invoice"
	"invoice-api/invoice/render"
	"invoice-api/pdf"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeRoutes(ctx *pdf.RemoteChromeDevToolsContext) *mux.Router {
	router := mux.NewRouter()

	invoiceRoutes := router.PathPrefix("/invoice").Subrouter()
	invoiceRoutes.HandleFunc("/generate", makeInvoiceGenerateHandler(ctx)).Methods("POST")

	return router
}

func makeInvoiceGenerateHandler(ctx *pdf.RemoteChromeDevToolsContext) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// Parse input JSON
		inv := invoice.NewInvoice("")

		err := json.NewDecoder(req.Body).Decode(inv)
		if err != nil {
			log.Printf("JSON Decoding failed: %v", err)
			renderError(res, errors.New("failed to parse request"), http.StatusBadRequest)
			return
		}

		// Render HTML
		html, err := render.RenderInvoiceHTML(inv)
		if err != nil {
			log.Printf("PDF generation failed: %v", err)
			renderError(res, errors.New("failed to generate PDF"), http.StatusInternalServerError)
			return
		}

		// Render PDF
		options := pdf.DefaultChromeDevToolsPDFRenderOptions()
		options.PrintBackground = true

		r := pdf.NewRemoteChromeDevToolsRenderer(ctx, options)

		pdf, err := r.RenderPDF(html)
		if err != nil {
			log.Printf("PDF generation failed: %v", err)
			renderError(res, errors.New("failed to generate PDF"), http.StatusInternalServerError)
			return
		}

		// Send PDF file as download
		filename := fmt.Sprintf("%s #%s.pdf", inv.Title, inv.Number)

		res.Header().Set("Content-Type", "application/pdf")
		res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

		res.WriteHeader(200)
		res.Write(pdf)
	}
}
