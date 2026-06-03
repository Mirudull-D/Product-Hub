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

func (s *Store) GetUserById(ctx context.Context, id int) (*generated.User, error) {
	u, err := s.queries.GetByUserId(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *Store) CreateUser(ctx context.Context, payload types.RegisterUserPayload) error {
	_, err := s.queries.CreateUser(ctx, generated.CreateUserParams{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	})
	return err
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*generated.User, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	return &user, err
}
