package mt5api

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

type Order struct {
	Ticket            int64      `json:"ticket"`
	Profit            float64    `json:"profit"`
	Swap              float64    `json:"swap"`
	Commission        float64    `json:"commission"`
	Fee               float64    `json:"fee"`
	ClosePrice        float64    `json:"closePrice"`
	CloseTimestampUTC uint64     `json:"closeTimestampUTC"`
	CloseLots         float64    `json:"closeLots"`
	CloseComment      string     `json:"closeComment"`
	OpenPrice         float64    `json:"openPrice"`
	OpenTimestampUTC  uint64     `json:"openTimestampUTC"`
	Lots              float64    `json:"lots"`
	ContractSize      float64    `json:"contractSize"`
	ExpertId          int64      `json:"expertId"`
	PlacedType        PlacedType `json:"placedType"`
	OrderType         OrderType  `json:"orderType"`
	Symbol            string     `json:"symbol"`
	Comment           string     `json:"comment"`
	State             OrderState `json:"state"`
	StopLoss          float64    `json:"stopLoss"`
	TakeProfit        float64    `json:"takeProfit"`
	RequestId         int32      `json:"requestId"`
	Digits            int32      `json:"digits"`
	ProfitRate        float64    `json:"profitRate"`
	StopLimitPrice    float64    `json:"stopLimitPrice"`
}

// SortType for ordering results
type SortType string

const (
	SortByOpenTime  SortType = "OpenTime"
	SortByCloseTime SortType = "CloseTime"
)

// OpenedOrders gets list of opened orders
func (c *Client) OpenedOrders(ctx context.Context, sort SortType, ascending bool) ([]Order, error) {
	params := url.Values{}
	if sort != "" {
		params.Add("sort", string(sort))
	}
	params.Add("ascending", strconv.FormatBool(ascending))

	body, err := c.doRequest(ctx, "GET", "/OpenedOrders", params)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

// OpenedOrder gets opened order by ticket
func (c *Client) OpenedOrder(ctx context.Context, ticket int64) (*Order, error) {
	params := url.Values{}
	params.Add("ticket", strconv.FormatInt(ticket, 10))

	body, err := c.doRequest(ctx, "GET", "/OpenedOrder", params)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
