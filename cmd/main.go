package main

import (
	"auth/backend/db"
	"auth/backend/handlers/middleware"
	item "auth/backend/items"
	"auth/backend/users"
	"fmt"
	"log"
	"net/http"
)

func main() {
	pool := db.InitPool()
	defer pool.Close()
	userRepo := users.NewUserRepository(pool)
	itemRepo := item.NewItemRepository(pool)

	http.HandleFunc("/user", middleware.UserHandler(userRepo))
	http.HandleFunc("/item", middleware.ItemsMiddleware(itemRepo))

	fmt.Println("✅ |  Listen server on: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
