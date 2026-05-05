package main

import (
	"net/http"

	pb "github.com/OMSV1/common/api"
)

type Handler struct {
	client pb.OmsServiceClient
}

func NewHandler(client pb.OmsServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		// fill in the request fields based on the incoming HTTP request
	})
}
