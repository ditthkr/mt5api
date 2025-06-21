package mt5api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client represents the MT5 API client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string // Session token from Connect
}

// NewClient creates a new MT5 API client
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: strings.TrimSuffix(baseURL, "/"),
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SetToken sets the session token for authenticated requests
func (c *Client) SetToken(token string) {
	c.Token = token
}

// doRequest performs HTTP request with common error handling
func (c *Client) doRequest(ctx context.Context, method, endpoint string, params url.Values) ([]byte, error) {
	// Add token to params if available and not already present
	if c.Token != "" && params.Get("id") == "" {
		params.Set("id", c.Token)
	}

	var reqURL string
	if method == "GET" && len(params) > 0 {
		reqURL = fmt.Sprintf("%s%s?%s", c.BaseURL, endpoint, params.Encode())
	} else {
		reqURL = fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	// Handle API errors (status code 201 indicates exception)
	if resp.StatusCode == 201 {
		var apiErr ExceptionResult
		if err := json.Unmarshal(body, &apiErr); err == nil {
			return nil, fmt.Errorf("API error [%s]: %s", apiErr.Code, apiErr.Message)
		}
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
