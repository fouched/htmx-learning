package handlers

import (
	"fmt"
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"net/http"
)

func (m *HandlerConfig) BasicFormLanding(w http.ResponseWriter, r *http.Request) {

	templates := []string{"/pages/basicForm/landing.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{})
}

func (m *HandlerConfig) BasicFormPost(w http.ResponseWriter, r *http.Request) {

	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
	}

	firstName := r.Form.Get("firstName")
	println("firstName:" + firstName)
	println("lastName:" + r.Form.Get("lastName"))
	println("email:" + r.Form.Get("email"))
	println("password:" + r.Form.Get("password"))

	boolMap := make(map[string]bool)
	if firstName == "Fouche" {
		boolMap["success"] = true
	} else {
		boolMap["success"] = false
	}

	templates := []string{"/pages/basicForm/landing.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		BoolMap: boolMap,
	})
}
