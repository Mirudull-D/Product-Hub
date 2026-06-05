package products

import (
	"Product-Hub/db/generated"
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	db      *sql.DB
	queries *generated.Queries
	redis   *redis.Client
}

func (s *Store) UpdateProduct(ctx context.Context, product generated.UpdateProductParams) error {
	err := s.queries.UpdateProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func NewStore(db *sql.DB, rdb *redis.Client) *Store {
	return &Store{
		db:      db,
		queries: generated.New(db),
		redis:   rdb,
	}
}

func (s *Store) GetProducts(ctx context.Context) ([]generated.Product, error) {
	cached, err := s.redis.Get(
		ctx,
		"products",
	).Result()

	if err == nil {
		var products []generated.Product
		err = json.Unmarshal(
			[]byte(cached),
			&products,
		)
		if err == nil {
			return products, nil
		}
	}
	products, err := s.queries.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(products)

	s.redis.Set(
		ctx,
		"products",
		data,
		time.Minute*10,
	)
	return products, err
}
func (s *Store) GetProductsById(ctx context.Context, id []int32) ([]generated.Product, error) {
	products, err := s.queries.GetProductsById(ctx, id)
	if err != nil {
		return nil, err
	}
	return products, err
}
