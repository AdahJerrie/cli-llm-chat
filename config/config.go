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
func LoadConfig() (Config, error) {
	if os.Getenv("BASE_URL") == "" {
		return Config{}, errors.New("BASE_URL environment variable is required")
	}
	if os.Getenv("MODEL") == "" {
		return Config{}, errors.New("MODEL environment variable is required")
	}

	return Config{
		BaseURL: os.Getenv("BASE_URL"),
		Model:   os.Getenv("MODEL"),
		APIKey:  os.Getenv("API_KEY"),
	}, nil
}
