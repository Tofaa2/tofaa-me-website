package handlers

import (
	"html/template"
	"net/http"

	"me.tofaa/internal/pkg/session"
)

// Admin credentials (in production, these should be stored securely)
const (
	adminUsername = "admin"
	adminPassword = "your-secure-password-here" // Change this!
)

// HandleAdminLogin handles the admin login page and form submission
func HandleAdminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == adminUsername && password == adminPassword {
			// Create session
			sess, err := session.CreateSession(username)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Set session cookie
			session.SetSessionCookie(w, sess)

			// Redirect to dashboard
			http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			return
		}

		// Invalid credentials
		data := struct {
			Error string
		}{
			Error: "Invalid username or password",
		}
		tmpl := template.Must(template.ParseFiles(
			"web/templates/layouts/base.html",
			"web/templates/admin/login.html",
		))
		tmpl.Execute(w, data)
		return
	}

	// Show login page
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/admin/login.html",
	))
	tmpl.Execute(w, nil)
}

// HandleAdminDashboard handles the admin dashboard page
func HandleAdminDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/admin/dashboard.html",
	))
	tmpl.Execute(w, nil)
}

// HandleAdminLogout handles the admin logout
func HandleAdminLogout(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSessionFromRequest(r)
	if sess != nil {
		session.DeleteSession(sess.ID)
	}
	session.ClearSessionCookie(w)
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}
