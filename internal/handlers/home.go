package handlers

import (
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"net/http"
)

// Home is the home page handler
func (m *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "/home.gohtml", &models.TemplateData{})
}
