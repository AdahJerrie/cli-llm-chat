package llm

import (
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

// a constructor that creates the client from scratch
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}
