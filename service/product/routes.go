package product

import (
	"net/http"

	"github.com/flexGURU/goAPI/types"
	"github.com/flexGURU/goAPI/utils"
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
	prods, err := h.Store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusBadGateway, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, prods )
}
