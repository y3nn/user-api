package main

import (
	"auth/backend/db"
	"auth/backend/handlers/middleware"
	"auth/backend/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	pool := db.InitPool()
	defer pool.Close()

	userRepo := repository.NewUserRepository(pool)
	// itemRepo := repository.NewItemRepository(pool)

	r.Post("/user", middleware.CreateUserHandler(userRepo)) // create
	r.Get("/user/{id}", middleware.ReadUser(userRepo))      // watch(read)

	fmt.Println("✅ |  Listen server on: http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
