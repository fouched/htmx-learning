package handlers

import (
	"fmt"
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func (m *HandlerConfig) State(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["topMsg"] = "Hello, GO & HTMX!"

	boolMap := make(map[string]bool)
	boolMap["stateMsg"] = true

	templates := []string{"/state/landing.gohtml", "/components/input.gohtml"}
	render.MultipleTemplates(w, r, templates, &models.TemplateData{
		StringMap:    stringMap,
		BoolMap:      boolMap,
		Data:         GetInitialState(),
		ComponentMap: GetComponents(),
	})
}

func (m *HandlerConfig) StateInputChange(w http.ResponseWriter, r *http.Request) {

	// which field are we working with
	id := chi.URLParam(r, "id")

	// get the fields data
	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
	}

	value := r.Form.Get(id)
	stringMap := make(map[string]string)
	stringMap["value"] = value

	// return a template snippet to be rendered elsewhere on the page
	render.TemplateSnippet(w, r, "/state/input.change.gohtml", &models.TemplateData{
		StringMap: stringMap,
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

func GetInitialState() map[string]interface{} {

	layout := "2006-01-02"
	birthDate, _ := time.Parse(layout, "1997-02-03")
	data := make(map[string]interface{})
	data["1"] = models.Person{
		ID:        1,
		FirstName: "Mary",
		LastName:  "Jones",
		DOB:       birthDate,
	}

	birthDate, _ = time.Parse(layout, "1990-04-06")
	data["2"] = models.Person{
		ID:        2,
		FirstName: "Jack",
		LastName:  "Smith",
		DOB:       birthDate,
	}

	return data
}

func GetComponents() map[string]interface{} {

	componentMap := make(map[string]interface{})

	componentMap["firstName"] = models.Input{
		Name:         "firstName",
		Title:        "First Name",
		Type:         "text",
		Class:        "form-control",
		AutoComplete: "firstNameNew",
		Value:        "",
		HxTrigger:    "keyup delay:100ms changed",
		HxPost:       "/state/input/firstName",
		HxTarget:     "#fn",
		HxSwap:       "innerHTML",
	}
	componentMap["lastName"] = models.Input{
		Name:         "lastName",
		Title:        "Last Name",
		Type:         "text",
		Class:        "form-control",
		AutoComplete: "lastNameNew",
		Value:        "",
		HxTrigger:    "keyup delay:100ms changed",
		HxPost:       "/state/input/lastName",
		HxTarget:     "#ln",
		HxSwap:       "innerHTML",
	}
	componentMap["dob"] = models.Input{
		Name:         "dob",
		Title:        "Date of Birth",
		Type:         "date",
		Class:        "form-control",
		AutoComplete: "dobNew",
		Value:        "",
		HxPost:       "/state/input/dob",
		HxTarget:     "#db",
		HxSwap:       "innerHTML",
	}

	return componentMap
}
