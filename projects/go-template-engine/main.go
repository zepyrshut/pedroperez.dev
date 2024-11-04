package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type TemplateData struct {
	Title string
	Name  string
}

// renderTemplate se encarga de obtener la ruta de la plantilla, hacer el
// procesamiento y su ejecución.
func renderTemplate(w http.ResponseWriter, templateName string, data TemplateData) {
	path := filepath.Join("templates", templateName)
	tmpl, _ := template.ParseFiles(path)
	tmpl.Execute(w, data)
}

func main() {

	templateData := TemplateData{
		Title: "Hola Mundo",
		Name:  "Gopher",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html", templateData)
	})

	http.ListenAndServe(":8080", nil)
}
