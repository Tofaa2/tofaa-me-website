package main

import (
	"fmt"
	"log"
	"net/http"

	"me.tofaa/internal/handlers"
	"me.tofaa/internal/middleware"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve templates
	http.HandleFunc("/", handlers.HandleLandingPage)
	http.HandleFunc("/about", handlers.HandleAboutPage)
	http.HandleFunc("/projects", handlers.HandleProjectsPage)
	http.HandleFunc("/blogs", handlers.BlogsHandler)
	http.HandleFunc("/contact", handlers.HandleContactPage)
	http.HandleFunc("/api/projects", handlers.HandleProjectsAPI)

	// Admin routes
	http.HandleFunc("/admin/login", handlers.HandleAdminLogin)
	http.HandleFunc("/admin/logout", handlers.HandleAdminLogout)
	http.HandleFunc("/admin/dashboard", middleware.RequireAuth(handlers.HandleAdminDashboard))

	fmt.Println("Server starting on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
