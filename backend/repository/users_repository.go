package repository

import (
	"auth/backend/model"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) CreateUser(ctx context.Context, u *model.User) error {
	query := `
		INSERT INTO users (name,email)
		VALUES ($1, $2)
		RETURNING id
	`
	return r.db.QueryRow(ctx, query, u.Name, u.Email).Scan(&u.ID)
}

func (r *UsersRepository) GetUser(ctx context.Context, id int64, u *model.User) error {
	query := ` 
		SELECT * FROM users WHERE id = $1
	`
	return r.db.QueryRow(ctx, query, id).Scan(&u.ID, &u.Name, &u.Email)
}
