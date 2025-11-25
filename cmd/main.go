package main

import (
	"auth/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	r :=  handlers.NewRouter()
	fmt.Println("listen server on: http://localhost:8080")
	http.ListenAndServe(":8080",r)
}