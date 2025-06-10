package handlers

import (
	"net/http"
	"text/template"
)

func HandleContactPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Handle form submission
		name := r.FormValue("name")
		email := r.FormValue("email")
		subject := r.FormValue("subject")
		message := r.FormValue("message")

		// TODO: Add your email sending logic here
		// For now, we'll just print the values
		println("Received contact form submission:")
		println("Name:", name)
		println("Email:", email)
		println("Subject:", subject)
		println("Message:", message)

		// Redirect to thank you page or show success message
		http.Redirect(w, r, "/contact?success=true", http.StatusSeeOther)
		return
	}

	// Parse both the base layout and the contact template
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/contact.html",
	))

	// Execute the template with some basic data
	data := struct {
		Title string
	}{
		Title: "Contact",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
