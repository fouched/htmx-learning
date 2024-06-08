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
	mux.Get("/basicForm", handlers.Instance.BasicFormLanding)
	mux.Post("/basicForm", handlers.Instance.BasicFormPost)
	mux.Get("/component", handlers.Instance.GetComponentPage)
	mux.Post("/component/add", handlers.Instance.AddPerson)
	mux.Get("/component/input/{id}", handlers.Instance.ComponentInputChange)
	mux.Get("/component/toggle/{isTrue}", handlers.Instance.ComponentToggle)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
