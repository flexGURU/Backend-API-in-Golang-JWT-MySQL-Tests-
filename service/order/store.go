package order

import (
	"database/sql"

	"github.com/flexGURU/goAPI/types"
)

type Store struct {
	db *sql.DB
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	statement := "INSERT INTO orders (userId, total, status, address) VALUES ($1, $2, $3, $4)" 
				
	result, err := s.db.Exec(statement, order.UserID, order.Total, order.Status, order.Address)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error  {
	statement := "INSERT INTO order_items (orderId, productsId, quantity, price) VALUES ($1, $2, $3, $4)"
	_, err := s.db.Exec(statement, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Quantity, orderItem.Price)
	return err
	
}