package types

import (
	"time"

)

type ProductStore interface {
	GetProducts() ([]Product, error)
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) (error)
}


type User struct {
	ID        int `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`

}

type RegisterUserPayload struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=10"`
}

type LoginUserPayload struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=10"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) (error)

}
type Product struct {
	ID int  `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	CreatedAt time.Time `json:"createdat"`
}

type Order struct {
	ID int `json:"id"`
	UserID int `json:"userId"`
	Total float64 `json:"total"`
	Status string `json:"status"`
	Address string `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID int `json:"id"`
	OrderID int `json:"orderId"`
	ProductID int `json:"productId"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}


