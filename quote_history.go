package mt5api

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

// Bar represents OHLC data
type Bar struct {
	Time       string  `json:"time"`
	OpenPrice  float64 `json:"openPrice"`
	HighPrice  float64 `json:"highPrice"`
	LowPrice   float64 `json:"lowPrice"`
	ClosePrice float64 `json:"closePrice"`
	TickVolume int64   `json:"tickVolume"`
	Spread     int32   `json:"spread"`
	Volume     int64   `json:"volume"`
}

// BarsForSymbol represents bars for a specific symbol
type BarsForSymbol struct {
	Symbol    string `json:"symbol"`
	Bars      []Bar  `json:"bars"`
	Exception string `json:"exception"`
}

// QuoteHistoryEventArgs represents quote history event
type QuoteHistoryEventArgs struct {
	Symbol string `json:"symbol"`
	Bars   []Bar  `json:"bars"`
}

// PriceHistory gets price history for a date range
func (c *Client) PriceHistory(ctx context.Context, symbol string, from, to time.Time, timeFrame int) ([]Bar, error) {
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("from", from.Format("2006-01-02T15:04:05"))
	params.Add("to", to.Format("2006-01-02T15:04:05"))
	params.Add("timeFrame", strconv.Itoa(timeFrame))

	body, err := c.doRequest(ctx, "GET", "/PriceHistory", params)
	if err != nil {
		return nil, err
	}

	var bars []Bar
	if err := json.Unmarshal(body, &bars); err != nil {
		return nil, err
	}

	return bars, nil
}

// PriceHistoryMany gets price history for multiple symbols
func (c *Client) PriceHistoryMany(ctx context.Context, symbols []string, from, to time.Time, timeFrame int) ([]BarsForSymbol, error) {
	params := url.Values{}
	for _, symbol := range symbols {
		params.Add("symbol", symbol)
	}
	params.Add("from", from.Format("2006-01-02T15:04:05"))
	params.Add("to", to.Format("2006-01-02T15:04:05"))
	params.Add("timeFrame", strconv.Itoa(timeFrame))

	body, err := c.doRequest(ctx, "GET", "/PriceHistoryMany", params)
	if err != nil {
		return nil, err
	}

	var barsForSymbols []BarsForSymbol
	if err := json.Unmarshal(body, &barsForSymbols); err != nil {
		return nil, err
	}

	return barsForSymbols, nil
}

// PriceHistoryToday gets price history for today
func (c *Client) PriceHistoryToday(ctx context.Context, symbol string, timeFrame int) ([]Bar, error) {
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("timeFrame", strconv.Itoa(timeFrame))

	body, err := c.doRequest(ctx, "GET", "/PriceHistoryToday", params)
	if err != nil {
		return nil, err
	}

	var bars []Bar
	if err := json.Unmarshal(body, &bars); err != nil {
		return nil, err
	}

	return bars, nil
}

// PriceHistoryMonth gets price history for 30 days
func (c *Client) PriceHistoryMonth(ctx context.Context, symbol string, year, month, day, timeFrame int) ([]Bar, error) {
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("year", strconv.Itoa(year))
	params.Add("month", strconv.Itoa(month))
	params.Add("day", strconv.Itoa(day))
	params.Add("timeFrame", strconv.Itoa(timeFrame))

	body, err := c.doRequest(ctx, "GET", "/PriceHistoryMonth", params)
	if err != nil {
		return nil, err
	}

	var bars []Bar
	if err := json.Unmarshal(body, &bars); err != nil {
		return nil, err
	}

	return bars, nil
}

// PriceHistoryEx gets price history from specified date for several bars back
func (c *Client) PriceHistoryEx(ctx context.Context, symbol string, from time.Time, numBars, timeFrame int) ([]Bar, error) {
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("from", from.Format("2006-01-02T15:04:05"))
	params.Add("numBars", strconv.Itoa(numBars))
	params.Add("timeFrame", strconv.Itoa(timeFrame))

	body, err := c.doRequest(ctx, "GET", "/PriceHistoryEx", params)
	if err != nil {
		return nil, err
	}

	var bars []Bar
	if err := json.Unmarshal(body, &bars); err != nil {
		return nil, err
	}

	return bars, nil
}
