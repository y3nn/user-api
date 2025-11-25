package routers

import "github.com/go-chi/chi"



func NewRouter() *chi.Mux  {
	r := chi.NewRouter()
	r.Get("/hello")
	r.Get("/time")
	r.Get("/status",)

	return  r 
}