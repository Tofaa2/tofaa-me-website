package discord

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// DiscordUser represents a Discord user's data
type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
}

// CacheEntry represents a cached avatar URL with expiration
type CacheEntry struct {
	URL       string
	ExpiresAt time.Time
}

// Cache is a thread-safe cache for Discord avatars
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

// GetAvatarURL fetches a Discord user's avatar URL with caching
func GetAvatarURL(identifier string) (string, error) {
	// Try to get from cache first
	if url, ok := defaultCache.get(identifier); ok {
		return url, nil
	}

	// If not in cache, fetch from Discord
	url, err := fetchAvatarURL(identifier)
	if err != nil {
		return "", err
	}

	// Store in cache
	defaultCache.set(identifier, url)
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

// fetchAvatarURL fetches the avatar URL from Discord
func fetchAvatarURL(identifier string) (string, error) {
	// Check if the identifier is a Discord ID (17-19 digits)
	if len(identifier) >= 17 && len(identifier) <= 19 {
		return getAvatarByID(identifier)
	}

	// Otherwise, treat it as a username
	return getAvatarByUsername(identifier)
}

// getAvatarByID fetches avatar URL using Discord ID
func getAvatarByID(id string) (string, error) {
	url := fmt.Sprintf("https://discord.com/api/v10/users/%s", id)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch Discord user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("discord API returned status: %d", resp.StatusCode)
	}

	var user DiscordUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", fmt.Errorf("failed to decode Discord user data: %w", err)
	}

	return formatAvatarURL(user.ID, user.Avatar), nil
}

// getAvatarByUsername fetches avatar URL using username
// Note: This is less reliable as usernames can change
func getAvatarByUsername(username string) (string, error) {
	// First, we need to get the user's ID
	// This requires a bot token and proper authentication
	// For now, we'll return an error
	return "", fmt.Errorf("fetching by username requires bot authentication")
}

// formatAvatarURL formats the Discord avatar URL
func formatAvatarURL(userID, avatarHash string) string {
	if avatarHash == "" {
		// Return default Discord avatar
		return fmt.Sprintf("https://cdn.discordapp.com/embed/avatars/%d.png", 0)
	}

	// Check if the avatar is animated (starts with 'a_')
	format := "png"
	if len(avatarHash) > 2 && avatarHash[:2] == "a_" {
		format = "gif"
	}

	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.%s", userID, avatarHash, format)
}
