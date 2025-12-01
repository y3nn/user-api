package repository

import (
	"auth/backend/model"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemsRepository struct {
	db *pgxpool.Pool
}

func NewItemRepository(db *pgxpool.Pool) *ItemsRepository {
	return &ItemsRepository{db: db}
}

func (r *ItemsRepository) CreateItem(ctx context.Context, i *model.Item) error {
	query := `
		INSERT INTO items (item_name,item_price)
		VALUES ($1,$2)
		RETURNING id
		`
	return r.db.QueryRow(ctx, query, i.Name, i.Price).Scan(&i.ID)
}
