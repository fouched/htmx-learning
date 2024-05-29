package render

import (
	"fmt"
	"github.com/fouched/htmx-learning/internal/config"
	"github.com/fouched/htmx-learning/internal/models"
	"html/template"
	"net/http"
)

var pathToTemplates = "./templates"
var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

// Templates can render multiple templates. "Parent" templates should be defined first
func Templates(w http.ResponseWriter, r *http.Request, tmpl []string, addBaseTemplate bool, td *models.TemplateData) {

	td = AddDefaultData(td, r)

	for i, t := range tmpl {
		tmpl[i] = pathToTemplates + t
	}

	if addBaseTemplate {
		tmpl = append(tmpl, pathToTemplates+"/base.layout.gohtml")
	}

	parsedTemplate, _ := template.ParseFiles(tmpl...)
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
