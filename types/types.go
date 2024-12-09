package types

import (
	"time"

)

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

type ProductStore interface {
	GetProducts() ([]Product, error)
}