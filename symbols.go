package mt5api

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// SymbolInfo represents symbol information
type SymbolInfo struct {
	UpdateTime      int64   `json:"updateTime"`
	Currency        string  `json:"currency"`
	ISIN            string  `json:"isin"`
	Description     string  `json:"description"`
	Basis           string  `json:"basis"`
	RefToSite       string  `json:"refToSite"`
	Custom          int32   `json:"custom"`
	ProfitCurrency  string  `json:"profitCurrency"`
	MarginCurrency  string  `json:"marginCurrency"`
	Precision       int32   `json:"precision"`
	BkgndColor      int32   `json:"bkgndColor"`
	Digits          int32   `json:"digits"`
	Points          float64 `json:"points"`
	LimitPoints     float64 `json:"limitPoints"`
	Id              int32   `json:"id"`
	DepthOfMarket   int32   `json:"depthOfMarket"`
	Spread          int32   `json:"spread"`
	TickValue       float64 `json:"tickValue"`
	TickSize        float64 `json:"tickSize"`
	ContractSize    float64 `json:"contractSize"`
	SettlementPrice float64 `json:"settlementPrice"`
	LowerLimit      float64 `json:"lowerLimit"`
	UpperLimit      float64 `json:"upperLimit"`
	FaceValue       float64 `json:"faceValue"`
	AccruedInterest float64 `json:"accruedInterest"`
	FirstTradeTime  int64   `json:"firstTradeTime"`
	LastTradeTime   int64   `json:"lastTradeTime"`
	BidTickValue    float64 `json:"bid_tickvalue"`
	AskTickValue    float64 `json:"ask_tickvalue"`
}

// SymGroup represents symbol group information
type SymGroup struct {
	GroupName         string  `json:"groupName"`
	DeviationRate     int32   `json:"deviationRate"`
	RoundRate         int32   `json:"roundRate"`
	TradeMode         string  `json:"tradeMode"`
	SL                int32   `json:"sl"`
	TP                int32   `json:"tp"`
	TradeType         string  `json:"tradeType"`
	FillPolicy        string  `json:"fillPolicy"`
	Expiration        string  `json:"expiration"`
	OrderFlags        int32   `json:"orderFlags"`
	PriceTimeout      int32   `json:"priceTimeout"`
	RequoteTimeout    int32   `json:"requoteTimeout"`
	RequestLots       int32   `json:"requestLots"`
	MinVolume         int64   `json:"minVolume"`
	MaxVolume         int64   `json:"maxVolume"`
	VolumeStep        int64   `json:"volumeStep"`
	InitialMargin     float64 `json:"initialMargin"`
	MaintenanceMargin float64 `json:"maintenanceMargin"`
	HedgedMargin      float64 `json:"hedgedMargin"`
	SwapType          string  `json:"swapType"`
	SwapLong          float64 `json:"swapLong"`
	SwapShort         float64 `json:"swapShort"`
	ThreeDaysSwap     string  `json:"threeDaysSwap"`
	MinLots           float64 `json:"minLots"`
	MaxLots           float64 `json:"maxLots"`
	LotsStep          float64 `json:"lotsStep"`
}

// SymbolParams represents symbol parameters
type SymbolParams struct {
	Symbol      string     `json:"symbol"`
	SymbolInfo  SymbolInfo `json:"symbolInfo"`
	SymbolGroup SymGroup   `json:"symbolGroup"`
}

// Quote represents a price quote
type Quote struct {
	Symbol       string  `json:"symbol"`
	Bid          float64 `json:"bid"`
	Ask          float64 `json:"ask"`
	TimestampUTC uint64  `json:"timestampUTC"`
	Last         float64 `json:"last"`
	Volume       int64   `json:"volume"`
}

// SessionState represents session state for symbol
type SessionState struct {
	Symbol string `json:"symbol"`
	Active bool   `json:"active"`
}

// Symbols gets list of available symbols with parameters
func (c *Client) Symbols(ctx context.Context) (map[string]SymbolInfo, error) {
	body, err := c.doRequest(ctx, "GET", "/Symbols", url.Values{})
	if err != nil {
		return nil, err
	}

	var symbols map[string]SymbolInfo
	if err := json.Unmarshal(body, &symbols); err != nil {
		return nil, err
	}

	return symbols, nil
}

// SymbolList gets list of available symbol names
func (c *Client) SymbolList(ctx context.Context) ([]string, error) {
	body, err := c.doRequest(ctx, "GET", "/SymbolList", url.Values{})
	if err != nil {
		return nil, err
	}

	var symbols []string
	if err := json.Unmarshal(body, &symbols); err != nil {
		return nil, err
	}

	return symbols, nil
}

// GetQuote gets latest quote for the specified symbol
func (c *Client) GetQuote(ctx context.Context, symbol string, msNotOlder int) (*Quote, error) {
	params := url.Values{}
	params.Add("symbol", symbol)
	if msNotOlder > 0 {
		params.Add("msNotOlder", strconv.Itoa(msNotOlder))
	}

	body, err := c.doRequest(ctx, "GET", "/GetQuote", params)
	if err != nil {
		return nil, err
	}

	var quote Quote
	if err := json.Unmarshal(body, &quote); err != nil {
		return nil, err
	}

	return &quote, nil
}

// GetQuoteMany gets latest quotes for multiple symbols
func (c *Client) GetQuoteMany(ctx context.Context, symbols []string, msNotOlder int) ([]Quote, error) {
	params := url.Values{}
	for _, symbol := range symbols {
		params.Add("symbols", symbol)
	}
	if msNotOlder > 0 {
		params.Add("msNotOlder", strconv.Itoa(msNotOlder))
	}

	body, err := c.doRequest(ctx, "GET", "/GetQuoteMany", params)
	if err != nil {
		return nil, err
	}

	var quotes []Quote
	if err := json.Unmarshal(body, &quotes); err != nil {
		return nil, err
	}

	return quotes, nil
}

// SymbolParams gets full information about symbol and its group
func (c *Client) SymbolParams(ctx context.Context, symbol string) (*SymbolParams, error) {
	params := url.Values{}
	params.Add("symbol", symbol)

	body, err := c.doRequest(ctx, "GET", "/SymbolParams", params)
	if err != nil {
		return nil, err
	}

	var symbolParams SymbolParams
	if err := json.Unmarshal(body, &symbolParams); err != nil {
		return nil, err
	}

	return &symbolParams, nil
}

// IsTradeSession checks if market is open for specified symbol
func (c *Client) IsTradeSession(ctx context.Context, symbol string) (bool, error) {
	params := url.Values{}
	params.Add("symbol", symbol)

	body, err := c.doRequest(ctx, "GET", "/IsTradeSession", params)
	if err != nil {
		return false, err
	}

	var isOpen bool
	if err := json.Unmarshal(body, &isOpen); err != nil {
		return false, err
	}

	return isOpen, nil
}

// RequiredMargin calculates required margin for a trade
func (c *Client) RequiredMargin(ctx context.Context, symbol string, lots float64, orderType OrderType, price float64) (float64, error) {
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("lots", strconv.FormatFloat(lots, 'f', -1, 64))
	params.Add("type", string(orderType))
	if price > 0 {
		params.Add("price", strconv.FormatFloat(price, 'f', -1, 64))
	}

	body, err := c.doRequest(ctx, "GET", "/RequiredMargin", params)
	if err != nil {
		return 0, err
	}

	var margin float64
	if err := json.Unmarshal(body, &margin); err != nil {
		return 0, err
	}

	return margin, nil
}
