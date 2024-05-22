package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"order-api/config"
	"order-api/server/api"
	"order-api/utils"
)

func Run(cfg config.HttpConfig, handler *api.Handler) {
	log.Info(fmt.Sprintf("running http api..."))
	r := gin.Default()

	r.POST(config.UserBalanceEmitPath, handler.UserBalanceEmit)
	r.POST(config.CreateOrderPath, handler.CreateOrder)
	r.DELETE(config.RemoveOrderPath, handler.RemoveOrder)
	r.GET(config.UserBalanceInfoPath, handler.UserBalanceInfo)
	r.GET(config.ExchangeRatePath, handler.ExchangeRate)
	r.GET(config.UserOrdersPath, handler.UserOrders)
	go handler.GetUserBalanceConsumer.ConsumeMessages()
	go handler.EmitUserBalanceConsumer.ConsumeMessages()
	go handler.OrderEventConsumer.ConsumeMessages()
	go handler.GetUserOrdersConsumer.ConsumeMessages()
	go handler.GetExchangeRateConsumer.ConsumeMessages()
	err := r.Run(fmt.Sprintf(":%d", cfg.Port))
	utils.IsError(err, config.ErrRunServer)
}
