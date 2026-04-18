package main

import "net/http"

type Handler struct {
	// add here the service that we will use to handle the requests
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// here we will add the code to handle the request to create an order
}