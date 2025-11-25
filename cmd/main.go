package main

import (
	"auth/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r :=  routers.NewRouter()
	fmt.Println("✅ '  listen server on: http://localhost:8080")
	if err := http.ListenAndServe(":8080",r); err != nil { 
		log.Fatal(err)
	}
}