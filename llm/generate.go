package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const generateEndpoint = "/api/generate"

// the generate method: this compiles the needed request to the llm and the response from the llm
func (c *Client) Generate(ctx context.Context, prompt string, model string) (string, error) {

	//build request struct from the prompt
	reqBody := GenerateRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}

	//marshal it into json bytes
	requestBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("encoding request body: %w", err)
	}

	//construct/build the http request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+generateEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		return "", fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	//send the request and read the response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("retrieving response: %w", err)
	}
	defer resp.Body.Close()

	//drain the entire response from resp.Body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("generate request failed: %d %s", resp.StatusCode, string(responseBody))
	}

	//write from the body into existing variable.
	var genResp GenerateResponse
	if err := json.Unmarshal(responseBody, &genResp); err != nil {
		return "", fmt.Errorf("decoding response body: %w", err)
	}

	return genResp.Response, nil
}
