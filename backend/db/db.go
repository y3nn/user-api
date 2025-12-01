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
		connstr = "postgresql://yan:123123123@localhost:5432/yoyo?sslmode=disable"
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
