package config

import (
	"errors"
	"os"
)

// config struct defines the configuration for the application.
type Config struct {
	BaseURL string
	Model   string
	APIKey  string
}

// build/load the configuration from environment variables or a configuration file.
func LoadConfig() (*Config, error) {
	baseURL := os.Getenv("BASE_URL")
	model := os.Getenv("MODEL")

	if baseURL == "" {
		return nil, errors.New("BASE_URL environment variable is required")
	}
	if model == "" {
		return nil, errors.New("MODEL environment variable is required")
	}

	return &Config{
		BaseURL: baseURL,
		Model:   model,
		APIKey:  os.Getenv("API_KEY"),
	}, nil
}
