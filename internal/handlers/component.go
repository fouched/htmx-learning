package handlers

import (
	"fmt"
	"github.com/fouched/htmx-learning/internal/models"
	"github.com/fouched/htmx-learning/internal/render"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

var data = make(map[string]interface{})

func (m *HandlerConfig) GetComponentPage(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["topMsg"] = "Hello, GO & HTMX!"

	boolMap := make(map[string]bool)
	boolMap["stateMsg"] = true

	InitInitialState()

	templates := []string{"/pages/component/landing.gohtml", "/components/input.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		StringMap:    stringMap,
		BoolMap:      boolMap,
		Data:         data,
		ComponentMap: GetComponents(),
	})
}

func (m *HandlerConfig) ComponentInputChange(w http.ResponseWriter, r *http.Request) {

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

	templates := []string{"/snippets/component/input.gohtml"}
	render.Templates(w, r, templates, false, &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *HandlerConfig) ComponentToggle(w http.ResponseWriter, r *http.Request) {

	isTrue := chi.URLParam(r, "isTrue")

	boolMap := make(map[string]bool)
	if isTrue == "true" {
		boolMap["stateMsg"] = false
	} else {
		boolMap["stateMsg"] = true
	}

	templates := []string{"/snippets/component/toggle.gohtml"}
	render.Templates(w, r, templates, false, &models.TemplateData{
		BoolMap: boolMap,
	})
}

func (m *HandlerConfig) AddPerson(w http.ResponseWriter, r *http.Request) {

	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
	}

	layout := "2006-01-02"
	birthDate, _ := time.Parse(layout, r.Form.Get("dob"))
	id := len(data) + 1
	ids := strconv.Itoa(id)
	data[ids] = models.Person{
		ID:        id,
		FirstName: r.Form.Get("firstName"),
		LastName:  r.Form.Get("lastName"),
		DOB:       birthDate,
	}

	templates := []string{"/snippets/component/grid.gohtml", "/components/input.gohtml"}
	render.Templates(w, r, templates, false, &models.TemplateData{
		Data:         data,
		ComponentMap: GetComponents(),
	})

}

func InitInitialState() {

	if len(data) == 0 {
		layout := "2006-01-02"

		birthDate, _ := time.Parse(layout, "1997-02-03")
		data[strconv.Itoa(len(data)+1)] = models.Person{
			ID:        1,
			FirstName: "Mary",
			LastName:  "Jones",
			DOB:       birthDate,
		}

		birthDate, _ = time.Parse(layout, "1990-04-06")
		data[strconv.Itoa(len(data)+1)] = models.Person{
			ID:        2,
			FirstName: "Jack",
			LastName:  "Smith",
			DOB:       birthDate,
		}
	}
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
		HxGet:        "/component/input/firstName",
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
		HxGet:        "/component/input/lastName",
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
		HxGet:        "/component/input/dob",
		HxTarget:     "#db",
		HxSwap:       "innerHTML",
	}

	return componentMap
}
