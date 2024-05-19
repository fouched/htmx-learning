package handlers

import (
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"net/http"
)

// Home is the home page handler
func (m *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {
	templates := []string{"/home.gohtml", "/components/helloworld.top.gohtml", "/components/helloworld.bottom.gohtml"}
	stringMap := make(map[string]string)
	stringMap["topMsg"] = "Hello, prop"
	stringMap["bottomMsg"] = "Hello, prop again"
	render.MultipleTemplates(w, r, templates, &models.TemplateData{
		StringMap: stringMap,
	})
}
