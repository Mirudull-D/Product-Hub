package types

import "Product-Hub/db/generated"

type UserStore interface {
	GetUserByEmail(email string) (generated.User, error)
	GetUserById(id int) (generated.User, error)
	CreateUser(payload RegisterUserPayload) error
}
type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
