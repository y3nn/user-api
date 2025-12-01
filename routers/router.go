package routers

import (
	handler "auth/backend/handlers"

	chi "github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/time", handler.TimeHandler)
	r.Get("/status", handler.StatusHandler)
	r.Post("/json", handler.AcceptAndGiveJSON)
	return r
}
