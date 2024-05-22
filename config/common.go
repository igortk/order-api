package config

// Path
const (
	RemoveOrderPath     = "/v1/order/remove"
	CreateOrderPath     = "/v1/order/create"
	UserOrdersPath      = "/v1/user/orders"
	UserBalanceEmitPath = "/v1/user/balance/emit"
	UserBalanceInfoPath = "/v1/user/balance/info/:id"
	ExchangeRatePath    = "/v1/currency/exchange/rate"
)

// Error
const (
	ErrRunServer    = "error run server"
	ErrLoadConfig   = "error load configuration"
	ErrParseLog     = "error parse log level"
	ErrCreateSender = "error create message sender"
	ErrConnectDb    = "error connect db"
	ErrBindJson     = "error bind JSON"
)

// Response
const (
	PositiveHttpResp = "Request was send to server"
)

// RMQ URL pattern
const (
	RmqUrlConnectionPattern = "amqp://%s:%s@%s:%d/"
)

// Exchanges
const (
	RabbitEventExchange        = "e.events.forward"
	RabbitOrderExchange        = "e.orders.forward"
	RabbitBalanceExchange      = "e.balances.forward"
	RabbitExchangeRateExchange = "e.exchange.forward"
)

// RMQ rk
const (
	UpdatedOrderEventRoutingKey       = "r.event.order.OrderUpdateEvent"
	CreateOrderRequestRoutingKey      = "r.order-processing-service.order.order-api.CreateOrderRequest"
	RemoveOrderRequestRoutingKey      = "r.order-processing-service.order.order-api.RemoveOrderRequest"
	GetUserOrdersRequestRoutingKey    = "r.order-processing-service.order.order-api.GetUserOrdersRequest"
	GetUserOrdersResponseRoutingKey   = "r.order-processing-service.order.GetUserOrdersResponse"
	GetExchangeRateResponseRoutingKey = "r.quote-manager.rate.GetExchangeRateResponse"

	GetExchangeRateRequestRoutingKey  = "r.quote-manager.rate.order-api.GetExchangeRateRequest"
	EmmitUserBalanceRequestRoutingKey = "r.balance-service.balances.order-api.EmmitUserBalanceRequest"
	GetUserBalanceResponseRoutingKey  = "r.balance.GetBalanceByUserIdResponse"
	EmitUserBalanceResponseRoutingKey = "r.balance.EmitUserBalanceResponse"
	GetUserBalanceRequestRoutingKey   = "r.balance-service.balances.order-api.GetBalanceByUserIdRequest"
)

// ExchangeTypes
const (
	Topic = "topic"
)
