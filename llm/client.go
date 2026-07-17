package llm

import (
	"bytes"
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

	//construct/build the http request
	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/api/generate", bytes.NewReader(dataByte))
	if err != nil {
		return "", fmt.Errorf("buiding request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	//send the request and read the response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "",  fmt.Errorf("Retrieving response: %w", err)
	}
	defer resp.Body.Close()

	return, nil
}
