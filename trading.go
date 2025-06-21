package mt5api

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// OrderSendRequest represents order send parameters
type OrderSendRequest struct {
	Symbol         string     `json:"symbol"`
	Operation      OrderType  `json:"operation"`
	Volume         float64    `json:"volume"`
	Price          float64    `json:"price,omitempty"`
	Slippage       int64      `json:"slippage,omitempty"`
	StopLoss       float64    `json:"stoploss,omitempty"`
	TakeProfit     float64    `json:"takeprofit,omitempty"`
	Comment        string     `json:"comment,omitempty"`
	ExpertId       int64      `json:"expertId,omitempty"`
	StopLimitPrice float64    `json:"stopLimitPrice,omitempty"`
	PlacedType     PlacedType `json:"placedType,omitempty"`
}

// OrderModifyRequest represents order modify parameters
type OrderModifyRequest struct {
	Ticket     int64   `json:"ticket"`
	StopLoss   float64 `json:"stoploss"`
	TakeProfit float64 `json:"takeprofit"`
	Price      float64 `json:"price,omitempty"`
	StopLimit  float64 `json:"stoplimit,omitempty"`
}

// OrderCloseRequest represents order close parameters
type OrderCloseRequest struct {
	Ticket   int64   `json:"ticket"`
	Lots     float64 `json:"lots,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Slippage int64   `json:"slippage,omitempty"`
}

// OrderSend sends market or pending order
func (c *Client) OrderSend(ctx context.Context, req OrderSendRequest) (*Order, error) {
	params := url.Values{}
	params.Add("symbol", req.Symbol)
	params.Add("operation", string(req.Operation))
	params.Add("volume", strconv.FormatFloat(req.Volume, 'f', -1, 64))

	if req.Price > 0 {
		params.Add("price", strconv.FormatFloat(req.Price, 'f', -1, 64))
	}
	if req.Slippage > 0 {
		params.Add("slippage", strconv.FormatInt(req.Slippage, 10))
	}
	if req.StopLoss > 0 {
		params.Add("stoploss", strconv.FormatFloat(req.StopLoss, 'f', -1, 64))
	}
	if req.TakeProfit > 0 {
		params.Add("takeprofit", strconv.FormatFloat(req.TakeProfit, 'f', -1, 64))
	}
	if req.Comment != "" {
		params.Add("comment", req.Comment)
	}
	if req.ExpertId > 0 {
		params.Add("expertId", strconv.FormatInt(req.ExpertId, 10))
	}
	if req.StopLimitPrice > 0 {
		params.Add("stopLimitPrice", strconv.FormatFloat(req.StopLimitPrice, 'f', -1, 64))
	}
	if req.PlacedType != "" {
		params.Add("placedType", string(req.PlacedType))
	}

	body, err := c.doRequest(ctx, "GET", "/OrderSend", params)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

// OrderModify modifies market or pending order
func (c *Client) OrderModify(ctx context.Context, req OrderModifyRequest) (*Order, error) {
	params := url.Values{}
	params.Add("ticket", strconv.FormatInt(req.Ticket, 10))
	params.Add("stoploss", strconv.FormatFloat(req.StopLoss, 'f', -1, 64))
	params.Add("takeprofit", strconv.FormatFloat(req.TakeProfit, 'f', -1, 64))

	if req.Price > 0 {
		params.Add("price", strconv.FormatFloat(req.Price, 'f', -1, 64))
	}
	if req.StopLimit > 0 {
		params.Add("stoplimit", strconv.FormatFloat(req.StopLimit, 'f', -1, 64))
	}

	body, err := c.doRequest(ctx, "GET", "/OrderModify", params)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

// OrderClose closes market or pending order
func (c *Client) OrderClose(ctx context.Context, req OrderCloseRequest) (*Order, error) {
	params := url.Values{}
	params.Add("ticket", strconv.FormatInt(req.Ticket, 10))

	if req.Lots > 0 {
		params.Add("lots", strconv.FormatFloat(req.Lots, 'f', -1, 64))
	}
	if req.Price > 0 {
		params.Add("price", strconv.FormatFloat(req.Price, 'f', -1, 64))
	}
	if req.Slippage > 0 {
		params.Add("slippage", strconv.FormatInt(req.Slippage, 10))
	}

	body, err := c.doRequest(ctx, "GET", "/OrderClose", params)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
