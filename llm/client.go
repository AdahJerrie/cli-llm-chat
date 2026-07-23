// Package llm provides a client for communicating with an LLM server.
package llm

import (
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

// 1.  a constructor func that creates the client from scratch
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{ /*Timeout: 15 * time.Second,*/ },
	}
}
