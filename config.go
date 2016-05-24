package alchemyapi

import (
	"net/http"
)

// Config for the AlchemyAPI clients.
type Config struct {
	HTTPClient *http.Client
	APIKey     string
}

// NewConfig with the given access token.
func NewConfig(apiKey string) *Config {
	return &Config{
		HTTPClient: http.DefaultClient,
		APIKey:     apiKey,
	}
}
