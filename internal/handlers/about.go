package handlers

import (
	"log"
	"net/http"
	"text/template"

	"me.tofaa/internal/pkg/github"
)

func HandleAboutPage(w http.ResponseWriter, r *http.Request) {
	// Get GitHub avatar URL
	avatarURL, err := github.GetAvatarURL("Tofaa2")
	if err != nil {
		// If there's an error, use a default avatar
		avatarURL = "/static/images/default-avatar.svg"
		log.Println(err)
	}

	// Parse both the base layout and the about template
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/about.html",
	))

	// Execute the template with the avatar URL
	data := struct {
		Title     string
		AvatarURL string
	}{
		Title:     "About",
		AvatarURL: avatarURL,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
