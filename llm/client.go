package llm

import (
	"encoding/json"
	"net/http"
)
// struct for request body
type Client struct {
	BaseURL string `json:"baseurl"`
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
	httpCLient *http.Client 
}

func NewClient(c *Client) (string, error) {
	bytedata, error := json.Marshal(c.Prompt)
}
