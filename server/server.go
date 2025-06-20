package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"order-api/config"
	"order-api/di"
	"order-api/services/handlers"
	"order-api/services/rmq"
	"order-api/utils"
)

type Handlers struct {
	createOrderHandler,
	userBalanceInfoHandler,
	removeOrderHandler,
	userBalanceEmitHandler,
	exchangeRateHandler,
	userOrdersHandler gin.HandlerFunc
}

var allHandlers = &Handlers{}

func Run() {
	defer rmq.Close()

	cfg := di.Get[*config.Config]("config").HttpConfig

	log.Info(fmt.Sprintf("running http api..."))
	r := gin.Default()

	createRestApi(r)

	err := r.Run(fmt.Sprintf(":%d", cfg.Port))
	utils.IsError(err, config.ErrRunServer)
}

func createRestApi(r *gin.Engine) {
	go di.Get[*rmq.Consumer]("UserBalanceEmitConsumer").ConsumeMessages()

	r.POST(config.UserBalanceEmitPath, allHandlers.userBalanceEmitHandler)
	r.POST(config.CreateOrderPath, allHandlers.createOrderHandler)
	r.DELETE(config.RemoveOrderPath, allHandlers.removeOrderHandler)
	r.GET(config.UserBalanceInfoPath, allHandlers.userBalanceInfoHandler)
	r.GET(config.ExchangeRatePath, allHandlers.exchangeRateHandler)
	r.GET(config.UserOrdersPath, allHandlers.userOrdersHandler)
}

func init() {
	initSenders()
	initConsumers()
	initHandlers()
}

func initConsumers() {
	conn := di.Get[*amqp.Connection]("RmqConnection")

	getUserBalanceConsumer := rmq.NewConsumer(conn,
		config.RabbitBalanceExchange,
		config.GetUserBalanceResponseRoutingKey,
		"q.get.user.balance.order.api",
		make(chan []byte))

	orderEventConsumer := rmq.NewConsumer(conn,
		config.RabbitEventExchange,
		config.UpdatedOrderEventRoutingKey,
		"q.order.create.user.order.api",
		make(chan []byte))

	getUserOrdersConsumer := rmq.NewConsumer(conn,
		config.RabbitOrderExchange,
		config.GetUserOrdersResponseRoutingKey,
		"q.order.get.user.orders.api",
		make(chan []byte))

	getExchangeRateConsumer := rmq.NewConsumer(conn,
		config.RabbitExchangeRateExchange,
		config.GetExchangeRateResponseRoutingKey,
		"q.order-api.exchange.rate.get.api",
		make(chan []byte))

	userBalanceEmitConsumer := rmq.NewConsumer(conn,
		config.RabbitBalanceExchange,
		"r.balance-service.EmitUserBalanceResponse", //config.EmitUserBalanceResponseRoutingKey,
		"q.order-api.user.balance.emit.response",
		make(chan []byte))

	di.Provide("GetUserBalanceConsumer", getUserBalanceConsumer)
	di.Provide("OrderEventConsumer", orderEventConsumer)
	di.Provide("GetUserOrdersConsumer", getUserOrdersConsumer)
	di.Provide("GetExchangeRateConsumer", getExchangeRateConsumer)
	di.Provide("UserBalanceEmitConsumer", userBalanceEmitConsumer)
}

func initHandlers() {
	sender := di.Get[*rmq.MessageSender]("sender")

	getUserBalanceConsumer := di.Get[*rmq.Consumer]("GetUserBalanceConsumer")
	orderEventConsumer := di.Get[*rmq.Consumer]("OrderEventConsumer")
	emitUserBalanceConsumer := di.Get[*rmq.Consumer]("UserBalanceEmitConsumer")
	getExchangeRateConsumer := di.Get[*rmq.Consumer]("GetExchangeRateConsumer")
	getUserOrdersConsumer := di.Get[*rmq.Consumer]("GetUserOrdersConsumer")

	di.Provide("UserBalanceInfoHandler", handlers.NewUserBalanceInfoHandler(sender, getUserBalanceConsumer))
	di.Provide("CreateOrderHandler", handlers.NewCreateOrderHandler(sender, orderEventConsumer))
	di.Provide("RemoveOrderHandler", handlers.NewRemoveOrderHandler(sender, orderEventConsumer))
	di.Provide("UserBalanceEmitHandler", handlers.NewUserBalanceEmitHandler(sender, emitUserBalanceConsumer))
	di.Provide("ExchangeRateHandler", handlers.NewExchangeRateHandler(sender, getExchangeRateConsumer))
	di.Provide("UserOrdersHandler", handlers.NewUserOrdersHandler(sender, getUserOrdersConsumer))

	allHandlers = &Handlers{
		createOrderHandler:     di.Get[handlers.Handler]("CreateOrderHandler").Action,
		userBalanceInfoHandler: di.Get[handlers.Handler]("UserBalanceInfoHandler").Action,
		removeOrderHandler:     di.Get[handlers.Handler]("RemoveOrderHandler").Action,
		userBalanceEmitHandler: di.Get[handlers.Handler]("UserBalanceEmitHandler").Action,
		exchangeRateHandler:    di.Get[handlers.Handler]("ExchangeRateHandler").Action,
		userOrdersHandler:      di.Get[handlers.Handler]("UserOrdersHandler").Action,
	}
}

func initSenders() {
	sender := rmq.NewMessageSender()
	di.Provide("sender", sender)
}
