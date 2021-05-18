package main

import (
	"net/http"

	"github.com/Christoffer-lindow/portfolio/internal/config"
	"github.com/Christoffer-lindow/portfolio/internal/handlers"
	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
	Mux := chi.NewMux()
	Mux.Get("/api/properties/{city}", handlers.Repo.Properties)
	Mux.Get("/api/properties/single/{id}/city/{city}", handlers.Repo.SingleProperty)
	Mux.Get("/api/properties/max_price/{city}", handlers.Repo.MostExpensiveProperty)
	return Mux
}
