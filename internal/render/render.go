package render

import (
	"fmt"
	"github.com/fouched/htmx-learning/internal/config"
	"github.com/fouched/htmx-learning/internal/models"
	"html/template"
	"net/http"
)

var pathToTemplates = "./templates"
var pathToTemplateSnippet = "./templates/snippets"
var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

// MultipleTemplates renders parent and children templates. The parent should be defined first
func MultipleTemplates(w http.ResponseWriter, r *http.Request, tmpl []string, td *models.TemplateData) {

	td = AddDefaultData(td, r)

	for i, t := range tmpl {
		tmpl[i] = pathToTemplates + t
	}
	tmpl = append(tmpl, pathToTemplates+"/base.layout.gohtml")

	parsedTemplate, _ := template.ParseFiles(tmpl...)
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}

}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	td = AddDefaultData(td, r)

	parsedTemplate, _ := template.ParseFiles(pathToTemplates+tmpl, pathToTemplates+"/base.layout.gohtml")
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}

}

func TemplateSnippet(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	parsedTemplate, _ := template.ParseFiles(pathToTemplateSnippet + tmpl)
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {

	//td.Success = app.Session.PopString(r.Context(), "success")
	//td.Warning = app.Session.PopString(r.Context(), "warning")
	//td.Error = app.Session.PopString(r.Context(), "error")

	return td
}
