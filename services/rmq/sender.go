package rmq

import (
	"fmt"
	gitProto "github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"order-api/config"
	"order-api/dto/proto"
	"order-api/utils"
)

type MessageSender struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewMessageSender(cfg *config.RabbitConfig) (*MessageSender, error) {
	log.Info("creating rmq sender...")

	log.Info("connecting to rmq...")
	conn, err := amqp.Dial(fmt.Sprintf(config.RmqUrlConnectionPattern, cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	utils.IsError(err, "err connect to rmq")

	log.Info("creating rmq chanel...")
	channel, err := conn.Channel()
	utils.IsError(err, "err create rmq chanel")
	log.Info("rmq chanel was created")

	return &MessageSender{
		Connection: conn,
		Channel:    channel,
	}, nil
}

func (ms *MessageSender) SendEmmitUserBalanceRequest(message *proto.EmmitBalanceByUserIdRequest) {
	req, err := gitProto.Marshal(message)
	utils.IsError(err, "err serialize message")
	ms.publishMessage(config.RabbitBalanceExchange, config.EmmitUserBalanceRequestRoutingKey, req)
}

func (ms *MessageSender) SendCreateOrderRequest(message *proto.CreateOrderRequest) {
	req, err := gitProto.Marshal(message)
	utils.IsError(err, "err serialize message")
	ms.publishMessage(config.RabbitOrderExchange, config.CreateOrderRequestRoutingKey, req)
}

func (ms *MessageSender) SendDeleteOrderRequest(message *proto.RemoveOrderRequest) {
	req, err := gitProto.Marshal(message)
	utils.IsError(err, "err serialize message")
	ms.publishMessage(config.RabbitOrderExchange, config.CreateOrderRequestRoutingKey, req)
}

func (ms *MessageSender) SendRemoveOrderRequest(message *proto.RemoveOrderRequest) {
	req, err := gitProto.Marshal(message)
	utils.IsError(err, "err serialize message")
	ms.publishMessage(config.RabbitOrderExchange, config.RemoveOrderRequestRoutingKey, req)
}

func (ms *MessageSender) SendGetUserBalanceRequest(message *proto.GetBalanceByUserIdRequest) {
	req, err := gitProto.Marshal(message)
	utils.IsError(err, "err serialize message")
	ms.publishMessage(config.RabbitBalanceExchange, config.GetUserBalanceRequestRoutingKey, req)
}

func (ms *MessageSender) SendGetUserOrdersRequest(message *proto.GetUserOrdersRequest) {
	req, err := gitProto.Marshal(message)
	utils.IsError(err, "err serialize message")
	ms.publishMessage(config.RabbitOrderExchange, config.GetUserOrdersRequestRoutingKey, req)
}

func (ms *MessageSender) SendGetExchangeRateRequest(message *proto.GetExchangeRateRequest) {
	req, err := gitProto.Marshal(message)
	utils.IsError(err, "err serialize message")
	ms.publishMessage(config.RabbitExchangeRateExchange, config.GetExchangeRateRequestRoutingKey, req)
}

func (ms *MessageSender) Close() {
	err := ms.Channel.Close()
	utils.IsError(err, "err close the channel")

	err = ms.Connection.Close()
	utils.IsError(err, "err close the connection")
}

func (ms *MessageSender) publishMessage(exchange, routingKey string, message []byte) {
	log.Info(fmt.Sprintf("publishing to exchange [%s] by rk[%s] ...", exchange, routingKey))

	err := ms.Channel.ExchangeDeclare(
		exchange,
		config.Topic,
		true,
		false,
		false,
		false,
		nil,
	)
	utils.IsError(err, fmt.Sprintf("err exchange[%s] declare ", exchange))

	err = ms.Channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Body: message,
		},
	)
	utils.IsError(err, fmt.Sprintf("err publish msg to exchange[%s] by rk[%s] ", exchange, routingKey))

	log.Info(fmt.Sprintf("msg was publish to exchange[%s] by rk[%s] ", exchange, routingKey))
}
