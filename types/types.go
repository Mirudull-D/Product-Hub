package types

import (
	"Product-Hub/db/generated"
	"context"
)

type UserStore interface {
	GetUserByEmail(ctx context.Context, email string) (*generated.User, error)
	GetUserById(ctx context.Context, id int) (*generated.User, error)
	CreateUser(ctx context.Context, payload RegisterUserPayload) error
	PingDb(ctx context.Context) error
}
type ProductStore interface {
	GetProducts(ctx context.Context) ([]generated.Product, error)
	GetProductsById(ctx context.Context, id []int32) ([]generated.Product, error)
	UpdateProduct(ctx context.Context, product generated.UpdateProductParams) error
}
type OrderStore interface {
	CreateOrder(ctx context.Context, order generated.CreateOrderParams) (int64, error)
	CreateOrderItems(ctx context.Context, order generated.CreateOrderItemsParams) error
}
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=20"`
}
type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CartItem struct {
	ProductId int `json:"productId" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type CartCheckoutPayload struct {
	Items []CartItem `json:"items" validate:"required"`
}
