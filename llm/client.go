package llm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

// create the generate structs
type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}
type GenerateResponse struct {
	Model      string `json:"model"`
	Created_at string `json:"created_at"`
	Response   string `json:"response"`
	Done       bool   `json:"done"`
}

// a constructor func that creates the client from scratch
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// the generate method: this compiles the needed request to the llm and the response from the llm
func (c *Client) Generate(prompt string) (string, error) {

	//build request struct from the prompt
	reqBody := GenerateRequest{
		Model:  "llama3",
		Prompt: prompt,
		Stream: false,
	}

	//marshal it into json bytes
	dataByte, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("encoding request body: %w", err)
	}
	return "", nil
}
