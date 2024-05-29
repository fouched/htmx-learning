package handlers

import (
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (m *HandlerConfig) Fun(w http.ResponseWriter, r *http.Request) {
	color := "green"
	stringMap := make(map[string]string)
	stringMap["color"] = color

	templates := []string{"/pages/fun/landing.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		StringMap: stringMap,
	})

}

func (m *HandlerConfig) FunChange(w http.ResponseWriter, r *http.Request) {

	color := chi.URLParam(r, "color")
	if color == "green" {
		color = "red"
	} else if color == "red" {
		color = "yellow"
	} else {
		color = "red"
	}

	stringMap := make(map[string]string)
	stringMap["color"] = color

	templates := []string{"/snippets/fun/change.gohtml"}
	render.Templates(w, r, templates, false, &models.TemplateData{
		StringMap: stringMap,
	})
}
