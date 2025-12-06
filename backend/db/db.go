package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPool() *pgxpool.Pool {
	connstr := os.Getenv("DATABASE_URL")
	if connstr == "" {
		dbUser := os.Getenv("POSTGRES_USER")
		dbPass := os.Getenv("POSTGRES_PASSWORD")
		dbName := os.Getenv("POSTGRES_DB")
		connstr = fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?sslmode=disable", dbUser, dbPass, dbName)
	}

	pool, err := pgxpool.New(context.Background(), connstr)
	if err != nil {
		log.Fatalf("🚫 | Ошибка иницилизации  пула: %v \n", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("🚫 | Не удалось проверить соединение с БД (Ping): %v", err)

	}
	fmt.Println("✅ | Пул {PostgreSQL} иницилизирован! ")
	return pool
}
