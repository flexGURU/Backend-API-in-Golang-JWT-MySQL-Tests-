package carts

import (
	"net/http"

	"github.com/flexGURU/goAPI/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.OrderStore
}


func NewHandler(store types.OrderStore) *Handler {
	return &Handler{store: store}
	
}

func (h *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout)
	
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	
}