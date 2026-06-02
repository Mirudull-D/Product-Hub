package types

import (
	"Product-Hub/db/generated"
	"context"
)

type UserStore interface {
	GetUserByEmail(ctx context.Context, email string) (*generated.User, error)
	GetUserById(ctx context.Context, id int) (*generated.User, error)
	CreateUser(ctx context.Context, payload RegisterUserPayload) error
}
type ProductStore interface {
	GetProducts(ctx context.Context) ([]generated.GetProductsRow, error)
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
