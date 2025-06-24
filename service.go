package mt5api

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// Company represents search result
type Company struct {
	CompanyName string   `json:"companyName"`
	Results     []Result `json:"results"`
}

// Result represents broker result
type Result struct {
	Name   string   `json:"name"`
	Access []string `json:"access"`
}

// PingHost pings a host and returns response time in milliseconds
func (c *Client) PingHost(ctx context.Context, host string, port int) (int, error) {
	params := url.Values{}
	params.Add("host", host)
	if port > 0 {
		params.Add("port", strconv.Itoa(port))
	}

	body, err := c.doRequest(ctx, "GET", "/PingHost", params)
	if err != nil {
		return 0, err
	}

	var pingTime int
	if err := json.Unmarshal(body, &pingTime); err != nil {
		return 0, err
	}

	return pingTime, nil
}

// Search searches for broker by company name
func (c *Client) Search(ctx context.Context, company string) ([]Company, error) {
	params := url.Values{}
	params.Add("company", company)

	body, err := c.doRequest(ctx, "GET", "/Search", params)
	if err != nil {
		return nil, err
	}

	var companies []Company
	if err := json.Unmarshal(body, &companies); err != nil {
		return nil, err
	}

	return companies, nil
}

func (c *Client) ServerTimezone(ctx context.Context) (int, error) {
	body, err := c.doRequest(ctx, "GET", "/ServerTimezone", url.Values{})
	if err != nil {
		return 0, err
	}
	timeZone, err := strconv.ParseFloat(string(body), 64)
	if err != nil {
		return 0, err
	}
	return int(timeZone), nil
}
