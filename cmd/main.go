package main

import (
	log "github.com/sirupsen/logrus"
	"order-api/config"
	"order-api/server"
	"order-api/server/api"
	"order-api/services/rmq"
	"order-api/utils"
)

func main() {
	cfg, err := config.GetConfig()
	utils.IsError(err, config.ErrLoadConfig)

	logLvl, err := log.ParseLevel(cfg.LoggerConfig.Level)
	utils.IsError(err, config.ErrParseLog)
	log.SetLevel(logLvl)

	sender, err := rmq.NewMessageSender(&cfg.RabbitConfig)
	utils.IsError(err, config.ErrCreateSender)

	getUserBalanceConsumer := rmq.NewConsumer(sender.Connection,
		config.RabbitBalanceExchange,
		config.GetUserBalanceResponseRoutingKey,
		"q.get.user.balance.order.api",
		make(chan []byte))

	UserBalanceEmitConsumer := rmq.NewConsumer(sender.Connection,
		config.RabbitBalanceExchange,
		config.EmitUserBalanceResponseRoutingKey,
		"q.emit.user.balance.order.api",
		make(chan []byte))

	OrderEventConsumer := rmq.NewConsumer(sender.Connection,
		config.RabbitEventExchange,
		config.UpdatedOrderEventRoutingKey,
		"q.order.create.user.order.api",
		make(chan []byte))

	GetUserOrdersConsumer := rmq.NewConsumer(sender.Connection,
		config.RabbitOrderExchange,
		config.GetUserOrdersResponseRoutingKey,
		"q.order.get.user.orders.api",
		make(chan []byte))

	GetExchangeRateConsumer := rmq.NewConsumer(sender.Connection,
		config.RabbitExchangeRateExchange,
		config.GetExchangeRateResponseRoutingKey,
		"q.order-api.exchange.rate.get.api",
		make(chan []byte))

	server.Run(cfg.HttpConfig, api.NewAPIHandler(sender,
		getUserBalanceConsumer,
		UserBalanceEmitConsumer,
		OrderEventConsumer,
		GetUserOrdersConsumer,
		GetExchangeRateConsumer))
}
