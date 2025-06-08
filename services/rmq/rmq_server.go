package rmq

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"order-api/config"
	"order-api/di"
	"order-api/utils"
)

func Close() {
	log.Info("Closing amqp connection...")
	err := di.Get[*amqp.Connection]("RmqConnection").Close()
	utils.IsError(err, "failed close channel")

	log.Info("Closing amqp channel...")
	err = di.Get[*amqp.Channel]("RmqChannel").Close()
	utils.IsError(err, "failed close connection")
}

func init() {
	cfg := di.Get[*config.Config]("config").RabbitConfig

	log.Info("connecting to rmq...")
	conn, err := amqp.Dial(fmt.Sprintf(config.RmqUrlConnectionPattern, cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	di.Provide("RmqConnection", conn)

	utils.IsError(err, "err connect to rmq")

	log.Info("creating rmq chanel...")
	channel, err := conn.Channel()
	di.Provide("RmqChannel", channel)

	utils.IsError(err, "err create rmq chanel")
}
