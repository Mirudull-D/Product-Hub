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
type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
