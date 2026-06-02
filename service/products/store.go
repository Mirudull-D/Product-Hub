package products

import (
	"Product-Hub/db/generated"
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

func (s *Store) GetProducts(ctx context.Context) ([]generated.GetProductsRow, error) {
	products, err := s.queries.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, err
}
