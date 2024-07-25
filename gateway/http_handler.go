package main

import "net/http"

type handler struct {
	// Gateway layer
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux){
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request){
	
}