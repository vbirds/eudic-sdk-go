package dict

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const baseURL = "https://api.frdic.com/api/open/v1"

// Client represents the Eudic API client
type Client struct {
	httpClient *http.Client
	token      string
}

// NewClient creates a new Eudic API client
func NewClient(token string) *Client {
	return &Client{
		httpClient: &http.Client{},
		token:      token,
	}
}

func (c *Client) doRequest(method, path string, body interface{}) (*http.Response, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	var url = baseURL + path
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.token)
	req.Header.Set("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}
