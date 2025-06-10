package middleware

import (
	"net/http"

	"me.tofaa/internal/pkg/session"
)

// RequireAuth is a middleware that ensures the user is authenticated
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess := session.GetSessionFromRequest(r)
		if !session.IsValid(sess) {
			// Redirect to login page if not authenticated
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}
