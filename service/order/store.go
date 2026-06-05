package order

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

func (s *Store) CreateOrder(ctx context.Context, order generated.CreateOrderParams) (int64, error) {
	createdOrder, err := s.queries.CreateOrder(ctx, order)
	if err != nil {
		return 0, err
	}
	return createdOrder.ID, nil
}
func (s *Store) CreateOrderItems(ctx context.Context, order generated.CreateOrderItemsParams) error {
	_, err := s.queries.CreateOrderItems(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
