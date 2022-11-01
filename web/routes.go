package web

import (
	"github.com/gorilla/mux"
)

func makeRoutes() *mux.Router {
	router := mux.NewRouter()

	invoiceRoutes := router.PathPrefix("/invoice").Subrouter()
	invoiceRoutes.HandleFunc("/generate", InvoiceGenerateHandler).Methods("POST")

	return router
}
