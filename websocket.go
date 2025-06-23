package mt5api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"strings"
	"time"
)

// WebSocketConnection represents a WebSocket connection
type WebSocketConnection struct {
	conn     *websocket.Conn
	endpoint string
}

// ConnectWebSocket establishes a WebSocket connection for the specified endpoint
func (c *Client) ConnectWebSocket(ctx context.Context, endpoint string) (*WebSocketConnection, error) {
	// Convert HTTP URL to WebSocket URL
	wsURL := strings.Replace(c.BaseURL, "http://", "ws://", 1)
	wsURL = strings.Replace(wsURL, "https://", "wss://", 1)

	u := url.URL{
		Scheme:   "ws",
		Host:     strings.TrimPrefix(wsURL, "ws://"),
		Path:     endpoint,
		RawQuery: fmt.Sprintf("id=%s", c.Token),
	}

	// Handle wss scheme
	if strings.HasPrefix(c.BaseURL, "https://") {
		u.Scheme = "wss"
		u.Host = strings.TrimPrefix(wsURL, "wss://")
	}

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("websocket dial error: %w", err)
	}

	return &WebSocketConnection{
		conn:     conn,
		endpoint: endpoint,
	}, nil
}

// Close closes the WebSocket connection
func (ws *WebSocketConnection) Close() error {
	if ws.conn != nil {
		return ws.conn.Close()
	}
	return nil
}

// ReadMessage reads a message from the WebSocket
func (ws *WebSocketConnection) ReadMessage() ([]byte, error) {
	if ws.conn == nil {
		return nil, fmt.Errorf("websocket not connected")
	}

	_, message, err := ws.conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("websocket read error: %w", err)
	}

	return message, nil
}

// SetReadDeadline sets read deadline for the connection
func (ws *WebSocketConnection) SetReadDeadline(t time.Time) error {
	if ws.conn == nil {
		return fmt.Errorf("websocket not connected")
	}
	return ws.conn.SetReadDeadline(t)
}

// WebSocket Methods for Client

// OnOrderUpdate connects to order update WebSocket
func (c *Client) OnOrderUpdate(ctx context.Context) (*WebSocketConnection, error) {
	return c.ConnectWebSocket(ctx, "/OnOrderUpdate")
}

// OnQuote connects to quote WebSocket
func (c *Client) OnQuote(ctx context.Context) (*WebSocketConnection, error) {
	return c.ConnectWebSocket(ctx, "/OnQuote")
}

// OnOrderProfit connects to order profit WebSocket
func (c *Client) OnOrderProfit(ctx context.Context) (*WebSocketConnection, error) {
	return c.ConnectWebSocket(ctx, "/OnOrderProfit")
}

// OnOhlc connects to OHLC WebSocket
func (c *Client) OnOhlc(ctx context.Context) (*WebSocketConnection, error) {
	return c.ConnectWebSocket(ctx, "/OnOhlc")
}

// OnTickHistory connects to tick history WebSocket
func (c *Client) OnTickHistory(ctx context.Context) (*WebSocketConnection, error) {
	return c.ConnectWebSocket(ctx, "/OnTickHistory")
}

// OnMarketWatch connects to market watch WebSocket
func (c *Client) OnMarketWatch(ctx context.Context) (*WebSocketConnection, error) {
	return c.ConnectWebSocket(ctx, "/OnMarketWatch")
}

// OnOrderBook connects to order book WebSocket
func (c *Client) OnOrderBook(ctx context.Context) (*WebSocketConnection, error) {
	return c.ConnectWebSocket(ctx, "/OnOrderBook")
}

// OnTickValue connects to tick value WebSocket
func (c *Client) OnTickValue(ctx context.Context) (*WebSocketConnection, error) {
	return c.ConnectWebSocket(ctx, "/OnTickValue")
}

// OnMail connects to mail WebSocket
func (c *Client) OnMail(ctx context.Context) (*WebSocketConnection, error) {
	return c.ConnectWebSocket(ctx, "/OnMail")
}

func (c *Client) SocketOnQuote(ctx context.Context, callback func(*Quote)) {
	backoff := time.Second
	maxBackoff := 5 * time.Minute

	for {
		select {
		case <-ctx.Done():
			return
		default:
			wsConn, err := c.ConnectWebSocket(ctx, "/OnQuote")
			if err != nil {
				select {
				case <-ctx.Done():
					return
				case <-time.After(backoff):
					backoff *= 2
					if backoff > maxBackoff {
						backoff = maxBackoff
					}
					continue
				}
			}

			backoff = time.Second
			done := make(chan struct{})
			go func() {
				defer close(done)
				for {
					_, rawMessage, err := wsConn.conn.ReadMessage()
					if err != nil {
						return
					}

					var quote Quote
					if err := json.Unmarshal(rawMessage, &quote); err == nil && quote.Symbol != "" {
						callback(&quote)
						continue
					}

					var response struct {
						Type string          `json:"type"`
						Data json.RawMessage `json:"data"`
					}
					if err := json.Unmarshal(rawMessage, &response); err == nil {
						switch response.Type {
						case "Quote":
							if err := json.Unmarshal(response.Data, &quote); err == nil {
								callback(&quote)
							}
						}
					}
				}
			}()

			ticker := time.NewTicker(30 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if err := wsConn.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						return
					}
				case <-done:
					goto reconnect
				}
			}
		reconnect:
			time.Sleep(time.Second)
		}
	}
}

func (c *Client) SocketOnOrderUpdate(ctx context.Context, callback func(*OrderUpdateSummary)) {
	backoff := time.Second
	maxBackoff := 5 * time.Minute

	for {
		select {
		case <-ctx.Done():
			return
		default:
			wsConn, err := c.ConnectWebSocket(ctx, "/OnOrderUpdate")
			if err != nil {
				select {
				case <-ctx.Done():
					return
				case <-time.After(backoff):
					backoff *= 2
					if backoff > maxBackoff {
						backoff = maxBackoff
					}
					continue
				}
			}

			backoff = time.Second
			done := make(chan struct{})
			go func() {
				defer close(done)
				for {
					_, rawMessage, err := wsConn.conn.ReadMessage()
					if err != nil {
						return
					}

					var orderUpdate OrderUpdateSummary
					if err := json.Unmarshal(rawMessage, &orderUpdate); err == nil {
						callback(&orderUpdate)
						continue
					}

					var response struct {
						Type string          `json:"type"`
						Data json.RawMessage `json:"data"`
					}
					if err := json.Unmarshal(rawMessage, &response); err == nil {
						switch response.Type {
						case "OrderUpdate":
							if err := json.Unmarshal(response.Data, &orderUpdate); err == nil {
								callback(&orderUpdate)
							}
						}
					}
				}
			}()

			ticker := time.NewTicker(30 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if err := wsConn.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						return
					}
				case <-done:
					goto reconnect
				}
			}
		reconnect:
			time.Sleep(time.Second)
		}
	}
}

func (c *Client) SocketOnOrderProfit(ctx context.Context, callback func(*ProfitUpdate)) {
	backoff := time.Second
	maxBackoff := 5 * time.Minute

	for {
		select {
		case <-ctx.Done():
			return
		default:
			wsConn, err := c.ConnectWebSocket(ctx, "/OnOrderProfit")
			if err != nil {
				select {
				case <-ctx.Done():
					return
				case <-time.After(backoff):
					backoff *= 2
					if backoff > maxBackoff {
						backoff = maxBackoff
					}
					continue
				}
			}

			backoff = time.Second
			done := make(chan struct{})
			go func() {
				defer close(done)
				for {
					_, rawMessage, err := wsConn.conn.ReadMessage()
					if err != nil {
						return
					}

					var profitUpdate ProfitUpdate
					if err := json.Unmarshal(rawMessage, &profitUpdate); err == nil {
						callback(&profitUpdate)
						continue
					}

					var response struct {
						Type string          `json:"type"`
						Data json.RawMessage `json:"data"`
					}
					if err := json.Unmarshal(rawMessage, &response); err == nil {
						switch response.Type {
						case "ProfitUpdate":
							if err := json.Unmarshal(response.Data, &profitUpdate); err == nil {
								callback(&profitUpdate)
							}
						}
					}
				}
			}()

			ticker := time.NewTicker(30 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if err := wsConn.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						return
					}
				case <-done:
					goto reconnect
				}
			}
		reconnect:
			time.Sleep(time.Second)
		}
	}
}

func (c *Client) SocketOnOHLC(ctx context.Context, callback func(*OhlcSubscription)) {
	backoff := time.Second
	maxBackoff := 5 * time.Minute

	for {
		select {
		case <-ctx.Done():
			return
		default:
			wsConn, err := c.ConnectWebSocket(ctx, "/OnOhlc")
			if err != nil {
				select {
				case <-ctx.Done():
					return
				case <-time.After(backoff):
					backoff *= 2
					if backoff > maxBackoff {
						backoff = maxBackoff
					}
					continue
				}
			}

			backoff = time.Second

			done := make(chan struct{})
			go func() {
				defer close(done)
				for {
					_, rawMessage, err := wsConn.conn.ReadMessage()
					if err != nil {
						return
					}

					var ohlc OhlcSubscription
					if err := json.Unmarshal(rawMessage, &ohlc); err == nil {
						callback(&ohlc)
						continue
					}

					var response struct {
						Type string          `json:"type"`
						Data json.RawMessage `json:"data"`
					}
					if err := json.Unmarshal(rawMessage, &response); err == nil {
						switch response.Type {
						case "Ohlc":
							if err := json.Unmarshal(response.Data, &ohlc); err == nil {
								callback(&ohlc)
							}
						}
					}
				}
			}()

			ticker := time.NewTicker(30 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if err := wsConn.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						return
					}
				case <-done:
					goto reconnect
				}
			}
		reconnect:
			time.Sleep(time.Second)
		}
	}
}

func (c *Client) SocketOnMarketWatch(ctx context.Context, callback func(MarketWatch)) {
	backoff := time.Second
	maxBackoff := 5 * time.Minute

	for {
		select {
		case <-ctx.Done():
			return
		default:
			wsConn, err := c.ConnectWebSocket(ctx, "/OnMarketWatch")
			if err != nil {
				select {
				case <-ctx.Done():
					return
				case <-time.After(backoff):
					backoff *= 2
					if backoff > maxBackoff {
						backoff = maxBackoff
					}
					continue
				}
			}

			backoff = time.Second
			done := make(chan struct{})
			go func() {
				defer close(done)
				for {
					_, rawMessage, err := wsConn.conn.ReadMessage()
					if err != nil {
						return
					}

					var marketWatch MarketWatch
					if err := json.Unmarshal(rawMessage, &marketWatch); err == nil {
						callback(marketWatch)
						continue
					}

					var response struct {
						Type string          `json:"type"`
						Data json.RawMessage `json:"data"`
					}
					if err := json.Unmarshal(rawMessage, &response); err == nil {
						switch response.Type {
						case "MarketWatch":
							if err := json.Unmarshal(response.Data, &marketWatch); err == nil {
								callback(marketWatch)
							}
						}
					}
				}
			}()

			ticker := time.NewTicker(30 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if err := wsConn.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						return
					}
				case <-done:
					goto reconnect
				}
			}
		reconnect:
			time.Sleep(time.Second)
		}
	}
}

// OrderUpdateSummary represents order update summary
type OrderUpdateSummary struct {
	OpenedOrders []Order     `json:"openedOrders"`
	Update       OrderUpdate `json:"update"`
	Balance      float64     `json:"balance"`
	Equity       float64     `json:"equity"`
	Margin       float64     `json:"margin"`
	FreeMargin   float64     `json:"freeMargin"`
	Profit       float64     `json:"profit"`
	MarginLevel  float64     `json:"marginLevel"`
	Credit       float64     `json:"credit"`
	User         int64       `json:"user"`
}

// OrderUpdate represents order update details
type OrderUpdate struct {
	Trans         TransactionInfo `json:"trans"`
	OrderInternal OrderInternal   `json:"orderInternal"`
	Deal          DealInternal    `json:"deal"`
	OppositeDeal  DealInternal    `json:"oppositeDeal"`
	Order         Order           `json:"order"`
	Type          string          `json:"type"`
	CloseByTicket int64           `json:"closeByTicket"`
}

// TransactionInfo represents transaction information
type TransactionInfo struct {
	UpdateId     int32      `json:"updateId"`
	Action       int32      `json:"action"`
	TicketNumber int64      `json:"ticketNumber"`
	Currency     string     `json:"currency"`
	Id           int32      `json:"id"`
	S58          OrderType  `json:"s58"`
	OrderState   OrderState `json:"orderState"`
	OpenPrice    float64    `json:"openPrice"`
	OrderPrice   float64    `json:"orderPrice"`
	StopLoss     float64    `json:"stopLoss"`
	TakeProfit   float64    `json:"takeProfit"`
	Volume       int64      `json:"volume"`
}

// ProfitUpdate represents profit update message
type ProfitUpdate struct {
	Balance     float64 `json:"balance"`
	Credit      float64 `json:"credit"`
	Equity      float64 `json:"equity"`
	Margin      float64 `json:"margin"`
	FreeMargin  float64 `json:"freeMargin"`
	Profit      float64 `json:"profit"`
	Orders      []Order `json:"orders"`
	MarginLevel float64 `json:"marginLevel"`
	User        int64   `json:"user"`
}

// ReadOrderUpdate reads order update from WebSocket
func (ws *WebSocketConnection) ReadOrderUpdate() (*OrderUpdateSummary, error) {
	data, err := ws.ReadMessage()
	if err != nil {
		return nil, err
	}

	var update OrderUpdateSummary
	if err := json.Unmarshal(data, &update); err != nil {
		return nil, err
	}

	return &update, nil
}

// ReadQuote reads quote from WebSocket
func (ws *WebSocketConnection) ReadQuote() (*Quote, error) {
	data, err := ws.ReadMessage()
	if err != nil {
		return nil, err
	}
	type QuoteWrapper struct {
		Type  string `json:"type"`
		Quote Quote  `json:"data"`
	}
	var wrapper QuoteWrapper
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}

	return &wrapper.Quote, nil
}

// ReadProfitUpdate reads profit update from WebSocket
func (ws *WebSocketConnection) ReadProfitUpdate() (*ProfitUpdate, error) {
	data, err := ws.ReadMessage()
	if err != nil {
		return nil, err
	}

	var update ProfitUpdate
	if err := json.Unmarshal(data, &update); err != nil {
		return nil, err
	}

	return &update, nil
}

// ReadOhlcUpdate reads OHLC update from WebSocket
func (ws *WebSocketConnection) ReadOhlcUpdate() (*OhlcSubscription, error) {
	data, err := ws.ReadMessage()
	if err != nil {
		return nil, err
	}

	var ohlc OhlcSubscription
	if err := json.Unmarshal(data, &ohlc); err != nil {
		return nil, err
	}

	return &ohlc, nil
}

// ReadTickHistory reads tick history from WebSocket
func (ws *WebSocketConnection) ReadTickHistory() (*TickHistoryEventArgs, error) {
	data, err := ws.ReadMessage()
	if err != nil {
		return nil, err
	}

	var tickHistory TickHistoryEventArgs
	if err := json.Unmarshal(data, &tickHistory); err != nil {
		return nil, err
	}

	return &tickHistory, nil
}

// ReadMarketWatch reads market watch update from WebSocket
func (ws *WebSocketConnection) ReadMarketWatch() (*MarketWatch, error) {
	data, err := ws.ReadMessage()
	if err != nil {
		return nil, err
	}

	var marketWatch MarketWatch
	if err := json.Unmarshal(data, &marketWatch); err != nil {
		return nil, err
	}

	return &marketWatch, nil
}

// ReadTickValue reads tick value update from WebSocket
func (ws *WebSocketConnection) ReadTickValue() (*SymbolTickValue, error) {
	data, err := ws.ReadMessage()
	if err != nil {
		return nil, err
	}

	var tickValue SymbolTickValue
	if err := json.Unmarshal(data, &tickValue); err != nil {
		return nil, err
	}

	return &tickValue, nil
}

// ReadMail reads mail message from WebSocket
func (ws *WebSocketConnection) ReadMail() (*MailMessage, error) {
	data, err := ws.ReadMessage()
	if err != nil {
		return nil, err
	}

	var mail MailMessage
	if err := json.Unmarshal(data, &mail); err != nil {
		return nil, err
	}

	return &mail, nil
}

// Additional WebSocket types that might be missing

// OhlcSubscription represents OHLC subscription data
type OhlcSubscription struct {
	Symbol        string    `json:"symbol"`
	Timeframe     int32     `json:"timeframe"`
	Open          float64   `json:"open"`
	High          float64   `json:"high"`
	Low           float64   `json:"low"`
	Close         float64   `json:"close"`
	Time          time.Time `json:"time"`
	Volume        int64     `json:"volume"`
	TickVolume    int64     `json:"tickVolume"`
	LastQuoteTime time.Time `json:"lastQuoteTime"`
}

// TickHistoryEventArgs represents tick history event
type TickHistoryEventArgs struct {
	Symbol string    `json:"symbol"`
	Bars   []TickBar `json:"bars"`
}

// TickBar represents tick data
type TickBar struct {
	Time   time.Time `json:"time"`
	Bid    float64   `json:"bid"`
	Ask    float64   `json:"ask"`
	Last   float64   `json:"last"`
	Volume int64     `json:"volume"`
}

// MarketWatch represents market watch data
type MarketWatch struct {
	Symbol      string  `json:"symbol"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	OpenPrice   float64 `json:"openPrice"`
	ClosePrice  float64 `json:"closePrice"`
	DailyChange float64 `json:"dailyChange"`
	Bid         float64 `json:"bid"`
	Ask         float64 `json:"ask"`
	Spread      int32   `json:"spread"`
	Volume      int64   `json:"volume"`
}

// SymbolTickValue represents tick value update
type SymbolTickValue struct {
	Symbol    string  `json:"symbol"`
	TickValue float64 `json:"tickValue"`
	TickSize  float64 `json:"tickSize"`
}

// MailMessage represents mail message
type MailMessage struct {
	Id      int64     `json:"id"`
	Time    time.Time `json:"time"`
	From    string    `json:"from"`
	To      string    `json:"to"`
	Subject string    `json:"subject"`
	Body    string    `json:"body"`
}
