package product

import (
	"net/http"

	"github.com/flexGURU/goAPI/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	Store types.ProductStore
}

func NewHandler(s types.ProductStore) *Handler  {

	return &Handler{Store: s}
	
}


func (h *Handler) RegisterRoute(router *mux.Router) {

	router.HandleFunc("/products", h.productHandler)
	
}

func (h *Handler) productHandler(w http.ResponseWriter, r *http.Request) {
	h.Store.GetProducts()
	
}
