package handlers

import (
	"net/http"
	"text/template"

	"me.tofaa/internal/types"
)

func HandleProjectsPage(w http.ResponseWriter, r *http.Request) {
	// Parse both the base layout and the projects template
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/projects.html",
	))

	// Execute the template with some basic data
	data := struct {
		Title string
	}{
		Title: "Projects",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleProjectsAPI(w http.ResponseWriter, r *http.Request) {
	// Parse the project card template
	tmpl := template.Must(template.ParseFiles(
		"web/templates/components/project_card.html",
	))

	// Get projects from the type system
	projects := types.GetProjects()

	// Render each project card
	for _, project := range projects {
		if err := tmpl.ExecuteTemplate(w, "project-card", project); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
