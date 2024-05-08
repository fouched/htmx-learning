package handlers

import (
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"net/http"
)

// Fun is the home page handler
func (m *HandlerConfig) Fun(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "/fun.gohtml", &models.TemplateData{})
}
