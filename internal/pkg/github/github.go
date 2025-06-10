package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// GitHubUser represents a GitHub user's data
type GitHubUser struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

// CacheEntry represents a cached avatar URL with expiration
type CacheEntry struct {
	URL       string
	ExpiresAt time.Time
}

// Cache is a thread-safe cache for GitHub avatars
type Cache struct {
	entries map[string]CacheEntry
	mu      sync.RWMutex
	ttl     time.Duration
}

var (
	// Default cache with 1 hour TTL
	defaultCache = NewCache(1 * time.Hour)
)

// NewCache creates a new cache with the specified TTL
func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		entries: make(map[string]CacheEntry),
		ttl:     ttl,
	}
}

// GetAvatarURL fetches a GitHub user's avatar URL with caching
func GetAvatarURL(username string) (string, error) {
	// Try to get from cache first
	if url, ok := defaultCache.get(username); ok {
		return url, nil
	}

	// If not in cache, fetch from GitHub
	url, err := fetchAvatarURL(username)
	if err != nil {
		return "", err
	}

	// Store in cache
	defaultCache.set(username, url)
	return url, nil
}

// get retrieves a URL from the cache if it exists and is not expired
func (c *Cache) get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.entries[key]
	if !exists {
		return "", false
	}

	if time.Now().After(entry.ExpiresAt) {
		// Entry has expired, remove it
		c.mu.RUnlock()
		c.mu.Lock()
		delete(c.entries, key)
		c.mu.Unlock()
		c.mu.RLock()
		return "", false
	}

	return entry.URL, true
}

// set stores a URL in the cache with the configured TTL
func (c *Cache) set(key, url string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = CacheEntry{
		URL:       url,
		ExpiresAt: time.Now().Add(c.ttl),
	}
}

// fetchAvatarURL fetches the avatar URL from GitHub
func fetchAvatarURL(username string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	// Create a new request with proper headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add required headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch GitHub user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("github API returned status: %d", resp.StatusCode)
	}

	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", fmt.Errorf("failed to decode GitHub user data: %w", err)
	}

	return user.AvatarURL, nil
}
