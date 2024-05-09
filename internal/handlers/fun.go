package handlers

import (
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"net/http"
)

func (m *HandlerConfig) Fun(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "/fun.gohtml", &models.TemplateData{})
}

func (m *HandlerConfig) FunChange(w http.ResponseWriter, r *http.Request) {
	render.TemplateSnippet(w, r, "/fun.change.gohtml", &models.TemplateData{})
}
