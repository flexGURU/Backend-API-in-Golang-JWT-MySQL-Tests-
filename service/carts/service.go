package carts

import (
	"fmt"

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


func (h *Handler) createOrder (products []types.Product, items []types.CartItem, userID int ) (int, float64, error) {
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
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce quantity of products in the db
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)

	}

	// create the order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID: userID,
		Total: totalPrice,
		Status: "pending",
		Address: "some addy",
		
	})
	if err != nil {
		return 0, 0, err
	}

	// create the order items
	for _, items := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID: orderID,
			ProductID: items.ProductID,
			Quantity: items.Quantity,
			Price: productMap[items.ProductID].Price,
		})
	}


	return orderID, totalPrice, nil

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

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64  {
	var totalPrice float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		totalPrice += product.Price * float64(item.Quantity)
	}

	return totalPrice
}