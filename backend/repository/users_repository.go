package repository

import (
	"auth/backend/model"
	"context"
	"fmt"
	"log"

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

func (r *UsersRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	query := `SELECT id,name,email FROM users`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Rows - %v", err)
	}

	defer rows.Close()

	users := make([]*model.User, 0)
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, fmt.Errorf("ERROR: SCAN ROWS - %v", err)
		}
		users = append(users, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ERROR: [FOR] ROWS - %v", err)
	}
	return users, nil
}

func (r *UsersRepository) DeleteUser(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *UsersRepository) UpdateUser(ctx context.Context, u *model.User) error {
	query := `
	UPDATE users
	SET name = $1, email = $2 WHERE id = $3
	`
	result, err := r.db.Exec(ctx, query, u.Name, u.Email, u.ID)
	if err != nil {
		return fmt.Errorf("ERROR: failed to update user - %v", err)
	}
	if result.RowsAffected() == 0 {
		log.Printf("⚠️ | Не удалось обновить пользователя с id %d (пользователь не найден или данные не изменились)", u.ID)
		return fmt.Errorf("ERROR: user with id[%d] not found or no changes", u.ID)
	}
	return nil
}
