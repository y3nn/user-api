package routers

import (
	handler "auth/internal/handlers"
	"auth/internal/handlers/middleware"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux  {
	r := chi.NewRouter()
	r.Get("/hello",middleware.HelloMiddleWare(handler.HelloHandler))
	r.Get("/time",handler.TimeHandler)
	r.Get("/status",handler.StatusHandler)
	r.Post("/json",handler.AcceptAndGiveJSON)
	return  r 
}