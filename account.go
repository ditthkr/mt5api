package mt5api

import (
	"context"
	"encoding/json"
	"net/url"
)

// AccountRec represents account details
type AccountRec struct {
	Login       int64   `json:"login"`
	Type        string  `json:"type"`
	UserName    string  `json:"userName"`
	TradeFlags  int32   `json:"tradeFlags"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	ZipCode     string  `json:"zipCode"`
	UserAddress string  `json:"userAddress"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	Balance     float64 `json:"balance"`
	Credit      float64 `json:"credit"`
	Blocked     float64 `json:"blocked"`
	Leverage    int32   `json:"leverage"`
}

// AccountSummary represents account trading summary
type AccountSummary struct {
	Balance     float64       `json:"balance"`
	Credit      float64       `json:"credit"`
	Profit      float64       `json:"profit"`
	Equity      float64       `json:"equity"`
	Margin      float64       `json:"margin"`
	FreeMargin  float64       `json:"freeMargin"`
	MarginLevel float64       `json:"marginLevel"`
	Leverage    float64       `json:"leverage"`
	Currency    string        `json:"currency"`
	Method      AccountMethod `json:"method"`
	Type        string        `json:"type"`
	IsInvestor  bool          `json:"isInvestor"`
}

// AccountDetails represents detailed account information
type AccountDetails struct {
	ServerName string `json:"serverName"`
	User       int64  `json:"user"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       int32  `json:"port"`
	//ServerTime      time.Time     `json:"serverTime"`
	//ServerTimeZone  int32         `json:"serverTimeZone"`
	Company         string        `json:"company"`
	Currency        string        `json:"currency"`
	AccountName     string        `json:"accountName"`
	Group           string        `json:"group"`
	AccountType     string        `json:"accountType"`
	AccountLeverage int32         `json:"accountLeverage"`
	AccountMethod   AccountMethod `json:"accountMethod"`
	IsInvestor      bool          `json:"isInvestor"`
}

// Account gets account details
func (c *Client) Account(ctx context.Context) (*AccountRec, error) {
	body, err := c.doRequest(ctx, "GET", "/Account", url.Values{})
	if err != nil {
		return nil, err
	}

	var account AccountRec
	if err := json.Unmarshal(body, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

// AccountSummary gets account trading summary
func (c *Client) AccountSummary(ctx context.Context) (*AccountSummary, error) {
	body, err := c.doRequest(ctx, "GET", "/AccountSummary", url.Values{})
	if err != nil {
		return nil, err
	}

	var summary AccountSummary
	if err := json.Unmarshal(body, &summary); err != nil {
		return nil, err
	}

	return &summary, nil
}

// AccountDetails gets detailed account information
func (c *Client) AccountDetails(ctx context.Context) (*AccountDetails, error) {
	body, err := c.doRequest(ctx, "GET", "/AccountDetails", url.Values{})
	if err != nil {
		return nil, err
	}

	var details AccountDetails
	if err := json.Unmarshal(body, &details); err != nil {
		return nil, err
	}

	return &details, nil
}
