package handlers

import (
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (m *HandlerConfig) State(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["topMsg"] = "Hello, again!"

	boolMap := make(map[string]bool)
	boolMap["stateMsg"] = true

	render.Template(w, r, "/state/landing.gohtml", &models.TemplateData{
		StringMap: stringMap,
		BoolMap:   boolMap,
	})
}

func (m *HandlerConfig) StateToggle(w http.ResponseWriter, r *http.Request) {

	//TODO: this stuff should really just be using the session?
	// So much for stateless, but it does reduce network load...
	isTrue := chi.URLParam(r, "isTrue")

	boolMap := make(map[string]bool)
	if isTrue == "true" {
		boolMap["stateMsg"] = false
	} else {
		boolMap["stateMsg"] = true
	}

	render.TemplateSnippet(w, r, "/state/toggle.state.gohtml", &models.TemplateData{
		BoolMap: boolMap,
	})
}
