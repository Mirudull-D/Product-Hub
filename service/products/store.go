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

func (s *Store) UpdateProduct(ctx context.Context, product generated.Product) error {
	err := s.queries.UpdateProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		queries: generated.New(db),
	}
}

func (s *Store) GetProducts(ctx context.Context) ([]generated.Product, error) {
	products, err := s.queries.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, err
}
func (s *Store) GetProductsById(ctx context.Context, id []int32) ([]generated.Product, error) {
	products, err := s.queries.GetProductsById(ctx, id)
	if err != nil {
		return nil, err
	}
	return products, err
}
