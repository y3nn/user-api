package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)


type Repository struct { 
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *Repository { 
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, u*User) error { 
	query := `
		INSERT INTO users (name,email)
		VALUES ($1, $2)
		RETURNING id
	`
	return r.db.QueryRow(ctx,query,u.Name,u.Email).Scan(&u.ID)
}
