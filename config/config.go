package config

import (
	"os"
)

// config struct defines the configuration for the application.
type Config struct {
	BaseURL string
	Model   string
	APIKey  string
}

const defaultBaseURL, defaultModel string = "http://localhost:11434", "llama3"

// build/load the configuration from environment variables or a configuration file.
func Load() (*Config, error) {
	baseURL := os.Getenv("BASE_URL")
	model := os.Getenv("MODEL")

	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	if model == "" {
		model = defaultModel
	}

	return &Config{
		BaseURL: baseURL,
		Model:   model,
		APIKey:  os.Getenv("API_KEY"),
	}, nil
}
