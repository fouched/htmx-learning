package main

import (
	"github.com/fouched/htmx-learning/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Instance.Home)
	mux.Get("/fun", handlers.Instance.Fun)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
