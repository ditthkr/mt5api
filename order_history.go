package mt5api

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

// OrderHistoryEventArgs represents order history response
type OrderHistoryEventArgs struct {
	Orders          []Order         `json:"orders"`
	InternalDeals   []DealInternal  `json:"internalDeals"`
	InternalOrders  []OrderInternal `json:"internalOrders"`
	Action          int32           `json:"action"`
	PartialResponse bool            `json:"partialResponse"`
}

// DealInternal represents internal deal information
type DealInternal struct {
	TicketNumber   int64      `json:"ticketNumber"`
	Id             string     `json:"id"`
	Login          int64      `json:"login"`
	HistoryTime    int64      `json:"historyTime"`
	OrderTicket    int64      `json:"orderTicket"`
	OpenTime       int64      `json:"openTime"`
	Symbol         string     `json:"symbol"`
	Type           string     `json:"type"`
	Direction      string     `json:"direction"`
	OpenPrice      float64    `json:"openPrice"`
	Price          float64    `json:"price"`
	StopLoss       float64    `json:"stopLoss"`
	TakeProfit     float64    `json:"takeProfit"`
	Volume         int64      `json:"volume"`
	Profit         float64    `json:"profit"`
	ProfitRate     float64    `json:"profitRate"`
	VolumeRate     float64    `json:"volumeRate"`
	Commission     float64    `json:"commission"`
	Fee            float64    `json:"fee"`
	Swap           float64    `json:"swap"`
	ExpertId       int64      `json:"expertId"`
	PositionTicket int64      `json:"positionTicket"`
	Comment        string     `json:"comment"`
	ContractSize   float64    `json:"contractSize"`
	Digits         int32      `json:"digits"`
	MoneyDigits    int32      `json:"moneyDigits"`
	FreeProfit     float64    `json:"freeProfit"`
	TrailRounder   float64    `json:"trailRounder"`
	OpenTimeMs     int64      `json:"openTimeMs"`
	PlacedType     PlacedType `json:"placedType"`
	//OpenTimeAsDateTime time.Time  `json:"openTimeAsDateTime"`
	Lots float64 `json:"lots"`
}

// OrderInternal represents internal order information
type OrderInternal struct {
	TicketNumber   int64      `json:"ticketNumber"`
	Id             string     `json:"id"`
	Login          int64      `json:"login"`
	Symbol         string     `json:"symbol"`
	HistoryTime    int64      `json:"historyTime"`
	OpenTime       int64      `json:"openTime"`
	ExpirationTime int64      `json:"expirationTime"`
	ExecutionTime  int64      `json:"executionTime"`
	Type           OrderType  `json:"type"`
	FillPolicy     string     `json:"fillPolicy"`
	PlacedType     PlacedType `json:"placedType"`
	OpenPrice      float64    `json:"openPrice"`
	StopLimitPrice float64    `json:"stopLimitPrice"`
	Price          float64    `json:"price"`
	StopLoss       float64    `json:"stopLoss"`
	TakeProfit     float64    `json:"takeProfit"`
	Volume         int64      `json:"volume"`
	RequestVolume  int64      `json:"requestVolume"`
	State          OrderState `json:"state"`
	ExpertId       int64      `json:"expertId"`
	DealTicket     int64      `json:"dealTicket"`
	Comment        string     `json:"comment"`
	ContractSize   float64    `json:"contractSize"`
	Digits         int32      `json:"digits"`
	BaseDigits     int32      `json:"baseDigits"`
	ProfitRate     float64    `json:"profitRate"`
	OpenTimeMs     int64      `json:"openTimeMs"`
	Ticket         int64      `json:"ticket"`
	Lots           float64    `json:"lots"`
	RequestLots    float64    `json:"requestLots"`
	//OpenTimeMsAsDateTime time.Time  `json:"openTimeMsAsDateTime"`
	//OpenTimeAsDateTime   time.Time  `json:"openTimeAsDateTime"`
}

// PaginationReply represents paginated response
type PaginationReply struct {
	PagesCount int     `json:"pagesCount"`
	PageNumber int     `json:"pageNumber"`
	Orders     []Order `json:"orders"`
}

// OrderHistory gets order history for a date range
func (c *Client) OrderHistory(ctx context.Context, from, to time.Time, sort SortType, ascending bool, filter []string) (*OrderHistoryEventArgs, error) {
	params := url.Values{}
	params.Add("from", from.Format("2006-01-02T15:04:05"))
	params.Add("to", to.Format("2006-01-02T15:04:05"))
	if sort != "" {
		params.Add("sort", string(sort))
	}
	params.Add("ascending", strconv.FormatBool(ascending))
	for _, f := range filter {
		params.Add("filter", f)
	}

	body, err := c.doRequest(ctx, "GET", "/OrderHistory", params)
	if err != nil {
		return nil, err
	}

	var history OrderHistoryEventArgs
	if err := json.Unmarshal(body, &history); err != nil {
		return nil, err
	}

	return &history, nil
}

// OrderHistoryPagination gets order history with pagination
func (c *Client) OrderHistoryPagination(ctx context.Context, from, to time.Time, ordersPerPage, pageNumber int, requestAgain bool, sort SortType, ascending bool, tickets []int64, ignoreDepositWithdraw bool) (*PaginationReply, error) {
	params := url.Values{}
	params.Add("from", from.Format("2006-01-02T15:04:05"))
	params.Add("to", to.Format("2006-01-02T15:04:05"))
	params.Add("ordersPerPage", strconv.Itoa(ordersPerPage))
	params.Add("pageNumber", strconv.Itoa(pageNumber))
	params.Add("requestAgain", strconv.FormatBool(requestAgain))
	if sort != "" {
		params.Add("sort", string(sort))
	}
	params.Add("ascending", strconv.FormatBool(ascending))
	for _, ticket := range tickets {
		params.Add("tickets", strconv.FormatInt(ticket, 10))
	}
	params.Add("ignoreDepositWithdraw", strconv.FormatBool(ignoreDepositWithdraw))

	body, err := c.doRequest(ctx, "GET", "/OrderHistoryPagination", params)
	if err != nil {
		return nil, err
	}

	var reply PaginationReply
	if err := json.Unmarshal(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}

// HistoryDealsByPositionId gets history deals by position Id
func (c *Client) HistoryDealsByPositionId(ctx context.Context, ticket int64) ([]DealInternal, error) {
	params := url.Values{}
	params.Add("ticket", strconv.FormatInt(ticket, 10))

	body, err := c.doRequest(ctx, "GET", "/HistoryDealsByPositionId", params)
	if err != nil {
		return nil, err
	}

	var deals []DealInternal
	if err := json.Unmarshal(body, &deals); err != nil {
		return nil, err
	}

	return deals, nil
}

// HistoryPositions gets history positions by ticket numbers
func (c *Client) HistoryPositions(ctx context.Context, tickets []int64) ([]Order, error) {
	params := url.Values{}
	for _, ticket := range tickets {
		params.Add("tickets", strconv.FormatInt(ticket, 10))
	}

	body, err := c.doRequest(ctx, "GET", "/HistoryPositions", params)
	if err != nil {
		return nil, err
	}

	var positions []Order
	if err := json.Unmarshal(body, &positions); err != nil {
		return nil, err
	}

	return positions, nil
}

// HistoryPositionsByCloseTime gets history positions by close time
func (c *Client) HistoryPositionsByCloseTime(ctx context.Context, from, to time.Time) ([]Order, error) {
	params := url.Values{}
	params.Add("from", from.Format("2006-01-02T15:04:05"))
	params.Add("to", to.Format("2006-01-02T15:04:05"))

	body, err := c.doRequest(ctx, "GET", "/HistoryPositionsByCloseTime", params)
	if err != nil {
		return nil, err
	}

	var positions []Order
	if err := json.Unmarshal(body, &positions); err != nil {
		return nil, err
	}

	return positions, nil
}

// OrderHistoryDownloadComplete checks if order history download is complete
func (c *Client) OrderHistoryDownloadComplete(ctx context.Context) (bool, error) {
	body, err := c.doRequest(ctx, "GET", "/OrderHistoryDownloadComplete", url.Values{})
	if err != nil {
		return false, err
	}

	var complete bool
	if err := json.Unmarshal(body, &complete); err != nil {
		return false, err
	}

	return complete, nil
}
