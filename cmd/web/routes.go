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
	mux.Get("/fun/{color}", handlers.Instance.FunChange)
	mux.Get("/state", handlers.Instance.GetStatePage)
	mux.Post("/state/add", handlers.Instance.AddPerson)
	mux.Get("/state/input/{id}", handlers.Instance.StateInputChange)
	mux.Get("/state/toggle/{isTrue}", handlers.Instance.StateToggle)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
