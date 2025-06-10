package session

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
	"time"
)

// Session represents a user session
type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// Manager handles session management
type Manager struct {
	sessions map[string]*Session
	mu       sync.RWMutex
	// In a production environment, you'd want to use a more persistent storage
	// like Redis or a database
}

var (
	manager = NewManager()
)

// NewManager creates a new session manager
func NewManager() *Manager {
	return &Manager{
		sessions: make(map[string]*Session),
	}
}

// CreateSession creates a new session for a user
func CreateSession(userID string) (*Session, error) {
	// Generate a random session ID
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	sessionID := base64.URLEncoding.EncodeToString(b)

	// Create session
	session := &Session{
		ID:        sessionID,
		UserID:    userID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour), // Sessions expire after 24 hours
	}

	// Store session
	manager.mu.Lock()
	manager.sessions[sessionID] = session
	manager.mu.Unlock()

	return session, nil
}

// GetSession retrieves a session by ID
func GetSession(sessionID string) *Session {
	manager.mu.RLock()
	defer manager.mu.RUnlock()
	return manager.sessions[sessionID]
}

// DeleteSession removes a session
func DeleteSession(sessionID string) {
	manager.mu.Lock()
	delete(manager.sessions, sessionID)
	manager.mu.Unlock()
}

// IsValid checks if a session is valid
func IsValid(session *Session) bool {
	if session == nil {
		return false
	}
	return time.Now().Before(session.ExpiresAt)
}

// GetSessionFromRequest gets the session from a request's cookies
func GetSessionFromRequest(r *http.Request) *Session {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil
	}
	return GetSession(cookie.Value)
}

// SetSessionCookie sets the session cookie in the response
func SetSessionCookie(w http.ResponseWriter, session *Session) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // Only send cookie over HTTPS
		SameSite: http.SameSiteStrictMode,
		Expires:  session.ExpiresAt,
	})
}

// ClearSessionCookie removes the session cookie
func ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})
}
