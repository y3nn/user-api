package main

import (
	"auth/internal/db"
	"auth/internal/handlers/middleware"
	"auth/internal/repository"
	routers "auth/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := routers.NewRouter()

	pool := db.InitPool()
	defer pool.Close()

	userRepo := repository.NewUserRepository(pool)
	// itemRepo := repository.NewItemRepository(pool)

	r.Post("/user", middleware.CreateUserHandler(userRepo)) // create
	r.Get("/user/{id}", middleware.ReadUser(userRepo))      // watch(read)
	r.Get("/users", middleware.ListMiddleWare(*userRepo))   // list
	r.Delete("/user/{id}", middleware.DeleteUser(userRepo)) // delete
	r.Patch("/user/{id}", middleware.UpdateUser(*userRepo)) // update

	fmt.Println("✅ |  Listen server on: http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
