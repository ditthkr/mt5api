package mt5api

// ExceptionResult Common exception result structure
type ExceptionResult struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	StackTrace string `json:"stackTrace"`
}

// OrderType Order types and enums
type OrderType string

const (
	OrderBuy           OrderType = "Buy"
	OrderSell          OrderType = "Sell"
	OrderBuyLimit      OrderType = "BuyLimit"
	OrderSellLimit     OrderType = "SellLimit"
	OrderBuyStop       OrderType = "BuyStop"
	OrderSellStop      OrderType = "SellStop"
	OrderBuyStopLimit  OrderType = "BuyStopLimit"
	OrderSellStopLimit OrderType = "SellStopLimit"
	OrderCloseBy       OrderType = "CloseBy"
	OrderBalance       OrderType = "Balance"
	OrderCredit        OrderType = "Credit"
)

type OrderState string

const (
	StateStarted           OrderState = "Started"
	StatePlaced            OrderState = "Placed"
	StateCancelled         OrderState = "Cancelled"
	StatePartial           OrderState = "Partial"
	StateFilled            OrderState = "Filled"
	StateRejected          OrderState = "Rejected"
	StateExpired           OrderState = "Expired"
	StateRequestAdding     OrderState = "RequestAdding"
	StateRequestModifying  OrderState = "RequestModifying"
	StateRequestCancelling OrderState = "RequestCancelling"
)

type PlacedType string

const (
	PlacedManually        PlacedType = "Manually"
	PlacedByExpert        PlacedType = "ByExpert"
	PlacedByDealer        PlacedType = "ByDealer"
	PlacedOnSL            PlacedType = "OnSL"
	PlacedOnTP            PlacedType = "OnTP"
	PlacedOnStopOut       PlacedType = "OnStopOut"
	PlacedOnRollover      PlacedType = "OnRollover"
	PlacedOnVmargin       PlacedType = "OnVmargin"
	PlacedGateway         PlacedType = "Gateway"
	PlacedSignal          PlacedType = "Signal"
	PlacedSettlement      PlacedType = "Settlement"
	PlacedTransfer        PlacedType = "Transfer"
	PlacedSync            PlacedType = "Sync"
	PlacedExternalService PlacedType = "ExternalService"
	PlacedMigration       PlacedType = "Migration"
	PlacedMobile          PlacedType = "Mobile"
	PlacedWeb             PlacedType = "Web"
	PlacedOnSplit         PlacedType = "OnSplit"
	PlacedDefault         PlacedType = "Default"
)

type ExpirationType string

const (
	ExpirationGTC          ExpirationType = "GTC"
	ExpirationToday        ExpirationType = "Today"
	ExpirationSpecified    ExpirationType = "Specified"
	ExpirationSpecifiedDay ExpirationType = "SpecifiedDay"
)

// AccountMethod Account related types
type AccountMethod string

const (
	AccountDefault AccountMethod = "Default"
	AccountNetting AccountMethod = "Netting"
	AccountHedging AccountMethod = "Hedging"
)

// ProxyType Proxy types
type ProxyType string

const (
	ProxyNone   ProxyType = "None"
	ProxyHttps  ProxyType = "Https"
	ProxySocks4 ProxyType = "Socks4"
	ProxySocks5 ProxyType = "Socks5"
)
