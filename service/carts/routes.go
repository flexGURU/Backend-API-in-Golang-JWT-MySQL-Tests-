package carts

import (
	"fmt"
	"net/http"

	"github.com/flexGURU/goAPI/types"
	"github.com/flexGURU/goAPI/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.OrderStore
	productStore types.ProductStore

}


func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store}
	
}

func (h *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout)
	
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	var cart types.CartCheckoutPayload
	userID := 0

	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadGateway, err)
		return
	} 

	if err := utils.Validate.Struct(cart); err != nil {
		utils.WriteError(w, http.StatusBadGateway, fmt.Errorf("bad cart payload: %v",err))
		return
	}

	productIDs, err := getCartItemIds(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err )
		return
	}

	ps, err := h.productStore.GetProductsByIDs(productIDs)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err )
		return
	}

	orderID, price, err := h.createOrder(ps, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return 
	}
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"total_price" : price,
		"order_id" : orderID,
	})



	






}