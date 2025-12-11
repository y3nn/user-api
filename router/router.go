package routers

// import go-chi/chi package for routing
import (
	chi "github.com/go-chi/chi"
)

// NewRouter returns new chi router
func NewRouter() *chi.Mux {
	r := chi.NewRouter() // create new router instance
	return r             // return router
}
