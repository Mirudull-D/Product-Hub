package user

import (
	"Product-Hub/db/generated"
	"Product-Hub/types"
	"context"
	"database/sql"
)

type Store struct {
	db      *sql.DB
	queries *generated.Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		queries: generated.New(db),
	}
}

func (s *Store) GetUserById(id int) (generated.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) CreateUser(payload types.RegisterUserPayload) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (generated.User, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	return user, err
}
