package handlers

import (
	"net/http"
	"text/template"
)

func HandleLandingPage(w http.ResponseWriter, r *http.Request) {
	// Parse both the base layout and the index template
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/index.html",
	))

	// Execute the template with some basic data
	data := struct {
		Title string
	}{
		Title: "Home",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
