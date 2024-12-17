package carts

import (
	"fmt"

	"github.com/flexGURU/goAPI/service/product"
	"github.com/flexGURU/goAPI/types"
)

// Obtain the prodct IDs
func getCartItemIds(items []types.CartItem) ([]int, error) {

	productIDs := make([]int, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("no available items for product %d ", item.ProductID)
		}

		productIDs[i] =item.ProductID
	}

	return productIDs, nil

}


func (h *Handler) createOrder (products []types.Product, items []types.CartItem, user types.User ) (int, float64, error) {
	// create a product Map

	productMap := make(map[int]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	// check if all products are actually in stock

	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err

	}
	 
	// calculate the total price

	totalPrice := calculateTotalPrice()
	// reduce quantity of products in the db
	// create the order
	// create the order items




}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	// check of the cart item is empty
	if len(cartItems) == 0 {
		return fmt.Errorf("cart items are zero")
	}

	// if not empty check of product actually exists by looping through the cart items
	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available", item.ProductID)
		}
		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantiy requested", product.Name)
		}
	}

	return nil


	
}