package handlers

import (
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func (m *HandlerConfig) State(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["topMsg"] = "Hello, again!"

	boolMap := make(map[string]bool)
	boolMap["stateMsg"] = true

	data := make(map[string]interface{})
	layout := "2006-01-02"

	birthDate, _ := time.Parse(layout, "1997-02-03")
	data["1"] = models.Person{
		FirstName: "Mary",
		LastName:  "Jones",
		DOB:       birthDate,
	}

	birthDate, _ = time.Parse(layout, "1990-04-06")
	data["2"] = models.Person{
		FirstName: "Jack",
		LastName:  "Smith",
		DOB:       birthDate,
	}

	render.Template(w, r, "/state/landing.gohtml", &models.TemplateData{
		StringMap: stringMap,
		BoolMap:   boolMap,
		Data:      data,
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
