package mt5api

import (
	"context"
	"net/url"
	"strconv"
)

// ConnectRequest represents connection parameters
type ConnectRequest struct {
	User                    int64  `json:"user"`
	Password                string `json:"password"`
	Host                    string `json:"host"`
	Port                    int    `json:"port"`
	Id                      string `json:"id,omitempty"`
	HardwareId              string `json:"hardwareId,omitempty"`
	OTP                     string `json:"otp,omitempty"`
	ConnectTimeoutSeconds   int    `json:"connectTimeoutSeconds,omitempty"`
	DownloadOrderHistory    bool   `json:"downloadOrderHistory,omitempty"`
	ReconnectOnSymbolUpdate bool   `json:"reconnectOnSymbolUpdate,omitempty"`
}

// ConnectExRequest represents extended connection parameters
type ConnectExRequest struct {
	User                               int64  `json:"user"`
	Password                           string `json:"password"`
	Server                             string `json:"server"`
	Id                                 string `json:"id,omitempty"`
	HardwareId                         string `json:"hardwareId,omitempty"`
	OTP                                string `json:"otp,omitempty"`
	ConnectTimeoutSeconds              int    `json:"connectTimeoutSeconds,omitempty"`
	ConnectTimeoutClusterMemberSeconds int    `json:"connectTimeoutClusterMemberSeconds,omitempty"`
	DownloadOrderHistory               bool   `json:"downloadOrderHistory,omitempty"`
	ReconnectOnSymbolUpdate            bool   `json:"reconnectOnSymbolUpdate,omitempty"`
}

// ConnectProxyRequest represents proxy connection parameters
type ConnectProxyRequest struct {
	ConnectRequest
	ProxyUser     string    `json:"proxyUser,omitempty"`
	ProxyPassword string    `json:"proxyPassword,omitempty"`
	ProxyHost     string    `json:"proxyHost"`
	ProxyPort     int       `json:"proxyPort"`
	ProxyType     ProxyType `json:"proxyType"`
}

// Connect establishes connection with user, password, host, port
func (c *Client) Connect(ctx context.Context, req ConnectRequest) (string, error) {
	params := url.Values{}
	params.Add("user", strconv.FormatInt(req.User, 10))
	params.Add("password", req.Password)
	params.Add("host", req.Host)
	params.Add("port", strconv.Itoa(req.Port))

	if req.Id != "" {
		params.Add("id", req.Id)
	}
	if req.HardwareId != "" {
		params.Add("hardwareId", req.HardwareId)
	}
	if req.OTP != "" {
		params.Add("otp", req.OTP)
	}
	if req.ConnectTimeoutSeconds > 0 {
		params.Add("connectTimeoutSeconds", strconv.Itoa(req.ConnectTimeoutSeconds))
	}
	if req.DownloadOrderHistory {
		params.Add("downloadOrderHistory", "true")
	}
	if req.ReconnectOnSymbolUpdate {
		params.Add("reconnectOnSymbolUpdate", "true")
	}

	body, err := c.doRequest(ctx, "GET", "/Connect", params)
	if err != nil {
		return "", err
	}

	token := string(body)
	c.SetToken(token)
	return token, nil
}

// ConnectEx establishes connection using server name instead of host/port
func (c *Client) ConnectEx(ctx context.Context, req ConnectExRequest) (string, error) {
	params := url.Values{}
	params.Add("user", strconv.FormatInt(req.User, 10))
	params.Add("password", req.Password)
	params.Add("server", req.Server)

	if req.Id != "" {
		params.Add("id", req.Id)
	}
	if req.HardwareId != "" {
		params.Add("hardwareId", req.HardwareId)
	}
	if req.OTP != "" {
		params.Add("otp", req.OTP)
	}
	if req.ConnectTimeoutSeconds > 0 {
		params.Add("connectTimeoutSeconds", strconv.Itoa(req.ConnectTimeoutSeconds))
	}
	if req.ConnectTimeoutClusterMemberSeconds > 0 {
		params.Add("connectTimeoutClusterMemberSeconds", strconv.Itoa(req.ConnectTimeoutClusterMemberSeconds))
	}
	if req.DownloadOrderHistory {
		params.Add("downloadOrderHistory", "true")
	}
	if req.ReconnectOnSymbolUpdate {
		params.Add("reconnectOnSymbolUpdate", "true")
	}

	body, err := c.doRequest(ctx, "GET", "/ConnectEx", params)
	if err != nil {
		return "", err
	}

	token := string(body)
	c.SetToken(token)
	return token, nil
}

// ConnectProxy establishes connection through proxy
func (c *Client) ConnectProxy(ctx context.Context, req ConnectProxyRequest) (string, error) {
	params := url.Values{}
	params.Add("user", strconv.FormatInt(req.User, 10))
	params.Add("password", req.Password)
	params.Add("host", req.Host)
	params.Add("port", strconv.Itoa(req.Port))
	params.Add("proxyHost", req.ProxyHost)
	params.Add("proxyPort", strconv.Itoa(req.ProxyPort))
	params.Add("proxyType", string(req.ProxyType))

	if req.ProxyUser != "" {
		params.Add("proxyUser", req.ProxyUser)
	}
	if req.ProxyPassword != "" {
		params.Add("proxyPassword", req.ProxyPassword)
	}
	if req.Id != "" {
		params.Add("id", req.Id)
	}
	if req.HardwareId != "" {
		params.Add("hardwareId", req.HardwareId)
	}
	if req.OTP != "" {
		params.Add("otp", req.OTP)
	}
	if req.ConnectTimeoutSeconds > 0 {
		params.Add("connectTimeoutSeconds", strconv.Itoa(req.ConnectTimeoutSeconds))
	}
	if req.DownloadOrderHistory {
		params.Add("downloadOrderHistory", "true")
	}
	if req.ReconnectOnSymbolUpdate {
		params.Add("reconnectOnSymbolUpdate", "true")
	}

	body, err := c.doRequest(ctx, "GET", "/ConnectProxy", params)
	if err != nil {
		return "", err
	}

	token := string(body)
	c.SetToken(token)
	return token, nil
}

// CheckConnect checks connection state and reconnects if connection lost
func (c *Client) CheckConnect(ctx context.Context) (string, error) {
	body, err := c.doRequest(ctx, "GET", "/CheckConnect", url.Values{})
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Disconnect disconnects from account
func (c *Client) Disconnect(ctx context.Context) (string, error) {
	body, err := c.doRequest(ctx, "GET", "/Disconnect", url.Values{})
	if err != nil {
		return "", err
	}

	// Clear token on successful disconnect
	c.Token = ""
	return string(body), nil
}
