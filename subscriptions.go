package mt5api

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// Subscribe subscribes to real-time quotes for a symbol
func (c *Client) Subscribe(ctx context.Context, symbol string, interval int) (string, error) {
	params := url.Values{}
	params.Add("symbol", symbol)
	if interval > 0 {
		params.Add("interval", strconv.Itoa(interval))
	}

	body, err := c.doRequest(ctx, "GET", "/Subscribe", params)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// SubscribeMany subscribes to real-time quotes for multiple symbols
func (c *Client) SubscribeMany(ctx context.Context, symbols []string, interval int) (string, error) {
	params := url.Values{}
	for _, symbol := range symbols {
		params.Add("symbols", symbol)
	}
	if interval > 0 {
		params.Add("interval", strconv.Itoa(interval))
	}

	body, err := c.doRequest(ctx, "GET", "/SubscribeMany", params)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// UnSubscribe unsubscribes from real-time quotes for a symbol
func (c *Client) UnSubscribe(ctx context.Context, symbol string) (string, error) {
	params := url.Values{}
	params.Add("symbol", symbol)

	body, err := c.doRequest(ctx, "GET", "/UnSubscribe", params)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// UnSubscribeMany unsubscribes from real-time quotes for multiple symbols
func (c *Client) UnSubscribeMany(ctx context.Context, symbols []string) (string, error) {
	params := url.Values{}
	for _, symbol := range symbols {
		params.Add("symbols", symbol)
	}

	body, err := c.doRequest(ctx, "GET", "/UnSubscribeMany", params)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// SubscribeOrderProfit subscribes for order profit updates
func (c *Client) SubscribeOrderProfit(ctx context.Context, interval int) ([]Order, error) {
	params := url.Values{}
	if interval > 0 {
		params.Add("interval", strconv.Itoa(interval))
	}

	body, err := c.doRequest(ctx, "GET", "/SubscribeOrderProfit", params)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

// SubscribeOhlc subscribes to OHLC price updates for symbol
func (c *Client) SubscribeOHLC(ctx context.Context, symbol string, timeframe, interval int) (string, error) {
	params := url.Values{}
	if symbol != "" {
		params.Add("symbol", symbol)
	}
	if timeframe > 0 {
		params.Add("timeframe", strconv.Itoa(timeframe))
	}
	if interval > 0 {
		params.Add("interval", strconv.Itoa(interval))
	}

	body, err := c.doRequest(ctx, "GET", "/SubscribeOhlc", params)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// UnsubscribeOhlc unsubscribes from OHLC updates
func (c *Client) UnsubscribeOHLC(ctx context.Context, symbol string, timeframe int) (string, error) {
	params := url.Values{}
	if symbol != "" {
		params.Add("symbol", symbol)
	}
	if timeframe > 0 {
		params.Add("timeframe", strconv.Itoa(timeframe))
	}

	body, err := c.doRequest(ctx, "GET", "/UnsubscribeOhlc", params)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
