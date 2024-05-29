package handlers

import (
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"net/http"
)

// Home is the home page handler
func (m *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["topMsg"] = "Hello, prop"
	stringMap["bottomMsg"] = "Hello, prop again"

	templates := []string{"/pages/home.gohtml", "/components/helloworld.top.gohtml", "/components/helloworld.bottom.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		StringMap: stringMap,
	})
}
